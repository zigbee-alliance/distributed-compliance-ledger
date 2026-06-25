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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

// runUpgrade160ToMaster builds the master image, hands the binary to each
// localnet node via cosmovisor add-upgrade, then proposes/approves an
// upgrade plan whose name is the master git short hash. After upgrade
// applies, seeds master-era state.
//
//nolint:funlen
func runUpgrade160ToMaster(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_6_0)
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// Build master image, extract binary, compute plan name, push to nodes.
	// ------------------------------------------------------------------
	var planName string
	MustRun(t, "BuildAndDistributeMasterBinary", func(t *testing.T) {
		t.Helper()
		DockerCleanup(MasterUpgradeContainerName)

		require.NoError(t, BuildMasterImage(), "docker build dcld-build-master")
		require.NoError(t, CreateMasterContainer(), "create master container")
		t.Cleanup(func() { DockerCleanup(MasterUpgradeContainerName) })

		require.NoError(t, ExtractMasterBinary(DcldMasterBinaryPath), "extract dcld_master")

		pn, perr := GetMasterPlanName()
		require.NoError(t, perr, "git rev-parse short HEAD inside master image")
		require.NotEmpty(t, pn, "master plan name should be non-empty git hash")
		planName = pn
		t.Logf("master upgrade plan name = %s", planName)

		require.NoError(t,
			PrepareCosmovisorUpgradeOnLocalnetNodes(planName, DcldMasterBinaryPath),
			"prepare cosmovisor add-upgrade on localnet nodes",
		)
	})

	// ------------------------------------------------------------------
	// Upgrade flow.
	// ------------------------------------------------------------------
	MustRun(t, "ProposeApproveMasterUpgrade", func(t *testing.T) {
		t.Helper()
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		// Master upgrade plan submission passes empty upgrade-info — the binary
		// was already seeded into cosmovisor manually above.
		tx, err := ProposeUpgrade(dcldOld, planName, planHeight, "", state.Trustee1)
		requireTxSuccess(t, tx, err)

		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveUpgrade(dcldOld, planName, who)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "approve %s: %s", who, tx.RawLog)
		}

		cliputils.WaitForHeight(t, planHeight+1, 300)

		// Use the new master binary for post-upgrade queries.
		out, _ := QueryUpgradePlan(DcldMasterBinaryPath)
		require.True(t, strings.Contains(string(out), "no upgrade scheduled"),
			"expected 'no upgrade scheduled', got: %s", string(out))

		out, err = QueryAppliedPlan(DcldMasterBinaryPath, planName)
		require.NoError(t, err, "query upgrade applied %s", planName)
		t.Logf("applied: %s", string(out))
	})

	// ------------------------------------------------------------------
	// Verify carry-over across all prior eras.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAcrossAllEras", func(t *testing.T) {
		t.Helper()
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4, state.VIDFor1_5_1} {
			out, qerr := QueryVendor(DcldMasterBinaryPath, vid)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// 1.5.2 pid_2 has 1.6.0 productLabel/partNumber (set in script 09).
		out, err := QueryGetModel(DcldMasterBinaryPath, VIDFor1_5_2, PID2For1_5_2)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_6_0)
		checkResponseContains(t, out, PartNumberFor1_6_0)
	})

	// Bulk readback — gap-fill queries for auth (single+all), compliance
	// (single+all), model bulk, pki (global/DA/NOC + revocation),
	// vendorinfo all-vendors, and validator all-nodes.
	MustRun(t, "VerifyPreservedListings_Master", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllVendors(DcldMasterBinaryPath)
		require.NoError(t, err)
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4, state.VIDFor1_5_1} {
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// ----- Auth -----
		out, err = QueryAllAccounts(DcldMasterBinaryPath)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address,
			state.User11Address, state.User14Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccounts(DcldMasterBinaryPath)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address,
			state.User12Address, state.User15Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccountsToRevoke(DcldMasterBinaryPath)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address,
			state.User11Address, state.User14Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllRevokedAccounts(DcldMasterBinaryPath)
		require.NoError(t, err)
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address,
			state.User10Address, state.User13Address,
		} {
			checkResponseContains(t, out, addr)
		}

		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address,
			state.User11Address, state.User14Address,
		} {
			out, err = QueryAccount(DcldMasterBinaryPath, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address,
			state.User12Address, state.User15Address,
		} {
			out, err = QueryProposedAccount(DcldMasterBinaryPath, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address,
			state.User11Address, state.User14Address,
		} {
			out, err = QueryProposedAccountToRevoke(DcldMasterBinaryPath, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address,
			state.User10Address, state.User13Address,
		} {
			out, err = QueryRevokedAccount(DcldMasterBinaryPath, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		// ----- Model bulk listings -----
		out, err = QueryAllModels(DcldMasterBinaryPath)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_4)

		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4, state.VIDFor1_5_1, VIDFor1_5_2} {
			_, err = QueryVendorModels(DcldMasterBinaryPath, vid)
			require.NoError(t, err)
		}

		_, err = QueryAllModelVersions(DcldMasterBinaryPath, VIDFor1_4_4, PID1For1_4_4)
		require.NoError(t, err)

		_, err = QueryModelVersion(DcldMasterBinaryPath, VIDFor1_4_4, PID1For1_4_4, SoftwareVersionFor1_4_4)
		require.NoError(t, err)

		// ----- Compliance single-record + all-* listings -----
		out, err = QueryCertifiedModel(DcldMasterBinaryPath, VIDFor1_4_4, PID1For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		_, err = QueryRevokedModel(DcldMasterBinaryPath, VIDFor1_4_4, PID2For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)

		_, err = QueryProvisionalModel(DcldMasterBinaryPath, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)

		_, err = QueryComplianceInfo(DcldMasterBinaryPath, VIDFor1_4_4, PID1For1_4_4, SoftwareVersionFor1_4_4, CertificationTypeFor1_4_4)
		require.NoError(t, err)

		for _, cdID := range []string{
			cdCertificateIDV012, CDCertificateIDFor1_2, CDCertificateIDFor1_4_3, CDCertificateIDFor1_4_4,
		} {
			out, err = QueryDeviceSoftwareCompliance(DcldMasterBinaryPath, cdID)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		_, err = QueryAllCertifiedModels(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllProvisionalModels(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllRevokedModels(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllComplianceInfo(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllDeviceSoftwareCompliance(DcldMasterBinaryPath)
		require.NoError(t, err)

		// ----- PKI single-record forms (global/DA/NOC) -----
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = QueryCert(DcldMasterBinaryPath, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			out, err = QueryX509Cert(DcldMasterBinaryPath, c.subj, c.kid)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			// DA-store subjects must not resolve in the NOC store: the singular
			// noc-x509-cert query exits non-zero with a "Not Found" body, so the
			// error is expected here. Mirrors the script's DA-vs-NOC separation.
			nocOut, _ := QueryNocX509Cert(DcldMasterBinaryPath, c.subj, c.kid)
			checkResponseContains(t, nocOut, "Not Found")
		}

		_, _ = QueryRevokedX509Cert(DcldMasterBinaryPath, IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2)
		_, _ = QueryRevokedNocX509RootCert(DcldMasterBinaryPath, NOCRootCert1SubjectFor1_4_3, NOCRootCert1SubjectKeyIDFor1_4_3)

		out, err = QueryRevocationPoint(DcldMasterBinaryPath, VIDFor1_2, ProductLabelFor1_2, IssuerSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		_, err = QueryRevocationPoints(DcldMasterBinaryPath, IssuerSubjectKeyID)
		require.NoError(t, err)

		_, err = QueryAllCerts(DcldMasterBinaryPath)
		require.NoError(t, err)

		// DA store (all-x509-certs) must exclude NOC-store SKIDs — mirrors the
		// script's "Get certificates (DA)" separation check.
		out, err = QueryAllX509Certs(DcldMasterBinaryPath)
		require.NoError(t, err)
		require.NotContains(t, string(out), NOCRootCert2V144SubjectKeyIDFor1_4_4,
			"all-x509-certs (DA store) leaked NOC root SKID: %s", string(out))
		require.NotContains(t, string(out), NOCICACert2V144SubjectKeyIDFor1_4_4,
			"all-x509-certs (DA store) leaked NOC ICA SKID: %s", string(out))

		_, err = QueryAllRevokedX509Certs(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllRevokedX509RootCerts(DcldMasterBinaryPath)
		require.NoError(t, err)

		// NOC store (all-noc-x509-certs) must exclude DA-store SKIDs — mirrors
		// the script's "Get certificates (NOC)" separation check.
		out, err = QueryAllNocX509Certs(DcldMasterBinaryPath)
		require.NoError(t, err)
		require.NotContains(t, string(out), DARootCert1SubjectKeyIDFor1_4_4,
			"all-noc-x509-certs (NOC store) leaked DA root SKID: %s", string(out))
		_, err = QueryAllRevokedNocX509RootCerts(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllRevokedNocX509IcaCerts(DcldMasterBinaryPath)
		require.NoError(t, err)
		_, err = QueryAllRevocationPoints(DcldMasterBinaryPath)
		require.NoError(t, err)

		for _, subj := range []string{
			RootCertWithVIDSubjectFor1_4_3, TestRootCertSubjectFor1_2, testRootCertSubject,
		} {
			_, _ = QueryAllSubjectCerts(DcldMasterBinaryPath, subj)
			_, _ = QueryAllSubjectX509Certs(DcldMasterBinaryPath, subj)
			_, _ = QueryAllNocSubjectX509Certs(DcldMasterBinaryPath, subj)
		}

		// ----- Validator (host-side) -----
		if state.ValidatorAddress != "" {
			out, err = QueryAllNodes(DcldMasterBinaryPath)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Seed master-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_Master", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, DcldMasterBinaryPath, VendorAccountForMaster, "Vendor",
			VIDForMaster, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddMasterUserKeys", func(t *testing.T) {
		t.Helper()
		u13, err := newUserKey(DcldMasterBinaryPath)
		require.NoError(t, err)
		u14, err := newUserKey(DcldMasterBinaryPath)
		require.NoError(t, err)
		u15, err := newUserKey(DcldMasterBinaryPath)
		require.NoError(t, err)
		state.User13Address, state.User13Pubkey = u13.address, u13.pubkey
		state.User14Address, state.User14Pubkey = u14.address, u14.pubkey
		state.User15Address, state.User15Pubkey = u15.address, u15.pubkey
	})

	MustRun(t, "VendorInfoFor_Master", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(DcldMasterBinaryPath, VendorArgs{VID: VIDForMaster, VendorName: VendorNameForMaster, CompanyLegalName: CompanyLegalNameForMaster, CompanyPreferredName: CompanyPreferredNameForMaster, VendorLandingPageURL: VendorLandingPageURLForMaster, From: VendorAccountForMaster})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateVendor(DcldMasterBinaryPath, VendorArgs{VID: VIDFor1_2, VendorName: VendorNameFor1_2, CompanyLegalName: CompanyLegalNameFor1_2, CompanyPreferredName: CompanyPreferredNameForMaster, VendorLandingPageURL: VendorLandingPageURLForMaster, From: state.VendorAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsAndVersionsFor_Master", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1ForMaster, PID2ForMaster, PID3ForMaster} {
			tx, err := AddModel(DcldMasterBinaryPath, AddModelArgs{VID: VIDForMaster, PID: pid, DeviceTypeID: DeviceTypeIDForMaster, ProductName: ProductNameForMaster, ProductLabel: ProductLabelForMaster, PartNumber: PartNumberForMaster, CommissioningCustomFlow: intPtr(CommissioningCustomFlow), From: VendorAccountForMaster})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(DcldMasterBinaryPath, AddModelVersionArgs{VID: VIDForMaster, PID: pid, SoftwareVersion: SoftwareVersionForMaster, SoftwareVersionString: SoftwareVersionStringForMaster, CDVersionNumber: CDVersionNumberForMaster, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionForMaster, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionForMaster, From: VendorAccountForMaster})
			requireTxSuccess(t, tx, err)
		}

		// Delete pid_3.
		tx, err := DeleteModel(DcldMasterBinaryPath, VIDForMaster, PID3ForMaster, VendorAccountForMaster)
		requireTxSuccess(t, tx, err)

		// Update carry-over 0.12 pid_2 with master-era values.
		tx, err = UpdateModel(DcldMasterBinaryPath, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: ProductLabelForMaster, PartNumber: PartNumberForMaster, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(DcldMasterBinaryPath, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionForMaster, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionForMaster, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceFor_Master", func(t *testing.T) {
		t.Helper()
		// certify pid_1.
		tx, err := CertifyModel(DcldMasterBinaryPath, CertifyModelArgs{
			VID: VIDForMaster, PID: PID1ForMaster, SoftwareVersion: SoftwareVersionForMaster,
			SoftwareVersionString: SoftwareVersionStringForMaster, CertificationType: CertificationTypeForMaster,
			CertificationDate: CertificationDateForMaster, CDCertificateID: CDCertificateIDForMaster,
			CDVersionNumber: CDVersionNumberForMaster, SpecificationVersion: SpecificationVersionForMaster, From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		// provision pid_2, certify pid_2, revoke pid_2. revoke-model does not
		// accept --cdCertificateId, so it is set only on provision/certify.
		tx, err = ProvisionModel(DcldMasterBinaryPath, ProvisionModelArgs{
			VID: VIDForMaster, PID: PID2ForMaster, SoftwareVersion: SoftwareVersionForMaster,
			SoftwareVersionString: SoftwareVersionStringForMaster, CertificationType: CertificationTypeForMaster,
			ProvisionalDate: ProvisionalDateForMaster, CDCertificateID: CDCertificateIDForMaster,
			CDVersionNumber: CDVersionNumberForMaster, SpecificationVersion: SpecificationVersionForMaster, From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		tx, err = CertifyModel(DcldMasterBinaryPath, CertifyModelArgs{
			VID: VIDForMaster, PID: PID2ForMaster, SoftwareVersion: SoftwareVersionForMaster,
			SoftwareVersionString: SoftwareVersionStringForMaster, CertificationType: CertificationTypeForMaster,
			CertificationDate: CertificationDateForMaster, CDCertificateID: CDCertificateIDForMaster,
			CDVersionNumber: CDVersionNumberForMaster, SpecificationVersion: SpecificationVersionForMaster, From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		tx, err = RevokeModel(DcldMasterBinaryPath, RevokeModelArgs{
			VID: VIDForMaster, PID: PID2ForMaster, SoftwareVersion: SoftwareVersionForMaster,
			SoftwareVersionString: SoftwareVersionStringForMaster, CertificationType: CertificationTypeForMaster,
			RevocationDate: CertificationDateForMaster, CDVersionNumber: CDVersionNumberForMaster,
			From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)
	})

	// Query-back the master-era vendorinfo + compliance just seeded above. This
	// is the post-seed verification block the bash script runs (originally the
	// "VENDORINFO" and "COMPLIANCE" sections of script 10); without it a broken
	// migration of these stores would pass undetected.
	MustRun(t, "VerifyMasterComplianceAndVendorInfo", func(t *testing.T) {
		t.Helper()
		VerifyMasterComplianceAndVendorInfo(t, state)
	})

	MustRun(t, "AccountFlowsFor_Master", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, DcldMasterBinaryPath, state.Trustee1, approvers,
			state.User13Address, state.User13Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, DcldMasterBinaryPath, state.Trustee1, approvers,
			state.User14Address, state.User14Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, DcldMasterBinaryPath, state.Trustee1, nil,
			state.User15Address, state.User15Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, DcldMasterBinaryPath, state.Trustee1, approvers, state.User13Address, true)
		revokeUserAccount(t, DcldMasterBinaryPath, state.Trustee1, nil, state.User14Address, false)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, DcldMasterBinaryPath,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade data is in place. Phase 4's add-new-node script
	// inherits this state.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNew_Master_Data", func(t *testing.T) {
		t.Helper()
		out, err := QueryVendor(DcldMasterBinaryPath, VIDForMaster)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDForMaster)
		checkResponseContains(t, out, CompanyLegalNameForMaster)

		out, err = QueryGetModel(DcldMasterBinaryPath, VIDForMaster, PID1ForMaster)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDForMaster)
		requireFieldEquals(t, out, "pid", PID1ForMaster)
		checkResponseContains(t, out, ProductLabelForMaster)

		// Master plan name is recorded in state for script 11 to verify
		// against (the new observer should eventually report this version).
		state.MasterPlanName = planName
	})
}

// VerifyMasterComplianceAndVendorInfo queries back the master-era vendorinfo
// and compliance state seeded by VendorInfoFor_Master and ComplianceFor_Master,
// asserting field-level correctness. It reproduces the post-seed VENDORINFO and
// COMPLIANCE verification blocks of script 10
// (10-test-upgrade-1.6.0-to-master.sh). The seeding flow certifies pid_1,
// then provisions -> certifies -> revokes pid_2, so pid_2 ends in the revoked
// state (and is no longer provisional).
//
//nolint:funlen
func VerifyMasterComplianceAndVendorInfo(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	// ----------------------------------------------------------------------
	// VENDORINFO
	// ----------------------------------------------------------------------
	// New master vendor record.
	out, err := QueryVendor(DcldMasterBinaryPath, VIDForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vendorID", VIDForMaster)
	checkResponseContains(t, out, CompanyLegalNameForMaster)

	// vid_for_1_2 vendor was updated with master-era preferred name + landing
	// page URL (companyLegalName/vendorName kept the 1.2-era values).
	out, err = QueryVendor(DcldMasterBinaryPath, VIDFor1_2)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vendorID", VIDFor1_2)
	checkResponseContains(t, out, VendorNameFor1_2)
	checkResponseContains(t, out, CompanyPreferredNameForMaster)
	checkResponseContains(t, out, VendorLandingPageURLForMaster)

	// all-vendors listing contains the master vendor.
	out, err = QueryAllVendors(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vendorID", VIDForMaster)
	checkResponseContains(t, out, CompanyLegalNameForMaster)
	checkResponseContains(t, out, VendorNameForMaster)

	// ----------------------------------------------------------------------
	// COMPLIANCE
	// ----------------------------------------------------------------------
	// pid_1 was certified -> certified-model value true.
	out, err = QueryCertifiedModel(DcldMasterBinaryPath, VIDForMaster, PID1ForMaster, SoftwareVersionForMaster, CertificationTypeForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "value", true)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)
	requireFieldEquals(t, out, "softwareVersion", SoftwareVersionForMaster)
	checkResponseContains(t, out, CertificationTypeForMaster)

	// pid_2 was provisioned -> certified -> revoked, so it ends revoked.
	out, err = QueryRevokedModel(DcldMasterBinaryPath, VIDForMaster, PID2ForMaster, SoftwareVersionForMaster, CertificationTypeForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "value", true)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID2ForMaster)

	// pid_2 is no longer provisional -> provisional-model value false.
	out, err = QueryProvisionalModel(DcldMasterBinaryPath, VIDForMaster, PID2ForMaster, SoftwareVersionForMaster, CertificationTypeForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "value", false)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID2ForMaster)

	// compliance-info for both pids.
	out, err = QueryComplianceInfo(DcldMasterBinaryPath, VIDForMaster, PID1ForMaster, SoftwareVersionForMaster, CertificationTypeForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)
	requireFieldEquals(t, out, "softwareVersion", SoftwareVersionForMaster)
	checkResponseContains(t, out, CertificationTypeForMaster)

	out, err = QueryComplianceInfo(DcldMasterBinaryPath, VIDForMaster, PID2ForMaster, SoftwareVersionForMaster, CertificationTypeForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID2ForMaster)
	requireFieldEquals(t, out, "softwareVersion", SoftwareVersionForMaster)
	checkResponseContains(t, out, CertificationTypeForMaster)

	// device-software-compliance for the master cdCertId (set on the pid_1
	// certify + pid_2 provision/certify txs).
	out, err = QueryDeviceSoftwareCompliance(DcldMasterBinaryPath, CDCertificateIDForMaster)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)
	checkResponseContains(t, out, CDCertificateIDForMaster)

	// all-* listings contain the master-era records.
	out, err = QueryAllCertifiedModels(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)

	out, err = QueryAllRevokedModels(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID2ForMaster)

	// No master-era provisional record persists (pid_2 ended revoked); assert
	// the carry-over provisional record (vid/pid_3 from the 0.12 era) is still
	// present, matching the script's all-provisional-models check.
	out, err = QueryAllProvisionalModels(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", state.VID)
	requireFieldEquals(t, out, "pid", pid3V012)

	out, err = QueryAllComplianceInfo(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)
	requireFieldEquals(t, out, "pid", PID2ForMaster)

	out, err = QueryAllDeviceSoftwareCompliance(DcldMasterBinaryPath)
	require.NoError(t, err)
	requireFieldEquals(t, out, "vid", VIDForMaster)
	requireFieldEquals(t, out, "pid", PID1ForMaster)
	checkResponseContains(t, out, CDCertificateIDForMaster)
}
