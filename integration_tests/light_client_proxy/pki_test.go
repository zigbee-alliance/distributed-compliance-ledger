// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lightclientproxy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// Bash pki.sh hardcodes vid=1, which is the genesis Vendor account's vid
// (provisioned via gentestaccounts.sh: "tu00001" gets vid=1). All cert
// fixtures (subjects/keyIDs/serials) come from integration_tests/constants/*
// pre-signed PEMs the suite ships.
const (
	pkiVID = 1

	pkiRootCertPath         = "integration_tests/constants/root_cert"
	pkiRootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	pkiRootCertSubjectKeyID = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	pkiRootCertSerialNumber = "442314047376310867378175982234956458728610743315"

	pkiIntermediateCertPath         = "integration_tests/constants/intermediate_cert"
	pkiIntermediateCertSubject      = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	pkiIntermediateCertSubjectKeyID = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"

	pkiLeafCertPath         = "integration_tests/constants/leaf_cert"
	pkiLeafCertSubject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	pkiLeafCertSubjectKeyID = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	pkiLeafCertSerialNumber = "143290473708569835418599774898811724528308722063"

	pkiUnknownCertSubject      = "Tz11bmtub3duLWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	pkiUnknownCertSubjectKeyID = "68:99:0E:76:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
)

// TestLightClientProxyPKI is the Go translation of
// integration_tests/light_client_proxy/pki.sh.
//
//nolint:funlen
func TestLightClientProxyPKI(t *testing.T) {
	skipIfDisabled(t)

	// 1. Every single-record cert query returns Not Found through the proxy
	//    before any certs are added. (pki.sh lines 30-83)
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		queries := []struct {
			name string
			args []string
		}{
			{"x509-cert", []string{"x509-cert",
				"--subject", pkiRootCertSubject, "--subject-key-id", pkiRootCertSubjectKeyID}},
			{"revoked-x509-cert", []string{"revoked-x509-cert",
				"--subject", pkiRootCertSubject, "--subject-key-id", pkiRootCertSubjectKeyID}},
			{"proposed-x509-root-cert", []string{"proposed-x509-root-cert",
				"--subject", pkiRootCertSubject, "--subject-key-id", pkiRootCertSubjectKeyID}},
			{"all-subject-x509-certs", []string{"all-subject-x509-certs",
				"--subject", pkiRootCertSubject}},
			{"all-x509-root-certs", []string{"all-x509-root-certs"}},
			{"all-revoked-x509-root-certs", []string{"all-revoked-x509-root-certs"}},
			{"proposed-x509-root-cert-to-revoke", []string{"proposed-x509-root-cert-to-revoke",
				"--subject", pkiRootCertSubject, "--subject-key-id", pkiRootCertSubjectKeyID}},
			{"all-child-x509-certs", []string{"all-child-x509-certs",
				"--subject", pkiRootCertSubject, "--subject-key-id", pkiRootCertSubjectKeyID}},
		}
		for _, q := range queries {
			args := append([]string{"query", "pki"}, q.args...)
			out, qerr := queryWithRetry(LightClientProxyAddr, args...)
			require.NoError(t, qerr, "%s", q.name)
			assertContains(t, out, "Not Found", q.name)
		}
	})

	// 2. The proxy rejects bulk list queries.
	mustRun(t, "ListQueries_Rejected", func(t *testing.T) {
		t.Helper()
		for _, q := range []string{
			"all-x509-certs", "all-revoked-x509-certs",
			"all-proposed-x509-root-certs", "all-proposed-x509-root-certs-to-revoke",
		} {
			out, qerr := queryWithRetry(LightClientProxyAddr, "query", "pki", q)
			assertRejectionContains(t, out, qerr, listQueryRejection, q)
		}
	})

	// 3. Build out the cert chain on the full node: vendor proposes (rejected),
	//    trustees propose+approve a root, then vendor adds intermediate and
	//    leaf, revokes the leaf, and Jack proposes revocation of the root.
	//
	//    Account name is suffixed with utils.RandString() — bash hardcodes
	//    "vendor_account_1" but the five Go tests share one init_pool so we
	//    need a unique keyring entry per test (see run-all.sh).
	vendorAccount := "pki_vendor_" + utils.RandString()

	mustRun(t, "Seed_CertChain", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, pkiVID)

		// Vendor (non-Trustee) propose-root: must fail with a non-zero code.
		// Bash pki.sh line 142 asserts response_does_not_contain "code": 0.
		tx, err := utils.ExecuteTx(
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", pkiRootCertPath,
			"--vid", fmt.Sprintf("%d", pkiVID),
			"--from", vendorAccount,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), tx.Code,
			"vendor must not be allowed to propose-add-x509-root-cert: %s", tx.RawLog)

		// Trustee jack proposes the root cert.
		tx, err = utils.ExecuteTx(
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", pkiRootCertPath,
			"--vid", fmt.Sprintf("%d", pkiVID),
			"--from", "jack",
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "jack propose-add-root: %s", tx.RawLog)

		// Trustee alice approves — root cert is now active.
		tx, err = utils.ExecuteTx(
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", pkiRootCertSubject,
			"--subject-key-id", pkiRootCertSubjectKeyID,
			"--from", "alice",
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "alice approve-add-root: %s", tx.RawLog)

		// Vendor adds intermediate, then leaf.
		for _, certPath := range []string{pkiIntermediateCertPath, pkiLeafCertPath} {
			tx, err = utils.ExecuteTx(
				"tx", "pki", "add-x509-cert",
				"--certificate", certPath,
				"--from", vendorAccount,
				"--node", FullNodeAddr,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "add-x509-cert %s: %s", certPath, tx.RawLog)
		}

		// Vendor revokes the leaf.
		tx, err = utils.ExecuteTx(
			"tx", "pki", "revoke-x509-cert",
			"--subject", pkiLeafCertSubject,
			"--subject-key-id", pkiLeafCertSubjectKeyID,
			"--from", vendorAccount,
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "revoke leaf: %s", tx.RawLog)

		// Jack proposes revocation of the root — left in proposed state.
		tx, err = utils.ExecuteTx(
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", pkiRootCertSubject,
			"--subject-key-id", pkiRootCertSubjectKeyID,
			"--from", "jack",
			"--node", FullNodeAddr,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "propose-revoke-root: %s", tx.RawLog)
	})

	// 4. Proxy now serves the cert chain. (pki.sh lines 210-255)
	mustRun(t, "Found_AfterSeed", func(t *testing.T) {
		t.Helper()
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "pki", "x509-cert",
			"--subject", pkiRootCertSubject,
			"--subject-key-id", pkiRootCertSubjectKeyID,
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "x509-cert.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "x509-cert.subjectKeyId")
		assertContains(t, out, pkiRootCertSerialNumber, "x509-cert.serialNumber")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "pki", "revoked-x509-cert",
			"--subject", pkiLeafCertSubject,
			"--subject-key-id", pkiLeafCertSubjectKeyID,
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiLeafCertSubject, "revoked-x509-cert.subject")
		assertContains(t, out, pkiLeafCertSubjectKeyID, "revoked-x509-cert.subjectKeyId")
		assertContains(t, out, pkiLeafCertSerialNumber, "revoked-x509-cert.serialNumber")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "pki", "all-subject-x509-certs",
			"--subject", pkiRootCertSubject,
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "all-subject-x509-certs.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "all-subject-x509-certs.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "pki", "all-x509-root-certs",
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "all-x509-root-certs.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "all-x509-root-certs.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "pki", "proposed-x509-root-cert-to-revoke",
			"--subject", pkiRootCertSubject,
			"--subject-key-id", pkiRootCertSubjectKeyID,
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "proposed-x509-root-cert-to-revoke.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "proposed-x509-root-cert-to-revoke.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			"query", "pki", "all-child-x509-certs",
			"--subject", pkiRootCertSubject,
			"--subject-key-id", pkiRootCertSubjectKeyID,
		)
		require.NoError(t, qerr)
		assertContains(t, out, pkiIntermediateCertSubject, "all-child-x509-certs.subject")
		assertContains(t, out, pkiIntermediateCertSubjectKeyID, "all-child-x509-certs.subjectKeyId")
	})

	// 5. Querying an unknown cert subject still returns Not Found (or, for
	//    all-revoked-x509-root-certs, an empty array since no roots have been
	//    fully revoked yet — Jack only proposed). (pki.sh lines 266-311)
	mustRun(t, "NotFound_UnknownCert", func(t *testing.T) {
		t.Helper()
		queries := []struct {
			name string
			args []string
		}{
			{"x509-cert", []string{"x509-cert",
				"--subject", pkiUnknownCertSubject, "--subject-key-id", pkiUnknownCertSubjectKeyID}},
			{"proposed-x509-root-cert", []string{"proposed-x509-root-cert",
				"--subject", pkiUnknownCertSubject, "--subject-key-id", pkiUnknownCertSubjectKeyID}},
			{"revoked-x509-cert", []string{"revoked-x509-cert",
				"--subject", pkiUnknownCertSubject, "--subject-key-id", pkiUnknownCertSubjectKeyID}},
			{"all-subject-x509-certs", []string{"all-subject-x509-certs",
				"--subject", pkiUnknownCertSubject}},
			{"proposed-x509-root-cert-to-revoke", []string{"proposed-x509-root-cert-to-revoke",
				"--subject", pkiUnknownCertSubject, "--subject-key-id", pkiUnknownCertSubjectKeyID}},
			{"all-child-x509-certs", []string{"all-child-x509-certs",
				"--subject", pkiUnknownCertSubject, "--subject-key-id", pkiUnknownCertSubjectKeyID}},
		}
		for _, q := range queries {
			args := append([]string{"query", "pki"}, q.args...)
			out, qerr := queryWithRetry(LightClientProxyAddr, args...)
			require.NoError(t, qerr, "%s", q.name)
			assertContains(t, out, "Not Found", q.name)
		}

		// all-revoked-x509-root-certs: empty-array response since no root has
		// been fully revoked yet (Jack only proposed at line 192).
		out, qerr := queryWithRetry(LightClientProxyAddr,
			"query", "pki", "all-revoked-x509-root-certs",
		)
		require.NoError(t, qerr)
		// Bash asserts "\[\]" — accept either an empty JSON array or a
		// "Not Found" response, depending on dcld version formatting.
		require.True(t,
			containsAnyLocal(out, "[]", "Not Found"),
			"expected empty array or Not Found, got: %s", string(out))
	})

	// 6. Write through the proxy is rejected. (pki.sh lines 322-325)
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		out, err := executeCLIWithNode(LightClientProxyAddr,
			"tx", "pki", "add-x509-cert",
			"--certificate", pkiIntermediateCertPath,
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		assertRejectionContains(t, out, err, writeRejection, "add-x509-cert")
	})
}
