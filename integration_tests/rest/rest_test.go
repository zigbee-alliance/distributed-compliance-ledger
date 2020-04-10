package rest

import (
	"encoding/json"
	"fmt"
	app "git.dsr-corporation.com/zb-ledger/zb-ledger"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
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
		* prepare config with `genlocalconfig.sh`
		* update `/.zbld/config/genesis.json` to set `administrator` role to the first account as described in Readme (#Genesis template)
		* run node with `zbld start`
		* run RPC service with `zblcli rest-server --chain-id zblchain --trust-node`

	TODO: prepare environment automatically
*/

const (
	ChainId                  = "zblchain"
	AdministratorAccountName = "jack"
	AccountName              = "alice"
	AccountPassphrase        = "test1234"
)

func TestDemo(t *testing.T) {
	// Get all model infos
	inputModelInfos := getModelInfos()

	// Get key info for Jack
	adminKeyInfo := getKeyInfo(AdministratorAccountName)

	// Get account info for Jack
	adminAccountInfo := getAccountInfo(adminKeyInfo.Address)
	require.Equal(t, adminAccountInfo.Roles, []string{string(authz.Administrator)})

	// Get account info for Jack
	adminAccountInfo2 := getAccountInfo(adminKeyInfo.Address)
	require.Equal(t, adminAccountInfo2.Roles, []string{string(authz.Administrator)})

	// Get key info for Alice
	aliceKeyInfo := getKeyInfo(AccountName)

	// Get account info for Alice
	aliceAccountInfo := getAccountInfo(aliceKeyInfo.Address)
	require.Equal(t, aliceAccountInfo.Roles, []string{})

	// Assign Manufacturer role to Alice
	newMsgAssignRole := authz.NewMsgAssignRole(
		aliceKeyInfo.Address,
		authz.Manufacturer,
		adminKeyInfo.Address,
	)

	// Sign and Broadcast AssignRole message
	signAndBroadcastMessage(adminKeyInfo.Name, adminAccountInfo, newMsgAssignRole)

	// Get account info for Alice
	aliceAccountInfo = getAccountInfo(aliceKeyInfo.Address)
	require.Equal(t, aliceAccountInfo.Roles, []string{string(authz.Manufacturer)})

	// Prepare model info
	id := utils.RandString()

	newMsgAddModelInfo := compliance.NewMsgAddModelInfo(
		id,
		test_constants.Name,
		aliceKeyInfo.Address,
		test_constants.Description,
		test_constants.Sku,
		test_constants.FirmwareVersion,
		test_constants.HardwareVersion,
		test_constants.CertificateID,
		test_constants.CertifiedDate,
		test_constants.TisOrTrpTestingCompleted,
		aliceKeyInfo.Address,
	)

	// Sign and Broadcast AddModelInfo message
	signAndBroadcastMessage(aliceKeyInfo.Name, aliceAccountInfo, newMsgAddModelInfo)

	// Check model is created
	receivedModelInfo := getModelInfo(id)
	require.Equal(t, receivedModelInfo.ID, newMsgAddModelInfo.ID)
	require.Equal(t, receivedModelInfo.Name, newMsgAddModelInfo.Name)

	// Get all model infos
	modelInfos := getModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+1, utils.ParseUint(modelInfos.Total))
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
	response := utils.SendPostRequest(uri, body)

	var result interface{}
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func broadcastMessage(message interface{}) {
	println("Broadcast response")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	utils.SendPostRequest(uri, body)
}

func getModelInfo(id string) compliance.ModelInfo {
	println("Get Model Info with ID: ", id)

	uri := fmt.Sprintf("%s/%s/%s", compliance.RouterKey, compliance.QueryModelInfo, id)
	response := utils.SendGetRequest(uri)

	var result compliance.ModelInfo
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func getModelInfos() utils.ModelInfoHeadersResult {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", compliance.RouterKey, compliance.QueryModelInfo)
	response := utils.SendGetRequest(uri)

	var result utils.ModelInfoHeadersResult
	_ = json.Unmarshal(responseResult(response), &result)

	return result
}

func responseResult(response []byte) json.RawMessage {
	var responseWrapper utils.ResponseWrapper
	_ = json.Unmarshal(response, &responseWrapper)
	return responseWrapper.Result
}
