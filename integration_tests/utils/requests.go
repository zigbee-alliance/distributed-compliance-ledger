package utils

import (
	"encoding/json"
	"fmt"
	app "git.dsr-corporation.com/zb-ledger/zb-ledger"
	constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	complianceRest "git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/client/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	compliancetestRest "git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/client/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	modelinfoRest "git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func GetKeyInfo(accountName string) KeyInfo {
	println("Get User Key Info")

	uri := fmt.Sprintf("key/%s", accountName)
	response := SendGetRequest(uri)

	var keyInfo KeyInfo
	_ = json.Unmarshal(response, &keyInfo)

	return keyInfo
}

func GetAccountInfo(address sdk.AccAddress) AccountInfo {
	println("Get Account Info")

	uri := fmt.Sprintf("%s/account/%s", authnext.RouterKey, address)
	response := SendGetRequest(uri)

	var result AccountInfo
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func SignAndBroadcastMessage(sender KeyInfo, message sdk.Msg) {
	senderAccountInfo := GetAccountInfo(sender.Address) // Refresh account info
	signResponse := SignMessage(sender.Name, senderAccountInfo, message)
	BroadcastMessage(signResponse)
}

func PublishModelInfo(model modelinfo.MsgAddModelInfo) json.RawMessage {
	println("Publish Model Info")

	request := modelinfoRest.ModelInfoRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    model.Signer.String(),
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
	response := SendPostRequest(uri, body, constants.AccountName, constants.AccountPassphrase)
	return removeResponseWrapper(response)
}

func SignMessage(accountName string, accountInfo AccountInfo, message sdk.Msg) interface{} {
	println("Sign Message")

	stdSigMsg := types.StdSignMsg{
		ChainID:       constants.ChainId,
		AccountNumber: ParseUint(accountInfo.AccountNumber),
		Sequence:      ParseUint(accountInfo.Sequence),
		Fee:           types.StdFee{Gas: 200000},
		Msgs:          []sdk.Msg{message},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s?name=%s&passphrase=%s", "tx", "sign", accountName, constants.AccountPassphrase)
	response := SendPostRequest(uri, body, "", "")

	var result interface{}
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func BroadcastMessage(message interface{}) {
	println("Broadcast Message")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	SendPostRequest(uri, body, "", "")
}

func GetModelInfo(vid int16, pid int16) modelinfo.ModelInfo {
	println(fmt.Sprintf("Get Model Info with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", modelinfo.RouterKey, "models", vid, pid)
	response := SendGetRequest(uri)

	var result modelinfo.ModelInfo
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func GetModelInfos() ModelInfoHeadersResult {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")
	response := SendGetRequest(uri)

	var result ModelInfoHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func GetVendors() VendorItemHeadersResult {
	println("Get the list of vendors")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "vendors")
	response := SendGetRequest(uri)

	var result VendorItemHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func GetVendorModels(vid int16) modelinfo.VendorProducts {
	println("Get the list of models for VID:", vid)

	uri := fmt.Sprintf("%s/%s/%v", modelinfo.RouterKey, "models", vid)
	response := SendGetRequest(uri)

	var result modelinfo.VendorProducts
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func PublishTestingResult(testingResult compliancetest.MsgAddTestingResult) json.RawMessage {
	println("Publish Testing Result")

	request := compliancetestRest.TestingResultRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    testingResult.Signer.String(),
		},
		VID:        testingResult.VID,
		PID:        testingResult.PID,
		TestResult: testingResult.TestResult,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", compliancetest.RouterKey, "testresults")
	response := SendPostRequest(uri, body, constants.AccountName, constants.AccountPassphrase)
	return removeResponseWrapper(response)
}

func GetTestingResult(vid int16, pid int16) compliancetest.TestingResults {
	println(fmt.Sprintf("Get Testing Result for Model with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", compliancetest.RouterKey, "testresults", vid, pid)
	response := SendGetRequest(uri)

	var result compliancetest.TestingResults
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func AssignRole(targetAddress sdk.AccAddress, sender KeyInfo, role authz.AccountRole) {
	// Assign TestHouse role to Jack
	newMsgAssignRole := authz.NewMsgAssignRole(
		targetAddress,
		role,
		sender.Address,
	)

	// Sign and Broadcast AssignRole message
	SignAndBroadcastMessage(sender, newMsgAssignRole)
}

func PublishCertifiedModel(certifyModel compliance.MsgCertifyModel) json.RawMessage {
	println("Publish Certified Model")

	request := complianceRest.CertifiedModelRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    certifyModel.Signer.String(),
		},
		VID:               certifyModel.VID,
		PID:               certifyModel.PID,
		CertificationDate: certifyModel.CertificationDate,
		CertificationType: certifyModel.CertificationType,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", compliance.RouterKey, "certified")
	response := SendPostRequest(uri, body, constants.AccountName, constants.AccountPassphrase)
	return removeResponseWrapper(response)
}

func GetCertifiedModel(vid int16, pid int16) compliance.CertifiedModel {
	println(fmt.Sprintf("Get Certification Data for Model with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", compliance.RouterKey, "certified", vid, pid)
	response := SendGetRequest(uri)

	var result compliance.CertifiedModel
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func GetAllCertifiedModels() CertifiedModelsHeadersResult {
	println("Get all certified models")

	uri := fmt.Sprintf("%s/%s", compliance.RouterKey, "certified")
	response := SendGetRequest(uri)

	var result CertifiedModelsHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result
}

func NewMsgAddModelInfo(owner sdk.AccAddress) modelinfo.MsgAddModelInfo {
	return modelinfo.NewMsgAddModelInfo(
		int16(RandInt()),
		int16(RandInt()),
		constants.CID,
		RandString(),
		RandString(),
		RandString(),
		constants.FirmwareVersion,
		constants.HardwareVersion,
		RandString(),
		constants.TisOrTrpTestingCompleted,
		owner,
	)
}

func NewMsgAddTestingResult(vid int16, pid int16, owner sdk.AccAddress) compliancetest.MsgAddTestingResult {
	return compliancetest.NewMsgAddTestingResult(
		vid,
		pid,
		RandString(),
		owner,
	)
}

func removeResponseWrapper(response []byte) json.RawMessage {
	var responseWrapper ResponseWrapper
	_ = json.Unmarshal(response, &responseWrapper)
	return responseWrapper.Result
}
