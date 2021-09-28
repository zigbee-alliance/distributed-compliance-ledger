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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	modelRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/client/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
	modelVersionRest "github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/client/rest"
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

func ProposeAddAccount(keyInfo KeyInfo, signer KeyInfo, roles auth.AccountRoles, vendorId uint16) (TxnResponse, int) {
	println("Propose Add Account for: ", keyInfo.Name)

	request := authRest.ProposeAddAccountRequest{
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    signer.Address.String(),
		},
		Address:  keyInfo.Address,
		Pubkey:   keyInfo.PublicKey,
		Roles:    roles,
		VendorId: vendorId,
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

func CreateNewAccount(roles auth.AccountRoles, vendorId uint16) KeyInfo {
	name := RandString()
	println("Register new account on the ledger: ", name)

	jackKeyInfo, _ := GetKeyInfo(constants.JackAccount)
	aliceKeyInfo, _ := GetKeyInfo(constants.AliceAccount)

	keyInfo, _ := CreateKey(name)

	ProposeAddAccount(keyInfo, jackKeyInfo, roles, vendorId)
	ApproveAddAccount(keyInfo, aliceKeyInfo)

	return keyInfo
}

func SignAndBroadcastMessage(sender KeyInfo, message sdk.Msg) (TxnResponse, int) {
	txn := types.StdTx{
		Msgs: []sdk.Msg{message},
		Fee:  types.StdFee{Gas: 2000000},
	}
	fmt.Printf("txn: %v\n", txn)
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

func AddModel(model model.MsgAddModel, sender KeyInfo) (TxnResponse, int) {
	println("Add Model Info")

	response, code := SendAddModelRequest(model, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func AddModelVersion(modelVersion modelversion.MsgAddModelVersion, sender KeyInfo) (TxnResponse, int) {
	println("Add Model Version")

	response, code := SendAddModelVersionRequest(modelVersion, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareAddModelTransaction(model model.MsgAddModel) (types.StdTx, int) {
	println("Prepare Add Model Info Transaction")

	response, code := SendAddModelRequest(model, "")

	return parseStdTxn(response, code)
}

func PrepareAddModelVersionTransaction(modelVersion modelversion.MsgAddModelVersion) (types.StdTx, int) {
	println("Prepare Add Model Version Transaction")

	response, code := SendAddModelVersionRequest(modelVersion, "")

	return parseStdTxn(response, code)
}

func SendAddModelRequest(msgAddModel model.MsgAddModel, account string) ([]byte, int) {
	request := modelRest.AddModelRequest{
		Model: msgAddModel.Model,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    msgAddModel.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", model.RouterKey, "models")

	return SendPostRequest(uri, body, account, constants.Passphrase)
}

func SendAddModelVersionRequest(msgAddModelVersion modelversion.MsgAddModelVersion, account string) ([]byte, int) {
	request := modelVersionRest.AddModelVersionRequest{
		ModelVersion: msgAddModelVersion.ModelVersion,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    msgAddModelVersion.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", modelversion.RouterKey, "version")

	return SendPostRequest(uri, body, account, constants.Passphrase)
}

func UpdateModel(model model.MsgUpdateModel, sender KeyInfo) (TxnResponse, int) {
	println("Update Model Info")

	response, code := SendUpdateModelRequest(model, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func UpdateModelVersion(modelVersion modelversion.MsgUpdateModelVersion, sender KeyInfo) (TxnResponse, int) {
	println("Update Model Version")

	response, code := SendUpdateModelVersionRequest(modelVersion, sender.Name)

	return parseWriteTxnResponse(response, code)
}

func PrepareUpdateModelTransaction(model model.MsgUpdateModel) (types.StdTx, int) {
	println("Prepare Update Model Info Transaction")

	response, code := SendUpdateModelRequest(model, "")

	return parseStdTxn(response, code)
}

func PrepareUpdateModelVersionTransaction(modelVersion modelversion.MsgUpdateModelVersion) (types.StdTx, int) {
	println("Prepare Update Model Version Transaction")

	response, code := SendUpdateModelVersionRequest(modelVersion, "")

	return parseStdTxn(response, code)
}

func SendUpdateModelRequest(msgUpdateModel model.MsgUpdateModel, account string) ([]byte, int) {
	request := modelRest.UpdateModelRequest{
		Model: msgUpdateModel.Model,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    msgUpdateModel.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", model.RouterKey, "models")

	return SendPutRequest(uri, body, account, constants.Passphrase)
}

func SendUpdateModelVersionRequest(msgUpdateModelVersion modelversion.MsgUpdateModelVersion, account string) ([]byte, int) {
	request := modelVersionRest.UpdateModelVersionRequest{
		ModelVersion: msgUpdateModelVersion.ModelVersion,
		BaseReq: restTypes.BaseReq{
			ChainID: constants.ChainID,
			From:    msgUpdateModelVersion.Signer.String(),
		},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", modelversion.RouterKey, "version")

	return SendPutRequest(uri, body, account, constants.Passphrase)
}

func GetModel(vid uint16, pid uint16) (model.Model, int) {
	println(fmt.Sprintf("Get Model Info with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", model.RouterKey, "models", vid, pid)
	response, code := SendGetRequest(uri)

	var result model.Model

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetModelVersion(vid uint16, pid uint16, softwareVersion uint32) (modelversion.ModelVersion, int) {
	println(fmt.Sprintf("Get Model Version with VID:%v PID:%v SV:%v", vid, pid, softwareVersion))

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", modelversion.RouterKey, "version", vid, pid, softwareVersion)
	response, code := SendGetRequest(uri)

	var result modelversion.ModelVersion

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetModels() (ModelHeadersResult, int) {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", model.RouterKey, "models")
	response, code := SendGetRequest(uri)

	var result ModelHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetVendors() (VendorItemHeadersResult, int) {
	println("Get the list of vendors")

	uri := fmt.Sprintf("%s/%s", model.RouterKey, "vendors")
	response, code := SendGetRequest(uri)

	var result VendorItemHeadersResult

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetVendorModels(vid uint16) (model.VendorProducts, int) {
	println("Get the list of models for VID:", vid)

	uri := fmt.Sprintf("%s/%s/%v", model.RouterKey, "models", vid)
	response, code := SendGetRequest(uri)

	var result model.VendorProducts

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
		VID:                   testingResult.VID,
		PID:                   testingResult.PID,
		SoftwareVersion:       testingResult.SoftwareVersion,
		SoftwareVersionString: testingResult.SoftwareVersionString,
		TestResult:            testingResult.TestResult,
		TestDate:              testingResult.TestDate,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", compliancetest.RouterKey, "testresults")

	return SendPostRequest(uri, body, name, constants.Passphrase)
}

func GetTestingResult(vid uint16, pid uint16, softwareVersion uint32) (compliancetest.TestingResults, int) {
	println(fmt.Sprintf("Get Testing Result for Model with VID:%v PID:%v SV:%v", vid, pid, softwareVersion))

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", compliancetest.RouterKey, "testresults", vid, pid, softwareVersion)
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

	uri := fmt.Sprintf("%s/%v/%v/%v/%v/%v", compliance.RouterKey, compliance.Certified,
		certifyModel.VID, certifyModel.PID, certifyModel.SoftwareVersion, certifyModel.CertificationType)

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

	uri := fmt.Sprintf("%s/%v/%v/%v/%v/%v", compliance.RouterKey, compliance.Revoked,
		revokeModel.VID, revokeModel.PID, revokeModel.SoftwareVersion, revokeModel.CertificationType)

	return SendPutRequest(uri, body, name, constants.Passphrase)
}

func GetComplianceInfo(vid uint16, pid uint16, softwareVersion uint32,
	certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	println(fmt.Sprintf("Get Compliance Info for Model with VID:%v PID:%v SV:%v", vid, pid, softwareVersion))

	return getComplianceInfo(vid, pid, softwareVersion, certificationType)
}

func GetCertifiedModel(vid uint16, pid uint16, softwareVersion uint32,
	certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v SV:%v Certified", vid, pid, softwareVersion))

	return getComplianceInfoInState(vid, pid, softwareVersion, certificationType, compliance.Certified)
}

func GetRevokedModel(vid uint16, pid uint16, softwareVersion uint32,
	certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v SV:%v revoked", vid, pid, softwareVersion))

	return getComplianceInfoInState(vid, pid, softwareVersion, certificationType, compliance.Revoked)
}

func getComplianceInfo(vid uint16, pid uint16, softwareVersion uint32,
	certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	uri := fmt.Sprintf("%s/%v/%v/%v/%v", compliance.RouterKey, vid, pid, softwareVersion, certificationType)
	response, code := SendGetRequest(uri)

	var result compliance.ComplianceInfo

	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func getComplianceInfoInState(vid uint16, pid uint16, softwareVersion uint32,
	certificationType compliance.CertificationType, state compliance.ComplianceState) (compliance.ComplianceInfoInState, int) {
	uri := fmt.Sprintf("%s/%v/%v/%v/%v/%v", compliance.RouterKey, state, vid, pid, softwareVersion, certificationType)

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

func NewMsgAddModel(owner sdk.AccAddress, vid uint16) model.MsgAddModel {
	newModel := model.Model{

		VID:                                      vid,
		PID:                                      common.RandUint16(),
		DeviceTypeID:                             constants.DeviceTypeID,
		ProductName:                              RandString(),
		ProductLabel:                             RandString(),
		PartNumber:                               RandString(),
		CommissioningCustomFlow:                  constants.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               constants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        constants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: constants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      constants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: constants.CommissioningModeSecondaryStepsInstruction,
		UserManualURL: constants.UserManualURL,
		SupportURL:    constants.SupportURL,
		ProductURL:    constants.ProductURL,
	}

	return model.NewMsgAddModel(
		newModel,
		owner,
	)
}

func NewMsgUpdateModel(vid uint16, pid uint16, owner sdk.AccAddress) model.MsgUpdateModel {
	newModel := model.Model{
		VID:                        vid,
		PID:                        pid,
		DeviceTypeID:               constants.DeviceTypeID + 1,
		ProductLabel:               RandString(),
		CommissioningCustomFlowURL: constants.CommissioningCustomFlowURL + "/new",
		UserManualURL:              constants.UserManualURL + "/new",
		SupportURL:                 constants.SupportURL + "/new",
		ProductURL:                 constants.ProductURL + "/new",
	}

	return model.NewMsgUpdateModel(
		newModel,
		owner,
	)
}

func NewMsgAddModelVersion(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string, owner sdk.AccAddress) modelversion.MsgAddModelVersion {
	newModelVersion := modelversion.ModelVersion{

		VID:                          vid,
		PID:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		FirmwareDigests:              constants.FirmwareDigests,
		OtaURL:                       constants.OtaURL,
		OtaFileSize:                  constants.OtaFileSize,
		OtaChecksum:                  constants.OtaChecksum,
		OtaChecksumType:              constants.OtaChecksumType,
		CDVersionNumber:              constants.CDVersionNumber,
		MinApplicableSoftwareVersion: constants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: constants.MaxApplicableSoftwareVersion,
		ReleaseNotesURL:              constants.ReleaseNotesURL,
	}

	return modelversion.NewMsgAddModelVersion(
		newModelVersion,
		owner,
	)
}

func NewMsgUpdateModelVersion(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string, owner sdk.AccAddress) modelversion.MsgUpdateModelVersion {
	updateModelVersion := modelversion.ModelVersion{
		VID:             vid,
		PID:             pid,
		SoftwareVersion: softwareVersion,
		OtaURL:          constants.OtaURL + "/new",
		ReleaseNotesURL: constants.ReleaseNotesURL + "/new",
	}

	return modelversion.NewMsgUpdateModelVersion(
		updateModelVersion,
		owner,
	)
}

func NewMsgAddTestingResult(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string,
	owner sdk.AccAddress) compliancetest.MsgAddTestingResult {
	return compliancetest.NewMsgAddTestingResult(
		vid,
		pid,
		softwareVersion,
		softwareVersionString,
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

func InitStartData() (KeyInfo, KeyInfo, model.MsgAddModel, modelversion.MsgAddModelVersion,
	ComplianceInfosHeadersResult, ComplianceInfosHeadersResult) {
	// Register new Vendor account
	vendor := CreateNewAccount(auth.AccountRoles{auth.Vendor}, constants.VID)

	// Register new ZBCertificationCenter account
	zb := CreateNewAccount(auth.AccountRoles{auth.ZBCertificationCenter}, 0)

	// Publish model info
	model := NewMsgAddModel(vendor.Address, constants.VID)
	txnResponse, errCode := AddModel(model, vendor)
	fmt.Printf("%v, %v", txnResponse, errCode)

	// Publish model version
	modelVersion := NewMsgAddModelVersion(model.VID, model.PID, constants.SoftwareVersion, constants.SoftwareVersionString, vendor.Address)
	txnResponse, errCode = AddModelVersion(modelVersion, vendor)
	fmt.Printf("%v, %v", txnResponse, errCode)

	// Get all certified models
	inputCertifiedModels, _ := GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := GetAllRevokedModels()

	return vendor, zb, model, modelVersion, inputCertifiedModels, inputRevokedModels
}
