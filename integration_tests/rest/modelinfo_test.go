package rest_test

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/common"
	"net/http"
	"testing"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

func TestModelinfoDemo(t *testing.T) {
	// Get all model infos
	inputModelInfos, _ := utils.GetModelInfos()

	// Get all vendors
	inputVendors, _ := utils.GetVendors()

	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.AccountName)

	// Get account info for Jack
	jackAccountInfo, _ := utils.GetAccountInfo(jackKeyInfo.Address)

	// Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, auth.Vendor)

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
	_, _ = utils.PublishModelInfo(secondModelInfo, jackKeyInfo)

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

func TestModelinfoDemo_Prepare_Sign_Broadcast(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.AccountName)

	// Register new Vendor account
	vendor, _ := utils.RegisterNewAccount()
	utils.AssignRole(vendor.Address, jackKeyInfo, auth.Vendor)

	// Prepare model info
	modelInfo := utils.NewMsgAddModelInfo(vendor.Address)

	// Prepare Sing Broadcast
	addModelTransaction, _ := utils.PrepareModelInfoTransaction(modelInfo)
	_, _ = utils.SignAndBroadcastTransaction(vendor, addModelTransaction)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(modelInfo.VID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
}

/* Error cases */

func Test_AddModelinfo_ByNonVendor(t *testing.T) {
	// register new account
	testAccount, _ := utils.RegisterNewAccount()

	// try to publish model info
	modelInfo := utils.NewMsgAddModelInfo(testAccount.Address)
	res, _ := utils.SignAndBroadcastMessage(testAccount, modelInfo)
	require.Equal(t, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
}

func Test_AddModelinfo_Twice(t *testing.T) {
	// register new account
	testAccount, _ := utils.RegisterNewAccount()

	// get jack account
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.JackAccount)

	// Assign Vendor role to test account
	utils.AssignRole(testAccount.Address, jackKeyInfo, auth.Vendor)

	// publish model info
	modelInfo := utils.NewMsgAddModelInfo(jackKeyInfo.Address)
	res, _ := utils.PublishModelInfo(modelInfo, jackKeyInfo)
	require.Equal(t, sdk.CodeOK, sdk.CodeType(res.Code))

	// publish second time
	res, _ = utils.PublishModelInfo(modelInfo, jackKeyInfo)
	require.Equal(t, modelinfo.CodeModelInfoAlreadyExists, sdk.CodeType(res.Code))
}

func Test_GetModelinfo_ForUnknown(t *testing.T) {
	_, code := utils.GetModelInfo(common.RandUint16(), common.RandUint16())
	require.Equal(t, http.StatusNotFound, code)
}

func Test_GetModelinfo_ForInvalidVidPid(t *testing.T) {
	// zero vid
	_, code := utils.GetModelInfo(0, common.RandUint16())
	require.Equal(t, http.StatusBadRequest, code)

	// zero pid
	_, code = utils.GetModelInfo(common.RandUint16(), 0)
	require.Equal(t, http.StatusBadRequest, code)
}
