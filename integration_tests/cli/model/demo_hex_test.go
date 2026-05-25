package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestModelDemoHex translates model-demo-hex.sh.
func TestModelDemoHex(t *testing.T) {
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vidHex := fmt.Sprintf("0x%X", vid)
	pidHex := fmt.Sprintf("0x%X", pid)

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	t.Run("QueryNonExistent", func(t *testing.T) {
		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = utils.ExecuteCLI("query", "model", "vendor-models",
			"--vid", vidHex, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
	})

	productLabel := "Device #1"

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", productLabel,
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"productLabel":"%s"`, productLabel))

		out, err = QueryAllModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))

		out, err = utils.ExecuteCLI("query", "model", "vendor-models",
			"--vid", vidHex, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
	})

	description := "New Device Description"

	t.Run("UpdateModel", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "update-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--enhancedSetupFlowOptions", "2",
			"--from", vendorAccount,
			"--productLabel", description,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"productLabel":"%s"`, description))
	})

	supportURL := "https://newsupporturl.test"

	t.Run("UpdateModelSupportURL", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "update-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--enhancedSetupFlowOptions", "2",
			"--from", vendorAccount,
			"--supportURL", supportURL,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"supportUrl":"%s"`, supportURL))
	})

	t.Run("DeleteModel", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "delete-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
