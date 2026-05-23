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

// runUpgrade160ToMaster is the Go translation of
// integration_tests/upgrade/10-test-upgrade-1.6.0-to-master.sh.
//
// Builds the master image, hands the binary to each localnet node via
// cosmovisor add-upgrade, then proposes/approves an upgrade plan whose name
// is the master git short hash. After upgrade applies, seeds master-era state.
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
	t.Run("BuildAndDistributeMasterBinary", func(t *testing.T) {
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
	t.Run("ProposeApproveMasterUpgrade", func(t *testing.T) {
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		// Master upgrade plan submission doesn't pass --upgrade-info — the
		// binary was already seeded into cosmovisor manually above.
		tx, err := ExecuteTxWithBin(dcldOld,
			"tx", "dclupgrade", "propose-upgrade",
			"--name", planName,
			"--upgrade-height", fmt.Sprintf("%d", planHeight),
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

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
	t.Run("VerifyPreservedAcrossAllEras", func(t *testing.T) {
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4, state.VIDFor1_5_1, VIDFor1_5_2, VIDFor1_6_0} {
			out, qerr := ExecuteCLIWithBin(DcldMasterBinaryPath,
				"query", "vendorinfo", "vendor",
				"--vid", fmt.Sprintf("%d", vid),
			)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// 1.5.2 pid_2 has 1.6.0 productLabel/partNumber (set in script 09).
		out, err := ExecuteCLIWithBin(DcldMasterBinaryPath,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_6_0)
		checkResponseContains(t, out, PartNumberFor1_6_0)
	})

	// ------------------------------------------------------------------
	// Seed master-era state.
	// ------------------------------------------------------------------
	t.Run("CreateVendor_Master", func(t *testing.T) {
		_ = CreateAndApproveAccount(t, DcldMasterBinaryPath, VendorAccountForMaster, "Vendor",
			VIDForMaster, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	t.Run("AddMasterUserKeys", func(t *testing.T) {
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

	t.Run("VendorInfoFor_Master", func(t *testing.T) {
		tx, err := ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDForMaster),
			"--vendorName", VendorNameForMaster,
			"--companyLegalName", CompanyLegalNameForMaster,
			"--companyPreferredName", CompanyPreferredNameForMaster,
			"--vendorLandingPageURL", VendorLandingPageURLForMaster,
			"--from", VendorAccountForMaster,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--vendorName", VendorNameFor1_2,
			"--companyLegalName", CompanyLegalNameFor1_2,
			"--companyPreferredName", CompanyPreferredNameForMaster,
			"--vendorLandingPageURL", VendorLandingPageURLForMaster,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ModelsAndVersionsFor_Master", func(t *testing.T) {
		for _, pid := range []int{PID1ForMaster, PID2ForMaster, PID3ForMaster} {
			tx, err := ExecuteTxWithBin(DcldMasterBinaryPath,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDForMaster),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDForMaster),
				"--productName", ProductNameForMaster,
				"--productLabel", ProductLabelForMaster,
				"--partNumber", PartNumberForMaster,
				"--from", VendorAccountForMaster,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(DcldMasterBinaryPath,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDForMaster),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForMaster),
				"--softwareVersionString", SoftwareVersionStringForMaster,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForMaster),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionForMaster),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionForMaster),
				"--from", VendorAccountForMaster,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDForMaster),
			"--pid", fmt.Sprintf("%d", PID3ForMaster),
			"--from", VendorAccountForMaster,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update carry-over 0.12 pid_2 with master-era values.
		tx, err = ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelForMaster,
			"--partNumber", PartNumberForMaster,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionForMaster),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionForMaster),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ComplianceFor_Master", func(t *testing.T) {
		// certify pid_1.
		tx, err := ExecuteTxWithBin(DcldMasterBinaryPath,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDForMaster),
			"--pid", fmt.Sprintf("%d", PID1ForMaster),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForMaster),
			"--softwareVersionString", SoftwareVersionStringForMaster,
			"--certificationType", CertificationTypeForMaster,
			"--certificationDate", CertificationDateForMaster,
			"--cdCertificateId", CDCertificateIDForMaster,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForMaster),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2, certify pid_2, revoke pid_2.
		for _, action := range []struct{ cmd, dateFlag, dateVal string }{
			{"provision-model", "--provisionalDate", ProvisionalDateForMaster},
			{"certify-model", "--certificationDate", CertificationDateForMaster},
			{"revoke-model", "--revocationDate", CertificationDateForMaster},
		} {
			tx, err = ExecuteTxWithBin(DcldMasterBinaryPath,
				"tx", "compliance", action.cmd,
				"--vid", fmt.Sprintf("%d", VIDForMaster),
				"--pid", fmt.Sprintf("%d", PID2ForMaster),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionForMaster),
				"--softwareVersionString", SoftwareVersionStringForMaster,
				"--certificationType", CertificationTypeForMaster,
				action.dateFlag, action.dateVal,
				"--cdCertificateId", CDCertificateIDForMaster,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberForMaster),
				"--from", CertificationCenterAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "%s pid_2: %s", action.cmd, tx.RawLog)
		}
	})

	t.Run("AccountFlowsFor_Master", func(t *testing.T) {
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

	t.Run("ValidatorDisableEnableFlow", func(t *testing.T) {
		RunValidatorDisableEnableFlow(t, state, DcldMasterBinaryPath,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade data is in place. Phase 4's add-new-node script
	// inherits this state.
	// ------------------------------------------------------------------
	t.Run("VerifyNew_Master_Data", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(DcldMasterBinaryPath,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDForMaster),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDForMaster)
		checkResponseContains(t, out, CompanyLegalNameForMaster)

		out, err = ExecuteCLIWithBin(DcldMasterBinaryPath,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDForMaster),
			"--pid", fmt.Sprintf("%d", PID1ForMaster),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDForMaster)
		requireFieldEquals(t, out, "pid", PID1ForMaster)
		checkResponseContains(t, out, ProductLabelForMaster)

		// Master plan name is recorded in state for script 11 to verify
		// against (the new observer should eventually report this version).
		state.MasterPlanName = planName
	})
}
