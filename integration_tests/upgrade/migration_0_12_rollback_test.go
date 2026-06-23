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

// runRollback012 submits a propose+approve sequence for an upgrade plan
// whose name has no registered cosmovisor handler ("wrong_plan_name").
// Once block height passes plan_height, the chain proceeds without
// applying any upgrade. Then seeds _for_rollback state (vendor 4705 +
// models + compliance + users).
//
//nolint:funlen
func runRollback012(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcld, err := EnsureBinary("0.12.0")
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// Wrong-plan-name upgrade: propose + approve, then verify NOT applied.
	// ------------------------------------------------------------------
	MustRun(t, "WrongPlanNameUpgrade", func(t *testing.T) {
		t.Helper()
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		upgradeInfo := UpgradeInfoForVersion(BinaryVersionV1_2, WrongPlanChecksumV12)

		tx, err := ProposeUpgrade(dcld, WrongPlanName, planHeight, upgradeInfo, state.Trustee1)
		requireTxSuccess(t, tx, err)

		for _, who := range []string{state.Trustee2, state.Trustee3} {
			tx, err = ApproveUpgrade(dcld, WrongPlanName, who)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "approve %s: %s", who, tx.RawLog)
		}

		cliputils.WaitForHeight(t, planHeight+1, 300)

		// No upgrade scheduled anymore.
		out, _ := QueryUpgradePlan(dcld)
		require.True(t, strings.Contains(string(out), "no upgrade scheduled"),
			"expected 'no upgrade scheduled', got: %s", string(out))

		// `applied` query must fail / return Not Found for the wrong plan name.
		// Either non-zero exit or a "not applied"/"not found"-ish output is
		// acceptable.
		out, err = QueryAppliedPlan(dcld, WrongPlanName)
		if err == nil {
			require.False(t, strings.Contains(string(out), `"height"`),
				"upgrade unexpectedly applied: %s", string(out))
		}
	})

	// ------------------------------------------------------------------
	// Verify carry-over data from script 01 is intact — post-rollback
	// readback of vendor info, models, compliance, pki, accounts, validator.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedV0_12", func(t *testing.T) {
		t.Helper()
		// ----- VendorInfo -----
		out, err := QueryVendor(dcld, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)

		out, err = QueryAllVendors(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)

		// ----- Model -----
		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryGetModel(dcld, state.VID, pid)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			checkResponseContains(t, out, state.ProductLabel)
		}

		out, err = QueryAllModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryVendorModels(dcld, state.VID)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryModelVersion(dcld, state.VID, pid, state.SoftwareVersion)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}

		// ----- Compliance -----
		out, err = QueryCertifiedModel(dcld, state.VID, pid1V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryRevokedModel(dcld, state.VID, state.PID2, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryProvisionalModel(dcld, state.VID, pid3V012, state.SoftwareVersion, certificationTypeV012)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = QueryComplianceInfo(dcld, state.VID, pid, state.SoftwareVersion, certificationTypeV012)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}

		out, err = QueryDeviceSoftwareCompliance(dcld, cdCertificateIDV012)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryAllCertifiedModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = QueryAllProvisionalModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		out, err = QueryAllRevokedModels(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryAllComplianceInfo(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = QueryAllDeviceSoftwareCompliance(dcld)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, cdCertificateIDV012)

		// ----- PKI -----
		out, err = QueryAllX509RootCerts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryAllRevokedX509RootCerts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, rootCertSubject)
		checkResponseContains(t, out, rootCertSubjectKeyID)

		out, err = QueryAllProposedX509RootCerts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		out, err = QueryAllProposedX509RootCertsToRevoke(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryX509Cert(dcld, testRootCertSubject, testRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = QueryProposedX509RootCert(dcld, googleRootCertSubject, googleRootCertSubjectKeyID)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		// ----- Auth -----
		out, err = QueryAllAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = QueryAllProposedAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address)

		out, err = QueryAllProposedAccountsToRevoke(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)

		out, err = QueryAllRevokedAccounts(dcld)
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address)

		// ----- Validator (host-side queries against the chain) -----
		// Skip silently if the validator-demo container wasn't initialized
		// in this test run.
		if state.ValidatorAddress != "" {
			out, err = QueryProposedDisableNode(dcld, state.ValidatorAddress)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)

			out, err = QueryAllNodes(dcld)
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade seed (chain is still v0.12 since rollback no-op'd).
	// ------------------------------------------------------------------
	MustRun(t, "CreateRollbackAccounts", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		_ = CreateAndApproveAccount(t, dcld, VendorAccountForRollback, "Vendor",
			VIDForRollback, state.Trustee1, approvers)
		_ = CreateAndApproveAccount(t, dcld, CertificationCenterAccountForRollback,
			"CertificationCenter", -1, state.Trustee1, approvers)
	})

	MustRun(t, "AddRollbackUserKeys", func(t *testing.T) {
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

	MustRun(t, "VendorInfoForRollback", func(t *testing.T) {
		t.Helper()
		tx, err := AddVendor(dcld, VendorArgs{VID: VIDForRollback, VendorName: VendorNameForRollback, CompanyLegalName: CompanyLegalNameForRollback, CompanyPreferredName: CompanyPreferredNameForRollback, VendorLandingPageURL: VendorLandingPageURLForRollback, From: VendorAccountForRollback})
		requireTxSuccess(t, tx, err)

		// Update 0.12 vendor with rollback-era values (companyPreferredName etc).
		tx, err = UpdateVendor(dcld, VendorArgs{VID: state.VID, VendorName: vendorNameV012, CompanyLegalName: companyLegalNameV012, CompanyPreferredName: CompanyPreferredNameForRollback, VendorLandingPageURL: VendorLandingPageURLForRollback, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ModelsForRollback", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1ForRollback, PID2ForRollback, PID3ForRollback} {
			tx, err := AddModel(dcld, AddModelArgs{VID: VIDForRollback, PID: pid, DeviceTypeID: DeviceTypeIDForRollback, ProductName: ProductNameForRollback, ProductLabel: ProductLabelForRollback, PartNumber: PartNumberForRollback, From: VendorAccountForRollback})
			requireTxSuccess(t, tx, err)

			tx, err = AddModelVersion(dcld, AddModelVersionArgs{VID: VIDForRollback, PID: pid, SoftwareVersion: SoftwareVersionForRollback, SoftwareVersionString: SoftwareVersionStringForRollback, CDVersionNumber: CDVersionNumberForRollback, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionForRollback, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionForRollback, From: VendorAccountForRollback})
			requireTxSuccess(t, tx, err)
		}

		// Delete pid_3.
		tx, err := DeleteModel(dcld, VIDForRollback, PID3ForRollback, VendorAccountForRollback)
		requireTxSuccess(t, tx, err)

		// Update 0.12 pid_2 with rollback productLabel/partNumber.
		tx, err = UpdateModel(dcld, UpdateModelArgs{VID: state.VID, PID: state.PID2, ProductName: state.ProductName, ProductLabel: ProductLabelForRollback, PartNumber: PartNumberForRollback, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)

		tx, err = UpdateModelVersion(dcld, UpdateModelVersionArgs{VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionForRollback, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionForRollback, From: state.VendorAccount})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "ComplianceForRollback", func(t *testing.T) {
		t.Helper()
		// certify pid_1
		tx, err := CertifyModel(dcld, CertifyModelArgs{VID: VIDForRollback, PID: PID1ForRollback, SoftwareVersion: SoftwareVersionForRollback, SoftwareVersionString: SoftwareVersionStringForRollback, CertificationType: CertificationTypeForRollback, CertificationDate: CertificationDateForRollback, CDCertificateID: CDCertificateIDForRollback, CDVersionNumber: CDVersionNumberForRollback, From: CertificationCenterAccountForRollback})
		requireTxSuccess(t, tx, err)

		// provision pid_2, certify pid_2, revoke pid_2. revoke-model does not
		// accept --cdCertificateId, so it is set only on provision/certify.
		tx, err = ProvisionModel(dcld, ProvisionModelArgs{
			VID: VIDForRollback, PID: PID2ForRollback, SoftwareVersion: SoftwareVersionForRollback,
			SoftwareVersionString: SoftwareVersionStringForRollback, CertificationType: CertificationTypeForRollback,
			ProvisionalDate: ProvisionalDateForRollback, CDCertificateID: CDCertificateIDForRollback,
			CDVersionNumber: CDVersionNumberForRollback, From: CertificationCenterAccountForRollback,
		})
		requireTxSuccess(t, tx, err)

		tx, err = CertifyModel(dcld, CertifyModelArgs{
			VID: VIDForRollback, PID: PID2ForRollback, SoftwareVersion: SoftwareVersionForRollback,
			SoftwareVersionString: SoftwareVersionStringForRollback, CertificationType: CertificationTypeForRollback,
			CertificationDate: CertificationDateForRollback, CDCertificateID: CDCertificateIDForRollback,
			CDVersionNumber: CDVersionNumberForRollback, From: CertificationCenterAccountForRollback,
		})
		requireTxSuccess(t, tx, err)

		tx, err = RevokeModel(dcld, RevokeModelArgs{
			VID: VIDForRollback, PID: PID2ForRollback, SoftwareVersion: SoftwareVersionForRollback,
			SoftwareVersionString: SoftwareVersionStringForRollback, CertificationType: CertificationTypeForRollback,
			RevocationDate: CertificationDateForRollback, CDVersionNumber: CDVersionNumberForRollback,
			From: CertificationCenterAccountForRollback,
		})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "AccountFlowsForRollback", func(t *testing.T) {
		t.Helper()
		// 2 trustee approvals: with 4 trustees on-chain, the threshold for
		// revoke-account (1/3 quorum) is still satisfied by 2 approvals.
		approvers2 := []string{state.Trustee2, state.Trustee3}

		proposeUserAccount(t, dcld, state.Trustee1, approvers2,
			state.User4Address, state.User4Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, approvers2,
			state.User5Address, state.User5Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, nil,
			state.User6Address, state.User6Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcld, state.Trustee1, approvers2, state.User4Address, true)
		revokeUserAccount(t, dcld, state.Trustee1, nil, state.User5Address, false)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		// Script 02 mirrors script 01's 2-trustee disable approval pattern.
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3})
	})
}
