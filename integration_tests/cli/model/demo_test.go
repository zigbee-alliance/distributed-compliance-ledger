package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelDemo(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	vidWithPids := vid + 1
	pid := rand.Intn(65534) + 1
	pidRanges := fmt.Sprintf("%d-%d", pid, pid)
	vendorAccountWithPids := fmt.Sprintf("vendor_account_%d", vidWithPids)
	cliputils.CreateVendorAccount(t, vendorAccountWithPids, vidWithPids, pidRanges)

	t.Run("QueryNonExistentModel", func(t *testing.T) {
		out, err := QueryModel(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryVendorModels(vid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
	})

	productLabel := "Device #1"
	sv := 1
	cdVersionNum := 10

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", productLabel,
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--schemaVersion", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelWithPidRanges", func(t *testing.T) {
		enhancedSetupFlowTCUrl := "https://example.org/file.txt"
		enhancedSetupFlowTCRevision := 1
		enhancedSetupFlowTCDigest := "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA=="
		enhancedSetupFlowTCFileSize := 1024
		enhancedSetupFlowMaintenanceUrl := "https://example.org"
		commissioningFallbackUrl := "https://url.commissioningfallbackurl.dclmodel"

		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vidWithPids),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", productLabel,
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "1",
			"--enhancedSetupFlowTCUrl", enhancedSetupFlowTCUrl,
			"--enhancedSetupFlowTCRevision", fmt.Sprintf("%d", enhancedSetupFlowTCRevision),
			"--enhancedSetupFlowTCDigest", enhancedSetupFlowTCDigest,
			"--enhancedSetupFlowTCFileSize", fmt.Sprintf("%d", enhancedSetupFlowTCFileSize),
			"--enhancedSetupFlowMaintenanceUrl", enhancedSetupFlowMaintenanceUrl,
			"--commissioningFallbackUrl", commissioningFallbackUrl,
			"--discoveryCapabilitiesBitmask", "1",
			"--from", vendorAccountWithPids,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		out, err := QueryModel(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"productLabel":"%s"`, productLabel))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), `"enhancedSetupFlowOptions":0`)
	})

	t.Run("AddModelVersions", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--minApplicableSoftwareVersion", "1",
			"--maxApplicableSoftwareVersion", "15",
			"--softwareVersionString", fmt.Sprintf("%d", sv),
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNum),
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vidWithPids),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--minApplicableSoftwareVersion", "1",
			"--maxApplicableSoftwareVersion", "15",
			"--softwareVersionString", fmt.Sprintf("%d", sv),
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNum),
			"--from", vendorAccountWithPids,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAllModelsAndVendorModels", func(t *testing.T) {
		out, err := QueryAllModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))

		out, err = QueryVendorModels(vid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
	})

	description := "New Device Description"
	newCommissioningModeInitialStepsHint := 8
	newCommissioningModeSecondaryStepsHint := 9
	newIcdUserActiveModeTriggerHint := 7
	newFactoryResetStepsHint := 6

	t.Run("UpdateModel", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--from", vendorAccount,
			"--productLabel", description,
			"--schemaVersion", "0",
			"--commissioningModeInitialStepsHint", fmt.Sprintf("%d", newCommissioningModeInitialStepsHint),
			"--commissioningModeSecondaryStepsHint", fmt.Sprintf("%d", newCommissioningModeSecondaryStepsHint),
			"--icdUserActiveModeTriggerHint", fmt.Sprintf("%d", newIcdUserActiveModeTriggerHint),
			"--enhancedSetupFlowOptions", "2",
			"--factoryResetStepsHint", fmt.Sprintf("%d", newFactoryResetStepsHint),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryUpdatedModel", func(t *testing.T) {
		out, err := QueryModel(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"productLabel":"%s"`, description))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), fmt.Sprintf(`"commissioningModeInitialStepsHint":%d`, newCommissioningModeInitialStepsHint))
		require.Contains(t, string(out), fmt.Sprintf(`"commissioningModeSecondaryStepsHint":%d`, newCommissioningModeSecondaryStepsHint))
		require.Contains(t, string(out), fmt.Sprintf(`"icdUserActiveModeTriggerHint":%d`, newIcdUserActiveModeTriggerHint))
		require.Contains(t, string(out), fmt.Sprintf(`"factoryResetStepsHint":%d`, newFactoryResetStepsHint))
		require.Contains(t, string(out), `"enhancedSetupFlowOptions":2`)
	})

	supportURL := "https://newsupporturl.test"

	t.Run("UpdateModelSupportURL", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--from", vendorAccount,
			"--supportURL", supportURL,
			"--enhancedSetupFlowOptions", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModel(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"supportUrl":"%s"`, supportURL))
	})

	t.Run("DeleteModels", func(t *testing.T) {
		txResult, err := DeleteModel(vid, pid, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = DeleteModel(vidWithPids, pid, vendorAccountWithPids)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterDeletion", func(t *testing.T) {
		out, err := QueryModel(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryModel(vidWithPids, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryModelVersion(vidWithPids, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
