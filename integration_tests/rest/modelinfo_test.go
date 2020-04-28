package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

func /*TestModelinfo*/ Demo(t *testing.T) {
	// Get all model infos
	inputModelInfos, _ := utils.GetModelInfos()

	// Get all vendors
	inputVendors, _ := utils.GetVendors()

	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(test_constants.AccountName)

	// Get account info for Jack
	jackAccountInfo, _ := utils.GetAccountInfo(jackKeyInfo.Address)

	// Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Prepare model info
	firstModelInfo := utils.NewMsgAddModelInfo(jackAccountInfo.Address)
	VID := firstModelInfo.VID

	// Sign and Broadcast AddModelInfo message
	utils.SignAndBroadcastMessage(jackKeyInfo, firstModelInfo)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(firstModelInfo.VID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, firstModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, firstModelInfo.Name)

	// Publish second model info using POST command with passing name and passphrase. Same Vendor
	secondModelInfo := utils.NewMsgAddModelInfo(jackAccountInfo.Address)
	secondModelInfo.VID = VID // Set same Vendor as for the first model
	_, _ = utils.PublishModelInfo(secondModelInfo)

	// Check model is created
	receivedModelInfo, _ = utils.GetModelInfo(secondModelInfo.VID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, secondModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, secondModelInfo.Name)

	// Get all model infos
	modelInfos, _ := utils.GetModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+2, utils.ParseUint(modelInfos.Total))

	// Get all vendors
	vendors, _ := utils.GetVendors()
	require.Equal(t, utils.ParseUint(inputVendors.Total)+1, utils.ParseUint(vendors.Total))

	// Get vendor models
	vendorModels, _ := utils.GetVendorModels(VID)
	require.Equal(t, uint64(2), uint64(len(vendorModels.Products)))
	require.Equal(t, firstModelInfo.PID, vendorModels.Products[0].PID)
	require.Equal(t, secondModelInfo.PID, vendorModels.Products[1].PID)
}
