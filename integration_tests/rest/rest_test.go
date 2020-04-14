package rest

import (
	"encoding/json"
	"fmt"
	app "git.dsr-corporation.com/zb-ledger/zb-ledger"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	modelinfoRest "git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/client/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* prepare config with `genlocalconfig.sh`
		* update `/.zbld/config/genesis.json` to set `administrator` role to the first account as described in Readme (#Genesis template)
		* run node with `zbld start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
*/

const (
	ChainId           = "zblchain"
	AccountName       = "jack"
	AccountPassphrase = "test1234"
)

func /*Test*/Demo(t *testing.T) {
	// Get all model infos
	inputModelInfos := getModelInfos()

	// Get all vendors
	inputVendors := getVendors()

	// Get key info for Jack
	jackKeyInfo := getKeyInfo(AccountName)

	// Get account info for Jack
	jackAccountInfo := getAccountInfo(jackKeyInfo.Address)
	require.Equal(t, jackAccountInfo.Roles, []string{string(authz.Administrator)})

	// Assign Vendor role to Jack
	newMsgAssignRole := authz.NewMsgAssignRole(
		jackAccountInfo.Address,
		authz.Vendor,
		jackKeyInfo.Address,
	)

	// Sign and Broadcast AssignRole message
	signAndBroadcastMessage(jackKeyInfo.Name, jackAccountInfo, newMsgAssignRole)

	// Get account info for Jack
	jackAccountInfo = getAccountInfo(jackAccountInfo.Address)
	require.Equal(t, jackAccountInfo.Roles, []string{string(authz.Administrator), string(authz.Vendor)})

	// Prepare model info
	vid := int16(utils.RandInt())
	pid := int16(utils.RandInt())

	newMsgAddModelInfo := modelinfo.NewMsgAddModelInfo(
		vid,
		pid,
		int16(utils.RandInt()),
		test_constants.Name,
		test_constants.Description,
		test_constants.Sku,
		test_constants.FirmwareVersion,
		test_constants.HardwareVersion,
		test_constants.Custom,
		test_constants.TisOrTrpTestingCompleted,
		jackAccountInfo.Address,
	)

	// Sign and Broadcast AddModelInfo message
	signAndBroadcastMessage(jackKeyInfo.Name, jackAccountInfo, newMsgAddModelInfo)

	// Check model is created
	receivedModelInfo := getModelInfo(vid, pid)
	require.Equal(t, receivedModelInfo.VID, newMsgAddModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, newMsgAddModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, newMsgAddModelInfo.Name)

	// Publish second model info using POST command with passing name and passphrase. Same Vendor
	pid = int16(utils.RandInt())
	secondModelInfo := modelinfo.NewModelInfo(
		vid,
		pid,
		int16(utils.RandInt()),
		"Second Model Name",
		jackAccountInfo.Address,
		"Other Description",
		test_constants.Sku,
		test_constants.FirmwareVersion,
		test_constants.HardwareVersion,
		test_constants.Custom,
		test_constants.TisOrTrpTestingCompleted,
	)

	publishModelInfo(jackAccountInfo, secondModelInfo)

	// Check model is created
	receivedModelInfo = getModelInfo(vid, pid)
	require.Equal(t, receivedModelInfo.VID, secondModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, secondModelInfo.Name)

	// Get all model infos
	modelInfos := getModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+2, utils.ParseUint(modelInfos.Total))

	// Get all vendors
	vendors := getVendors()
	require.Equal(t, utils.ParseUint(inputVendors.Total)+1, utils.ParseUint(vendors.Total))

	// Get vendor models
	vendorModels := getVendorModels(vendors.Items[0].VID)
	require.Equal(t, uint64(2), uint64(len(vendorModels.Products)))
	require.Equal(t, newMsgAddModelInfo.PID, vendorModels.Products[0].PID)
	require.Equal(t, secondModelInfo.PID, vendorModels.Products[1].PID)
}

func getKeyInfo(accountName string) utils.KeyInfo {
	println("Get User Key Info")

	uri := fmt.Sprintf("key/%s", accountName)
	response := utils.SendGetRequest(uri)

	var keyInfo utils.KeyInfo
	_ = json.Unmarshal(response, &keyInfo)

	return keyInfo
}

func getAccountInfo(address sdk.AccAddress) utils.AccountInfo {
	println("Get Account Info")

	uri := fmt.Sprintf("%s/account/%s", authnext.RouterKey, address)
	response := utils.SendGetRequest(uri)

	var result utils.AccountInfo
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func signAndBroadcastMessage(accountName string, accountInfo utils.AccountInfo, message sdk.Msg) {
	signResponse := signMessage(accountName, accountInfo, message)
	broadcastMessage(signResponse)
}

func publishModelInfo(accountInfo utils.AccountInfo, model modelinfo.ModelInfo) interface{} {
	println("Publish Model Info")

	request := modelinfoRest.ModelInfoRequest{
		BaseReq: rest.BaseReq{
			ChainID: ChainId,
			From:    accountInfo.Address.String(),
		},
		VID:                      model.VID,
		PID:                      model.PID,
		CID:                      model.CID,
		Name:                     model.Name,
		Description:              model.Description,
		SKU:                      model.SKU,
		FirmwareVersion:          model.FirmwareVersion,
		HardwareVersion:          model.HardwareVersion,
		Custom:                   model.Custom,
		TisOrTrpTestingCompleted: model.TisOrTrpTestingCompleted,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")
	response := utils.SendPostRequest(uri, body, AccountName, AccountPassphrase)

	var result interface{}
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func signMessage(accountName string, accountInfo utils.AccountInfo, message sdk.Msg) interface{} {
	println("Sign Message")

	stdSigMsg := types.StdSignMsg{
		ChainID:       ChainId,
		AccountNumber: utils.ParseUint(accountInfo.AccountNumber),
		Sequence:      utils.ParseUint(accountInfo.Sequence),
		Fee:           types.StdFee{Gas: 200000,},
		Msgs:          []sdk.Msg{message},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s?name=%s&passphrase=%s", "tx", "sign", accountName, AccountPassphrase)
	response := utils.SendPostRequest(uri, body, "", "")

	var result interface{}
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func broadcastMessage(message interface{}) {
	println("Broadcast Message")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	utils.SendPostRequest(uri, body, "", "")
}

func getModelInfo(vid int16, pid int16) modelinfo.ModelInfo {
	println(fmt.Sprintf("Get Model Info with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", modelinfo.RouterKey, "models", vid, pid)
	response := utils.SendGetRequest(uri)

	var result modelinfo.ModelInfo
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func getModelInfos() utils.ModelInfoHeadersResult {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")
	response := utils.SendGetRequest(uri)

	var result utils.ModelInfoHeadersResult
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func getVendors() utils.VendorItemHeadersResult {
	println("Get the list of vendors")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "vendors")
	response := utils.SendGetRequest(uri)

	var result utils.VendorItemHeadersResult
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func getVendorModels(vid int16) modelinfo.VendorProducts {
	println("Get the list of models for VID:", vid)

	uri := fmt.Sprintf("%s/%s/%v", modelinfo.RouterKey, "models", vid)
	response := utils.SendGetRequest(uri)

	var result modelinfo.VendorProducts
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func responseResult(response []byte) json.RawMessage {
	var responseWrapper utils.ResponseWrapper
	_ = json.Unmarshal(response, &responseWrapper)
	return responseWrapper.Result
}
