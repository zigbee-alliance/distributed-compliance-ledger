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
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// pkiVID=1 matches the genesis Vendor account's vid (provisioned via
// gentestaccounts.sh: "tu00001" gets vid=1). All cert fixtures
// (subjects/keyIDs/serials) come from integration_tests/constants/*
// pre-signed PEMs the suite ships.
const (
	pkiVID = 1

	// Cert fixture values are the same encoding the cli/pki package uses —
	// subjects are PrintableString-encoded (the old UTF8String DER produced
	// by an earlier openssl pinned different bytes; updates to the cert
	// files re-encoded everything as PrintableString). SKIDs / serials read
	// directly from the on-disk PEMs via `openssl x509 -ext
	// subjectKeyIdentifier` / `-serial`.
	pkiRootCertPath         = "integration_tests/constants/root_cert"
	pkiRootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	pkiRootCertSubjectKeyID = "DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"
	pkiRootCertSerialNumber = "81311506302208030248766861785118937702312370677"

	pkiIntermediateCertPath         = "integration_tests/constants/intermediate_cert"
	pkiIntermediateCertSubject      = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
	pkiIntermediateCertSubjectKeyID = "1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"

	pkiLeafCertPath         = "integration_tests/constants/leaf_cert"
	pkiLeafCertSubject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMQ0wCwYDVQQKEwRsZWFm"
	pkiLeafCertSubjectKeyID = "2A:31:8D:39:6E:50:DA:96:DF:95:C5:98:83:68:F0:58:B2:15:B3:3A"
	pkiLeafCertSerialNumber = "409691117370409054634487600348183880852961428328"

	pkiUnknownCertSubject      = "Tz11bmtub3duLWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	pkiUnknownCertSubjectKeyID = "68:99:0E:76:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
)

// pkiNotFoundQuery is a (label, args) pair for the recurring "every
// single-record cert query returns Not Found" loops. Args is the full
// `dcld query ...` command line (transport-agnostic — queryWithRetry adds
// the proxy node).
type pkiNotFoundQuery struct {
	label string
	args  []string
}

// pkiRootCertNotFoundQueries builds the loop body shared by NotFound_BeforeAdd
// and NotFound_UnknownCert — both ask the same set of single-record queries,
// only the (subject, subject-key-id) vary.
func pkiRootCertNotFoundQueries(subject, skid string) []pkiNotFoundQuery {
	bySubjAndSKID := []string{
		"x509-cert", "revoked-x509-cert",
		"proposed-x509-root-cert", "proposed-x509-root-cert-to-revoke",
		"all-child-x509-certs",
	}
	queries := make([]pkiNotFoundQuery, 0, len(bySubjAndSKID)+1)
	for _, cmd := range bySubjAndSKID {
		queries = append(queries, pkiNotFoundQuery{
			label: cmd,
			args:  PkiBySubjectAndSKID(cmd, subject, skid),
		})
	}
	queries = append(queries, pkiNotFoundQuery{
		label: "all-subject-x509-certs",
		args:  PkiBySubject("all-subject-x509-certs", subject),
	})

	return queries
}

// TestLightClientProxyPKI exercises the dcld pki module against the light
// client proxy.
func TestLightClientProxyPKI(t *testing.T) {
	skipIfDisabled(t)

	// 1. Every single-record cert query returns Not Found through the proxy
	//    before any certs are added. Also covers all-x509-root-certs /
	//    all-revoked-x509-root-certs which take no flags.
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		for _, q := range pkiRootCertNotFoundQueries(pkiRootCertSubject, pkiRootCertSubjectKeyID) {
			assertNotFoundOnProxy(t, q.label, q.args...)
		}
		assertNotFoundOnProxy(t, "all-x509-root-certs", PkiNoArgs("all-x509-root-certs")...)
		assertNotFoundOnProxy(t, "all-revoked-x509-root-certs", PkiNoArgs("all-revoked-x509-root-certs")...)
	})

	// 2. The proxy rejects bulk list queries.
	mustRun(t, "ListQueries_Rejected", func(t *testing.T) {
		t.Helper()
		assertListQueriesRejected(t, "pki",
			"all-x509-certs", "all-revoked-x509-certs",
			"all-proposed-x509-root-certs", "all-proposed-x509-root-certs-to-revoke")
	})

	// 3. Build out the cert chain on the full node: vendor proposes (rejected),
	//    trustees propose+approve a root, then vendor adds intermediate and
	//    leaf, revokes the leaf, and Jack proposes revocation of the root.
	//
	//    Account name is suffixed with utils.RandString() — the five tests in
	//    this package share one init_pool, so the keyring needs a unique
	//    entry per test (see run-all.sh).
	vendorAccount := "pki_vendor_" + utils.RandString()
	rootRef := CertRefArgs{Subject: pkiRootCertSubject, SubjectKeyID: pkiRootCertSubjectKeyID}
	leafRef := CertRefArgs{Subject: pkiLeafCertSubject, SubjectKeyID: pkiLeafCertSubjectKeyID}

	mustRun(t, "Seed_CertChain", func(t *testing.T) {
		t.Helper()
		_ = proposeVendorAccount(t, vendorAccount, pkiVID)

		// Vendor (non-Trustee) propose-root: must fail with a non-zero code —
		// only trustees are allowed to propose root certs.
		tx, err := ProposeAddX509RootCertArgs{
			Certificate: pkiRootCertPath, VID: pkiVID,
		}.Send(vendorAccount)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), tx.Code,
			"vendor must not be allowed to propose-add-x509-root-cert: %s", tx.RawLog)

		// Trustee jack proposes the root cert.
		tx, err = ProposeAddX509RootCertArgs{
			Certificate: pkiRootCertPath, VID: pkiVID,
		}.Send("jack")
		requireTxOK(t, tx, err, "jack propose-add-root")

		// Trustee alice approves — root cert is now active.
		tx, err = ApproveAddX509RootCertArgs{CertRefArgs: rootRef}.Send("alice")
		requireTxOK(t, tx, err, "alice approve-add-root")

		// Vendor adds intermediate, then leaf.
		for _, certPath := range []string{pkiIntermediateCertPath, pkiLeafCertPath} {
			tx, err = AddX509CertArgs{Certificate: certPath}.Send(vendorAccount)
			requireTxOK(t, tx, err, "add-x509-cert "+certPath)
		}

		// Vendor revokes the leaf.
		tx, err = RevokeX509CertArgs{CertRefArgs: leafRef}.Send(vendorAccount)
		requireTxOK(t, tx, err, "revoke leaf")

		// Jack proposes revocation of the root — left in proposed state.
		tx, err = ProposeRevokeX509RootCertArgs{CertRefArgs: rootRef}.Send("jack")
		requireTxOK(t, tx, err, "propose-revoke-root")
	})

	// 4. Proxy now serves the cert chain.
	//    Warm up by polling the *latest* write (propose-revoke-x509-root-cert,
	//    visible as proposed-x509-root-cert-to-revoke). Once that's visible
	//    every earlier write — root cert, intermediate, leaf, leaf revocation —
	//    is guaranteed visible too. Poll up to 30s.
	mustRun(t, "Found_AfterSeed", func(t *testing.T) {
		t.Helper()
		proposedRevokeArgs := PkiBySubjectAndSKID(
			"proposed-x509-root-cert-to-revoke", pkiRootCertSubject, pkiRootCertSubjectKeyID)
		_, qerr := queryUntilContains(LightClientProxyAddr, pkiRootCertSubjectKeyID,
			proposedRevokeArgs...)
		require.NoError(t, qerr)

		out, qerr := queryWithRetry(LightClientProxyAddr,
			PkiBySubjectAndSKID("x509-cert", pkiRootCertSubject, pkiRootCertSubjectKeyID)...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "x509-cert.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "x509-cert.subjectKeyId")
		assertContains(t, out, pkiRootCertSerialNumber, "x509-cert.serialNumber")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			PkiBySubjectAndSKID("revoked-x509-cert", pkiLeafCertSubject, pkiLeafCertSubjectKeyID)...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiLeafCertSubject, "revoked-x509-cert.subject")
		assertContains(t, out, pkiLeafCertSubjectKeyID, "revoked-x509-cert.subjectKeyId")
		assertContains(t, out, pkiLeafCertSerialNumber, "revoked-x509-cert.serialNumber")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			PkiBySubject("all-subject-x509-certs", pkiRootCertSubject)...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "all-subject-x509-certs.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "all-subject-x509-certs.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr, PkiNoArgs("all-x509-root-certs")...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "all-x509-root-certs.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "all-x509-root-certs.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr, proposedRevokeArgs...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiRootCertSubject, "proposed-x509-root-cert-to-revoke.subject")
		assertContains(t, out, pkiRootCertSubjectKeyID, "proposed-x509-root-cert-to-revoke.subjectKeyId")

		out, qerr = queryWithRetry(LightClientProxyAddr,
			PkiBySubjectAndSKID("all-child-x509-certs", pkiRootCertSubject, pkiRootCertSubjectKeyID)...)
		require.NoError(t, qerr)
		assertContains(t, out, pkiIntermediateCertSubject, "all-child-x509-certs.subject")
		assertContains(t, out, pkiIntermediateCertSubjectKeyID, "all-child-x509-certs.subjectKeyId")
	})

	// 5. Querying an unknown cert subject still returns Not Found (or, for
	//    all-revoked-x509-root-certs, an empty array since no roots have been
	//    fully revoked yet — Jack only proposed).
	mustRun(t, "NotFound_UnknownCert", func(t *testing.T) {
		t.Helper()
		for _, q := range pkiRootCertNotFoundQueries(pkiUnknownCertSubject, pkiUnknownCertSubjectKeyID) {
			assertNotFoundOnProxy(t, q.label, q.args...)
		}

		// all-revoked-x509-root-certs: empty-array response since no root has
		// been fully revoked yet (Jack only proposed).
		out, qerr := queryWithRetry(LightClientProxyAddr, PkiNoArgs("all-revoked-x509-root-certs")...)
		require.NoError(t, qerr)
		// Accept either an empty JSON array or a "Not Found" response,
		// depending on dcld version formatting.
		require.True(t,
			containsAnyLocal(out, "[]", "Not Found"),
			"expected empty array or Not Found, got: %s", string(out))
	})

	// 6. Write through the proxy is rejected.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		args := append(AddX509CertArgs{Certificate: pkiIntermediateCertPath}.Build(),
			"--from", vendorAccount)
		assertWriteRejected(t, "add-x509-cert", args...)
	})
}
