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

// runUpgrade012To12 runs the v0.12.0 → v1.2.2 cosmovisor upgrade, then
// seeds 1.2-era state: a new vendor account (VID=4701), models with
// compliance certifications, PKI root cert ladder, revocation point, and
// three new users.
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
		BinaryVersionNew: BinaryVersionV1_2,
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

	MustRun(t, "VerifyPreservedVendorInfo", func(t *testing.T) {
		t.Helper()
		out, err := QueryVendor(dcldNew, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)

		// `all-vendors` carries both vid (script 01) and vid_for_rollback
		// (added by script 02), along with their respective legal/vendor names.
		out, err = QueryAllVendors(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDForRollback)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, CompanyLegalNameForRollback)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, VendorNameForRollback)
	})

	MustRun(t, "VerifyPreservedModels", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, state.VID, pid1V012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, state.ProductLabel)

		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		// Script 02 (rollback) updated pid_2's productLabel/partNumber with
		// the `_for_rollback` values; assert those overwrites are intact.
		checkResponseContains(t, out, ProductLabelForRollback)
		checkResponseContains(t, out, PartNumberForRollback)

		out, err = QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryVendorModels(dcldNew, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryModelVersion(dcldNew, state.VID, pid, state.SoftwareVersion)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}
	})

	MustRun(t, "VerifyPreservedCompliance", func(t *testing.T) {
		t.Helper()
		out, err := QueryCertifiedModel(dcldNew, state.VID, pid1V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryRevokedModel(dcldNew, state.VID, state.PID2, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryProvisionalModel(dcldNew, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryComplianceInfo(dcldNew, state.VID, pid, state.SoftwareVersion, certificationTypeV012)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
		}

		out, err = QueryDeviceSoftwareCompliance(dcldNew, cdCertificateIDV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = QueryAllProposedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address)

		out, err = QueryAllProposedAccountsToRevoke(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = QueryAllRevokedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address)
	})

	// Bulk `all-*` listings and pki / validator readbacks. Gap-filling
	// assertions that confirm script 01's state survives the v0.12 → v1.2
	// upgrade in aggregate form.
	MustRun(t, "VerifyPreservedListings_1_2", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllCertifiedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryAllProvisionalModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		out, err = QueryAllRevokedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryAllComplianceInfo(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryAllDeviceSoftwareCompliance(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, cdCertificateIDV012)

		// PKI listings preserved from v0.12.
		out, err = QueryAllX509RootCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryAllRevokedX509RootCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, rootCertSubject)
		checkResponseContains(t, out, rootCertSubjectKeyID)

		out, err = QueryAllProposedX509RootCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		out, err = QueryAllProposedX509RootCertsToRevoke(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryX509Cert(dcldNew, testRootCertSubject, testRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryProposedX509RootCert(dcldNew, googleRootCertSubject, googleRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		// Validator queries — run host-side against the chain. Gated on the
		// container being initialized.
		if state.ValidatorAddress != "" {
			out, err = QueryProposedDisableNode(dcldNew, state.ValidatorAddress)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)

			out, err = QueryAllNodes(dcldNew)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.2-era accounts, vendor info, models, compliance,
	// PKI, revocation points, additional users.
	// ------------------------------------------------------------------

	MustRun(t, "CreatePostUpgradeAccounts", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		_ = CreateAndApproveAccount(t, dcldNew, state.VendorAccountFor1_2, "Vendor",
			VIDFor1_2, state.Trustee1, approvers)

		_ = CreateAndApproveAccount(t, dcldNew, CertificationCenterAccountFor1_2, "CertificationCenter",
			-1, state.Trustee1, approvers)

		_ = CreateAndApproveAccount(t, dcldNew, VendorAdminAccount, "VendorAdmin",
			-1, state.Trustee1, approvers)
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "VendorInfoAddAndUpdate", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcldNew, VendorArgs{VID: VIDFor1_2, VendorName: VendorNameFor1_2, CompanyLegalName: CompanyLegalNameFor1_2, CompanyPreferredName: CompanyPreferredNameFor1_2, VendorLandingPageURL: VendorLandingPageURLFor1_2, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// Update the original (v0.12-era) vendor record with 1.2-era fields.
		tx, err = UpdateVendor(dcldNew, VendorArgs{VID: state.VID, VendorName: vendorNameV012, CompanyLegalName: companyLegalNameV012, CompanyPreferredName: CompanyPreferredNameFor1_2, VendorLandingPageURL: VendorLandingPageURLFor1_2, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsAndVersionsFor1_2", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1For1_2, PID2For1_2, PID3For1_2} {
			tx, err := AddModel(dcldNew, AddModelArgs{VID: VIDFor1_2, PID: pid, DeviceTypeID: DeviceTypeIDFor1_2, ProductName: ProductNameFor1_2, ProductLabel: ProductLabelFor1_2, PartNumber: PartNumberFor1_2, From: state.VendorAccountFor1_2})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: VIDFor1_2, PID: pid, SoftwareVersion: SoftwareVersionFor1_2, SoftwareVersionString: SoftwareVersionStringFor1_2, CDVersionNumber: CDVersionNumberFor1_2, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_2, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_2, From: state.VendorAccountFor1_2})
			requireTxSuccess(t, tx, err)
		}

		// Delete the 1.2-era pid_3 model.
		tx, err := DeleteModel(dcldNew, VIDFor1_2, PID3For1_2, state.VendorAccountFor1_2)
		requireTxSuccess(t, tx, err)

		// Update the 0.12-era model's productLabel/partNumber to 1.2 values.
		tx, err = UpdateModel(dcldNew, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: ProductLabelFor1_2, PartNumber: PartNumberFor1_2, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(dcldNew, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_2, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_2, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceFor1_2", func(t *testing.T) {
		t.Helper()
		// certify pid_1 (1.2-era)
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_2, PID: PID1For1_2, SoftwareVersion: SoftwareVersionFor1_2, SoftwareVersionString: SoftwareVersionStringFor1_2, CertificationType: CertificationTypeFor1_2, CertificationDate: CertificationDateFor1_2, CDCertificateID: CDCertificateIDFor1_2, CDVersionNumber: CDVersionNumberFor1_2, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// provision pid_2
		tx, err = ProvisionModel(dcldNew, ProvisionModelArgs{VID: VIDFor1_2, PID: PID2For1_2, SoftwareVersion: SoftwareVersionFor1_2, SoftwareVersionString: SoftwareVersionStringFor1_2, CertificationType: CertificationTypeFor1_2, ProvisionalDate: ProvisionalDateFor1_2, CDCertificateID: CDCertificateIDFor1_2, CDVersionNumber: CDVersionNumberFor1_2, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// certify pid_2 (after provision)
		tx, err = CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_2, PID: PID2For1_2, SoftwareVersion: SoftwareVersionFor1_2, SoftwareVersionString: SoftwareVersionStringFor1_2, CertificationType: CertificationTypeFor1_2, CertificationDate: CertificationDateFor1_2, CDCertificateID: CDCertificateIDFor1_2, CDVersionNumber: CDVersionNumberFor1_2, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// revoke pid_2
		tx, err = RevokeModel(dcldNew, RevokeModelArgs{VID: VIDFor1_2, PID: PID2For1_2, SoftwareVersion: SoftwareVersionFor1_2, SoftwareVersionString: SoftwareVersionStringFor1_2, CertificationType: CertificationTypeFor1_2, RevocationDate: CertificationDateFor1_2, CDVersionNumber: CDVersionNumberFor1_2, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "AssignVidToTestRoot", func(t *testing.T) {
		t.Helper()
		// 1.2 introduces `assign-vid` for the v0.12-era test_root cert.
		tx, err := AssignVid(dcldNew, testRootCertSubject, testRootCertSubjectKeyID, TestRootCertVIDForAssign, VendorAdminAccount)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "PKIFor1_2", func(t *testing.T) {
		t.Helper()
		// 1.2-era root_cert ladder: propose + (approve + reject + approve x3 by
		// remaining trustees). Uses 4 trustees + trustee_5 (added in script 02)
		// for the 5-trustee quorum.
		tx, err := ProposeAddX509RootCert(dcldNew, RootCertPathFor1_2, RootCertRandomVIDFor1_2, state.Trustee1)
		requireTxSuccess(t, tx, err)

		// trustee_2 approves, then rejects (exercises both approval paths).
		tx, err = ApproveAddX509RootCert(dcldNew, RootCertSubjectFor1_2, RootCertSubjectKeyIDFor1_2, state.Trustee2)
		requireTxSuccess(t, tx, err)
		tx, err = RejectAddX509RootCert(dcldNew, RootCertSubjectFor1_2, RootCertSubjectKeyIDFor1_2, state.Trustee2)
		requireTxSuccess(t, tx, err)

		// trustee_3, trustee_4, trustee_5 approve.
		for _, who := range []string{state.Trustee3, state.Trustee4, state.Trustee5} {
			tx, err = ApproveAddX509RootCert(dcldNew, RootCertSubjectFor1_2, RootCertSubjectKeyIDFor1_2, who)
			requireTxSuccess(t, tx, err)
		}

		// test_root_cert (1.2): propose + 3 approvals.
		tx, err = ProposeAddX509RootCert(dcldNew, TestRootCertPathFor1_2, TestRootCertVIDFor1_2, state.Trustee1)
		requireTxSuccess(t, tx, err)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveAddX509RootCert(dcldNew, TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2, who)
			requireTxSuccess(t, tx, err)
		}

		// google_root_cert (1.2): propose only.
		tx, err = ProposeAddX509RootCert(dcldNew, GoogleRootCertPathFor1_2, GoogleRootCertRandomVIDFor1_2, state.Trustee1)
		requireTxSuccess(t, tx, err)

		// Intermediate cert add + revoke.
		tx, err = AddX509Cert(dcldNew, IntermediateCertPathFor1_2, state.VendorAccount)
		requireTxSuccess(t, tx, err)
		tx, err = RevokeX509Cert(dcldNew, IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2, "", state.VendorAccount)
		requireTxSuccess(t, tx, err)

		// Propose + 3 approvals revoke 1.2 root cert.
		tx, err = ProposeRevokeX509RootCert(dcldNew, RootCertSubjectFor1_2, RootCertSubjectKeyIDFor1_2, state.Trustee1)
		requireTxSuccess(t, tx, err)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveRevokeX509RootCert(dcldNew, RootCertSubjectFor1_2, RootCertSubjectKeyIDFor1_2, who)
			requireTxSuccess(t, tx, err)
		}

		// Propose revoke 1.2 test_root cert (no approvals — stays proposed).
		tx, err = ProposeRevokeX509RootCert(dcldNew, TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2, state.Trustee1)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "RevocationPoints", func(t *testing.T) {
		t.Helper()
		// Add → update → delete → add again (final state: one active revocation point).
		add := func(label, dataURL string) {
			tx, err := AddRevocationPoint(dcldNew, AddRevocationPointArgs{VID: VIDFor1_2, RevocationType: "1", IsPAA: true, Certificate: testRootCertPath, Label: label, DataURL: dataURL, IssuerSubjectKeyID: IssuerSubjectKeyID, From: state.VendorAccountFor1_2})
			requireTxSuccess(t, tx, err)
		}

		add(state.ProductLabel, TestDataURL)

		tx, err := UpdateRevocationPoint(dcldNew, UpdateRevocationPointArgs{VID: VIDFor1_2, Certificate: testRootCertPath, Label: state.ProductLabel, DataURL: TestDataURL + "/new", IssuerSubjectKeyID: IssuerSubjectKeyID, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)

		tx, err = DeleteRevocationPoint(dcldNew, VIDFor1_2, state.ProductLabel, IssuerSubjectKeyID, state.VendorAccountFor1_2)
		requireTxSuccess(t, tx, err)

		add(ProductLabelFor1_2, TestDataURL)
	})

	MustRun(t, "AccountFlowsFor1_2", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
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
	requireTxSuccess(t, tx, err)

	if !fullApprove {
		return
	}
	for _, who := range approvers {
		tx, err = ApproveAddAccount(binPath, address, who)
		requireTxSuccess(t, tx, err)
	}
}

// revokeUserAccount runs propose-revoke-account plus optional approvals.
func revokeUserAccount(t *testing.T, binPath, proposer string, approvers []string, address string, fullApprove bool) {
	t.Helper()

	tx, err := ProposeRevokeAccount(binPath, address, proposer)
	requireTxSuccess(t, tx, err)

	if !fullApprove {
		return
	}
	for _, who := range approvers {
		tx, err = ApproveRevokeAccount(binPath, address, who)
		requireTxSuccess(t, tx, err)
	}
}
