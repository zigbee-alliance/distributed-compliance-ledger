package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* prepare config with `genlocalconfig.sh`
		* update `/.zbld/config/genesis.json` to set `Administrator` role to the first account as described in Readme (#Genesis template)
		* run node with `zbld start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
*/

func TestModelinfoDemo(t *testing.T) {
	// Get all model infos
	inputModelInfos := utils.GetModelInfos()

	// Get all vendors
	inputVendors := utils.GetVendors()

	// Get key info for Jack
	jackKeyInfo := utils.GetKeyInfo(test_constants.AccountName)

	// Get account info for Jack
	jackAccountInfo := utils.GetAccountInfo(jackKeyInfo.Address)

	// Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Prepare model info
	VID := int16(utils.RandInt())

	firstModelInfo := modelinfo.NewMsgAddModelInfo(
		VID,
		int16(utils.RandInt()),
		int16(utils.RandInt()),
		utils.RandString(),
		test_constants.Description,
		test_constants.Sku,
		test_constants.FirmwareVersion,
		test_constants.HardwareVersion,
		test_constants.Custom,
		test_constants.TisOrTrpTestingCompleted,
		jackAccountInfo.Address,
	)

	// Sign and Broadcast AddModelInfo message
	utils.SignAndBroadcastMessage(jackKeyInfo, firstModelInfo)

	// Check model is created
	receivedModelInfo := utils.GetModelInfo(firstModelInfo.VID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, firstModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, firstModelInfo.Name)

	// Publish second model info using POST command with passing name and passphrase. Same Vendor
	secondModelInfo := utils.NewModelInfo(jackAccountInfo.Address)
	secondModelInfo.VID = VID // Set same Vendor as for the first model
	utils.PublishModelInfo(jackAccountInfo.Address, secondModelInfo)

	// Check model is created
	receivedModelInfo = utils.GetModelInfo(secondModelInfo.VID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, secondModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, secondModelInfo.Name)

	// Get all model infos
	modelInfos := utils.GetModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+2, utils.ParseUint(modelInfos.Total))

	// Get all vendors
	vendors := utils.GetVendors()
	require.Equal(t, utils.ParseUint(inputVendors.Total)+1, utils.ParseUint(vendors.Total))

	// Get vendor models
	vendorModels := utils.GetVendorModels(VID)
	require.Equal(t, uint64(2), uint64(len(vendorModels.Products)))
	require.Equal(t, firstModelInfo.PID, vendorModels.Products[0].PID)
	require.Equal(t, secondModelInfo.PID, vendorModels.Products[1].PID)
}
