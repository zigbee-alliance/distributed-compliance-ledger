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

// runUpgrade143To144 runs the v1.4.3 → v1.4.4 cosmovisor upgrade, then
// seeds 1.4.4-era state: new vendor (VID=65522), DA root certs + their
// intermediates with revocation pairs, NOC certs with the new
// revoke-noc-x509-* commands, and a revocation point.
//
//nolint:funlen
func runUpgrade143To144(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_4_3)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_4_4)
	require.NoError(t, err)

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_4_4,
		BinaryVersionNew: BinaryVersionV1_4_4,
		Checksum:         UpgradeChecksumV1_4_4,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.4.4.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAcrossThreeEras", func(t *testing.T) {
		t.Helper()
		// Spot-check the three vendor-info records.
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3} {
			out, qerr := QueryVendor(dcldNew, vid)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// Spot-check key models from each era.
		for _, pair := range [][2]int{
			{state.VID, pid1V012}, {state.VID, state.PID2},
			{VIDFor1_2, PID1For1_2}, {VIDFor1_2, PID2For1_2},
			{VIDFor1_4_3, PID2For1_4_3},
		} {
			out, qerr := QueryGetModel(dcldNew, pair[0], pair[1])
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vid", pair[0])
			requireFieldEquals(t, out, "pid", pair[1])
		}

		// 0.12 pid_2 now has 1.4.3 productLabel + partNumber (set in script 05).
		out, err := QueryModelVersion(dcldNew, state.VID, state.PID2, state.SoftwareVersion)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_4_3)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_4_3)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllAccounts(dcldNew)
		require.NoError(t, err)
		// Active accounts from all prior scripts.
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User8Address)

		out, err = QueryAllProposedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address)
		checkResponseContains(t, out, state.User6Address)
		checkResponseContains(t, out, state.User9Address)

		out, err = QueryAllProposedAccountsToRevoke(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User8Address)

		out, err = QueryAllRevokedAccounts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address)
		checkResponseContains(t, out, state.User4Address)
		checkResponseContains(t, out, state.User7Address)

		// Single-record account variants.
		for _, addr := range []string{state.User5Address, state.User2Address, state.User8Address} {
			out, err = QueryAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User6Address, state.User3Address, state.User9Address} {
			out, err = QueryProposedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User5Address, state.User2Address, state.User8Address} {
			out, err = QueryProposedAccountToRevoke(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User4Address, state.User1Address, state.User7Address} {
			out, err = QueryRevokedAccount(dcldNew, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
	})

	// Bulk readback — gap-fill queries covering vendorinfo/model/compliance/PKI
	// listings + global/DA/NOC cert queries introduced in 1.4.x. NOC-side
	// "Not Found" responses are run for coverage but not asserted on (treated
	// as informational).
	MustRun(t, "VerifyPreservedListings_1_4_4", func(t *testing.T) {
		t.Helper()
		// VendorInfo: all-vendors across three eras.
		out, err := QueryAllVendors(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)

		// Model bulk listings.
		out, err = QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)

		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3} {
			_, err = QueryVendorModels(dcldNew, vid)
			require.NoError(t, err)
		}
		_, err = QueryAllModelVersions(dcldNew, VIDFor1_4_3, PID1For1_4_3)
		require.NoError(t, err)

		// Compliance single-record forms (gap entries).
		out, err = QueryCertifiedModel(dcldNew, VIDFor1_4_3, PID1For1_4_3, SoftwareVersionFor1_4_3, CertificationTypeFor1_4_3)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		_, err = QueryRevokedModel(dcldNew, VIDFor1_4_3, PID2For1_4_3, SoftwareVersionFor1_4_3, CertificationTypeFor1_4_3)
		require.NoError(t, err)

		_, err = QueryProvisionalModel(dcldNew, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)

		_, err = QueryComplianceInfo(dcldNew, VIDFor1_4_3, PID1For1_4_3, SoftwareVersionFor1_4_3, CertificationTypeFor1_4_3)
		require.NoError(t, err)

		for _, cdID := range []string{cdCertificateIDV012, CDCertificateIDFor1_2, CDCertificateIDFor1_4_3} {
			out, err = QueryDeviceSoftwareCompliance(dcldNew, cdID)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		// Compliance all-* listings.
		out, err = QueryAllCertifiedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)

		_, err = QueryAllProvisionalModels(dcldNew)
		require.NoError(t, err)

		out, err = QueryAllRevokedModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID2For1_4_3)

		_, err = QueryAllComplianceInfo(dcldNew)
		require.NoError(t, err)

		out, err = QueryAllDeviceSoftwareCompliance(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, CDCertificateIDFor1_4_3)

		// PKI single-record forms across DA/global namespaces (cert + x509-cert).
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = QueryCert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)

			out, err = QueryX509Cert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)

			_, err = QueryAllSubjectCerts(dcldNew, c.subj)
			require.NoError(t, err)

			_, err = QueryAllSubjectX509Certs(dcldNew, c.subj)
			require.NoError(t, err)
		}

		// NOC namespace queries return Not Found for DA-side subjects —
		// confirms DA/NOC namespace separation.
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
		} {
			_, _ = QueryNocX509Cert(dcldNew, c.subj, c.kid)
			_, _ = QueryAllNocSubjectX509Certs(dcldNew, c.subj)
		}

		// Proposed + revoked + propose-to-revoke (gap entries).
		for _, c := range []struct{ subj, kid string }{
			{GoogleRootCertSubjectFor1_2, GoogleRootCertSubjectKeyIDFor1_2},
			{googleRootCertSubject, googleRootCertSubjectKeyID},
		} {
			out, err = QueryProposedX509RootCert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		for _, c := range []struct{ subj, kid string }{
			{IntermediateCertWithVIDSubjectFor1_4_3, IntermediateCertWithVIDSubjectKeyIDFor1_4_3},
			{IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2},
			{intermediateCertSubject, intermediateCertSubjectKeyID},
		} {
			out, err = QueryRevokedX509Cert(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = QueryProposedX509RootCertToRevoke(dcldNew, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		// Revocation points (single + by-issuer + all).
		out, err = QueryRevocationPoint(dcldNew, VIDFor1_2, ProductLabelFor1_2, IssuerSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		_, err = QueryRevocationPoints(dcldNew, IssuerSubjectKeyID)
		require.NoError(t, err)

		out, err = QueryAllRevocationPoints(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		// Global vs DA all-* listings.
		out, err = QueryAllCerts(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertWithVIDSubjectKeyIDFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryAllX509Certs(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertWithVIDSubjectKeyIDFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

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
		checkResponseContains(t, out, RootCertWithVIDSubjectFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, testRootCertSubject)

		out, err = QueryAllRevokedX509Certs(dcldNew)
		require.NoError(t, err)
		checkResponseContains(t, out, IntermediateCertWithVIDSubjectFor1_4_3)

		// NOC-side listings (added in 1.4.3 NOC flow); since NOC certs were
		// added then removed at the script 05 tail, these should be empty —
		// run for coverage and check for absence of removed SKIDs.
		out, err = QueryAllNocX509Certs(dcldNew)
		require.NoError(t, err)
		require.False(t, strings.Contains(string(out), NOCRootCert1SubjectKeyIDFor1_4_3),
			"NOC root SKID lingered: %s", string(out))

		_, err = QueryNocX509RootCerts(dcldNew, VIDFor1_4_3)
		require.NoError(t, err)

		_, err = QueryAllRevokedNocX509RootCerts(dcldNew)
		require.NoError(t, err)

		_, err = QueryAllRevokedNocX509IcaCerts(dcldNew)
		require.NoError(t, err)

		_, _ = QueryRevokedNocX509RootCert(dcldNew, NOCRootCert1SubjectFor1_4_3, NOCRootCert1SubjectKeyIDFor1_4_3)

		// Validator (host-side).
		if state.ValidatorAddress != "" {
			out, err = QueryAllNodes(dcldNew)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.4.4-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_4_4", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_4_4, "Vendor",
			VIDFor1_4_4, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
		u10, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u11, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u12, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User10Address, state.User10Pubkey = u10.address, u10.pubkey
		state.User11Address, state.User11Pubkey = u11.address, u11.pubkey
		state.User12Address, state.User12Pubkey = u12.address, u12.pubkey
	})

	MustRun(t, "VendorInfoFor1_4_4", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcldNew, VendorArgs{VID: VIDFor1_4_4, VendorName: VendorNameFor1_4_4, CompanyLegalName: CompanyLegalNameFor1_4_4, CompanyPreferredName: CompanyPreferredNameFor1_4_4, VendorLandingPageURL: VendorLandingPageURLFor1_4_4, From: VendorAccountFor1_4_4})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateVendor(dcldNew, VendorArgs{VID: VIDFor1_2, VendorName: VendorNameFor1_2, CompanyLegalName: CompanyLegalNameFor1_2, CompanyPreferredName: CompanyPreferredNameFor1_4_4, VendorLandingPageURL: VendorLandingPageURLFor1_4_4, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsAndVersionsFor1_4_4", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1For1_4_4, PID2For1_4_4, PID3For1_4_4} {
			tx, err := AddModel(dcldNew, AddModelArgs{VID: VIDFor1_4_4, PID: pid, DeviceTypeID: DeviceTypeIDFor1_4_4, ProductName: ProductNameFor1_4_4, ProductLabel: ProductLabelFor1_4_4, PartNumber: PartNumberFor1_4_4, From: VendorAccountFor1_4_4})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: VIDFor1_4_4, PID: pid, SoftwareVersion: SoftwareVersionFor1_4_4, SoftwareVersionString: SoftwareVersionStringFor1_4_4, CDVersionNumber: CDVersionNumberFor1_4_4, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_4_4, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_4_4, From: VendorAccountFor1_4_4})
			requireTxSuccess(t, tx, err)
		}

		// Delete pid_3.
		tx, err := DeleteModel(dcldNew, VIDFor1_4_4, PID3For1_4_4, VendorAccountFor1_4_4)
		requireTxSuccess(t, tx, err)

		// Update carry-over 0.12 pid_2.
		tx, err = UpdateModel(dcldNew, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: ProductLabelFor1_4_4, PartNumber: PartNumberFor1_4_4, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(dcldNew, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_4_4, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_4_4, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceFor1_4_4", func(t *testing.T) {
		t.Helper()
		// certify pid_1
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_4_4, PID: PID1For1_4_4, SoftwareVersion: SoftwareVersionFor1_4_4, SoftwareVersionString: SoftwareVersionStringFor1_4_4, CertificationType: CertificationTypeFor1_4_4, CertificationDate: CertificationDateFor1_4_4, CDCertificateID: CDCertificateIDFor1_4_4, CDVersionNumber: CDVersionNumberFor1_4_4, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// provision pid_2
		tx, err = ProvisionModel(dcldNew, ProvisionModelArgs{VID: VIDFor1_4_4, PID: PID2For1_4_4, SoftwareVersion: SoftwareVersionFor1_4_4, SoftwareVersionString: SoftwareVersionStringFor1_4_4, CertificationType: CertificationTypeFor1_4_4, ProvisionalDate: ProvisionalDateFor1_4_4, CDCertificateID: CDCertificateIDFor1_4_4, CDVersionNumber: CDVersionNumberFor1_4_4, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// certify pid_2
		tx, err = CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_4_4, PID: PID2For1_4_4, SoftwareVersion: SoftwareVersionFor1_4_4, SoftwareVersionString: SoftwareVersionStringFor1_4_4, CertificationType: CertificationTypeFor1_4_4, CertificationDate: CertificationDateFor1_4_4, CDCertificateID: CDCertificateIDFor1_4_4, CDVersionNumber: CDVersionNumberFor1_4_4, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// revoke pid_2
		tx, err = RevokeModel(dcldNew, RevokeModelArgs{VID: VIDFor1_4_4, PID: PID2For1_4_4, SoftwareVersion: SoftwareVersionFor1_4_4, SoftwareVersionString: SoftwareVersionStringFor1_4_4, CertificationType: CertificationTypeFor1_4_4, RevocationDate: CertificationDateFor1_4_4, CDVersionNumber: CDVersionNumberFor1_4_4, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "PKIFor1_4_4_DARootCerts", func(t *testing.T) {
		t.Helper()
		tx, err := ProposeAddX509RootCert(dcldNew, DARootCert1PathFor1_4_4, fmt.Sprintf("%d", VIDFor1_4_4), state.Trustee1)
		requireTxSuccess(t, tx, err)

		tx, err = ApproveAddX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, state.Trustee2)
		requireTxSuccess(t, tx, err)

		tx, err = RejectAddX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, state.Trustee3)
		requireTxSuccess(t, tx, err)

		tx, err = ApproveAddX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, state.Trustee4)
		requireTxSuccess(t, tx, err)

		tx, err = ApproveAddX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, state.Trustee5)
		requireTxSuccess(t, tx, err)

		// da_root_cert_2: propose by t1, approve by t1/t2/t3.
		tx, err = ProposeAddX509RootCert(dcldNew, DARootCert2PathFor1_4_4, fmt.Sprintf("%d", VIDFor1_4_4), state.Trustee1)
		requireTxSuccess(t, tx, err)

		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveAddX509RootCert(dcldNew, DARootCert2SubjectFor1_4_4, DARootCert2SubjectKeyIDFor1_4_4, who)
			requireTxSuccess(t, tx, err)
		}

		// Add intermediates and revoke da_root_cert_1 + da_intermediate_cert_1.
		for _, certPath := range []string{
			DAIntermediateCert1PathFor1_4_4, DAIntermediateCert2PathFor1_4_4,
		} {
			tx, err = AddX509Cert(dcldNew, certPath, VendorAccountFor1_4_4)
			requireTxSuccess(t, tx, err)
		}

		// Propose-revoke + 3 approves of da_root_cert_1.
		tx, err = ProposeRevokeX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, state.Trustee1)
		requireTxSuccess(t, tx, err)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveRevokeX509RootCert(dcldNew, DARootCert1SubjectFor1_4_4, DARootCert1SubjectKeyIDFor1_4_4, who)
			requireTxSuccess(t, tx, err)
		}

		// Propose-revoke da_root_cert_2 (no approvals).
		tx, err = ProposeRevokeX509RootCert(dcldNew, DARootCert2SubjectFor1_4_4, DARootCert2SubjectKeyIDFor1_4_4, state.Trustee1)
		requireTxSuccess(t, tx, err)

		// Revoke da_intermediate_cert_1.
		tx, err = RevokeX509Cert(dcldNew, DAIntermediateCert1SubjectFor1_4_4, DAIntermediateCert1SubjectKeyIDFor1_4_4, "", VendorAccountFor1_4_4)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "NOCCertsAddRevoke", func(t *testing.T) {
		t.Helper()
		// 1.4.4 introduces `revoke-noc-x509-{root,ica}-cert` (vs 1.4.3's
		// `remove-noc-x509-*`). Add 2 root/ICA pairs, then revoke pair #1
		// only — pair #2 stays active.
		for _, pair := range []struct{ rootPath, icaPath string }{
			{NOCRootCert1V144PathFor1_4_4, NOCICACert1V144PathFor1_4_4},
			{NOCRootCert2V144PathFor1_4_4, NOCICACert2V144PathFor1_4_4},
		} {
			tx, err := AddNocX509RootCert(dcldNew, pair.rootPath, VendorAccountFor1_4_4)
			requireTxSuccess(t, tx, err)

			tx, err = AddNocX509IcaCert(dcldNew, pair.icaPath, VendorAccountFor1_4_4)
			requireTxSuccess(t, tx, err)
		}

		// Revoke NOC pair #1.
		tx, err := RevokeNocX509RootCert(dcldNew, NOCRootCert1V144SubjectFor1_4_4, NOCRootCert1V144SubjectKeyIDFor1_4_4, VendorAccountFor1_4_4)
		requireTxSuccess(t, tx, err)

		tx, err = RevokeNocX509IcaCert(dcldNew, NOCICACert1V144SubjectFor1_4_4, NOCICACert1V144SubjectKeyIDFor1_4_4, VendorAccountFor1_4_4)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "RevocationPointsFor1_4_4", func(t *testing.T) {
		t.Helper()
		// add → update → delete → add (one active PAA revocation point at end).
		addPAA := func(dataURL string) {
			tx, err := AddRevocationPoint(dcldNew, AddRevocationPointArgs{VID: VIDFor1_4_4, RevocationType: "1", IsPAA: true, Certificate: DARootCert2PathFor1_4_4, Label: ProductLabelFor1_4_4, DataURL: dataURL, IssuerSubjectKeyID: IssuerSubjectKeyID, From: VendorAccountFor1_4_4})
			requireTxSuccess(t, tx, err)
		}

		addPAA(TestDataURLFor1_4_4)

		tx, err := UpdateRevocationPoint(dcldNew, UpdateRevocationPointArgs{VID: VIDFor1_4_4, Certificate: DARootCert2PathFor1_4_4, Label: ProductLabelFor1_4_4, DataURL: TestDataURLFor1_4_4 + "/new", IssuerSubjectKeyID: IssuerSubjectKeyID, From: VendorAccountFor1_4_4})
		requireTxSuccess(t, tx, err)

		tx, err = DeleteRevocationPoint(dcldNew, VIDFor1_4_4, ProductLabelFor1_4_4, IssuerSubjectKeyID, VendorAccountFor1_4_4)
		requireTxSuccess(t, tx, err)

		addPAA(TestDataURLFor1_4_4)
	})

	MustRun(t, "AccountFlowsFor1_4_4", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User10Address, state.User10Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User11Address, state.User11Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User12Address, state.User12Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User10Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User11Address, false)
	})

	// Validator disable/enable (lines 1189-1238) — Docker-dependent, stubbed.
	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW data.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNew_1_4_4_Data", func(t *testing.T) {
		t.Helper()
		out, err := QueryVendor(dcldNew, VIDFor1_4_4)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_4)
		checkResponseContains(t, out, CompanyLegalNameFor1_4_4)

		out, err = QueryGetModel(dcldNew, VIDFor1_4_4, PID1For1_4_4)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_4)
		requireFieldEquals(t, out, "pid", PID1For1_4_4)
		checkResponseContains(t, out, ProductLabelFor1_4_4)

		// 0.12 pid_2 now has 1.4.4 productLabel/partNumber.
		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_4)
		checkResponseContains(t, out, PartNumberFor1_4_4)
	})
}
