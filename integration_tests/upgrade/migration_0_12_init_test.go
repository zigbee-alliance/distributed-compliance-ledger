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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Local constants used by script 01's body. Kept as constants rather than
// state fields because nothing in scripts 03-09 references them by name.
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

// Cert subjects / key IDs used by the PKI portion of script 01.
const (
	rootCertPath         = "integration_tests/constants/root_cert"
	rootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	rootCertSubjectKeyID = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"

	testRootCertPath         = "integration_tests/constants/test_root_cert"
	testRootCertSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	testRootCertSubjectKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

	googleRootCertPath         = "integration_tests/constants/google_root_cert"
	googleRootCertSubject      = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	googleRootCertSubjectKeyID = "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"

	intermediateCertPath = "integration_tests/constants/intermediate_cert"
)

// runInitV0_12 is the Go translation of
// integration_tests/upgrade/01-test-upgrade-initialize-0.12.sh.
//
// On entry: a clean localnet must already be running with dcld v0.12.0. On
// exit, the chain has the full menu of seeded state subsequent upgrade scripts
// depend on (accounts, vendor info, models, compliance state, x509 certs,
// validator-node).
//
//nolint:funlen
func runInitV0_12(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcld, err := EnsureBinary("0.12.0")
	require.NoError(t, err, "fetch dcld v0.12.0")

	// Mirror bash line 28: the script flips into broadcast-mode block.
	require.NoError(t, ConfigureClient(dcld), "configure dcld v0.12.0 client")

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
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	// Bash doesn't approve the vendor account: only a single trustee approval is
	// needed for the Vendor role under the v0.12.0 schema. We follow the bash.

	// CertificationCenter account (used to certify/revoke/provision models).
	ccAddr, ccPub, err := CreateKey(dcld, certificationCenterAccount)
	require.NoError(t, err, "create certification center key")
	tx, err = ProposeAddAccount(dcld, ccAddr, ccPub, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	tx, err = ApproveAddAccount(dcld, ccAddr, state.Trustee2)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

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
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	tx, err = ApproveAddAccount(dcld, trustee4Addr, state.Trustee2)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// Trustee5: jack proposes, alice + trustee4 approve.
	tx, err = ProposeAddAccount(dcld, trustee5Addr, trustee5Pub, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "Trustee",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	tx, err = ApproveAddAccount(dcld, trustee5Addr, state.Trustee2)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	tx, err = ApproveAddAccount(dcld, trustee5Addr, state.Trustee4)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- VENDOR_INFO -------------------------------------------------

	tx, err = ExecuteTxWithBin(dcld,
		"tx", "vendorinfo", "add-vendor",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--vendorName", vendorNameV012,
		"--companyLegalName", companyLegalNameV012,
		"--companyPreferredName", companyPreferredNameV012,
		"--vendorLandingPageURL", vendorLandingPageURLV012,
		"--from", state.VendorAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- MODEL / MODEL_VERSION --------------------------------------

	for _, pid := range []int{pid1V012, state.PID2, pid3V012} {
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", fmt.Sprintf("%d", deviceTypeIDV012),
			"--productName", state.ProductName,
			"--productLabel", state.ProductLabel,
			"--partNumber", state.PartNumber,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model pid=%d: %s", pid, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcld,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--softwareVersionString", softwareVersionStrV012,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumberV012),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", minSWVerV012),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", maxSWVerV012),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, "add-model-version pid=%d: %s", pid, tx.RawLog)
	}

	// Delete the third model (it'll be re-provisioned below).
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "model", "delete-model",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--pid", fmt.Sprintf("%d", pid3V012),
		"--from", state.VendorAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- COMPLIANCE -------------------------------------------------

	tx, err = ExecuteTxWithBin(dcld,
		"tx", "compliance", "certify-model",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--pid", fmt.Sprintf("%d", pid1V012),
		"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		"--softwareVersionString", softwareVersionStrV012,
		"--certificationType", certificationTypeV012,
		"--certificationDate", certificationDateV012,
		"--cdCertificateId", cdCertificateIDV012,
		"--from", certificationCenterAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	tx, err = ExecuteTxWithBin(dcld,
		"tx", "compliance", "certify-model",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--pid", fmt.Sprintf("%d", state.PID2),
		"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		"--softwareVersionString", softwareVersionStrV012,
		"--certificationType", certificationTypeV012,
		"--certificationDate", certificationDateV012,
		"--cdCertificateId", cdCertificateIDV012,
		"--from", certificationCenterAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	tx, err = ExecuteTxWithBin(dcld,
		"tx", "compliance", "revoke-model",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--pid", fmt.Sprintf("%d", state.PID2),
		"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		"--softwareVersionString", softwareVersionStrV012,
		"--certificationType", certificationTypeV012,
		"--revocationDate", certificationDateV012,
		"--from", certificationCenterAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	tx, err = ExecuteTxWithBin(dcld,
		"tx", "compliance", "provision-model",
		"--vid", fmt.Sprintf("%d", state.VID),
		"--pid", fmt.Sprintf("%d", pid3V012),
		"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		"--softwareVersionString", softwareVersionStrV012,
		"--certificationType", certificationTypeV012,
		"--provisionalDate", provisionalDateV012,
		"--cdCertificateId", cdCertificateIDV012,
		"--from", certificationCenterAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- X509 PKI ---------------------------------------------------

	// root_cert: propose + 2 approvals.
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", rootCertPath,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", rootCertSubject,
			"--subject-key-id", rootCertSubjectKeyID,
			"--from", approver,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// test_root_cert: propose + 2 approvals.
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", testRootCertPath,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", testRootCertSubject,
			"--subject-key-id", testRootCertSubjectKeyID,
			"--from", approver,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// google_root_cert: propose then REJECT (note: cert remains in proposed
	// state with a recorded rejection — used by later query verifications).
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "propose-add-x509-root-cert",
		"--certificate", googleRootCertPath,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "reject-add-x509-root-cert",
		"--subject", googleRootCertSubject,
		"--subject-key-id", googleRootCertSubjectKeyID,
		"--from", state.Trustee2,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// Intermediate cert.
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "add-x509-cert",
		"--certificate", intermediateCertPath,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// Revoke root_cert (propose + 2 approvals).
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "propose-revoke-x509-root-cert",
		"--subject", rootCertSubject,
		"--subject-key-id", rootCertSubjectKeyID,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "pki", "approve-revoke-x509-root-cert",
			"--subject", rootCertSubject,
			"--subject-key-id", rootCertSubjectKeyID,
			"--from", approver,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// Propose revoke test_root_cert (no approval — left in proposed state).
	tx, err = ExecuteTxWithBin(dcld,
		"tx", "pki", "propose-revoke-x509-root-cert",
		"--subject", testRootCertSubject,
		"--subject-key-id", testRootCertSubjectKeyID,
		"--from", state.Trustee1,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- AUTH (user_1..3 add + revoke flows) ------------------------

	// user_1: proposed + 2 approvals + propose-revoke + 2 approve-revokes.
	tx, err = ProposeAddAccount(dcld, user1.address, user1.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddAccount(dcld, user1.address, approver)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// user_2: proposed + 2 approvals (active CertCenter).
	tx, err = ProposeAddAccount(dcld, user2.address, user2.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveAddAccount(dcld, user2.address, approver)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// user_3: proposed but NOT approved (left in proposed state).
	tx, err = ProposeAddAccount(dcld, user3.address, user3.pubkey, state.Trustee1, ProposeAddAccountArgs{
		VID: -1, Roles: "CertificationCenter",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// Revoke user_1 (propose + 2 approvals).
	tx, err = ProposeRevokeAccount(dcld, user1.address, state.Trustee1)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	for _, approver := range []string{state.Trustee2, state.Trustee3} {
		tx, err = ApproveRevokeAccount(dcld, user1.address, approver)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}

	// Propose revoke user_2 (no approval — left proposed).
	tx, err = ProposeRevokeAccount(dcld, user2.address, state.Trustee1)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	// --- VALIDATOR_NODE --------------------------------------------
	MustRun(t, "AddValidatorNode", func(t *testing.T) {
		AddValidatorNode(t, state, dcld)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		// Script 01 uses 2 trustee approvals for disable-node (trustee_4 not
		// yet active in pool quorum until later sections).
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3})
	})

	// --- Final query verifications ---------------------------------

	out, err := ExecuteCLIWithBin(dcld,
		"query", "pki", "x509-cert",
		"--subject", testRootCertSubject,
		"--subject-key-id", testRootCertSubjectKeyID,
	)
	require.NoError(t, err)
	checkResponseContains(t, out, testRootCertSubject)
	checkResponseContains(t, out, testRootCertSubjectKeyID)

	out, err = ExecuteCLIWithBin(dcld,
		"query", "pki", "proposed-x509-root-cert",
		"--subject", googleRootCertSubject,
		"--subject-key-id", googleRootCertSubjectKeyID,
	)
	require.NoError(t, err)
	checkResponseContains(t, out, googleRootCertSubject)
	checkResponseContains(t, out, googleRootCertSubjectKeyID)
}

// userKey bundles the address + pubkey of a randomly-named account.
type userKey struct {
	name    string
	address string
	pubkey  string
}

// newUserKey generates a random user account locally on the test keyring.
func newUserKey(binPath string) (userKey, error) {
	name := RandomString()
	addr, pub, err := CreateKey(binPath, name)
	if err != nil {
		return userKey{}, err
	}

	return userKey{name: name, address: addr, pubkey: pub}, nil
}
