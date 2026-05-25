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

// runRollback12 is the Go translation of
// integration_tests/upgrade/04-test-upgrade-1.2-rollback.sh.
//
// Same shape as runRollback012, but submitted against a v1.2 chain with
// "wrong_plan_name_2" — verifying the chain doesn't accidentally upgrade to
// v1.4.3 just because a plan with that target binary was approved.
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
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		upgradeInfo := UpgradeInfoForVersion("v"+BinaryVersionV1_4_3, WrongPlanChecksumV143)

		tx, err := ProposeUpgrade(dcld, WrongPlanName2, planHeight, upgradeInfo, state.Trustee1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

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
	// Verify carry-over from scripts 01/02/03 is intact.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAfterRollback1_2", func(t *testing.T) {
		// VendorInfo from 03 update of 0.12 vendor.
		out, err := ExecuteCLIWithBin(dcld,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, CompanyPreferredNameFor1_2)

		// 1.2 vendor.
		out, err = ExecuteCLIWithBin(dcld,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		// 0.12 pid_2 has 1.2 productLabel/partNumber from 03's update.
		out, err = ExecuteCLIWithBin(dcld,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_2)
		checkResponseContains(t, out, PartNumberFor1_2)
	})

	// ------------------------------------------------------------------
	// Seed _for_1_2_r2 state (still on v1.2).
	// ------------------------------------------------------------------
	MustRun(t, "CreateR2Accounts", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}
		_ = CreateAndApproveAccount(t, dcld, VendorAccountFor1_2R2, "Vendor",
			VIDFor1_2R2, state.Trustee1, approvers)
	})

	MustRun(t, "AddR2UserKeys", func(t *testing.T) {
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
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--vendorName", VendorNameFor1_2R2,
			"--companyLegalName", CompanyLegalNameFor1_2R2,
			"--companyPreferredName", CompanyPreferredNameFor1_2R2,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_2R2,
			"--from", VendorAccountFor1_2R2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ModelsForR2", func(t *testing.T) {
		for _, pid := range []int{PID1For1_2R2, PID2For1_2R2, PID3For1_2R2} {
			tx, err := ExecuteTxWithBin(dcld,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_2R2),
				"--productName", ProductNameFor1_2R2,
				"--productLabel", ProductLabelFor1_2R2,
				"--partNumber", PartNumberFor1_2R2,
				"--from", VendorAccountFor1_2R2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcld,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
				"--softwareVersionString", SoftwareVersionStringFor1_2R2,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_2R2),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_2R2),
				"--from", VendorAccountFor1_2R2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--pid", fmt.Sprintf("%d", PID3For1_2R2),
			"--from", VendorAccountFor1_2R2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ComplianceForR2", func(t *testing.T) {
		// certify pid_1.
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--pid", fmt.Sprintf("%d", PID1For1_2R2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
			"--softwareVersionString", SoftwareVersionStringFor1_2R2,
			"--certificationType", CertificationTypeFor1_2R2,
			"--certificationDate", CertificationDateFor1_2R2,
			"--cdCertificateId", CDCertificateIDFor1_2R2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2, certify pid_2, revoke pid_2.
		for _, action := range []struct {
			cmd, dateFlag, dateVal string
		}{
			{"provision-model", "--provisionalDate", ProvisionalDateFor1_2R2},
			{"certify-model", "--certificationDate", CertificationDateFor1_2R2},
			{"revoke-model", "--revocationDate", CertificationDateFor1_2R2},
		} {
			tx, err = ExecuteTxWithBin(dcld,
				"tx", "compliance", action.cmd,
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", PID2For1_2R2),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
				"--softwareVersionString", SoftwareVersionStringFor1_2R2,
				"--certificationType", CertificationTypeFor1_2R2,
				action.dateFlag, action.dateVal,
				"--cdCertificateId", CDCertificateIDFor1_2R2,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
				"--from", CertificationCenterAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "%s pid_2: %s", action.cmd, tx.RawLog)
		}
	})

	MustRun(t, "AccountFlowsForR2", func(t *testing.T) {
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
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})
}
