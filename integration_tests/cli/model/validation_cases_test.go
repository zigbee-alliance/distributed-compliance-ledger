package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestModelValidationCases translates model-validation-cases.sh.
func TestModelValidationCases(t *testing.T) {
	vid1 := rand.Intn(65534) + 1
	vendorAccount1 := fmt.Sprintf("vendor_account_%d", vid1)
	cliputils.CreateVendorAccount(t, vendorAccount1, vid1)

	pid1 := rand.Intn(65534) + 1
	pid2 := rand.Intn(65534) + 1
	pid3 := rand.Intn(65534) + 1

	t.Run("AddModel_MinimumFields", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "Test Product",
			"--partNumber", "1",
			"--enhancedSetupFlowOptions", "0",
			"--commissioningCustomFlow", "0",
			"--from", vendorAccount1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid1))
		require.Contains(t, string(out), `"productName":"TestProduct"`)
		require.Contains(t, string(out), `"partNumber":"1"`)
		require.Contains(t, string(out), `"commissioningCustomFlow":0`)
	})

	t.Run("AddModel_AllFields", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid2),
			"--deviceTypeID", "2",
			"--productName", "Test Product with All Fields",
			"--productLabel", "Test Product with All fields",
			"--partNumber", "23.456",
			"--commissioningCustomFlow", "1",
			"--commissioningCustomFlowURL", "https://customflow.url.info",
			"--commissioningModeInitialStepsHint", "1",
			"--commissioningModeInitialStepsInstruction", "Initial Instructions",
			"--commissioningModeSecondaryStepsHint", "2",
			"--commissioningModeSecondaryStepsInstruction", "Secondary Steps Instruction",
			"--icdUserActiveModeTriggerHint", "4",
			"--icdUserActiveModeTriggerInstruction", "ICD User Active Mode Trigger Instruction",
			"--factoryResetStepsHint", "3",
			"--factoryResetStepsInstruction", "Factory Reset Steps Instruction",
			"--userManualURL", "https://usermanual.url",
			"--productURL", "https://product.url.info",
			"--lsfURL", "https://lsf.url.info",
			"--supportURL", "https://support.url.info",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModel(vid1, pid2)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), `"productName":"Test Product with All Fields"`)
		require.Contains(t, string(out), `"deviceTypeId":2`)
	})

	t.Run("AddModel_WithSchemaVersion", func(t *testing.T) {
		// schemaVersion must be 0 for add-model (the only valid value)
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid3),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "Test Product",
			"--partNumber", "1",
			"--enhancedSetupFlowOptions", "0",
			"--commissioningCustomFlow", "0",
			"--schemaVersion", "0",
			"--from", vendorAccount1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModel(vid1, pid3)
		require.NoError(t, err)
		require.Contains(t, string(out), `"schemaVersion":0`)
	})
}
