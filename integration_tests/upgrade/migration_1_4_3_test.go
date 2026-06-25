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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// runUpgrade12To143 runs the v1.2 → v1.4.3 cosmovisor upgrade, then seeds
// 1.4.3-era state including the Issue #593 ghost-version pre-setup, a new
// vendor (VID=65521), DA + PAA root certificates, NOC cert add/remove,
// and a revocation point.
//
// Assumes the chain is currently running v1.2 with state from the preceding steps.
//
//nolint:funlen
func runUpgrade12To143(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_2)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_4_3)
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// ISSUE #593 pre-upgrade: add a model with 2 versions, delete one. The
	// resulting ghost-version state is what the 1.6.0 step verifies cleanup of.
	// State fields populated here come from DefaultBashState() (vid_for_1_2 = 4701
	// and pid_3_for_1_6_0 = 160).
	// ------------------------------------------------------------------
	MustRun(t, "Issue593PreUpgradeGhostSetup", func(t *testing.T) {
		t.Helper()
		vid := state.VIDFor1_6_0FromScript5
		pid := state.PID3For1_6_0FromScript5
		sv1 := state.SoftwareVersion1For1_6_0FromScript5
		sv2 := state.SoftwareVersion2For1_6_0FromScript5

		// Add the model (1.6.0-era values reused as defaults).
		tx, err := AddModel(dcldOld, AddModelArgs{VID: vid, PID: pid, DeviceTypeID: DeviceTypeIDForIssue593, ProductName: ProductNameForIssue593, ProductLabel: ProductLabelForIssue593, PartNumber: PartNumberForIssue593, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// Two model versions: sv1 (will be deleted to create ghost state) and sv2.
		for _, sv := range []int{sv1, sv2} {
			tx, err = AddModelVersion(dcldOld, AddModelVersionArgs{VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: SoftwareVersionStringIssue593, CDVersionNumber: CDVersionNumberIssue593, MinApplicableSoftwareVersion: MinSWVerIssue593, MaxApplicableSoftwareVersion: MaxSWVerIssue593, From: state.VendorAccountFor1_2})
			requireTxSuccess(t, tx, err)
		}

		// Delete sv1 — sv2 stays as a ghost-pointer the migration must clean.
		tx, err = DeleteModelVersion(dcldOld, vid, pid, sv1, state.VendorAccountFor1_2)
		requireTxSuccess(t, tx, err)
	})

	// ------------------------------------------------------------------
	// Upgrade 1.2 → 1.4.
	// ------------------------------------------------------------------
	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_4,
		BinaryVersionNew: BinaryVersionV1_4_3,
		Checksum:         UpgradeChecksumV1_4,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.4.3.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedVendorInfoAndModels", func(t *testing.T) {
		t.Helper()
		// VendorInfo for the v0.12 vendor — the 1.2 step updated companyPreferredName/landing URL to 1.2 values.
		out, err := QueryVendor(dcldNew, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, CompanyPreferredNameFor1_2)
		checkResponseContains(t, out, VendorLandingPageURLFor1_2)

		// VendorInfo for the 1.2 vendor.
		out, err = QueryVendor(dcldNew, VIDFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		out, err = QueryAllVendors(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		// Carry-over models from the v0.12 and v1.2 steps.
		for _, pair := range [][2]int{
			{state.VID, pid1V012}, {state.VID, state.PID2},
			{VIDFor1_2, PID1For1_2}, {VIDFor1_2, PID2For1_2},
		} {
			out, err = QueryGetModel(dcldNew, pair[0], pair[1])
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", pair[0])
			requireFieldEquals(t, out, "pid", pair[1])
		}

		// Updated 0.12 pid_2 has 1.2 min/max applicable software version.
		out, err = QueryModelVersion(dcldNew, state.VID, state.PID2, state.SoftwareVersion)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_2)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_2)
	})

	MustRun(t, "VerifyPreservedCompliance", func(t *testing.T) {
		t.Helper()
		// Certified 0.12 pid_1.
		out, err := QueryCertifiedModel(dcldNew, state.VID, pid1V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		// Certified 1.2 pid_1.
		out, err = QueryCertifiedModel(dcldNew, VIDFor1_2, PID1For1_2, SoftwareVersionFor1_2, CertificationTypeFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		// Revoked models persist.
		out, err = QueryRevokedModel(dcldNew, state.VID, state.PID2, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)

		out, err = QueryRevokedModel(dcldNew, VIDFor1_2, PID2For1_2, SoftwareVersionFor1_2, CertificationTypeFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address) // active from the v0.12 step
		checkResponseContains(t, out, state.User5Address) // active from the v1.2 step

		out, err = QueryAllProposedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address) // proposed from the v0.12 step
		checkResponseContains(t, out, state.User6Address) // proposed from the v1.2 step

		out, err = QueryAllProposedAccountsToRevoke(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)

		out, err = QueryAllRevokedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address) // revoked in the v0.12 step
		checkResponseContains(t, out, state.User4Address) // revoked in the v1.2 step

		// Single-record account variants.
		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = QueryAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User6Address, state.User3Address} {
			out, err = QueryProposedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = QueryProposedAccountToRevoke(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User4Address, state.User1Address} {
			out, err = QueryRevokedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
	})

	// Bulk readback — gap-fill queries covering compliance/PKI/model listings
	// and the remaining single-record pki/compliance forms not exercised by
	// the blocks above.
	MustRun(t, "VerifyPreservedListings_1_4_3", func(t *testing.T) {
		t.Helper()
		// ----- Model bulk listings -----
		out, err := QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		for _, vid := range []int{state.VID, VIDFor1_2} {
			_, err = QueryVendorModels(dcldNew, vid)
			require.NoError(t, err)
		}

		out, err = QueryAllModelVersions(dcldNew, VIDFor1_2, PID1For1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		// ----- Compliance: single-record forms not covered by VerifyPreservedCompliance -----
		out, err = QueryProvisionalModel(dcldNew, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryComplianceInfo(dcldNew, state.VID, pid, state.SoftwareVersion, certificationTypeV012)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
		}
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = QueryComplianceInfo(dcldNew, VIDFor1_2, pid, SoftwareVersionFor1_2, CertificationTypeFor1_2)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
		}

		for _, cdID := range []string{cdCertificateIDV012, CDCertificateIDFor1_2} {
			out, err = QueryDeviceSoftwareCompliance(dcldNew, cdID)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		// Compliance all-* listings.
		out, err = QueryAllCertifiedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		out, err = QueryAllProvisionalModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		out, err = QueryAllRevokedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = QueryAllComplianceInfo(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "vid", VIDFor1_2)

		out, err = QueryAllDeviceSoftwareCompliance(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, cdCertificateIDV012)
		checkResponseContains(t, out, CDCertificateIDFor1_2)

		// ----- PKI single-record + listings -----
		for _, c := range []struct{ subj, kid string }{
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = QueryX509Cert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)

			out, err = QueryAllSubjectX509Certs(dcldNew, c.subj)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)

			out, err = QueryProposedX509RootCertToRevoke(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)
		}

		for _, c := range []struct{ subj, kid string }{
			{GoogleRootCertSubjectFor1_2, GoogleRootCertSubjectKeyIDFor1_2},
			{googleRootCertSubject, googleRootCertSubjectKeyID},
		} {
			out, err = QueryProposedX509RootCert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)
		}

		for _, c := range []struct{ subj, kid string }{
			{IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2},
			{intermediateCertSubject, intermediateCertSubjectKeyID},
		} {
			out, err = QueryRevokedX509Cert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)
		}

		// Revocation points (single + by-issuer + all).
		out, err = QueryRevocationPoint(dcldNew, VIDFor1_2, ProductLabelFor1_2, IssuerSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)
		checkResponseContains(t, out, ProductLabelFor1_2)

		out, err = QueryRevocationPoints(dcldNew, IssuerSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		out, err = QueryAllRevocationPoints(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		// PKI all-* listings.
		out, err = QueryAllProposedX509RootCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, googleRootCertSubject)

		out, err = QueryAllRevokedX509RootCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertSubjectFor1_2)
		checkResponseContains(t, out, rootCertSubject)

		out, err = QueryAllProposedX509RootCertsToRevoke(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, testRootCertSubject)

		out, err = QueryAllX509Certs(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, testRootCertSubject)

		// ----- Validator (host-side) -----
		if state.ValidatorAddress != "" {
			out, err = QueryAllNodes(dcldNew)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.4.3-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_4_3", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_4_3, "Vendor",
			VIDFor1_4_3, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
		u7, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u8, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u9, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User7Address, state.User7Pubkey = u7.address, u7.pubkey
		state.User8Address, state.User8Pubkey = u8.address, u8.pubkey
		state.User9Address, state.User9Pubkey = u9.address, u9.pubkey
	})

	MustRun(t, "VendorInfoFor1_4_3", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcldNew, VendorArgs{VID: VIDFor1_4_3, VendorName: VendorNameFor1_4_3, CompanyLegalName: CompanyLegalNameFor1_4_3, CompanyPreferredName: CompanyPreferredNameFor1_4_3, VendorLandingPageURL: VendorLandingPageURLFor1_4_3, From: VendorAccountFor1_4_3})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateVendor(dcldNew, VendorArgs{VID: VIDFor1_2, VendorName: VendorNameFor1_2, CompanyLegalName: CompanyLegalNameFor1_2, CompanyPreferredName: CompanyPreferredNameFor1_4_3, VendorLandingPageURL: VendorLandingPageURLFor1_4_3, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsAndVersionsFor1_4_3", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1For1_4_3, PID2For1_4_3, PID3For1_4_3} {
			tx, err := AddModel(dcldNew, AddModelArgs{VID: VIDFor1_4_3, PID: pid, DeviceTypeID: DeviceTypeIDFor1_4_3, ProductName: ProductNameFor1_4_3, ProductLabel: ProductLabelFor1_4_3, PartNumber: PartNumberFor1_4_3, From: VendorAccountFor1_4_3})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: VIDFor1_4_3, PID: pid, SoftwareVersion: SoftwareVersionFor1_4_3, SoftwareVersionString: SoftwareVersionStringFor1_4_3, CDVersionNumber: CDVersionNumberFor1_4_3, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_4_3, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_4_3, From: VendorAccountFor1_4_3})
			requireTxSuccess(t, tx, err)
		}

		// Delete pid_3.
		tx, err := DeleteModel(dcldNew, VIDFor1_4_3, PID3For1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		// Update the carry-over model.
		tx, err = UpdateModel(dcldNew, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: ProductLabelFor1_4_3, PartNumber: PartNumberFor1_4_3, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(dcldNew, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_4_3, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_4_3, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceFor1_4_3", func(t *testing.T) {
		t.Helper()
		// certify pid_1 (CertCenter account survives across upgrades)
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{
			VID: VIDFor1_4_3, PID: PID1For1_4_3,
			SoftwareVersion: SoftwareVersionFor1_4_3, SoftwareVersionString: SoftwareVersionStringFor1_4_3,
			CertificationType: CertificationTypeFor1_4_3, CertificationDate: CertificationDateFor1_4_3,
			CDCertificateID: CDCertificateIDFor1_4_3, CDVersionNumber: CDVersionNumberFor1_4_3,
			From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		// provision pid_2
		tx, err = ProvisionModel(dcldNew, ProvisionModelArgs{VID: VIDFor1_4_3, PID: PID2For1_4_3, SoftwareVersion: SoftwareVersionFor1_4_3, SoftwareVersionString: SoftwareVersionStringFor1_4_3, CertificationType: CertificationTypeFor1_4_3, ProvisionalDate: ProvisionalDateFor1_4_3, CDCertificateID: CDCertificateIDFor1_4_3, CDVersionNumber: CDVersionNumberFor1_4_3, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// certify pid_2
		tx, err = CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_4_3, PID: PID2For1_4_3, SoftwareVersion: SoftwareVersionFor1_4_3, SoftwareVersionString: SoftwareVersionStringFor1_4_3, CertificationType: CertificationTypeFor1_4_3, CertificationDate: CertificationDateFor1_4_3, CDCertificateID: CDCertificateIDFor1_4_3, CDVersionNumber: CDVersionNumberFor1_4_3, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// revoke pid_2
		tx, err = RevokeModel(dcldNew, RevokeModelArgs{VID: VIDFor1_4_3, PID: PID2For1_4_3, SoftwareVersion: SoftwareVersionFor1_4_3, SoftwareVersionString: SoftwareVersionStringFor1_4_3, CertificationType: CertificationTypeFor1_4_3, RevocationDate: CertificationDateFor1_4_3, CDVersionNumber: CDVersionNumberFor1_4_3, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "PKIFor1_4_3", func(t *testing.T) {
		t.Helper()
		tx, err := ProposeAddX509RootCert(dcldNew, RootCertWithVIDPathFor1_4_3, fmt.Sprintf("%d", VIDFor1_4_3), state.Trustee1)
		requireTxSuccess(t, tx, err)

		// t2 approves.
		tx, err = ApproveAddX509RootCert(dcldNew, RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3, state.Trustee2)
		requireTxSuccess(t, tx, err)
		// t3 rejects.
		tx, err = RejectAddX509RootCert(dcldNew, RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3, state.Trustee3)
		requireTxSuccess(t, tx, err)
		// t4 approves.
		tx, err = ApproveAddX509RootCert(dcldNew, RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3, state.Trustee4)
		requireTxSuccess(t, tx, err)

		tx, err = ApproveAddX509RootCert(dcldNew, RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3, state.Trustee5)
		requireTxSuccess(t, tx, err)

		// paa_cert_no_vid: propose + 3 approvals.
		tx, err = ProposeAddX509RootCert(dcldNew, PaaCertNoVIDPathFor1_4_3, fmt.Sprintf("%d", VIDFor1_4_3), state.Trustee1)
		requireTxSuccess(t, tx, err)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveAddX509RootCert(dcldNew, PaaCertNoVIDSubjectFor1_4_3, PaaCertNoVIDSubjectKeyIDFor1_4_3, who)
			requireTxSuccess(t, tx, err)
		}

		// Propose-only root_cert (no approvals).
		tx, err = ProposeAddX509RootCert(dcldNew, RootCertPathFor1_4_3, fmt.Sprintf("%d", VIDFor1_4_3), state.Trustee1)
		requireTxSuccess(t, tx, err)

		// Add intermediate_cert_with_vid, then revoke via serial-number (new in 1.4).
		tx, err = AddX509Cert(dcldNew, IntermediateCertWithVIDPathFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		tx, err = RevokeX509Cert(dcldNew, IntermediateCertWithVIDSubjectFor1_4_3, IntermediateCertWithVIDSubjectKeyIDFor1_4_3, IntermediateCertWithVIDSerialNumberFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		// Revoke paa_cert_no_vid (propose + 3 approvals).
		tx, err = ProposeRevokeX509RootCert(dcldNew, PaaCertNoVIDSubjectFor1_4_3, PaaCertNoVIDSubjectKeyIDFor1_4_3, state.Trustee1)
		requireTxSuccess(t, tx, err)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveRevokeX509RootCert(dcldNew, PaaCertNoVIDSubjectFor1_4_3, PaaCertNoVIDSubjectKeyIDFor1_4_3, who)
			requireTxSuccess(t, tx, err)
		}

		// Propose revoke root_cert_with_vid (no approvals).
		tx, err = ProposeRevokeX509RootCert(dcldNew, RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3, state.Trustee1)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "NOCCertsAddRemove", func(t *testing.T) {
		t.Helper()
		// 1.4 introduces the NOC certificate flow. Add then immediately remove.
		tx, err := AddNocX509RootCert(dcldNew, NOCRootCert1PathFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		tx, err = AddNocX509IcaCert(dcldNew, NOCICACert1PathFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		tx, err = RemoveNocX509RootCert(dcldNew, NOCRootCert1SubjectFor1_4_3, NOCRootCert1SubjectKeyIDFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		tx, err = RemoveNocX509IcaCert(dcldNew, NOCICACert1SubjectFor1_4_3, NOCICACert1SubjectKeyIDFor1_4_3, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "RevocationPointsFor1_4_3", func(t *testing.T) {
		t.Helper()
		// add → update → delete → add.
		addPAA := func(label, dataURL string) {
			tx, err := AddRevocationPoint(dcldNew, AddRevocationPointArgs{VID: VIDFor1_4_3, RevocationType: "1", IsPAA: true, Certificate: RootCertWithVIDPathFor1_4_3, Label: label, DataURL: dataURL, IssuerSubjectKeyID: IssuerSubjectKeyID, From: VendorAccountFor1_4_3})
			requireTxSuccess(t, tx, err)
		}

		addPAA(ProductLabelFor1_4_3, TestDataURLFor1_4_3)

		tx, err := UpdateRevocationPoint(dcldNew, UpdateRevocationPointArgs{VID: VIDFor1_4_3, Certificate: RootCertWithVIDPathFor1_4_3, Label: ProductLabelFor1_4_3, DataURL: TestDataURLFor1_4_3 + "/new", IssuerSubjectKeyID: IssuerSubjectKeyID, From: VendorAccountFor1_4_3})
		requireTxSuccess(t, tx, err)

		tx, err = DeleteRevocationPoint(dcldNew, VIDFor1_4_3, ProductLabelFor1_4_3, IssuerSubjectKeyID, VendorAccountFor1_4_3)
		requireTxSuccess(t, tx, err)

		addPAA(ProductLabelFor1_4_3, TestDataURLFor1_4_3)

		// CRL signer revocation point delegated by PAI (new in 1.4).
		tx, err = AddRevocationPoint(dcldNew, AddRevocationPointArgs{VID: VIDFor1_4_3, RevocationType: "1", IsPAA: false, Certificate: CRLSignerDelegatedByPAI1, CertificateDelegator: DelegatorCertWithVID65521Path, Label: ProductLabelFor1_4_3, DataURL: TestDataURLFor1_4_3, IssuerSubjectKeyID: DelegatorCertWithVIDSubjectKeyID, From: VendorAccountFor1_4_3})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateRevocationPoint(dcldNew, UpdateRevocationPointArgs{VID: VIDFor1_4_3, Certificate: CRLSignerDelegatedByPAI1, CertificateDelegator: DelegatorCertWithVID65521Path, Label: ProductLabelFor1_4_3, DataURL: TestDataURLFor1_4_3 + "/new", IssuerSubjectKeyID: DelegatorCertWithVIDSubjectKeyID, From: VendorAccountFor1_4_3})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "AccountFlowsFor1_4_3", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User7Address, state.User7Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User8Address, state.User8Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User9Address, state.User9Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User7Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User8Address, false)
	})

	// Validator disable/enable — depends on the Docker validator-demo container
	// that the validator-node setup would create. Stubbed alongside the other
	// validator work.
	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		// disable-node is approved with 3 trustees (trustee_4 active).
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW data.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNewVendorAndModels", func(t *testing.T) {
		t.Helper()
		out, err := QueryVendor(dcldNew, VIDFor1_4_3)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)
		checkResponseContains(t, out, CompanyLegalNameFor1_4_3)

		out, err = QueryGetModel(dcldNew, VIDFor1_4_3, PID1For1_4_3)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)
		checkResponseContains(t, out, ProductLabelFor1_4_3)

		// Updated 0.12 pid_2 now has 1.4 productLabel/partNumber.
		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_3)
		checkResponseContains(t, out, PartNumberFor1_4_3)
	})

	MustRun(t, "VerifyNOCCertsEmptyAfterRemove", func(t *testing.T) {
		t.Helper()
		// After add+remove on the NOC certs, queries should return Not Found
		// and must NOT contain the removed subject key IDs.
		out, err := QueryNocX509RootCerts(dcldNew, VIDFor1_4_3)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found for cleaned NOC root, got: %s", string(out))
		require.False(t, strings.Contains(string(out), NOCRootCert1SubjectKeyIDFor1_4_3),
			"NOC root cert SKID lingered after remove: %s", string(out))

		out, err = QueryNocX509Certs(dcldNew, NOCRootCert1SubjectKeyIDFor1_4_3, VIDFor1_4_3)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found, got: %s", string(out))

		out, err = QueryNocX509Certs(dcldNew, NOCICACert1SubjectKeyIDFor1_4_3, VIDFor1_4_3)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found, got: %s", string(out))
	})
}
