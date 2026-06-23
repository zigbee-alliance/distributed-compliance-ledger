package dclauth

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/model"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
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
	err := cliputils.AddKey(name)
	require.NoError(t, err)

	userAddr, err := cliputils.GetAddress(name)
	require.NoError(t, err)
	userPubkey, err := cliputils.GetPubkey(name)
	require.NoError(t, err)

	jackAddr, err := cliputils.GetAddress(jack)
	require.NoError(t, err)

	t.Run("ProposeVendorAccountWithHexVID", func(t *testing.T) {
		txResult, err := ProposeAccount(userAddr, userPubkey, "Vendor", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VIDHex: vidHex})
		cliputils.RequireTxOK(t, txResult, err)

		// Vendor accounts need only 1/3 approvals — should be active immediately.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, userAddr, acc.Address)
		require.Len(t, acc.Approvals, 1)
		require.Equal(t, jackAddr, acc.Approvals[0].Address)
		require.Equal(t, "Jack is proposing this account", acc.Approvals[0].Info)

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))
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
		cliputils.RequireTxOK(t, txResult, err)
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
		txResult, err := model.UpdateModel(model.UpdateModelOpts{
			VIDHex: vidHex, PIDHex: pidHex, From: userAddr,
			ProductName:              productName,
			ProductLabel:             "Device Description",
			PartNumber:               "12",
			EnhancedSetupFlowOptions: 2,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		m, err := model.GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, "Device #1", m.ProductName)
	})
}
