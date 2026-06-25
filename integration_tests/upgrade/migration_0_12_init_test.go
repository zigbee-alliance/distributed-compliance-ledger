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

package upgrade

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Local constants used by the initial v0.12.0 seed. Kept as constants rather
// than state fields because nothing in later upgrade steps references them by name.
const (
	pid1V012               = 1
	pid3V012               = 3
	deviceTypeIDV012       = 12345
	softwareVersionStrV012 = "1.0"
	cdVersionNumberV012    = 312
	minSWVerV012           = 1
	maxSWVerV012           = 1000

	certificationTypeV012 = "zigbee"
	certificationDateV012 = "2020-01-01T00:00:00Z"
	provisionalDateV012   = "2019-12-12T00:00:00Z"
	cdCertificateIDV012   = "12345678910abcdefgh"

	vendorNameV012             = "VendorName"
	companyLegalNameV012       = "LegalCompanyName"
	companyPreferredNameV012   = "CompanyPreferredName"
	vendorLandingPageURLV012   = "https://www.example.com"
	certificationCenterAccount = "certification_center_account_"
)

// Cert subjects / key IDs used by the PKI portion of the initial v0.12.0 seed.
const (
	rootCertPath         = "integration_tests/constants/root_cert"
	rootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	rootCertSubjectKeyID = "DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"

	testRootCertPath         = "integration_tests/constants/test_root_cert"
	testRootCertSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	testRootCertSubjectKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

	googleRootCertPath         = "integration_tests/constants/google_root_cert"
	googleRootCertSubject      = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	googleRootCertSubjectKeyID = "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"

	intermediateCertPath         = "integration_tests/constants/intermediate_cert"
	intermediateCertSubject      = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
	intermediateCertSubjectKeyID = "1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"
)

// runInitV0_12 seeds the initial chain state at dcld v0.12.0 — accounts
// (trustees, vendor, certification center, three users), vendor info,
// three models with versions, compliance state (certify/revoke/provision),
// x509 root certificates, and a validator-node.
//
// On entry: a clean localnet must already be running with dcld v0.12.0. On
// exit, the chain has the full menu of seeded state subsequent upgrade
// steps depend on.
//
//nolint:funlen
func runInitV0_12(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcld, err := EnsureBinary("0.12.0")
	require.NoError(t, err, "fetch dcld v0.12.0")
	// EnsureBinary now applies ConfigureClient automatically (sets chain-id,
	// node, keyring-backend, broadcast-mode based on the binary version).

	// Genesis-provisioned trustees are already on-chain.
	state.Trustee1, state.Trustee2, state.Trustee3 = "jack", "alice", "bob"
	state.VID, state.PID2 = 1, 2
	state.VendorAccount = "vendor_account"
	state.ProductName, state.ProductLabel, state.PartNumber = "ProductName", "ProductLabel", "RCU2205A"
	state.SoftwareVersion = 1

	// --- Account creation --------------------------------------------

	// Three random "user" accounts — proposed below to be CertificationCenter.
	user1, err := newUserKey(dcld)
	require.NoError(t, err)
	user2, err := newUserKey(dcld)
	require.NoError(t, err)
	user3, err := newUserKey(dcld)
	require.NoError(t, err)
	state.User1Address, state.User1Pubkey = user1.address, user1.pubkey
	state.User2Address, state.User2Pubkey = user2.address, user2.pubkey
	state.User3Address, state.User3Pubkey = user3.address, user3.pubkey

	// Vendor account (named "vendor_account").
	vendorAddr, vendorPub, err := CreateKey(dcld, state.VendorAccount)
	require.NoError(t, err, "create vendor key")
	tx, err := ProposeAddAccount(dcld, vendorAddr, vendorPub, state.Trustee1, ProposeAddAccountArgs{
		VID: state.VID, Roles: "Vendor",
	})
	requireTxSuccess(t, tx, err)
	// Vendor role uses the 1/3 quorum, so a single trustee propose-add
	// already satisfies the threshold on the 3-trustee genesis chain — no
	// explicit approval needed.

	// CertificationCenter account (used to certify/revoke/provision models).
	ccAddr, ccPub, err := CreateKey(dcld, certificationCenterAccount)
	require.NoError(t, err, "create certification center key")
	tx, err = ProposeAddAccount(dcld, ccAddr, ccPub, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	requireTxSuccess(t, tx, err)
	tx, err = ApproveAddAccount(dcld, ccAddr, state.Trustee2)
	requireTxSuccess(t, tx, err)

	// Trustee4 + Trustee5 (random names).
	state.Trustee4 = RandomString()
	state.Trustee5 = RandomString()

	trustee4Addr, trustee4Pub, err := CreateKey(dcld, state.Trustee4)
	require.NoError(t, err)
	trustee5Addr, trustee5Pub, err := CreateKey(dcld, state.Trustee5)
	require.NoError(t, err)

	// Trustee4: jack proposes, alice approves (genesis threshold = 2 trustees).
	tx, err = ProposeAddAccount(dcld, trustee4Addr, trustee4Pub, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "Trustee",
	})
	requireTxSuccess(t, tx, err)
	tx, err = ApproveAddAccount(dcld, trustee4Addr, state.Trustee2)
	requireTxSuccess(t, tx, err)

	// Trustee5: jack proposes, alice + trustee4 approve.
	tx, err = ProposeAddAccount(dcld, trustee5Addr, trustee5Pub, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "Trustee",
	})
	requireTxSuccess(t, tx, err)
	tx, err = ApproveAddAccount(dcld, trustee5Addr, state.Trustee2)
	requireTxSuccess(t, tx, err)
	tx, err = ApproveAddAccount(dcld, trustee5Addr, state.Trustee4)
	requireTxSuccess(t, tx, err)

	// --- VENDOR_INFO -------------------------------------------------

	tx, err = AddVendor(dcld, VendorArgs{VID: state.VID, VendorName: vendorNameV012, CompanyLegalName: companyLegalNameV012, CompanyPreferredName: companyPreferredNameV012, VendorLandingPageURL: vendorLandingPageURLV012, From: state.VendorAccount})
	requireTxSuccess(t, tx, err)

	// --- MODEL / MODEL_VERSION --------------------------------------

	for _, pid := range []int{pid1V012, state.PID2, pid3V012} {
		tx, err = AddModel(dcld, AddModelArgs{VID: state.VID, PID: pid, DeviceTypeID: deviceTypeIDV012, ProductName: state.ProductName, ProductLabel: state.ProductLabel, PartNumber: state.PartNumber, From: state.VendorAccount})
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model pid=%d: %s", pid, tx.RawLog)

		tx, err = AddModelVersion(dcld, AddModelVersionArgs{VID: state.VID, PID: pid, SoftwareVersion: state.SoftwareVersion, SoftwareVersionString: softwareVersionStrV012, CDVersionNumber: cdVersionNumberV012, MinApplicableSoftwareVersion: minSWVerV012, MaxApplicableSoftwareVersion: maxSWVerV012, From: state.VendorAccount})
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model-version pid=%d: %s", pid, tx.RawLog)
	}

	// Delete the third model (it'll be re-provisioned below).
	tx, err = DeleteModel(dcld, state.VID, pid3V012, state.VendorAccount)
	requireTxSuccess(t, tx, err)

	// --- COMPLIANCE -------------------------------------------------

	tx, err = CertifyModel(dcld, CertifyModelArgs{VID: state.VID, PID: pid1V012, SoftwareVersion: state.SoftwareVersion, SoftwareVersionString: softwareVersionStrV012, CertificationType: certificationTypeV012, CertificationDate: certificationDateV012, CDCertificateID: cdCertificateIDV012, From: certificationCenterAccount})
	requireTxSuccess(t, tx, err)

	tx, err = CertifyModel(dcld, CertifyModelArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, SoftwareVersionString: softwareVersionStrV012, CertificationType: certificationTypeV012, CertificationDate: certificationDateV012, CDCertificateID: cdCertificateIDV012, From: certificationCenterAccount})
	requireTxSuccess(t, tx, err)

	tx, err = RevokeModel(dcld, RevokeModelArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, SoftwareVersionString: softwareVersionStrV012, CertificationType: certificationTypeV012, RevocationDate: certificationDateV012, From: certificationCenterAccount})
	requireTxSuccess(t, tx, err)

	tx, err = ProvisionModel(dcld, ProvisionModelArgs{VID: state.VID, PID: pid3V012, SoftwareVersion: state.SoftwareVersion, SoftwareVersionString: softwareVersionStrV012, CertificationType: certificationTypeV012, ProvisionalDate: provisionalDateV012, CDCertificateID: cdCertificateIDV012, From: certificationCenterAccount})
	requireTxSuccess(t, tx, err)

	// --- X509 PKI ---------------------------------------------------

	// root_cert: propose + 2 approvals.
	tx, err = ProposeAddX509RootCert(dcld, rootCertPath, "", state.Trustee1)
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddX509RootCert(dcld, rootCertSubject, rootCertSubjectKeyID, approver)
		requireTxSuccess(t, tx, err)
	}

	// test_root_cert: propose + 2 approvals.
	tx, err = ProposeAddX509RootCert(dcld, testRootCertPath, "", state.Trustee1)
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddX509RootCert(dcld, testRootCertSubject, testRootCertSubjectKeyID, approver)
		requireTxSuccess(t, tx, err)
	}

	// google_root_cert: propose then REJECT (note: cert remains in proposed
	// state with a recorded rejection — used by later query verifications).
	tx, err = ProposeAddX509RootCert(dcld, googleRootCertPath, "", state.Trustee1)
	requireTxSuccess(t, tx, err)
	tx, err = RejectAddX509RootCert(dcld, googleRootCertSubject, googleRootCertSubjectKeyID, state.Trustee2)
	requireTxSuccess(t, tx, err)

	// Intermediate cert.
	tx, err = AddX509Cert(dcld, intermediateCertPath, state.Trustee1)
	requireTxSuccess(t, tx, err)

	// Revoke root_cert (propose + 2 approvals).
	tx, err = ProposeRevokeX509RootCert(dcld, rootCertSubject, rootCertSubjectKeyID, state.Trustee1)
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveRevokeX509RootCert(dcld, rootCertSubject, rootCertSubjectKeyID, approver)
		requireTxSuccess(t, tx, err)
	}

	// Propose revoke test_root_cert (no approval — left in proposed state).
	tx, err = ProposeRevokeX509RootCert(dcld, testRootCertSubject, testRootCertSubjectKeyID, state.Trustee1)
	requireTxSuccess(t, tx, err)

	// --- AUTH (user_1..3 add + revoke flows) ------------------------

	// user_1: proposed + 2 approvals + propose-revoke + 2 approve-revokes.
	tx, err = ProposeAddAccount(dcld, user1.address, user1.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddAccount(dcld, user1.address, approver)
		requireTxSuccess(t, tx, err)
	}

	// user_2: proposed + 2 approvals (active CertCenter).
	tx, err = ProposeAddAccount(dcld, user2.address, user2.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddAccount(dcld, user2.address, approver)
		requireTxSuccess(t, tx, err)
	}

	// user_3: proposed but NOT approved (left in proposed state).
	tx, err = ProposeAddAccount(dcld, user3.address, user3.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	requireTxSuccess(t, tx, err)

	// Revoke user_1 (propose + 2 approvals).
	tx, err = ProposeRevokeAccount(dcld, user1.address, state.Trustee1)
	requireTxSuccess(t, tx, err)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveRevokeAccount(dcld, user1.address, approver)
		requireTxSuccess(t, tx, err)
	}

	// Propose revoke user_2 (no approval — left proposed).
	tx, err = ProposeRevokeAccount(dcld, user2.address, state.Trustee1)
	requireTxSuccess(t, tx, err)

	// --- VALIDATOR_NODE --------------------------------------------
	MustRun(t, "AddValidatorNode", func(t *testing.T) {
		t.Helper()
		AddValidatorNode(t, state, dcld)

		// Once the validator has joined the pool, assert the new owner has a
		// queryable last-power entry.
		if state.ValidatorAddress != "" {
			_, err := QueryLastPower(dcld, state.ValidatorAddress)
			require.NoError(t, err)
		}
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		// Uses 2 trustee approvals for disable-node (trustee_4 not
		// yet active in pool quorum until later sections).
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3})
	})

	// --- Final query verifications ---------------------------------

	out, err := QueryX509Cert(dcld, testRootCertSubject, testRootCertSubjectKeyID)
	require.NoError(t, err)
	checkResponseContains(t, out, testRootCertSubject)
	checkResponseContains(t, out, testRootCertSubjectKeyID)

	out, err = QueryProposedX509RootCert(dcld, googleRootCertSubject, googleRootCertSubjectKeyID)
	require.NoError(t, err)
	checkResponseContains(t, out, googleRootCertSubject)
	checkResponseContains(t, out, googleRootCertSubjectKeyID)
}
