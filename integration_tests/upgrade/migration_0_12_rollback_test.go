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
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

// runRollback012 is the Go translation of
// integration_tests/upgrade/02-test-upgrade-0.12-rollback.sh.
//
// Submits a propose+approve sequence for an upgrade plan whose name has no
// registered cosmovisor handler ("wrong_plan_name"). Once block height passes
// plan_height, the chain proceeds without applying any upgrade. Then seeds
// _for_rollback state (vendor 4705 + models + compliance + users).
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
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		upgradeInfo := UpgradeInfoForVersion("v"+BinaryVersionV1_2, WrongPlanChecksumV12)

		tx, err := ProposeUpgrade(dcld, WrongPlanName, planHeight, upgradeInfo, state.Trustee1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

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
		out, err = QueryAppliedPlan(dcld, WrongPlanName)
		// Bash uses `! query upgrade applied` so either non-zero exit or
		// "not applied"/"not found"-ish output is acceptable.
		if err == nil {
			require.False(t, strings.Contains(string(out), `"height"`),
				"upgrade unexpectedly applied: %s", string(out))
		}
	})

	// ------------------------------------------------------------------
	// Verify carry-over data from script 01 is intact.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedV0_12", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcld,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)

		// Spot-check models and compliance.
		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		out, err = ExecuteCLIWithBin(dcld, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address) // revoked in 01
	})

	// ------------------------------------------------------------------
	// Post-upgrade seed (chain is still v0.12 since rollback no-op'd).
	// ------------------------------------------------------------------
	MustRun(t, "CreateRollbackAccounts", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		_ = CreateAndApproveAccount(t, dcld, VendorAccountForRollback, "Vendor",
			VIDForRollback, state.Trustee1, approvers)
		_ = CreateAndApproveAccount(t, dcld, CertificationCenterAccountForRollback,
			"CertificationCenter", -1, state.Trustee1, approvers)
	})

	MustRun(t, "AddRollbackUserKeys", func(t *testing.T) {
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
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDForRollback),
			"--vendorName", VendorNameForRollback,
			"--companyLegalName", CompanyLegalNameForRollback,
			"--companyPreferredName", CompanyPreferredNameForRollback,
			"--vendorLandingPageURL", VendorLandingPageURLForRollback,
			"--from", VendorAccountForRollback,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update 0.12 vendor with rollback-era values (companyPreferredName etc).
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--vendorName", vendorNameV012,
			"--companyLegalName", companyLegalNameV012,
			"--companyPreferredName", CompanyPreferredNameForRollback,
			"--vendorLandingPageURL", VendorLandingPageURLForRollback,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ModelsForRollback", func(t *testing.T) {
		for _, pid := range []int{PID1ForRollback, PID2ForRollback, PID3ForRollback} {
			tx, err := ExecuteTxWithBin(dcld,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDForRollback),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDForRollback),
				"--productName", ProductNameForRollback,
				"--productLabel", ProductLabelForRollback,
				"--partNumber", PartNumberForRollback,
				"--from", VendorAccountForRollback,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcld,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDForRollback),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForRollback),
				"--softwareVersionString", SoftwareVersionStringForRollback,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForRollback),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionForRollback),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionForRollback),
				"--from", VendorAccountForRollback,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDForRollback),
			"--pid", fmt.Sprintf("%d", PID3ForRollback),
			"--from", VendorAccountForRollback,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update 0.12 pid_2 with rollback productLabel/partNumber.
		tx, err = ExecuteTxWithBin(dcld,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelForRollback,
			"--partNumber", PartNumberForRollback,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcld,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionForRollback),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionForRollback),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ComplianceForRollback", func(t *testing.T) {
		// certify pid_1
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDForRollback),
			"--pid", fmt.Sprintf("%d", PID1ForRollback),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForRollback),
			"--softwareVersionString", SoftwareVersionStringForRollback,
			"--certificationType", CertificationTypeForRollback,
			"--certificationDate", CertificationDateForRollback,
			"--cdCertificateId", CDCertificateIDForRollback,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForRollback),
			"--from", CertificationCenterAccountForRollback,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2, certify pid_2, revoke pid_2.
		// revoke-model does not accept --cdCertificateId, so it is appended
		// only for the provision/certify actions.
		for _, action := range []struct {
			cmd, dateFlag, dateVal string
		}{
			{"provision-model", "--provisionalDate", ProvisionalDateForRollback},
			{"certify-model", "--certificationDate", CertificationDateForRollback},
			{"revoke-model", "--revocationDate", CertificationDateForRollback},
		} {
			args := []string{
				"tx", "compliance", action.cmd,
				"--vid", fmt.Sprintf("%d", VIDForRollback),
				"--pid", fmt.Sprintf("%d", PID2ForRollback),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForRollback),
				"--softwareVersionString", SoftwareVersionStringForRollback,
				"--certificationType", CertificationTypeForRollback,
				action.dateFlag, action.dateVal,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForRollback),
				"--from", CertificationCenterAccountForRollback,
			}
			if action.cmd != "revoke-model" {
				args = append(args, "--cdCertificateId", CDCertificateIDForRollback)
			}
			tx, err = ExecuteTxWithBin(dcld, args...)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "%s pid_2: %s", action.cmd, tx.RawLog)
		}
	})

	MustRun(t, "AccountFlowsForRollback", func(t *testing.T) {
		// Note: bash only uses 2-trustee approvals here, but with the rollback
		// users being created post-trustee_4 (already 4 trustees on-chain),
		// the threshold for revoke-account is still satisfied by 2 approvals.
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
		// Script 02 mirrors script 01's 2-trustee disable approval pattern.
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3})
	})
}
