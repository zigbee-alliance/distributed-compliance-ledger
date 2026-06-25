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

// runUpgrade144To151 runs the v1.4.4 → v1.5.1 cosmovisor upgrade, then
// seeds 1.5.1-era state: new vendor (VID=65529), models that exercise the
// new commissioningModeSecondaryStepsHint field.
//
// This is the final Phase 2 step — after it runs the chain is at v1.5.1
// and Phase 1's 08/09 subtests can take over.
//
//nolint:funlen
func runUpgrade144To151(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_4_4)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_5_1)
	require.NoError(t, err)

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_5_1,
		BinaryVersionNew: BinaryVersionV1_5_1,
		Checksum:         UpgradeChecksumV1_5_1,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.5.1.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAcrossFourEras", func(t *testing.T) {
		t.Helper()
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4} {
			out, qerr := QueryVendor(dcldNew, vid)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// 0.12 pid_2 now carries 1.4.4 productLabel/partNumber (set in script 06).
		out, err := QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_4)
		checkResponseContains(t, out, PartNumberFor1_4_4)

		out, err = QueryModelVersion(dcldNew, state.VID, state.PID2, state.SoftwareVersion)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_4_4)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_4_4)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllAccounts(dcldNew)
		require.NoError(t, err)
		// Active accounts across all prior scripts.
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccounts(dcldNew)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address, state.User12Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccountsToRevoke(dcldNew)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllRevokedAccounts(dcldNew)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address, state.User10Address,
		} {
			checkResponseContains(t, out, addr)
		}

		// Single-record account variants.
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			out, err = QueryAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address, state.User12Address,
		} {
			out, err = QueryProposedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			out, err = QueryProposedAccountToRevoke(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address, state.User10Address,
		} {
			out, err = QueryRevokedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
	})

	// Bulk readback — gap-fill compliance/model/pki listings + remaining
	// single-record forms, spanning the four pre-1.5.1 eras.
	MustRun(t, "VerifyPreservedListings_1_5_1", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllVendors(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_4)

		// Model bulk listings.
		out, err = QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_4)

		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4} {
			_, err = QueryVendorModels(dcldNew, vid)
			require.NoError(t, err)
		}
		_, err = QueryAllModelVersions(dcldNew, VIDFor1_4_4, PID1For1_4_4)
		require.NoError(t, err)

		// Compliance single-record forms.
		out, err = QueryCertifiedModel(dcldNew, VIDFor1_4_4, PID1For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		_, err = QueryRevokedModel(dcldNew, VIDFor1_4_4, PID2For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)

		_, err = QueryProvisionalModel(dcldNew, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)

		_, err = QueryComplianceInfo(dcldNew, VIDFor1_4_4, PID1For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)

		for _, cdID := range []string{
			cdCertificateIDV012, CDCertificateIDFor1_2, CDCertificateIDFor1_4_3, CDCertificateIDFor1_4_4,
		} {
			out, err = QueryDeviceSoftwareCompliance(dcldNew, cdID)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		// Compliance all-* listings.
		_, err = QueryAllCertifiedModels(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllProvisionalModels(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevokedModels(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllComplianceInfo(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllDeviceSoftwareCompliance(dcldNew)
		require.NoError(t, err)

		// PKI single-record forms.
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = QueryCert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			out, err = QueryX509Cert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			_, _ = QueryNocX509Cert(dcldNew, c.subj, c.kid)
		}

		// Revoked + revocation points.
		_, _ = QueryRevokedX509Cert(dcldNew, IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2)

		_, _ = QueryRevokedNocX509RootCert(dcldNew, NOCRootCert1SubjectFor1_4_3, NOCRootCert1SubjectKeyIDFor1_4_3)

		out, err = QueryRevocationPoint(dcldNew, VIDFor1_2, ProductLabelFor1_2, IssuerSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		_, err = QueryRevocationPoints(dcldNew, IssuerSubjectKeyID)
		require.NoError(t, err)

		// PKI all-* listings.
		_, err = QueryAllCerts(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllX509Certs(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevokedX509Certs(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevokedX509RootCerts(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllNocX509Certs(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevokedNocX509RootCerts(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevokedNocX509IcaCerts(dcldNew)
		require.NoError(t, err)
		_, err = QueryAllRevocationPoints(dcldNew)
		require.NoError(t, err)

		// Subject-based listings.
		for _, c := range []struct{ subj string }{
			{RootCertWithVIDSubjectFor1_4_3},
			{TestRootCertSubjectFor1_2},
			{testRootCertSubject},
		} {
			_, _ = QueryAllSubjectCerts(dcldNew, c.subj)
			_, _ = QueryAllSubjectX509Certs(dcldNew, c.subj)
			_, _ = QueryAllNocSubjectX509Certs(dcldNew, c.subj)
		}

		// Validator (host-side).
		if state.ValidatorAddress != "" {
			out, err = QueryAllNodes(dcldNew)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.5.1-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_5_1", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_5_1, "Vendor",
			state.VIDFor1_5_1, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
		u13, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u14, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u15, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User13Address, state.User13Pubkey = u13.address, u13.pubkey
		state.User14Address, state.User14Pubkey = u14.address, u14.pubkey
		state.User15Address, state.User15Pubkey = u15.address, u15.pubkey
	})

	MustRun(t, "VendorInfoFor1_5_1", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcldNew, VendorArgs{VID: state.VIDFor1_5_1, VendorName: VendorNameFor1_5_1, CompanyLegalName: CompanyLegalNameFor1_5_1, CompanyPreferredName: CompanyPreferredNameFor1_5_1, VendorLandingPageURL: VendorLandingPageURLFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateVendor(dcldNew, VendorArgs{VID: VIDFor1_2, VendorName: VendorNameFor1_2, CompanyLegalName: CompanyLegalNameFor1_2, CompanyPreferredName: CompanyPreferredNameFor1_5_1, VendorLandingPageURL: VendorLandingPageURLFor1_5_1, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsAndVersionsFor1_5_1", func(t *testing.T) {
		t.Helper()
		// pid_1 with full 1.5-era field set (ICD, factory-reset, commissioning sec hint).
		tx, err := AddModel(dcldNew, AddModelArgs{VID: state.VIDFor1_5_1, PID: state.PID1For1_5_1, DeviceTypeID: DeviceTypeIDFor1_5_1, ProductName: ProductNameFor1_5_1, ProductLabel: state.ProductLabelFor1_5_1, PartNumber: state.PartNumberFor1_5_1, CommissioningModeSecondaryStepsHint: state.CommissioningModeSecondaryStepsHintFor1_5_1, IcdUserActiveModeTriggerHint: ICDUserActiveModeTriggerHintFor1_5_1, IcdUserActiveModeTriggerInstruction: ICDUserActiveModeTriggerInstructionFor1_5_1, FactoryResetStepsHint: FactoryResetStepsHintFor1_5_1, FactoryResetStepsInstruction: FactoryResetStepsInstructionFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)

		// pid_1 version with specificationVersion (1.5-era).
		tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: state.VIDFor1_5_1, PID: state.PID1For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, MinApplicableSoftwareVersion: state.MinApplicableSoftwareVersionFor1_5_1, MaxApplicableSoftwareVersion: state.MaxApplicableSoftwareVersionFor1_5_1, SpecificationVersion: SpecificationVersionFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)

		// pid_2 (no new fields).
		tx, err = AddModel(dcldNew, AddModelArgs{VID: state.VIDFor1_5_1, PID: state.PID2For1_5_1, DeviceTypeID: DeviceTypeIDFor1_5_1, ProductName: ProductNameFor1_5_1, ProductLabel: state.ProductLabelFor1_5_1, PartNumber: state.PartNumberFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)
		tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: state.VIDFor1_5_1, PID: state.PID2For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, MinApplicableSoftwareVersion: state.MinApplicableSoftwareVersionFor1_5_1, MaxApplicableSoftwareVersion: state.MaxApplicableSoftwareVersionFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)

		// pid_3 add + delete.
		tx, err = AddModel(dcldNew, AddModelArgs{VID: state.VIDFor1_5_1, PID: PID3For1_5_1, DeviceTypeID: DeviceTypeIDFor1_5_1, ProductName: ProductNameFor1_5_1, ProductLabel: state.ProductLabelFor1_5_1, PartNumber: state.PartNumberFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)
		tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: state.VIDFor1_5_1, PID: PID3For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, MinApplicableSoftwareVersion: state.MinApplicableSoftwareVersionFor1_5_1, MaxApplicableSoftwareVersion: state.MaxApplicableSoftwareVersionFor1_5_1, From: VendorAccountFor1_5_1})
		requireTxSuccess(t, tx, err)
		tx, err = DeleteModel(dcldNew, state.VIDFor1_5_1, PID3For1_5_1, VendorAccountFor1_5_1)
		requireTxSuccess(t, tx, err)

		// Update the 0.12 pid_2 model with 1.5.1 productLabel/partNumber.
		tx, err = UpdateModel(dcldNew, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: state.ProductLabelFor1_5_1, PartNumber: state.PartNumberFor1_5_1, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(dcldNew, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: state.MinApplicableSoftwareVersionFor1_5_1, MaxApplicableSoftwareVersion: state.MaxApplicableSoftwareVersionFor1_5_1, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceFor1_5_1", func(t *testing.T) {
		t.Helper()
		// certify pid_1
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{VID: state.VIDFor1_5_1, PID: state.PID1For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CertificationType: CertificationTypeFor1_5_1, CertificationDate: CertificationDateFor1_5_1, CDCertificateID: CDCertificateIDFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// provision pid_2
		tx, err = ProvisionModel(dcldNew, ProvisionModelArgs{VID: state.VIDFor1_5_1, PID: state.PID2For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CertificationType: CertificationTypeFor1_5_1, ProvisionalDate: ProvisionalDateFor1_5_1, CDCertificateID: CDCertificateIDFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// certify pid_2
		tx, err = CertifyModel(dcldNew, CertifyModelArgs{VID: state.VIDFor1_5_1, PID: state.PID2For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CertificationType: CertificationTypeFor1_5_1, CertificationDate: CertificationDateFor1_5_1, CDCertificateID: CDCertificateIDFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// revoke pid_2
		tx, err = RevokeModel(dcldNew, RevokeModelArgs{VID: state.VIDFor1_5_1, PID: state.PID2For1_5_1, SoftwareVersion: state.SoftwareVersionFor1_5_1, SoftwareVersionString: SoftwareVersionStringFor1_5_1, CertificationType: CertificationTypeFor1_5_1, RevocationDate: CertificationDateFor1_5_1, CDVersionNumber: CDVersionNumberFor1_5_1, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "AccountFlowsFor1_5_1", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User13Address, state.User13Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User14Address, state.User14Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User15Address, state.User15Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User13Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User14Address, false)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW 1.5.1 data. The Phase 1 subtests
	// (08/09) rely on this state being present.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNew_1_5_1_Data", func(t *testing.T) {
		t.Helper()
		out, err := QueryVendor(dcldNew, state.VIDFor1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VIDFor1_5_1)
		checkResponseContains(t, out, CompanyLegalNameFor1_5_1)

		// pid_1 has full 1.5-era fields.
		out, err = QueryGetModel(dcldNew, state.VIDFor1_5_1, state.PID1For1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint",
			state.CommissioningModeSecondaryStepsHintFor1_5_1)
		requireFieldEquals(t, out, "icdUserActiveModeTriggerHint",
			ICDUserActiveModeTriggerHintFor1_5_1)
		checkResponseContains(t, out, ICDUserActiveModeTriggerInstructionFor1_5_1)
		requireFieldEquals(t, out, "factoryResetStepsHint",
			FactoryResetStepsHintFor1_5_1)
		checkResponseContains(t, out, FactoryResetStepsInstructionFor1_5_1)

		// pid_2 with defaults.
		out, err = QueryGetModel(dcldNew, state.VIDFor1_5_1, state.PID2For1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)
		// Migration/back-fill default for this field is 1.
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", 1)

		// 0.12 pid_2 now has 1.5.1 productLabel/partNumber.
		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		checkResponseContains(t, out, state.PartNumberFor1_5_1)

		// Model version specificationVersion.
		out, err = QueryModelVersion(dcldNew, state.VIDFor1_5_1, state.PID1For1_5_1, state.SoftwareVersionFor1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_5_1)
	})
}
