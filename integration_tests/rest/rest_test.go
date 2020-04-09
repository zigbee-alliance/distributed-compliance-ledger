package rest

import (
	"encoding/json"
	"fmt"
	app "git.dsr-corporation.com/zb-ledger/zb-ledger"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* init ledger
		* run RPC service with `zblcli rest-server --chain-id zblchain --trust-node`

	TODO: prepare environment automatically?
	TODO: Generic response deserialization
*/

const (
	ChainId           = "zblchain"
	AccountName       = "jack"
	AccountPassphrase = "test1234"
)

func Demo(t *testing.T) {
	// Get all model infos
	inputModelInfos := getModelInfos()

	// Get key info
	keyInfo := getKeyInfo()

	// Get account address account
	accountInfo := getAccountInfo(keyInfo.Address)

	// Prepare model info
	id := utils.RandString()

	newMsgAddModelInfo := compliance.NewMsgAddModelInfo(
		id,
		test_constants.Name,
		keyInfo.Address,
		test_constants.Description,
		test_constants.Sku,
		test_constants.FirmwareVersion,
		test_constants.HardwareVersion,
		test_constants.CertificateID,
		test_constants.CertifiedDate,
		test_constants.TisOrTrpTestingCompleted,
		keyInfo.Address,
	)

	// Sign and Broadcast AddModelInfo message
	signAndBroadcastMessage(accountInfo, newMsgAddModelInfo)

	// Check model is created
	receivedModelInfo := getModelInfo(id)
	require.Equal(t, receivedModelInfo.ID, newMsgAddModelInfo.ID)
	require.Equal(t, receivedModelInfo.Name, newMsgAddModelInfo.Name)

	// Get all model infos
	modelInfos := getModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+1, utils.ParseUint(modelInfos.Total))
}

func getKeyInfo() utils.KeyInfo {
	println("Get User Key Info")

	uri := fmt.Sprintf("key/%s", AccountName)
	response := utils.SendGetRequest(uri)

	var keyInfo utils.KeyInfo
	_ = json.Unmarshal(response, &keyInfo)
	return keyInfo
}

func getAccountInfo(address sdk.AccAddress) utils.AccountInfo {
	println("Get Account Info")

	uri := fmt.Sprintf("%s/account/%s", authnext.RouterKey, address)
	response := utils.SendGetRequest(uri)

	var accountInfo utils.GetAccountResponse
	_ = json.Unmarshal(response, &accountInfo)

	return accountInfo.Result
}

func signAndBroadcastMessage(accountInfo utils.AccountInfo, message sdk.Msg) {
	signResponse := signMessage(accountInfo, message)
	broadcastMessage(signResponse)
}

func signMessage(accountInfo utils.AccountInfo, message sdk.Msg) utils.SignedMessage {
	println("Sign Message")

	stdSigMsg := types.StdSignMsg{
		ChainID:       ChainId,
		AccountNumber: utils.ParseUint(accountInfo.AccountNumber),
		Sequence:      utils.ParseUint(accountInfo.Sequence),
		Fee:           types.StdFee{Gas: 200000,},
		Msgs:          []sdk.Msg{message},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s?name=%s&passphrase=%s", "tx", "sign", AccountName, AccountPassphrase)
	response := utils.SendPostRequest(uri, body)

	var result utils.SignMessageResponse
	_ = json.Unmarshal(response, &result)

	return result.Result
}

func broadcastMessage(message utils.SignedMessage) {
	println("Broadcast response")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	utils.SendPostRequest(uri, body)
}

func getModelInfo(id string) compliance.ModelInfo {
	println("Get Model Info with ID: ", id)

	uri := fmt.Sprintf("%s/%s/%s", compliance.RouterKey, compliance.QueryModelInfo, id)
	response := utils.SendGetRequest(uri)

	var result utils.GetModelInfoResponse
	_ = json.Unmarshal(response, &result)

	return result.Result
}

func getModelInfos() utils.ModelInfoHeadersResult {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", compliance.RouterKey, compliance.QueryModelInfo)
	response := utils.SendGetRequest(uri)

	var result utils.GetListModelInfoResponse
	_ = json.Unmarshal(response, &result)

	return result.Result
}
