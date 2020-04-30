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
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	pkiRest "git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/client/rest"
	keyUtil "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"
	"net/http"
	"time"
)

func CreateKey(accountName string) (KeyInfo, int) {
	println("Create Key for: ", accountName)

	kb, _ := keyUtil.NewKeyBaseFromDir(app.DefaultCLIHome)

	entropySeed, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropySeed[:])
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

func CreateAccount(keyInfo KeyInfo, signer KeyInfo) (TxnResponse, int) {
	println("Create Account for: ", keyInfo.Name)

	msgAddAccount := authnext.NewMsgAddAccount(keyInfo.Address, keyInfo.PublicKey, signer.Address)
	return SignAndBroadcastMessage(signer, msgAddAccount)
}

func GetAccountInfo(address sdk.AccAddress) (AccountInfo, int) {
	println("Get Account Info")

	uri := fmt.Sprintf("%s/account/%s", authnext.RouterKey, address)
	response, code := SendGetRequest(uri)

	var result AccountInfo
	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func RegisterNewAccount() (KeyInfo, int) {
	name := RandString()
	println("Register new account on the ledger: ", name)

	jackKeyInfo, _ := GetKeyInfo(constants.JackAccount)
	testAccountKeyInfo, _ := CreateKey(name)
	_, code := CreateAccount(testAccountKeyInfo, jackKeyInfo)
	return testAccountKeyInfo, code
}

func SignAndBroadcastMessage(sender KeyInfo, message sdk.Msg) (TxnResponse, int) {
	senderAccountInfo, _ := GetAccountInfo(sender.Address) // Refresh account info
	signResponse, _ := SignMessage(sender.Name, senderAccountInfo, message)
	return BroadcastMessage(signResponse)
}

func PublishModelInfo(model modelinfo.MsgAddModelInfo, sender KeyInfo) (TxnResponse, int) {
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
	response, code := SendPostRequest(uri, body, sender.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func SignMessage(accountName string, accountInfo AccountInfo, message sdk.Msg) (json.RawMessage, int) {
	println("Sign Message")

	stdSigMsg := types.StdSignMsg{
		ChainID:       constants.ChainId,
		AccountNumber: ParseUint(accountInfo.AccountNumber),
		Sequence:      ParseUint(accountInfo.Sequence),
		Fee:           types.StdFee{Gas: 200000},
		Msgs:          []sdk.Msg{message},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s", "tx", "sign")
	response, code := SendPostRequest(uri, body, accountName, constants.Passphrase)
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

	request := compliancetestRest.TestingResultRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    testingResult.Signer.String(),
		},
		VID:        testingResult.VID,
		PID:        testingResult.PID,
		TestResult: testingResult.TestResult,
		TestDate:   testingResult.TestDate,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", compliancetest.RouterKey, "testresults")
	response, code := SendPostRequest(uri, body, sender.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func GetTestingResult(vid uint16, pid uint16) (compliancetest.TestingResults, int) {
	println(fmt.Sprintf("Get Testing Result for Model with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", compliancetest.RouterKey, "testresults", vid, pid)
	response, code := SendGetRequest(uri)

	var result compliancetest.TestingResults
	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
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

func PublishCertifiedModel(certifyModel compliance.MsgCertifyModel, sender KeyInfo) (TxnResponse, int) {
	println("Publish Certified Model")

	request := complianceRest.CertifyModelRequest{
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

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", compliance.RouterKey, "certified", certifyModel.VID, certifyModel.PID, certifyModel.CertificationType)
	response, code := SendPutRequest(uri, body, sender.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func PublishRevokedModel(revokeModel compliance.MsgRevokeModel, sender KeyInfo) (TxnResponse, int) {
	println("Publish Revoked Model")

	request := complianceRest.RevokeModelRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    revokeModel.Signer.String(),
		},
		VID:               revokeModel.VID,
		PID:               revokeModel.PID,
		RevocationDate:    revokeModel.RevocationDate,
		Reason:            revokeModel.Reason,
		CertificationType: revokeModel.CertificationType,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s/%v/%v/%v", compliance.RouterKey, "revoked", revokeModel.VID, revokeModel.PID, revokeModel.CertificationType)
	response, code := SendPutRequest(uri, body, sender.Name, constants.Passphrase)

	return parseWriteTxnResponse(response, code)
}

func GetComplianceInfo(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	println(fmt.Sprintf("Get Compliance Info for Model with VID:%v PID:%v", vid, pid))
	return getComplianceInfo(vid, pid, certificationType)
}

func GetCertifiedModel(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Certified", vid, pid))
	return getComplianceInfoInState(vid, pid, certificationType, "certified")
}

func GetRevokedModel(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, int) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Revoked", vid, pid))
	return getComplianceInfoInState(vid, pid, certificationType, "revoked")
}

func getComplianceInfo(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfo, int) {
	uri := fmt.Sprintf("%s/%v/%v/%v", compliance.RouterKey, vid, pid, certificationType)
	response, code := SendGetRequest(uri)

	var result compliance.ComplianceInfo
	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func getComplianceInfoInState(vid uint16, pid uint16, certificationType compliance.CertificationType, state string) (compliance.ComplianceInfoInState, int) {
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
		uri = fmt.Sprintf("%s", compliance.RouterKey)
	}

	response, code := SendGetRequest(uri)

	var result ComplianceInfosHeadersResult
	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func ProposeAddX509RootCert(proposeAddX509RootCert pki.MsgProposeAddX509RootCert, account string, passphrase string) (TxnResponse, int) {
	println("Propose X509 Root Certificate")

	request := pkiRest.AddCertificateRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    proposeAddX509RootCert.Signer.String(),
		},
		Certificate: proposeAddX509RootCert.Cert,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", pki.RouterKey, "certs/proposed/root")

	response, code := SendPostRequest(uri, body, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func ApproveAddX509RootCert(msgApproveAddX509RootCert pki.MsgApproveAddX509RootCert, account string, passphrase string) (TxnResponse, int) {
	println(fmt.Sprintf("Approve X509 Root Cert with subject=%s and subjectKeyId=%s", msgApproveAddX509RootCert.Subject, msgApproveAddX509RootCert.SubjectKeyId))

	request := pkiRest.ApproveCertificateRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    msgApproveAddX509RootCert.Signer.String(),
		},
		Subject:      msgApproveAddX509RootCert.Subject,
		SubjectKeyId: msgApproveAddX509RootCert.SubjectKeyId,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", pki.RouterKey, fmt.Sprintf("certs/proposed/root/%s/%s", msgApproveAddX509RootCert.Subject, msgApproveAddX509RootCert.SubjectKeyId))

	response, code := SendPatchRequest(uri, body, account, passphrase)

	return parseWriteTxnResponse(response, code)
}

func AddX509Cert(addX509Cert pki.MsgAddX509Cert, account string, passphrase string) (TxnResponse, int) {
	println("Add X509 Certificate")

	request := pkiRest.AddCertificateRequest{
		BaseReq: rest.BaseReq{
			ChainID: constants.ChainId,
			From:    addX509Cert.Signer.String(),
		},
		Certificate: addX509Cert.Cert,
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), request)

	uri := fmt.Sprintf("%s/%s", pki.RouterKey, "certs")

	response, code := SendPostRequest(uri, body, account, passphrase)

	return parseWriteTxnResponse(response, code)
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

func GetProposedX509RootCert(subject string, subjectKeyId string) (pki.ProposedCertificate, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/proposed/root/%s/%s", pki.RouterKey, subject, subjectKeyId))

	var result pki.ProposedCertificate
	parseGetReqResponse(removeResponseWrapper(response), &result, code)

	return result, code
}

func GetX509Cert(subject string, subjectKeyId string) (pki.Certificate, int) {
	response, code := SendGetRequest(fmt.Sprintf("%s/certs/%s/%s", pki.RouterKey, subject, subjectKeyId))

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

func NewMsgAddModelInfo(owner sdk.AccAddress) modelinfo.MsgAddModelInfo {
	return modelinfo.NewMsgAddModelInfo(
		uint16(RandInt()),
		uint16(RandInt()),
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

func parseGetReqResponse(response []byte, entity interface{}, code int) {
	if code == http.StatusOK {
		json.Unmarshal(response, entity)
	}
}
