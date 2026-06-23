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

// runRollback12 mirrors runRollback012's no-op-upgrade pattern, this time
// against a v1.2 chain with "wrong_plan_name_2" — verifying the chain
// doesn't accidentally upgrade to v1.4.3 just because a plan with that
// target binary was approved.
//
//nolint:funlen
func runRollback12(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcld, err := EnsureBinary(BinaryVersionV1_2)
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// Wrong-plan-name upgrade attempt.
	// ------------------------------------------------------------------
	MustRun(t, "WrongPlanName2Upgrade", func(t *testing.T) {
		t.Helper()
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		upgradeInfo := UpgradeInfoForVersion(BinaryVersionV1_4_3, WrongPlanChecksumV143)

		tx, err := ProposeUpgrade(dcld, WrongPlanName2, planHeight, upgradeInfo, state.Trustee1)
		requireTxSuccess(t, tx, err)

		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveUpgrade(dcld, WrongPlanName2, who)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "approve %s: %s", who, tx.RawLog)
		}

		cliputils.WaitForHeight(t, planHeight+1, 300)

		out, _ := QueryUpgradePlan(dcld)
		require.True(t, strings.Contains(string(out), "no upgrade scheduled"),
			"expected 'no upgrade scheduled', got: %s", string(out))

		out, err = QueryAppliedPlan(dcld, WrongPlanName2)
		if err == nil {
			require.False(t, strings.Contains(string(out), `"height"`),
				"upgrade unexpectedly applied: %s", string(out))
		}
	})

	// ------------------------------------------------------------------
	// Verify carry-over from scripts 01/02/03 is intact — post-rollback
	// readback of vendor info, models, compliance, pki, accounts, validator.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAfterRollback1_2", func(t *testing.T) {
		t.Helper()
		// ----- VendorInfo -----
		out, err := QueryVendor(dcld, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, CompanyPreferredNameFor1_2)
		checkResponseContains(t, out, VendorLandingPageURLFor1_2)

		out, err = QueryVendor(dcld, VIDFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		out, err = QueryAllVendors(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		// ----- Model: 0.12-era pid_1 + pid_2 -----
		out, err = QueryGetModel(dcld, state.VID, pid1V012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, state.ProductLabel)

		out, err = QueryGetModel(dcld, state.VID, state.PID2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, ProductLabelFor1_2)
		checkResponseContains(t, out, PartNumberFor1_2)

		// 1.2-era pid_1 + pid_2.
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = QueryGetModel(dcld, VIDFor1_2, pid)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
			checkResponseContains(t, out, ProductLabelFor1_2)
		}

		out, err = QueryAllModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = QueryVendorModels(dcld, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryVendorModels(dcld, VIDFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		// Model versions.
		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryModelVersion(dcld, state.VID, pid, state.SoftwareVersion)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = QueryModelVersion(dcld, VIDFor1_2, pid, SoftwareVersionFor1_2)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_2)
		}

		// ----- Compliance -----
		out, err = QueryCertifiedModel(dcld, state.VID, pid1V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		out, err = QueryCertifiedModel(dcld, VIDFor1_2, PID1For1_2, SoftwareVersionFor1_2, CertificationTypeFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		out, err = QueryRevokedModel(dcld, state.VID, state.PID2, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryRevokedModel(dcld, VIDFor1_2, PID2For1_2, SoftwareVersionFor1_2, CertificationTypeFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = QueryProvisionalModel(dcld, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryComplianceInfo(dcld, state.VID, pid, state.SoftwareVersion, certificationTypeV012)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
		}
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = QueryComplianceInfo(dcld, VIDFor1_2, pid, SoftwareVersionFor1_2, CertificationTypeFor1_2)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
		}

		out, err = QueryDeviceSoftwareCompliance(dcld, cdCertificateIDV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryDeviceSoftwareCompliance(dcld, CDCertificateIDFor1_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		// Compliance all-* listings.
		out, err = QueryAllCertifiedModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		out, err = QueryAllProvisionalModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		out, err = QueryAllRevokedModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = QueryAllComplianceInfo(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = QueryAllDeviceSoftwareCompliance(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, cdCertificateIDV012)
		checkResponseContains(t, out, CDCertificateIDFor1_2)

		// ----- PKI single-record + listings -----
		// 1.2-era root_cert: test_root_cert assigned vid in script 03's assign-vid.
		out, err = QueryX509Cert(dcld, TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		// 0.12-era test_root_cert.
		out, err = QueryX509Cert(dcld, testRootCertSubject, testRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryAllSubjectX509Certs(dcld, TestRootCertSubjectFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		out, err = QueryAllSubjectX509Certs(dcld, testRootCertSubject)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryProposedX509RootCert(dcld, GoogleRootCertSubjectFor1_2, GoogleRootCertSubjectKeyIDFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, GoogleRootCertSubjectKeyIDFor1_2)

		out, err = QueryProposedX509RootCert(dcld, googleRootCertSubject, googleRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		// Revoked intermediate certs (v0.12 + v1.2).
		out, err = QueryRevokedX509Cert(dcld, IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, IntermediateCertSubjectFor1_2)
		checkResponseContains(t, out, IntermediateCertSubjectKeyIDFor1_2)

		out, err = QueryRevokedX509Cert(dcld, intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, intermediateCertSubject)
		checkResponseContains(t, out, intermediateCertSubjectKeyID)

		// Proposed-to-revoke (both eras).
		out, err = QueryProposedX509RootCertToRevoke(dcld, TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		out, err = QueryProposedX509RootCertToRevoke(dcld, testRootCertSubject, testRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		// Revocation point (single + listing by issuer + all).
		out, err = QueryRevocationPoint(dcld, VIDFor1_2, ProductLabelFor1_2, IssuerSubjectKeyID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)
		checkResponseContains(t, out, ProductLabelFor1_2)
		checkResponseContains(t, out, TestDataURL)

		out, err = QueryRevocationPoints(dcld, IssuerSubjectKeyID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)
		checkResponseContains(t, out, ProductLabelFor1_2)

		out, err = QueryAllRevocationPoints(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		// All-* PKI listings.
		out, err = QueryAllProposedX509RootCerts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, GoogleRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		out, err = QueryAllRevokedX509RootCerts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertSubjectFor1_2)
		checkResponseContains(t, out, RootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, rootCertSubject)
		checkResponseContains(t, out, rootCertSubjectKeyID)

		out, err = QueryAllProposedX509RootCertsToRevoke(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryAllX509Certs(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		// ----- Auth: full single-record + listing coverage -----
		out, err = QueryAllAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User2Address)

		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = QueryAccount(dcld, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User6Address)
		checkResponseContains(t, out, state.User3Address)

		for _, addr := range []string{state.User6Address, state.User3Address} {
			out, err = QueryProposedAccount(dcld, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllProposedAccountsToRevoke(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User2Address)

		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = QueryProposedAccountToRevoke(dcld, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = QueryAllRevokedAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User4Address)
		checkResponseContains(t, out, state.User1Address)

		for _, addr := range []string{state.User4Address, state.User1Address} {
			out, err = QueryRevokedAccount(dcld, addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		// ----- Validator (host-side) -----
		if state.ValidatorAddress != "" {
			out, err = QueryAllNodes(dcld)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Seed _for_1_2_r2 state (still on v1.2).
	// ------------------------------------------------------------------
	MustRun(t, "CreateR2Accounts", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}
		_ = CreateAndApproveAccount(t, dcld, VendorAccountFor1_2R2, "Vendor",
			VIDFor1_2R2, state.Trustee1, approvers)
	})

	MustRun(t, "AddR2UserKeys", func(t *testing.T) {
		t.Helper()
		u4, err := newUserKey(dcld)
		require.NoError(t, err)
		u5, err := newUserKey(dcld)
		require.NoError(t, err)
		u6, err := newUserKey(dcld)
		require.NoError(t, err)
		state.User4Address, state.User4Pubkey = u4.address, u4.pubkey
		state.User5Address, state.User5Pubkey = u5.address, u5.pubkey
		state.User6Address, state.User6Pubkey = u6.address, u6.pubkey
	})

	MustRun(t, "VendorInfoForR2", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcld, VendorArgs{VID: VIDFor1_2R2, VendorName: VendorNameFor1_2R2, CompanyLegalName: CompanyLegalNameFor1_2R2, CompanyPreferredName: CompanyPreferredNameFor1_2R2, VendorLandingPageURL: VendorLandingPageURLFor1_2R2, From: VendorAccountFor1_2R2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsForR2", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1For1_2R2, PID2For1_2R2, PID3For1_2R2} {
			tx, err := AddModel(dcld, AddModelArgs{VID: VIDFor1_2R2, PID: pid, DeviceTypeID: DeviceTypeIDFor1_2R2, ProductName: ProductNameFor1_2R2, ProductLabel: ProductLabelFor1_2R2, PartNumber: PartNumberFor1_2R2, From: VendorAccountFor1_2R2})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(dcld, AddModelVersionArgs{VID: VIDFor1_2R2, PID: pid, SoftwareVersion: SoftwareVersionFor1_2R2, SoftwareVersionString: SoftwareVersionStringFor1_2R2, CDVersionNumber: CDVersionNumberFor1_2R2, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_2R2, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_2R2, From: VendorAccountFor1_2R2})
			requireTxSuccess(t, tx, err)
		}

		// Delete pid_3.
		tx, err := DeleteModel(dcld, VIDFor1_2R2, PID3For1_2R2, VendorAccountFor1_2R2)
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceForR2", func(t *testing.T) {
		t.Helper()
		// certify pid_1.
		tx, err := CertifyModel(dcld, CertifyModelArgs{VID: VIDFor1_2R2, PID: PID1For1_2R2, SoftwareVersion: SoftwareVersionFor1_2R2, SoftwareVersionString: SoftwareVersionStringFor1_2R2, CertificationType: CertificationTypeFor1_2R2, CertificationDate: CertificationDateFor1_2R2, CDCertificateID: CDCertificateIDFor1_2R2, CDVersionNumber: CDVersionNumberFor1_2R2, From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		// provision pid_2, certify pid_2, revoke pid_2. revoke-model does not
		// accept --cdCertificateId, so it is set only on provision/certify.
		tx, err = ProvisionModel(dcld, ProvisionModelArgs{
			VID: VIDFor1_2R2, PID: PID2For1_2R2, SoftwareVersion: SoftwareVersionFor1_2R2,
			SoftwareVersionString: SoftwareVersionStringFor1_2R2, CertificationType: CertificationTypeFor1_2R2,
			ProvisionalDate: ProvisionalDateFor1_2R2, CDCertificateID: CDCertificateIDFor1_2R2,
			CDVersionNumber: CDVersionNumberFor1_2R2, From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		tx, err = CertifyModel(dcld, CertifyModelArgs{
			VID: VIDFor1_2R2, PID: PID2For1_2R2, SoftwareVersion: SoftwareVersionFor1_2R2,
			SoftwareVersionString: SoftwareVersionStringFor1_2R2, CertificationType: CertificationTypeFor1_2R2,
			CertificationDate: CertificationDateFor1_2R2, CDCertificateID: CDCertificateIDFor1_2R2,
			CDVersionNumber: CDVersionNumberFor1_2R2, From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)

		tx, err = RevokeModel(dcld, RevokeModelArgs{
			VID: VIDFor1_2R2, PID: PID2For1_2R2, SoftwareVersion: SoftwareVersionFor1_2R2,
			SoftwareVersionString: SoftwareVersionStringFor1_2R2, CertificationType: CertificationTypeFor1_2R2,
			RevocationDate: CertificationDateFor1_2R2, CDVersionNumber: CDVersionNumberFor1_2R2,
			From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "AccountFlowsForR2", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcld, state.Trustee1, approvers,
			state.User4Address, state.User4Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, approvers,
			state.User5Address, state.User5Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, nil,
			state.User6Address, state.User6Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcld, state.Trustee1, approvers, state.User4Address, true)
		revokeUserAccount(t, dcld, state.Trustee1, nil, state.User5Address, false)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})
}
