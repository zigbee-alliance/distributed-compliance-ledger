package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

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

		out, err = QueryVendorModelsHex(vidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
	})

	productLabel := "Device #1"

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			ProductLabel: productLabel,
			From:         vendorAccount,
		})
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

		out, err = QueryVendorModelsHex(vidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
	})

	description := "New Device Description"

	t.Run("UpdateModel", func(t *testing.T) {
		txResult, err := UpdateModelHex(vidHex, pidHex, vendorAccount,
			"--enhancedSetupFlowOptions", "2",
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
		txResult, err := UpdateModelHex(vidHex, pidHex, vendorAccount,
			"--enhancedSetupFlowOptions", "2",
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
		txResult, err := DeleteModelHex(vidHex, pidHex, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
