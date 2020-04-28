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
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"time"
)

func GetKeyInfo(accountName string) (KeyInfo, error) {
	println("Get User Key Info: ", accountName)

	uri := fmt.Sprintf("key/%s", accountName)
	response, err := SendGetRequest(uri)
	if err != nil {
		return KeyInfo{}, err
	}

	var keyInfo KeyInfo
	_ = json.Unmarshal(response, &keyInfo)

	return keyInfo, nil
}

func GetAccountInfo(address sdk.AccAddress) (AccountInfo, error) {
	println("Get Account Info")

	uri := fmt.Sprintf("%s/account/%s", authnext.RouterKey, address)
	response, err := SendGetRequest(uri)
	if err != nil {
		return AccountInfo{}, err
	}

	var result AccountInfo
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func SignAndBroadcastMessage(sender KeyInfo, message sdk.Msg) {
	senderAccountInfo, _ := GetAccountInfo(sender.Address) // Refresh account info
	signResponse, _ := SignMessage(sender.Name, senderAccountInfo, message)
	_, _ = BroadcastMessage(signResponse)
}

func PublishModelInfo(model modelinfo.MsgAddModelInfo) (json.RawMessage, error) {
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
	response, err := SendPostRequest(uri, body, constants.AccountName, constants.Passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func SignMessage(accountName string, accountInfo AccountInfo, message sdk.Msg) (interface{}, error) {
	println("Sign Message")

	stdSigMsg := types.StdSignMsg{
		ChainID:       constants.ChainId,
		AccountNumber: ParseUint(accountInfo.AccountNumber),
		Sequence:      ParseUint(accountInfo.Sequence),
		Fee:           types.StdFee{Gas: 200000},
		Msgs:          []sdk.Msg{message},
	}

	body, _ := codec.MarshalJSONIndent(app.MakeCodec(), stdSigMsg)

	uri := fmt.Sprintf("%s/%s?name=%s&passphrase=%s", "tx", "sign", accountName, constants.Passphrase)
	response, err := SendPostRequest(uri, body, "", "")
	if err != nil {
		return json.RawMessage{}, err
	}

	var result interface{}
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func BroadcastMessage(message interface{}) ([]byte, error) {
	println("Broadcast Message")

	body, _ := json.Marshal(message)

	uri := fmt.Sprintf("%s/%s", "tx", "broadcast")
	return SendPostRequest(uri, body, "", "")
}

func GetModelInfo(vid uint16, pid uint16) (modelinfo.ModelInfo, error) {
	println(fmt.Sprintf("Get Model Info with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", modelinfo.RouterKey, "models", vid, pid)
	response, err := SendGetRequest(uri)
	if err != nil {
		return modelinfo.ModelInfo{}, err
	}

	var result modelinfo.ModelInfo
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func GetModelInfos() (ModelInfoHeadersResult, error) {
	println("Get the list of model infos")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "models")
	response, err := SendGetRequest(uri)
	if err != nil {
		return ModelInfoHeadersResult{}, err
	}

	var result ModelInfoHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func GetVendors() (VendorItemHeadersResult, error) {
	println("Get the list of vendors")

	uri := fmt.Sprintf("%s/%s", modelinfo.RouterKey, "vendors")
	response, err := SendGetRequest(uri)
	if err != nil {
		return VendorItemHeadersResult{}, err
	}

	var result VendorItemHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func GetVendorModels(vid uint16) (modelinfo.VendorProducts, error) {
	println("Get the list of models for VID:", vid)

	uri := fmt.Sprintf("%s/%s/%v", modelinfo.RouterKey, "models", vid)
	response, err := SendGetRequest(uri)
	if err != nil {
		return modelinfo.VendorProducts{}, err
	}

	var result modelinfo.VendorProducts
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func PublishTestingResult(testingResult compliancetest.MsgAddTestingResult) (json.RawMessage, error) {
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
	response, err := SendPostRequest(uri, body, constants.AccountName, constants.Passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func GetTestingResult(vid uint16, pid uint16) (compliancetest.TestingResults, error) {
	println(fmt.Sprintf("Get Testing Result for Model with VID:%v PID:%v", vid, pid))

	uri := fmt.Sprintf("%s/%s/%v/%v", compliancetest.RouterKey, "testresults", vid, pid)
	response, err := SendGetRequest(uri)
	if err != nil {
		return compliancetest.TestingResults{}, err
	}

	var result compliancetest.TestingResults
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
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

func PublishCertifiedModel(certifyModel compliance.MsgCertifyModel) (json.RawMessage, error) {
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
	response, err := SendPutRequest(uri, body, constants.AccountName, constants.Passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func PublishRevokedModel(revokeModel compliance.MsgRevokeModel) (json.RawMessage, error) {
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
	response, err := SendPutRequest(uri, body, constants.AccountName, constants.Passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func GetComplianceInfo(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfo, error) {
	println(fmt.Sprintf("Get Compliance Info for Model with VID:%v PID:%v", vid, pid))
	return getComplianceInfo(vid, pid, certificationType)
}

func GetCertifiedModel(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, error) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Certified", vid, pid))
	return getComplianceInfoInState(vid, pid, certificationType, "certified")
}

func GetRevokedModel(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfoInState, error) {
	println(fmt.Sprintf("Get if Model with VID:%v PID:%v Revoked", vid, pid))
	return getComplianceInfoInState(vid, pid, certificationType, "revoked")
}

func getComplianceInfo(vid uint16, pid uint16, certificationType compliance.CertificationType) (compliance.ComplianceInfo, error) {
	uri := fmt.Sprintf("%s/%v/%v/%v", compliance.RouterKey, vid, pid, certificationType)
	response, err := SendGetRequest(uri)
	if err != nil {
		return compliance.ComplianceInfo{}, err
	}

	var result compliance.ComplianceInfo
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func getComplianceInfoInState(vid uint16, pid uint16, certificationType compliance.CertificationType, state string) (compliance.ComplianceInfoInState, error) {
	uri := fmt.Sprintf("%s/%v/%v/%v/%v", compliance.RouterKey, state, vid, pid, certificationType)

	response, err := SendGetRequest(uri)
	if err != nil {
		return compliance.ComplianceInfoInState{}, err
	}

	var result compliance.ComplianceInfoInState
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func GetComplianceInfos() (ComplianceInfosHeadersResult, error) {
	println("Get all compliance info records")
	return GetAllComplianceInfos("")
}

func GetAllCertifiedModels() (ComplianceInfosHeadersResult, error) {
	println("Get all certified models")
	return GetAllComplianceInfos("certified")
}

func GetAllRevokedModels() (ComplianceInfosHeadersResult, error) {
	println("Get all revoked models")
	return GetAllComplianceInfos("revoked")
}

func GetAllComplianceInfos(state string) (ComplianceInfosHeadersResult, error) {
	var uri string

	if len(state) > 0 {
		uri = fmt.Sprintf("%s/%v", compliance.RouterKey, state)
	} else {
		uri = fmt.Sprintf("%s", compliance.RouterKey)
	}

	response, err := SendGetRequest(uri)
	if err != nil {
		return ComplianceInfosHeadersResult{}, err
	}

	var result ComplianceInfosHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func ProposeAddX509RootCert(proposeAddX509RootCert pki.MsgProposeAddX509RootCert, account string, passphrase string) (json.RawMessage, error) {
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

	response, err := SendPostRequest(uri, body, account, passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func ApproveAddX509RootCert(msgApproveAddX509RootCert pki.MsgApproveAddX509RootCert, account string, passphrase string) (json.RawMessage, error) {
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

	response, err := SendPatchRequest(uri, body, account, passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func AddX509Cert(addX509Cert pki.MsgAddX509Cert, account string, passphrase string) (json.RawMessage, error) {
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

	response, err := SendPostRequest(uri, body, account, passphrase)
	if err != nil {
		return json.RawMessage{}, err
	}

	return removeResponseWrapper(response), nil
}

func GetAllX509RootCerts() (CertificatesHeadersResult, error) {
	return getCertificates(fmt.Sprintf("%s/certs/root", pki.RouterKey))
}

func GetAllSubjectX509Certs(subject string) (CertificatesHeadersResult, error) {
	return getCertificates(fmt.Sprintf("%s/certs/%s", pki.RouterKey, subject))
}

func GetAllX509Certs() (CertificatesHeadersResult, error) {
	return getCertificates(fmt.Sprintf("%s/certs", pki.RouterKey))
}

func GetAllProposedX509RootCerts() (ProposedCertificatesHeadersResult, error) {
	return getProposedCertificates(fmt.Sprintf("%s/certs/proposed/root", pki.RouterKey))
}

func GetProposedX509RootCert(subject string, subjectKeyId string) (pki.ProposedCertificate, error) {
	response, err := SendGetRequest(fmt.Sprintf("%s/certs/proposed/root/%s/%s", pki.RouterKey, subject, subjectKeyId))
	if err != nil {
		return pki.ProposedCertificate{}, err
	}

	var result pki.ProposedCertificate
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func GetX509Cert(subject string, subjectKeyId string) (pki.Certificate, error) {
	response, err := SendGetRequest(fmt.Sprintf("%s/certs/%s/%s", pki.RouterKey, subject, subjectKeyId))
	if err != nil {
		return pki.Certificate{}, err
	}

	var result pki.Certificates
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	if len(result.Items) > 1 {
		return pki.Certificate{}, sdk.ErrInternal("Unexpected certificates number")
	}

	return result.Items[0], nil
}

func getCertificates(uri string) (CertificatesHeadersResult, error) {
	response, err := SendGetRequest(uri)
	if err != nil {
		return CertificatesHeadersResult{}, err
	}

	var result CertificatesHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
}

func getProposedCertificates(uri string) (ProposedCertificatesHeadersResult, error) {
	response, err := SendGetRequest(uri)
	if err != nil {
		return ProposedCertificatesHeadersResult{}, err
	}

	var result ProposedCertificatesHeadersResult
	_ = json.Unmarshal(removeResponseWrapper(response), &result)

	return result, nil
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
