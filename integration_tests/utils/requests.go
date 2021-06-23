// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	keyUtil "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	restTypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/libs/common"
	app "github.com/zigbee-alliance/distributed-compliance-ledger"
	constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	extRest "github.com/zigbee-alliance/distributed-compliance-ledger/restext/tx/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	authRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/client/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	complianceRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/client/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	compliancetestRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/client/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo"
	modelinfoRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/client/rest"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
	pkiRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/rest"
)

func CreateKey(accountName string) (KeyInfo, int) {
	println("Create Key for: ", accountName)

	kb, _ := keyUtil.NewKeyBaseFromDir(app.DefaultCLIHome)

	entropySeed, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropySeed)
	_, _ = kb.CreateAccount(accountName, mnemonic, "", constants.Passphrase, 0, 0)

	return GetKeyInfo(accountName)
}

func GetKeyInfo(accountName string) (KeyInfo, int) {
	println("Get User Key Info: ", accountName)

	uri := fmt.Sprintf("key/%s", accountName)
	response, code := SendGetRequest(uri)

	var keyInfo KeyInfo

	parseGetReqResponse(response, &keyInfo, code)

	return keyInfo, code
}

func ProposeAddAccount(keyInfo KeyInfo, signer KeyInfo, roles auth.AccountRoles) (TxnResponse, int) {
	println("Propose Add Account for: ", keyInfo.Name)

	request := authRest.ProposeAddAccountRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    signer.Address.String(),
		},
		Address: keyInfo.Address,
		Pubkey:  keyInfo.PublicKey,
		Roles:   roles,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", auth.RouterKey, "accounts/proposed")

	response, code := SendPostRequest(uri, body, signer.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func ApproveAddAccount(keyInfo KeyInfo, signer KeyInfo) (TxnResponse, int) {
	println("Approve Add Account for: ", keyInfo.Name)

	request := rest.BasicReq{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    signer.Address.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", auth.RouterKey, fmt.Sprintf("accounts/proposed/%v", keyInfo.Address.String()))

	response, code := SendPatchRequest(uri, body, signer.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func ProposeRevokeAccount(keyInfo KeyInfo, signer KeyInfo) (TxnResponse, int) {
	println("Propose Revoke Account for: ", keyInfo.Name)

	request := authRest.ProposeRevokeAccountRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    signer.Address.String(),
		},
		Address: keyInfo.Address,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", auth.RouterKey, "accounts/proposed/revoked")

	response, code := SendPostRequest(uri, body, signer.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func ApproveRevokeAccount(keyInfo KeyInfo, signer KeyInfo) (TxnResponse, int) {
	println("Approve Revoke Account for: ", keyInfo.Name)

	request := rest.BasicReq{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    signer.Address.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", auth.RouterKey, fmt.Sprintf("accounts/proposed/revoked/%v", keyInfo.Address.String()))

	response, code := SendPatchRequest(uri, body, signer.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func GetAccount(address sdk.AccAddress) (AccountInfo, int) {
	println("Get Account for: ", address)

	uri := fmt.Sprintf("%s/accounts/%s", auth.RouterKey, address.String())
	response, code := SendGetRequest(uri)

	var result AccountInfo

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetAccounts() (AccountHeadersResult, int) {
	println("Get Accounts")

	uri := fmt.Sprintf("%s/accounts", auth.RouterKey)
	response, code := SendGetRequest(uri)

	var result AccountHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetProposedAccounts() (ProposedAccountHeadersResult, int) {
	println("Get Proposed Accounts")

	uri := fmt.Sprintf("%s/accounts/proposed", auth.RouterKey)
	response, code := SendGetRequest(uri)

	var result ProposedAccountHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetProposedAccountsToRevoke() (ProposedAccountToRevokeHeadersResult, int) {
	println("Get Proposed Accounts to Revoke")

	uri := fmt.Sprintf("%s/accounts/proposed/revoked", auth.RouterKey)
	response, code := SendGetRequest(uri)

	var result ProposedAccountToRevokeHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func CreateNewAccount(roles auth.AccountRoles) KeyInfo {
	name := RandString()
	println("Register new account on the ledger: ", name)

	jackKeyInfo, _ := GetKeyInfo(constants.JackAccount)
	aliceKeyInfo, _ := GetKeyInfo(constants.AliceAccount)

	keyInfo, _ := CreateKey(name)

	ProposeAddAccount(keyInfo, jackKeyInfo, roles)
	ApproveAddAccount(keyInfo, aliceKeyInfo)

	return keyInfo
}

func SignAndBroadcastMessage(sender KeyInfo, message sdk.Msg) (TxnResponse, int) {
	txn := types.StdTx{
		Msgs: []sdk.Msg{message},
		Fee:  types.StdFee{Gas: 2000000},
	}
	signResponse, _ := SignMessage(sender, txn)

	return BroadcastMessage(signResponse)
}

func SignAndBroadcastTransaction(sender KeyInfo, txn types.StdTx) (TxnResponse, int) {
	signResponse, _ := SignMessage(sender, txn)

	return BroadcastMessage(signResponse)
}

func SignMessage(sender KeyInfo, txn types.StdTx) (json.RawMessage, int) {
	println("Sign prepared transaction")

	stdSigMsg := extRest.SignMessageRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    sender.Address.String(),
		},
		Txn: extRest.Txn{
			Value: txn,
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s", "tx", "sign")
	response, code := SendPostRequest(uri, body, sender.Name, constants.Passphrase)

	if code != http.StatusOK {
		return json.RawMessage{}, code
	}

	return removeResponseWrapper(response), code
}

func BroadcastMessage(message interface{}) (TxnResponse, int) {
	println("Broadcast Message")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	response, code := SendPostRequest(uri, body, "", "")

	return parseWriteTxnResponse(response, code)
}

func AddModelInfo(model modelinfo.MsgAddModelInfo, sender KeyInfo) (TxnResponse, int) {
	println("Add Model Info")

	response, code := SendAddModelInfoRequest(model, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareAddModelInfoTransaction(model modelinfo.MsgAddModelInfo) (types.StdTx, int) {
	println("Prepare Add Model Info Transaction")

	response, code := SendAddModelInfoRequest(model, "")

	return parseStdTxn(response, code)
}

func SendAddModelInfoRequest(model modelinfo.MsgAddModelInfo, account string) ([]byte, int) {
	request := modelinfoRest.AddModelInfoRequest{
		Model: model.Model,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    model.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")

	return SendPostRequest(uri, body, account, constants.Passphrase)
}

func UpdateModelInfo(model modelinfo.MsgUpdateModelInfo, sender KeyInfo) (TxnResponse, int) {
	println("Update Model Info")

	response, code := SendUpdateModelInfoRequest(model, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareUpdateModelInfoTransaction(model modelinfo.MsgUpdateModelInfo) (types.StdTx, int) {
	println("Prepare Update Model Info Transaction")

	response, code := SendUpdateModelInfoRequest(model, "")

	return parseStdTxn(response, code)
}

func SendUpdateModelInfoRequest(model modelinfo.MsgUpdateModelInfo, account string) ([]byte, int) {

	request := modelinfoRest.UpdateModelInfoRequest{
		Model: model.Model,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    model.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")

	return SendPutRequest(uri, body, account, constants.Passphrase)
}

func GetModelInfo(vid uint16, pid uint16) (modelinfo.ModelInfo, int) {
	println(fmt.Sprintf("Get Model Info with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", modelinfo.RouterKey, "models", vid, pid)
	response, code := SendGetRequest(uri)

	var result modelinfo.ModelInfo

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetModelInfos() (ModelInfoHeadersResult, int) {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")
	response, code := SendGetRequest(uri)

	var result ModelInfoHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetVendors() (VendorItemHeadersResult, int) {
	println("Get the list of vendors")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "vendors")
	response, code := SendGetRequest(uri)

	var result VendorItemHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetVendorModels(vid uint16) (modelinfo.VendorProducts, int) {
	println("Get the list of models for VID:", vid)

	uri := fmt.Sprintf("%s/%s/%v", modelinfo.RouterKey, "models", vid)
	response, code := SendGetRequest(uri)

	var result modelinfo.VendorProducts

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func PublishTestingResult(testingResult compliancetest.MsgAddTestingResult, sender KeyInfo) (TxnResponse, int) {
	println("Publish Testing Result")

	response, code := SendTestingResultRequest(testingResult, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareTestingResultTransaction(testingResult compliancetest.MsgAddTestingResult) (types.StdTx, int) {
	println("Prepare Testing Result Transaction")

	response, code := SendTestingResultRequest(testingResult, "")

	return parseStdTxn(response, code)
}

func SendTestingResultRequest(testingResult compliancetest.MsgAddTestingResult, name string) ([]byte, int) {
	request := compliancetestRest.TestingResultRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    testingResult.Signer.String(),
		},
		VID:        testingResult.VID,
		PID:        testingResult.PID,
		TestResult: testingResult.TestResult,
		TestDate:   testingResult.TestDate,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", compliancetest.RouterKey, "testresults")

	return SendPostRequest(uri, body, name, constants.Passphrase)
}

func GetTestingResult(vid uint16, pid uint16) (compliancetest.TestingResults, int) {
	println(fmt.Sprintf("Get Testing Result for Model with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", compliancetest.RouterKey, "testresults", vid, pid)
	response, code := SendGetRequest(uri)

	var result compliancetest.TestingResults

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func PublishCertifiedModel(certifyModel compliance.MsgCertifyModel, sender KeyInfo) (TxnResponse, int) {
	println("Publish Certified Model")

	response, code := SendCertifiedModelRequest(certifyModel, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareCertifiedModelTransaction(certifyModel compliance.MsgCertifyModel) (types.StdTx, int) {
	println("Prepare Certified Model Transaction")

	response, code := SendCertifiedModelRequest(certifyModel, "")

	return parseStdTxn(response, code)
}

func SendCertifiedModelRequest(certifyModel compliance.MsgCertifyModel, name string) ([]byte, int) {
	request := complianceRest.CertifyModelRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    certifyModel.Signer.String(),
		},
		CertificationDate: certifyModel.CertificationDate,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", compliance.RouterKey, "certified",
		certifyModel.VID, certifyModel.PID, certifyModel.CertificationType)

	return SendPutRequest(uri, body, name, constants.Passphrase)
}

func PublishRevokedModel(revokeModel compliance.MsgRevokeModel, sender KeyInfo) (TxnResponse, int) {
	println("Publish Revoked Model")

	response, code := SendRevokedModelRequest(revokeModel, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareRevokedModelTransaction(revokeModel compliance.MsgRevokeModel) (types.StdTx, int) {
	println("Prepare Revoked Model Transaction")

	response, code := SendRevokedModelRequest(revokeModel, "")

	return parseStdTxn(response, code)
}

func SendRevokedModelRequest(revokeModel compliance.MsgRevokeModel, name string) ([]byte, int) {
	request := complianceRest.RevokeModelRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    revokeModel.Signer.String(),
		},
		RevocationDate: revokeModel.RevocationDate,
		Reason:         revokeModel.Reason,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", compliance.RouterKey, "revoked",
		revokeModel.VID, revokeModel.PID, revokeModel.CertificationType)

	return SendPutRequest(uri, body, name, constants.Passphrase)
}

func GetComplianceInfo(vid uint16, pid uint16,
	certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	println(fmt.Sprintf("Get Compliance Info for Model with VID:%v PID:%v", vid, pid))

	return getComplianceInfo(vid, pid, certificationType)
}

func GetCertifiedModel(vid uint16, pid uint16,
	certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Certified", vid, pid))

	return getComplianceInfoInState(vid, pid, certificationType, "certified")
}

func GetRevokedModel(vid uint16, pid uint16,
	certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Revoked", vid, pid))

	return getComplianceInfoInState(vid, pid, certificationType, "revoked")
}

func getComplianceInfo(vid uint16, pid uint16,
	certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	uri := fmt.Sprintf("%s/%v/%v/%v", compliance.RouterKey, vid, pid, certificationType)
	response, code := SendGetRequest(uri)

	var result compliance.ComplianceInfo

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func getComplianceInfoInState(vid uint16, pid uint16,
	certificationType compliance.CertificationType, state string) (compliance.ComplianceInfoInState, int) {
	uri := fmt.Sprintf("%s/%v/%v/%v/%v", compliance.RouterKey, state, vid, pid, certificationType)

	response, code := SendGetRequest(uri)

	var result compliance.ComplianceInfoInState

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetComplianceInfos() (ComplianceInfosHeadersResult, int) {
	println("Get all compliance info records")

	return GetAllComplianceInfos("")
}

func GetAllCertifiedModels() (ComplianceInfosHeadersResult, int) {
	println("Get all certified models")

	return GetAllComplianceInfos("certified")
}

func GetAllRevokedModels() (ComplianceInfosHeadersResult, int) {
	println("Get all revoked models")

	return GetAllComplianceInfos("revoked")
}

func GetAllComplianceInfos(state string) (ComplianceInfosHeadersResult, int) {
	var uri string

	if len(state) > 0 {
		uri = fmt.Sprintf("%s/%v", compliance.RouterKey, state)
	} else {
		uri = compliance.RouterKey
	}

	response, code := SendGetRequest(uri)

	var result ComplianceInfosHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func ProposeAddX509RootCert(proposeAddX509RootCert pki.MsgProposeAddX509RootCert,
	account string, passphrase string) (TxnResponse, int) {
	println("Propose X509 Root Certificate")

	response, code := SendProposeAddX509RootCertRequest(proposeAddX509RootCert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareProposeAddX509RootCertTransaction(proposeAddX509RootCert pki.MsgProposeAddX509RootCert) (types.StdTx, int) {
	println("Prepare Propose X509 Root Certificate Transaction")

	response, code := SendProposeAddX509RootCertRequest(proposeAddX509RootCert, "", "")

	return parseStdTxn(response, code)
}

func SendProposeAddX509RootCertRequest(proposeAddX509RootCert pki.MsgProposeAddX509RootCert,
	account string, passphrase string) ([]byte, int) {
	request := pkiRest.ProposeAddRootCertificateRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    proposeAddX509RootCert.Signer.String(),
		},
		Cert: proposeAddX509RootCert.Cert,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", pki.RouterKey, "certs/proposed/root")

	return SendPostRequest(uri, body, account, passphrase)
}

func ApproveAddX509RootCert(approveAddX509RootCert pki.MsgApproveAddX509RootCert,
	account string, passphrase string) (TxnResponse, int) {
	println(fmt.Sprintf("Approve X509 Root Cert with subject=%s and subjectKeyID=%s",
		approveAddX509RootCert.Subject, approveAddX509RootCert.SubjectKeyID))

	response, code := SendApproveAddX509RootCertRequest(approveAddX509RootCert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareApproveAddX509RootCertTransaction(
	approveAddX509RootCert pki.MsgApproveAddX509RootCert) (types.StdTx, int) {
	println("Prepare Approve X509 Root Certificate Transaction")

	response, code := SendApproveAddX509RootCertRequest(approveAddX509RootCert, "", "")

	return parseStdTxn(response, code)
}

func SendApproveAddX509RootCertRequest(approveAddX509RootCert pki.MsgApproveAddX509RootCert,
	account string, passphrase string) ([]byte, int) {
	request := rest.BasicReq{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    approveAddX509RootCert.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)
	uri := fmt.Sprintf("%s/%s", pki.RouterKey, fmt.Sprintf("certs/proposed/root/%s/%s",
		approveAddX509RootCert.Subject, approveAddX509RootCert.SubjectKeyID))

	return SendPatchRequest(uri, body, account, passphrase)
}

func AddX509Cert(addX509Cert pki.MsgAddX509Cert, account string, passphrase string) (TxnResponse, int) {
	println("Add X509 Certificate")

	response, code := SendAddX509CertRequest(addX509Cert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareAddX509CertTransaction(addX509Cert pki.MsgAddX509Cert) (types.StdTx, int) {
	println("Prepare Add X509 Certificate Transaction")

	response, code := SendAddX509CertRequest(addX509Cert, "", "")

	return parseStdTxn(response, code)
}

func SendAddX509CertRequest(addX509Cert pki.MsgAddX509Cert, account string, passphrase string) ([]byte, int) {
	request := pkiRest.AddCertificateRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    addX509Cert.Signer.String(),
		},
		Cert: addX509Cert.Cert,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", pki.RouterKey, "certs")

	return SendPostRequest(uri, body, account, passphrase)
}

func ProposeRevokeX509RootCert(proposeRevokeX509RootCert pki.MsgProposeRevokeX509RootCert,
	account string, passphrase string) (TxnResponse, int) {
	println(fmt.Sprintf("Propose to Revoke X509 Root Cert with subject=%s and subjectKeyID=%s",
		proposeRevokeX509RootCert.Subject, proposeRevokeX509RootCert.SubjectKeyID))

	response, code := SendProposeRevokeX509RootCertRequest(proposeRevokeX509RootCert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareProposeRevokeX509RootCertTransaction(
	proposeRevokeX509RootCert pki.MsgProposeRevokeX509RootCert) (types.StdTx, int) {
	println("Prepare Propose to Revoke X509 Root Certificate Transaction")

	response, code := SendProposeRevokeX509RootCertRequest(proposeRevokeX509RootCert, "", "")

	return parseStdTxn(response, code)
}

func SendProposeRevokeX509RootCertRequest(proposeRevokeX509RootCert pki.MsgProposeRevokeX509RootCert,
	account string, passphrase string) ([]byte, int) {
	request := pkiRest.ProposeRevokeRootCertificateRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    proposeRevokeX509RootCert.Signer.String(),
		},
		Subject:      proposeRevokeX509RootCert.Subject,
		SubjectKeyID: proposeRevokeX509RootCert.SubjectKeyID,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/certs/proposed/revoked/root", pki.RouterKey)

	return SendPostRequest(uri, body, account, passphrase)
}

func ApproveRevokeX509RootCert(approveRevokeX509RootCert pki.MsgApproveRevokeX509RootCert,
	account string, passphrase string) (TxnResponse, int) {
	println(fmt.Sprintf("Approve to Revoke X509 Root Cert with subject=%s and subjectKeyID=%s",
		approveRevokeX509RootCert.Subject, approveRevokeX509RootCert.SubjectKeyID))

	response, code := SendApproveRevokeX509RootCertRequest(approveRevokeX509RootCert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareApproveRevokeX509RootCertTransaction(
	approveRevokeX509RootCert pki.MsgApproveRevokeX509RootCert) (types.StdTx, int) {
	println("Prepare Approve to Revoke X509 Root Certificate Transaction")

	response, code := SendApproveRevokeX509RootCertRequest(approveRevokeX509RootCert, "", "")

	return parseStdTxn(response, code)
}

func SendApproveRevokeX509RootCertRequest(approveRevokeX509RootCert pki.MsgApproveRevokeX509RootCert,
	account string, passphrase string) ([]byte, int) {
	request := rest.BasicReq{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    approveRevokeX509RootCert.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)
	uri := fmt.Sprintf("%s/certs/proposed/revoked/root/%s/%s",
		pki.RouterKey, approveRevokeX509RootCert.Subject, approveRevokeX509RootCert.SubjectKeyID)

	return SendPatchRequest(uri, body, account, passphrase)
}

func RevokeX509Cert(revokeX509Cert pki.MsgRevokeX509Cert, account string, passphrase string) (TxnResponse, int) {
	println(fmt.Sprintf("Revoke X509 Certificate with subject=%s and subjectKeyID=%s",
		revokeX509Cert.Subject, revokeX509Cert.SubjectKeyID))

	response, code := SendRevokeX509CertRequest(revokeX509Cert, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func PrepareRevokeX509CertTransaction(revokeX509Cert pki.MsgRevokeX509Cert) (types.StdTx, int) {
	println("Prepare Revoke X509 Certificate Transaction")

	response, code := SendRevokeX509CertRequest(revokeX509Cert, "", "")

	return parseStdTxn(response, code)
}

func SendRevokeX509CertRequest(revokeX509Cert pki.MsgRevokeX509Cert, account string, passphrase string) ([]byte, int) {
	request := rest.BasicReq{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    revokeX509Cert.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/certs/%s/%s", pki.RouterKey, revokeX509Cert.Subject, revokeX509Cert.SubjectKeyID)

	return SendDeleteRequest(uri, body, account, passphrase)
}

func GetAllX509RootCerts() (CertificatesHeadersResult, int) {
	return getCertificates(fmt.Sprintf("%s/certs/root", pki.RouterKey))
}

func GetAllSubjectX509Certs(subject string) (CertificatesHeadersResult, int) {
	return getCertificates(fmt.Sprintf("%s/certs/%s", pki.RouterKey, subject))
}

func GetAllX509Certs() (CertificatesHeadersResult, int) {
	return getCertificates(fmt.Sprintf("%s/certs", pki.RouterKey))
}

func GetAllProposedX509RootCerts() (ProposedCertificatesHeadersResult, int) {
	return getProposedCertificates(fmt.Sprintf("%s/certs/proposed/root", pki.RouterKey))
}

func GetProposedX509RootCert(subject string, subjectKeyID string) (pki.ProposedCertificate, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/proposed/root/%s/%s", pki.RouterKey, subject, subjectKeyID))

	var result pki.ProposedCertificate

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetX509Cert(subject string, subjectKeyID string) (pki.Certificate, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/%s/%s", pki.RouterKey, subject, subjectKeyID))

	var result pki.Certificates

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	if len(result.Items) > 1 {
		return pki.Certificate{}, http.StatusInternalServerError
	}

	return result.Items[0], code
}

func GetX509CertChain(subject string, subjectKeyID string) (pki.Certificates, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/chain/%s/%s", pki.RouterKey, subject, subjectKeyID))

	var result pki.Certificates

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetAllRevokedX509RootCerts() (CertificatesHeadersResult, int) {
	return getCertificates(fmt.Sprintf("%s/certs/revoked/root", pki.RouterKey))
}

func GetAllRevokedX509Certs() (CertificatesHeadersResult, int) {
	return getCertificates(fmt.Sprintf("%s/certs/revoked", pki.RouterKey))
}

func GetAllProposedX509RootCertsToRevoke() (ProposedCertificateRevocationsHeadersResult, int) {
	return getProposedCertificateRevocations(fmt.Sprintf("%s/certs/proposed/revoked/root", pki.RouterKey))
}

func GetProposedX509RootCertToRevoke(subject string, subjectKeyID string) (pki.ProposedCertificateRevocation, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/proposed/revoked/root/%s/%s",
		pki.RouterKey, subject, subjectKeyID))

	var result pki.ProposedCertificateRevocation

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetRevokedX509Cert(subject string, subjectKeyID string) (pki.Certificate, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/revoked/%s/%s", pki.RouterKey, subject, subjectKeyID))

	var result pki.Certificates

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	if len(result.Items) > 1 {
		return pki.Certificate{}, http.StatusInternalServerError
	}

	return result.Items[0], code
}

func getCertificates(uri string) (CertificatesHeadersResult, int) {
	response, code := SendGetRequest(uri)

	var result CertificatesHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func getProposedCertificates(uri string) (ProposedCertificatesHeadersResult, int) {
	response, code := SendGetRequest(uri)

	var result ProposedCertificatesHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func getProposedCertificateRevocations(uri string) (ProposedCertificateRevocationsHeadersResult, int) {
	response, code := SendGetRequest(uri)

	var result ProposedCertificateRevocationsHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func NewMsgAddModelInfo(owner sdk.AccAddress) modelinfo.MsgAddModelInfo {
	model := modelinfo.Model{

		VID:                                      common.RandUint16(),
		PID:                                      common.RandUint16(),
		CID:                                      constants.CID,
		Name:                                     RandString(),
		Description:                              RandString(),
		SKU:                                      RandString(),
		SoftwareVersion:                          constants.SoftwareVersion,
		SoftwareVersionString:                    constants.SoftwareVersionString,
		HardwareVersion:                          constants.HardwareVersion,
		HardwareVersionString:                    constants.HardwareVersionString,
		CDVersionNumber:                          constants.CDVersionNumber,
		FirmwareDigests:                          constants.FirmwareDigests,
		Revoked:                                  constants.Revoked,
		OtaURL:                                   constants.OtaURL,
		OtaChecksum:                              constants.OtaChecksum,
		OtaChecksumType:                          constants.OtaChecksumType,
		OtaBlob:                                  RandString(),
		CommissioningCustomFlow:                  constants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               constants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        constants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: constants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      constants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: constants.CommissioningModeSecondaryStepsInstruction,
		ReleaseNotesUrl: constants.ReleaseNotesUrl,
		UserManualUrl:   constants.UserManualUrl,
		SupportUrl:      constants.SupportUrl,
		ProductURL:      constants.ProductURL,
		ChipBlob:        constants.ChipBlob,
		VendorBlob:      constants.VendorBlob,
	}

	return modelinfo.NewMsgAddModelInfo(
		model,
		owner,
	)
}

func NewMsgUpdateModelInfo(vid uint16, pid uint16, owner sdk.AccAddress) modelinfo.MsgUpdateModelInfo {

	model := modelinfo.Model{

		VID:                        vid,
		PID:                        pid,
		CID:                        constants.CID + 1,
		Description:                RandString(),
		CDVersionNumber:            constants.CDVersionNumber + 1,
		Revoked:                    constants.Revoked,
		OtaURL:                     constants.OtaURL + "/new",
		OtaChecksum:                constants.OtaChecksum,
		OtaChecksumType:            constants.OtaChecksumType,
		OtaBlob:                    RandString(),
		CommissioningCustomFlowUrl: constants.CommissioningCustomFlowUrl + "/new",
		ReleaseNotesUrl:            constants.ReleaseNotesUrl + "/new",
		UserManualUrl:              constants.UserManualUrl + "/new",
		SupportUrl:                 constants.SupportUrl + "/new",
		ProductURL:                 constants.ProductURL + "/new",
		ChipBlob:                   constants.ChipBlob + "-new",
		VendorBlob:                 constants.VendorBlob + "-new",
	}

	return modelinfo.NewMsgUpdateModelInfo(
		model,
		owner,
	)
}

func NewMsgAddTestingResult(vid uint16, pid uint16, owner sdk.AccAddress) compliancetest.MsgAddTestingResult {
	return compliancetest.NewMsgAddTestingResult(
		vid,
		pid,
		RandString(),
		time.Now().UTC(),
		owner,
	)
}

func removeResponseWrapper(response []byte) json.RawMessage {
	var responseWrapper ResponseWrapper
	_ = json.Unmarshal(response, &responseWrapper)

	return responseWrapper.Result
}

func parseWriteTxnResponse(response []byte, code int) (TxnResponse, int) {
	if code != http.StatusOK {
		return TxnResponse{}, code
	}

	var result TxnResponse
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, code
}

func parseStdTxn(response []byte, code int) (types.StdTx, int) {
	if code != http.StatusOK {
		return types.StdTx{}, code
	}

	var txn types.StdTx
	_ = app.MakeCodec().UnmarshalJSON(response, &txn)

	return txn, code
}

func parseGetReqResponse(response []byte, entity interface{}, code int) {
	if code == http.StatusOK {
		_ = json.Unmarshal(response, entity)
	}
}

func InitStartData() (KeyInfo, KeyInfo, modelinfo.MsgAddModelInfo,
	ComplianceInfosHeadersResult, ComplianceInfosHeadersResult) {
	// Register new Vendor account
	vendor := CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Register new ZBCertificationCenter account
	zb := CreateNewAccount(auth.AccountRoles{auth.ZBCertificationCenter})

	// Publish model info
	modelInfo := NewMsgAddModelInfo(vendor.Address)
	_, _ = AddModelInfo(modelInfo, vendor)

	// Get all certified models
	inputCertifiedModels, _ := GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := GetAllRevokedModels()

	return vendor, zb, modelInfo, inputCertifiedModels, inputRevokedModels
}
