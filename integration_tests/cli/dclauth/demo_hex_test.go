package dclauth

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/model"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// It creates a Vendor account using a hex-format VID, adds a model with the hex VID,
// verifies that using a different VID is rejected, updates the model, and queries it.
func TestAuthDemoHex(t *testing.T) {
	jack := testconstants.JackAccount

	// Use random VID/PID expressed as hex to avoid collisions across test runs.
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vidHex := fmt.Sprintf("0x%X", vid)
	pidHex := fmt.Sprintf("0x%X", pid)

	// Generate and add key
	name := "hexvendor" + utils.RandString()
	err := AddKey(name)
	require.NoError(t, err)

	userAddr, err := GetAddress(name)
	require.NoError(t, err)
	userPubkey, err := GetPubkey(name)
	require.NoError(t, err)

	jackAddr, err := GetAddress(jack)
	require.NoError(t, err)

	t.Run("ProposeVendorAccountWithHexVID", func(t *testing.T) {
		txResult, err := ProposeAccount(userAddr, userPubkey, "Vendor", jack, ProposeAccountOpts{Info: "Jack is proposing this account", Extra: []string{"--vid", vidHex}})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Vendor accounts need only 1/3 approvals — should be active immediately
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), "Jack is proposing this account")

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
	})

	t.Run("AddModelWithHexVID", func(t *testing.T) {
		productName := "Device #1"
		txResult, err := model.AddModel(model.AddModelOpts{
			VIDHex:       vidHex,
			PIDHex:       pidHex,
			DeviceTypeID: 12,
			ProductName:  productName,
			ProductLabel: "Device Description",
			PartNumber:   "12",
			From:         userAddr,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelWithWrongVID_Fails", func(t *testing.T) {
		vidPlusOneHex := "0xA14"
		productName := "Device #1"

		txResult, err := model.AddModel(model.AddModelOpts{
			VIDHex:       vidPlusOneHex,
			PIDHex:       pidHex,
			DeviceTypeID: 12,
			ProductName:  productName,
			ProductLabel: "Device Description",
			PartNumber:   "12",
			From:         userAddr,
		})
		// With broadcast-mode sync, rejection might come at CLI level or on-chain.
		if err != nil {
			require.Contains(t, err.Error(), "vendorID")

			return
		}
		// Await on-chain result to drain tx from mempool before next subtest.
		txData, awaitErr := utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, awaitErr)
		require.Contains(t, string(txData), "vendorID")
	})

	t.Run("UpdateModel", func(t *testing.T) {
		productName := "Device #1"
		txResult, err := model.UpdateModelHex(vidHex, pidHex, userAddr,
			"--productName", productName,
			"--productLabel", "Device Description",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "2",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		out, err := model.QueryModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), `"vid":`+formatInt(vid))
		require.Contains(t, string(out), `"pid":`+formatInt(pid))
		require.Contains(t, string(out), "Device #1")
	})
}

func formatInt(n int) string {
	return itoa(n)
}
