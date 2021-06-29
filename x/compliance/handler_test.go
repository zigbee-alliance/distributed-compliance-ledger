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

//nolint:testpackage
package compliance

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo"
)

func TestHandler_CertifyModel(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid)

	// check
	checkCertifiedModel(t, receivedComplianceInfo, certifyModelMsg)

	certified, _ := queryCertifiedModel(setup, vid, pid)
	require.True(t, certified)
}

func TestHandler_CertifyModelByDifferentRoles(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	cases := []auth.AccountRole{
		auth.Vendor,
		auth.TestHouse,
	}

	for _, tc := range cases {
		address := constants.Address2
		account := auth.NewAccount(address, constants.PubKey1, auth.AccountRoles{tc})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to certify model
		certifyModelMsg := msgCertifyModel(address, vid, pid)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CertifyModelForUnknownModel(t *testing.T) {
	setup := Setup()

	// try to certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, constants.VID, constants.PID)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, modelinfo.CodeModelInfoDoesNotExist, result.Code)
}

func TestHandler_CertifyModelForModelWithoutTestingResults(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, constants.VID, constants.PID)

	// try to certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, compliancetest.CodeTestingResultDoesNotExist, result.Code)
}

func TestHandler_CertifyModelTwice(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// certify model second time
	secondCertifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	secondCertifyModelMsg.CertificationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code) // result is OK, BUT CertificationDate must be from the first message

	// check
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid)
	require.Equal(t, receivedComplianceInfo.Date, certifyModelMsg.CertificationDate)
}

func TestHandler_CertifyDifferentModels(t *testing.T) {
	setup := Setup()

	for i := uint16(1); i < uint16(5); i++ {
		// add model amd testing result
		vid, pid := addModel(setup, i, i)
		addTestingResult(setup, vid, pid)

		// add new testing result
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certified model
		receivedModel, _ := queryComplianceInfo(setup, vid, pid)

		// check
		checkCertifiedModel(t, receivedModel, certifyModelMsg)
	}
}

func TestHandler_CertifyModelForEmptyCertificationType(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationType = ""
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeUnknownRequest, result.Code)
}

func TestHandler_CertifyModelForNotZbCertificationType(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationType = "Other"
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeUnknownRequest, result.Code)
}

func TestHandler_RevokeModel(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query revoked model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, revokedModelMsg.VID, revokedModelMsg.PID)

	// check
	checkRevokedModel(t, receivedComplianceInfo, revokedModelMsg)

	revoked, _ := queryRevokedModel(setup, revokedModelMsg.VID, revokedModelMsg.PID)
	require.True(t, revoked)
}

func TestHandler_RevokeCertifiedModel(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	revokedModelMsg.RevocationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query revoked model
	revokedModel, _ := queryComplianceInfo(setup, vid, pid)

	// check
	checkRevokedModel(t, revokedModel, revokedModelMsg)
	require.Equal(t, 1, len(revokedModel.History))
	require.Equal(t, types.Certified, revokedModel.History[0].State)
	require.Equal(t, certifyModelMsg.CertificationDate, revokedModel.History[0].Date)

	revoked, _ := queryRevokedModel(setup, vid, pid)
	require.True(t, revoked)

	// query certified model
	_, err := queryCertifiedModel(setup, vid, pid)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestHandler_RevokeModelByDifferentRoles(t *testing.T) {
	setup := Setup()

	cases := []auth.AccountRole{
		auth.Vendor,
		auth.TestHouse,
	}

	for _, tc := range cases {
		address := constants.Address2
		account := auth.NewAccount(address, constants.PubKey1, auth.AccountRoles{tc})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to certify model
		revokeModelMsg := msgRevokedModel(address, constants.VID, constants.PID)
		result := setup.Handler(setup.Ctx, revokeModelMsg)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_RevokeModelTwice(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// certify model second time
	secondRevokeModelMsg := msgRevokedModel(setup.CertificationCenter, revokedModelMsg.VID, revokedModelMsg.PID)
	secondRevokeModelMsg.RevocationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, secondRevokeModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code) // result is OK, BUT RevocationDate must be from the first message

	// check
	receivedComplianceInfo, _ := queryComplianceInfo(setup, secondRevokeModelMsg.VID, secondRevokeModelMsg.PID)
	require.Equal(t, receivedComplianceInfo.Date, revokedModelMsg.RevocationDate)
}

func TestHandler_RevokeDifferentModels(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	for i := uint16(1); i < uint16(5); i++ {
		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, constants.VID, constants.PID)
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certified model
		receivedModel, _ := queryComplianceInfo(setup, revokedModelMsg.VID, revokedModelMsg.PID)

		// check
		checkRevokedModel(t, receivedModel, revokedModelMsg)
	}
}

func TestHandler_RevokeCertifiedModelForRevocationDateBeforeCertificationDate(t *testing.T) {
	setup := Setup()

	revocationDate := time.Now().UTC()
	certificationDate := revocationDate.AddDate(0, 0, 1)

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationDate = certificationDate
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	revokedModelMsg.RevocationDate = revocationDate
	result = setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, types.CodeInconsistentDates, result.Code)
}

func TestHandler_CertifyRevokedModelForCertificationDateBeforeRevocationDate(t *testing.T) {
	setup := Setup()

	certificationDate := time.Now().UTC()
	revocationDate := certificationDate.AddDate(0, 0, 1)

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	revokedModelMsg.RevocationDate = revocationDate
	result := setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationDate = certificationDate
	result = setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, types.CodeInconsistentDates, result.Code)
}

func TestHandler_CertifyRevokedModel(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// revoke model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid)
	revokedModelMsg.RevocationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query revoked model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid)
	require.Equal(t, types.Revoked, receivedComplianceInfo.State)
	require.Equal(t, 1, len(receivedComplianceInfo.History))

	// certify model again
	secondCertifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	secondCertifyModelMsg.CertificationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedComplianceInfo, _ = queryComplianceInfo(setup, vid, pid)

	// check
	checkCertifiedModel(t, receivedComplianceInfo, secondCertifyModelMsg)
	require.Equal(t, 2, len(receivedComplianceInfo.History))

	require.Equal(t, types.Certified, receivedComplianceInfo.History[0].State)
	require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

	require.Equal(t, types.Revoked, receivedComplianceInfo.History[1].State)
	require.Equal(t, revokedModelMsg.RevocationDate, receivedComplianceInfo.History[1].Date)

	// query revoked model
	_, err := queryRevokedModel(setup, vid, pid)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestHandler_CertifyRevokedModelForTrackRevocationStrategy(t *testing.T) {
	setup := Setup()

	// revoke non-existent model
	revokedModelMsg := msgRevokedModel(setup.CertificationCenter, constants.VID, constants.PID)
	revokedModelMsg.RevocationDate = time.Now().UTC()
	result := setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)

	// add model
	vid, pid := addModel(setup, constants.VID, constants.PID)

	// revoke model
	revokedModelMsg = msgRevokedModel(setup.CertificationCenter, vid, pid)
	revokedModelMsg.RevocationDate = time.Now().UTC()
	result = setup.Handler(setup.Ctx, revokedModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid)
	checkRevokedModel(t, receivedComplianceInfo, revokedModelMsg)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationDate = revokedModelMsg.RevocationDate.AddDate(0, 0, 1)
	result = setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_CheckZbCertificationDone(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create other account certification center
	account := auth.NewAccount(constants.Address3, constants.PubKey3, auth.AccountRoles{auth.ZBCertificationCenter})
	setup.authKeeper.SetAccount(setup.Ctx, account)

	secondCertifyModelMsg := msgCertifyModel(account.Address, vid, pid)
	result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
	require.Equal(t, types.CodeAlreadyCertifyed, result.Code)
}

func queryComplianceInfo(setup TestSetup, vid uint16, pid uint16) (types.ComplianceInfo, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{
			keeper.QueryComplianceInfo, fmt.Sprintf("%v", vid),
			fmt.Sprintf("%v", pid), fmt.Sprintf("%v", types.ZbCertificationType),
		},
		abci.RequestQuery{},
	)
	if err != nil {
		return types.ComplianceInfo{}, err
	}

	var model types.ComplianceInfo
	_ = setup.Cdc.UnmarshalJSON(result, &model)

	return model, nil
}

func queryCertifiedModel(setup TestSetup, vid uint16, pid uint16) (bool, sdk.Error) {
	return queryComplianceInfoInState(setup, vid, pid, keeper.QueryCertifiedModel)
}

func queryRevokedModel(setup TestSetup, vid uint16, pid uint16) (bool, sdk.Error) {
	return queryComplianceInfoInState(setup, vid, pid, keeper.QueryRevokedModel)
}

func queryComplianceInfoInState(setup TestSetup, vid uint16, pid uint16, state string) (bool, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{state, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", types.ZbCertificationType)},
		abci.RequestQuery{},
	)
	if err != nil {
		return false, err
	}

	var model types.ComplianceInfoInState
	_ = setup.Cdc.UnmarshalJSON(result, &model)

	return model.Value, nil
}

func addModel(setup TestSetup, vid uint16, pid uint16) (uint16, uint16) {
	modelInfo := modelinfo.ModelInfo{
		Model: modelinfo.Model{
			VID:                   vid,
			PID:                   pid,
			CID:                   constants.CID,
			Name:                  constants.Name,
			Description:           constants.Description,
			SKU:                   constants.SKU,
			SoftwareVersion:       constants.SoftwareVersion,
			SoftwareVersionString: constants.SoftwareVersionString,
			HardwareVersion:       constants.HardwareVersion,
			HardwareVersionString: constants.HardwareVersionString,
			CDVersionNumber:       constants.CDVersionNumber,
			OtaURL:                constants.OtaURL,
			OtaChecksum:           constants.OtaChecksum,
			OtaChecksumType:       constants.OtaChecksumType,
			Revoked:               constants.Revoked,
		},
		Owner: constants.Owner,
	}

	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)

	return vid, pid
}

func addTestingResult(setup TestSetup, vid uint16, pid uint16) (uint16, uint16) {
	testingResult := compliancetest.TestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: constants.TestResult,
		Owner:      constants.Owner,
	}

	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	return vid, pid
}

func msgCertifyModel(signer sdk.AccAddress, vid uint16, pid uint16) MsgCertifyModel {
	return MsgCertifyModel{
		VID:               vid,
		PID:               pid,
		CertificationDate: constants.CertificationDate,
		CertificationType: types.CertificationType(constants.CertificationType),
		Signer:            signer,
	}
}

func msgRevokedModel(signer sdk.AccAddress, vid uint16, pid uint16) MsgRevokeModel {
	return MsgRevokeModel{
		VID:               vid,
		PID:               pid,
		RevocationDate:    constants.RevocationDate,
		Reason:            constants.RevocationReason,
		CertificationType: types.CertificationType(constants.CertificationType),
		Signer:            signer,
	}
}

func checkCertifiedModel(t *testing.T, receivedComplianceInfo ComplianceInfo, certifyModelMsg MsgCertifyModel) {
	require.Equal(t, receivedComplianceInfo.VID, certifyModelMsg.VID)
	require.Equal(t, receivedComplianceInfo.PID, certifyModelMsg.PID)
	require.Equal(t, receivedComplianceInfo.State, types.Certified)
	require.Equal(t, receivedComplianceInfo.Date, certifyModelMsg.CertificationDate)
	require.Equal(t, receivedComplianceInfo.CertificationType, types.ZbCertificationType)
}

func checkRevokedModel(t *testing.T, receivedComplianceInfo ComplianceInfo, revokeModelMsg MsgRevokeModel) {
	require.Equal(t, receivedComplianceInfo.VID, revokeModelMsg.VID)
	require.Equal(t, receivedComplianceInfo.PID, revokeModelMsg.PID)
	require.Equal(t, receivedComplianceInfo.State, types.Revoked)
	require.Equal(t, receivedComplianceInfo.Date, revokeModelMsg.RevocationDate)
	require.Equal(t, receivedComplianceInfo.Reason, revokeModelMsg.Reason)
	require.Equal(t, receivedComplianceInfo.CertificationType, types.ZbCertificationType)
}
