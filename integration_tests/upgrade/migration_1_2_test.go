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

// runUpgrade012To12 is the Go translation of
// integration_tests/upgrade/03-test-upgrade-0.12-to-1.2.sh.
//
// Assumes the chain is currently running v0.12.0 with state from script 01.
//
//nolint:funlen
func runUpgrade012To12(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary("0.12.0")
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_2)
	require.NoError(t, err)

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_2,
		BinaryVersionNew: "v" + BinaryVersionV1_2,
		Checksum:         UpgradeChecksumV1_2,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		// Script 03 uses only 3 trustees for approval (genesis quorum on
		// v0.12.0): trustee_1 proposes, trustee_2 and trustee_3 approve.
		Trustees: []string{state.Trustee1, state.Trustee2, state.Trustee3},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.2.
	// ------------------------------------------------------------------

	t.Run("VerifyPreservedVendorInfo", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)

		// `all-vendors` would also check vid_for_rollback when script 02 has
		// run. Phase 2 skips 02, so we only assert the vid set up by 01.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "vendorinfo", "all-vendors")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
	})

	t.Run("VerifyPreservedModels", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, state.ProductLabel)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		// Bash script 02 (rollback) would have replaced productLabel/partNumber
		// with `_for_rollback` values; Phase 2 skips 02 so we keep the 01 values.
		checkResponseContains(t, out, state.ProductLabel)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "model", "model-version",
				"--vid", fmt.Sprintf("%d", state.VID),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}
	})

	t.Run("VerifyPreservedCompliance", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "provisional-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid3V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "compliance", "compliance-info",
				"--vid", fmt.Sprintf("%d", state.VID),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
				"--certificationType", certificationTypeV012,
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
		}

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "device-software-compliance",
			"--cdCertificateId", cdCertificateIDV012,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
	})

	t.Run("VerifyPreservedAccounts", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew, "query", "auth", "all-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address)
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.2-era accounts, vendor info, models, compliance,
	// PKI, revocation points, additional users.
	// ------------------------------------------------------------------

	t.Run("CreatePostUpgradeAccounts", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		_ = CreateAndApproveAccount(t, dcldNew, state.VendorAccountFor1_2, "Vendor",
			VIDFor1_2, state.Trustee1, approvers)

		_ = CreateAndApproveAccount(t, dcldNew, CertificationCenterAccountFor1_2, "CertificationCenter",
			-1, state.Trustee1, approvers)

		_ = CreateAndApproveAccount(t, dcldNew, VendorAdminAccount, "VendorAdmin",
			-1, state.Trustee1, approvers)
	})

	t.Run("AddPostUpgradeUserKeys", func(t *testing.T) {
		u4, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u5, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u6, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User4Address, state.User4Pubkey = u4.address, u4.pubkey
		state.User5Address, state.User5Pubkey = u5.address, u5.pubkey
		state.User6Address, state.User6Pubkey = u6.address, u6.pubkey
	})

	t.Run("VendorInfoAddAndUpdate", func(t *testing.T) {
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--vendorName", VendorNameFor1_2,
			"--companyLegalName", CompanyLegalNameFor1_2,
			"--companyPreferredName", CompanyPreferredNameFor1_2,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_2,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update the original (v0.12-era) vendor record with 1.2-era fields.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--vendorName", vendorNameV012,
			"--companyLegalName", companyLegalNameV012,
			"--companyPreferredName", CompanyPreferredNameFor1_2,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_2,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ModelsAndVersionsFor1_2", func(t *testing.T) {
		for _, pid := range []int{PID1For1_2, PID2For1_2, PID3For1_2} {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_2),
				"--productName", ProductNameFor1_2,
				"--productLabel", ProductLabelFor1_2,
				"--partNumber", PartNumberFor1_2,
				"--from", state.VendorAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
				"--softwareVersionString", SoftwareVersionStringFor1_2,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_2),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_2),
				"--from", state.VendorAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete the 1.2-era pid_3 model.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID3For1_2),
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update the 0.12-era model's productLabel/partNumber to 1.2 values.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelFor1_2,
			"--partNumber", PartNumberFor1_2,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_2),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_2),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ComplianceFor1_2", func(t *testing.T) {
		// certify pid_1 (1.2-era)
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID1For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--softwareVersionString", SoftwareVersionStringFor1_2,
			"--certificationType", CertificationTypeFor1_2,
			"--certificationDate", CertificationDateFor1_2,
			"--cdCertificateId", CDCertificateIDFor1_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID2For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--softwareVersionString", SoftwareVersionStringFor1_2,
			"--certificationType", CertificationTypeFor1_2,
			"--provisionalDate", ProvisionalDateFor1_2,
			"--cdCertificateId", CDCertificateIDFor1_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// certify pid_2 (after provision)
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID2For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--softwareVersionString", SoftwareVersionStringFor1_2,
			"--certificationType", CertificationTypeFor1_2,
			"--certificationDate", CertificationDateFor1_2,
			"--cdCertificateId", CDCertificateIDFor1_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// revoke pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "revoke-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID2For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--softwareVersionString", SoftwareVersionStringFor1_2,
			"--certificationType", CertificationTypeFor1_2,
			"--revocationDate", CertificationDateFor1_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("AssignVidToTestRoot", func(t *testing.T) {
		// 1.2 introduces `assign-vid` for the v0.12-era test_root cert.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "assign-vid",
			"--subject", testRootCertSubject,
			"--subject-key-id", testRootCertSubjectKeyID,
			"--vid", TestRootCertVIDForAssign,
			"--from", VendorAdminAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("PKIFor1_2", func(t *testing.T) {
		// 1.2-era root_cert ladder: propose + (approve + reject + approve x3 by remaining trustees).
		// Bash uses 4 trustees here, plus trustee_5 which is also in genesis quorum.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", RootCertPathFor1_2,
			"--vid", RootCertRandomVIDFor1_2,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// trustee_2 approves, then rejects (bash quirk that exercises both paths).
		for _, action := range []string{"approve-add-x509-root-cert", "reject-add-x509-root-cert"} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", action,
				"--subject", RootCertSubjectFor1_2,
				"--subject-key-id", RootCertSubjectKeyIDFor1_2,
				"--from", state.Trustee2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// trustee_3 and trustee_4 approve (bash also approves with trustee_5 —
		// skipped here since trustee_5 isn't propagated through state).
		for _, who := range []string{state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-add-x509-root-cert",
				"--subject", RootCertSubjectFor1_2,
				"--subject-key-id", RootCertSubjectKeyIDFor1_2,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// test_root_cert (1.2): propose + 3 approvals.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", TestRootCertPathFor1_2,
			"--vid", TestRootCertVIDFor1_2,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-add-x509-root-cert",
				"--subject", TestRootCertSubjectFor1_2,
				"--subject-key-id", TestRootCertSubjectKeyIDFor1_2,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// google_root_cert (1.2): propose only.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", GoogleRootCertPathFor1_2,
			"--vid", GoogleRootCertRandomVIDFor1_2,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Intermediate cert add + revoke.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "add-x509-cert",
			"--certificate", IntermediateCertPathFor1_2,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "revoke-x509-cert",
			"--subject", IntermediateCertSubjectFor1_2,
			"--subject-key-id", IntermediateCertSubjectKeyIDFor1_2,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Propose + 3 approvals revoke 1.2 root cert.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", RootCertSubjectFor1_2,
			"--subject-key-id", RootCertSubjectKeyIDFor1_2,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-revoke-x509-root-cert",
				"--subject", RootCertSubjectFor1_2,
				"--subject-key-id", RootCertSubjectKeyIDFor1_2,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Propose revoke 1.2 test_root cert (no approvals — stays proposed).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", TestRootCertSubjectFor1_2,
			"--subject-key-id", TestRootCertSubjectKeyIDFor1_2,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("RevocationPoints", func(t *testing.T) {
		// Add → update → delete → add again (final state: one active revocation point).
		add := func(label, dataURL string) {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-revocation-point",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--revocation-type", "1",
				"--is-paa", "true",
				"--certificate", testRootCertPath,
				"--label", label,
				"--data-url", dataURL,
				"--issuer-subject-key-id", IssuerSubjectKeyID,
				"--from", state.VendorAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		add(state.ProductLabel, TestDataURL)

		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "update-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--certificate", testRootCertPath,
			"--label", state.ProductLabel,
			"--data-url", TestDataURL+"/new",
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "delete-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--label", state.ProductLabel,
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		add(ProductLabelFor1_2, TestDataURL)
	})

	t.Run("AccountFlowsFor1_2", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		// user_4: propose + 3 approvals.
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User4Address, state.User4Pubkey, "CertificationCenter", true)

		// user_5: propose + 3 approvals.
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User5Address, state.User5Pubkey, "CertificationCenter", true)

		// user_6: propose only (left in proposed state).
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User6Address, state.User6Pubkey, "CertificationCenter", false)

		// Revoke user_4 (propose + 3 approvals).
		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User4Address, true)

		// Propose revoke user_5 (no approvals).
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User5Address, false)
	})

	t.Run("ValidatorDisableEnableFlow", func(t *testing.T) {
		// Script 03 uses 2 trustee approvals (the per-script pattern from 01).
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3})
	})
}

// proposeUserAccount runs propose-add-account plus optional approvals.
func proposeUserAccount(t *testing.T, binPath, proposer string, approvers []string, address, pubkey, role string, fullApprove bool) {
	t.Helper()

	tx, err := ProposeAddAccount(binPath, address, pubkey, proposer, ProposeAddAccountArgs{
		VID: -1, Roles: role,
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	if !fullApprove {
		return
	}
	for _, who := range approvers {
		tx, err = ApproveAddAccount(binPath, address, who)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}
}

// revokeUserAccount runs propose-revoke-account plus optional approvals.
func revokeUserAccount(t *testing.T, binPath, proposer string, approvers []string, address string, fullApprove bool) {
	t.Helper()

	tx, err := ProposeRevokeAccount(binPath, address, proposer)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)

	if !fullApprove {
		return
	}
	for _, who := range approvers {
		tx, err = ApproveRevokeAccount(binPath, address, who)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	}
}
