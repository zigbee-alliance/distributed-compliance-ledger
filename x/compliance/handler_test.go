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

//nolint:testpackage,lll,dupl
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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
)

func TestHandler_CertifyModel_Zigbee(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion,
		softwareVersionString, ZigbeeCertificationType)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, ZigbeeCertificationType)

	// check
	checkCertifiedModel(t, receivedComplianceInfo, certifyModelMsg, ZigbeeCertificationType)

	certified, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, ZigbeeCertificationType)
	require.True(t, certified)
}

func TestHandler_CertifyModel_Matter(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, MatterCertificationType)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, MatterCertificationType)

	// check
	checkCertifiedModel(t, receivedComplianceInfo, certifyModelMsg, MatterCertificationType)

	certified, _ := queryCertifiedModel(setup, vid, pid, softwareVersion, MatterCertificationType)
	require.True(t, certified)
}

func TestHandler_CertifyModelByDifferentRoles(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	cases := []auth.AccountRole{
		auth.Vendor,
		auth.TestHouse,
	}

	for _, tc := range cases {
		address := constants.Address2
		account := auth.NewAccount(address, constants.PubKey1, auth.AccountRoles{tc}, constants.VendorID1)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to certify model
		certifyModelMsg := msgCertifyModel(address, vid, pid, softwareVersion, softwareVersionString, ZigbeeCertificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)

		certifyModelMsg = msgCertifyModel(address, vid, pid, softwareVersion, softwareVersionString, MatterCertificationType)
		result = setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CertifyModelForUnknownModel(t *testing.T) {
	setup := Setup()

	// try to certify model

	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, constants.VID, constants.PID,
			constants.SoftwareVersion, constants.SoftwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, model.CodeModelVersionDoesNotExist, result.Code)
	}
}

func TestHandler_CertifyModelForModelWithoutTestingResults(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	// try to certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, compliancetest.CodeTestingResultDoesNotExist, result.Code)
	}
}

func TestHandler_CertifyModelWithWrongSoftwareVersionString(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString+"-modified", certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, types.CodeModelVersionStringDoesNotMatch, result.Code)
	}
}

func TestHandler_CertifyModelTwice(t *testing.T) {
	setup := Setup()

	// add model and testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion,
			softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// certify model second time
		secondCertifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid,
			softwareVersion, softwareVersionString, certificationType)
		secondCertifyModelMsg.CertificationDate = time.Now().UTC()
		result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code) // result is OK, BUT CertificationDate must be from the first message

		// check
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		require.Equal(t, receivedComplianceInfo.Date, certifyModelMsg.CertificationDate)
	}
}

func TestHandler_CertifyDifferentModels(t *testing.T) {
	setup := Setup()

	for i := uint16(1); i < uint16(5); i++ {
		// add model add testing result
		vid, pid := addModel(setup, constants.VID, constants.PID)
		// add model version
		_, _, softwareVersion, softwareVersionString :=
			addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

		addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

		for _, certificationType := range setup.CertificationTypes {
			// add new testing result
			certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
			result := setup.Handler(setup.Ctx, certifyModelMsg)
			require.Equal(t, sdk.CodeOK, result.Code)

			// query certified model
			receivedModel, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

			// check
			checkCertifiedModel(t, receivedModel, certifyModelMsg, certificationType)
		}
	}
}

func TestHandler_CertifyModelForEmptyCertificationType(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		certifyModelMsg.CertificationType = ""
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnknownRequest, result.Code)
	}
}

func TestHandler_CertifyModelForNotZbCertificationType(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	// certify model
	for _, certificationType := range setup.CertificationTypes {
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		certifyModelMsg.CertificationType = "Other"
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnknownRequest, result.Code)
	}
}

func TestHandler_RevokeModel(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query revoked model
		receivedComplianceInfo, _ := queryComplianceInfo(setup, revokedModelMsg.VID, revokedModelMsg.PID, revokedModelMsg.SoftwareVersion, certificationType)

		// check
		checkRevokedModel(t, receivedComplianceInfo, revokedModelMsg, certificationType)

		revoked, _ := queryRevokedModel(setup, revokedModelMsg.VID, revokedModelMsg.PID, revokedModelMsg.SoftwareVersion, certificationType)
		require.True(t, revoked)
	}
}

func TestHandler_RevokeCertifiedModel(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = time.Now().UTC()
		result = setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query revoked model
		revokedModel, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkRevokedModel(t, revokedModel, revokedModelMsg, certificationType)
		require.Equal(t, 1, len(revokedModel.History))
		require.Equal(t, types.CodeCertified, revokedModel.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, revokedModel.History[0].Date)

		revoked, _ := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.True(t, revoked)

		// query certified model
		_, err := queryCertifiedModel(setup, vid, pid, softwareVersion, certificationType)
		require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
	}
}

func TestHandler_RevokeModelByDifferentRoles(t *testing.T) {
	setup := Setup()

	cases := []auth.AccountRole{
		auth.Vendor,
		auth.TestHouse,
	}

	for _, tc := range cases {
		address := constants.Address2
		account := auth.NewAccount(address, constants.PubKey1, auth.AccountRoles{tc}, constants.VendorID1)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to certify model
		for _, certificationType := range setup.CertificationTypes {
			revokeModelMsg := msgRevokedModel(address, constants.VID, constants.PID, constants.SoftwareVersion,
				constants.SoftwareVersionString, certificationType)
			result := setup.Handler(setup.Ctx, revokeModelMsg)
			require.Equal(t, sdk.CodeUnauthorized, result.Code)
		}
	}
}

func TestHandler_RevokeModelTwice(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// certify model second time
		secondRevokeModelMsg := msgRevokedModel(setup.CertificationCenter, revokedModelMsg.VID, revokedModelMsg.PID,
			revokedModelMsg.SoftwareVersion, revokedModelMsg.SoftwareVersionString, certificationType)
		secondRevokeModelMsg.RevocationDate = time.Now().UTC()
		result = setup.Handler(setup.Ctx, secondRevokeModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code) // result is OK, BUT RevocationDate must be from the first message

		// check
		receivedComplianceInfo, _ := queryComplianceInfo(setup, secondRevokeModelMsg.VID, secondRevokeModelMsg.PID,
			secondRevokeModelMsg.SoftwareVersion, certificationType)
		require.Equal(t, receivedComplianceInfo.Date, revokedModelMsg.RevocationDate)
	}
}

func TestHandler_RevokeDifferentModels(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for i := uint16(1); i < uint16(5); i++ {
		for _, certificationType := range setup.CertificationTypes {
			// revoke model
			revokedModelMsg := msgRevokedModel(setup.CertificationCenter, constants.VID, constants.PID,
				constants.SoftwareVersion, constants.SoftwareVersionString, certificationType)
			result := setup.Handler(setup.Ctx, revokedModelMsg)
			require.Equal(t, sdk.CodeOK, result.Code)

			// query certified model
			receivedModel, _ := queryComplianceInfo(setup, revokedModelMsg.VID, revokedModelMsg.PID,
				revokedModelMsg.SoftwareVersion, certificationType)

			// check
			checkRevokedModel(t, receivedModel, revokedModelMsg, certificationType)
		}
	}
}

func TestHandler_RevokeCertifiedModelForRevocationDateBeforeCertificationDate(t *testing.T) {
	setup := Setup()

	revocationDate := time.Now().UTC()
	certificationDate := revocationDate.AddDate(0, 0, 1)

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		certifyModelMsg.CertificationDate = certificationDate
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = revocationDate
		result = setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, types.CodeInconsistentDates, result.Code)
	}
}

func TestHandler_CertifyRevokedModelForCertificationDateBeforeRevocationDate(t *testing.T) {
	setup := Setup()

	certificationDate := time.Now().UTC()
	revocationDate := certificationDate.AddDate(0, 0, 1)

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = revocationDate
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		certifyModelMsg.CertificationDate = certificationDate
		result = setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, types.CodeInconsistentDates, result.Code)
	}
}

func TestHandler_CertifyRevokedModel(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes { // certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = time.Now().UTC()
		result = setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query revoked model
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		require.Equal(t, types.CodeRevoked, receivedComplianceInfo.SoftwareVersionCertificationStatus)
		require.Equal(t, 1, len(receivedComplianceInfo.History))

		// certify model again
		secondCertifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		secondCertifyModelMsg.CertificationDate = time.Now().UTC()
		result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certified model
		receivedComplianceInfo, _ = queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)

		// check
		checkCertifiedModel(t, receivedComplianceInfo, secondCertifyModelMsg, certificationType)
		require.Equal(t, 2, len(receivedComplianceInfo.History))

		require.Equal(t, types.CodeCertified, receivedComplianceInfo.History[0].SoftwareVersionCertificationStatus)
		require.Equal(t, certifyModelMsg.CertificationDate, receivedComplianceInfo.History[0].Date)

		require.Equal(t, types.CodeRevoked, receivedComplianceInfo.History[1].SoftwareVersionCertificationStatus)
		require.Equal(t, revokedModelMsg.RevocationDate, receivedComplianceInfo.History[1].Date)

		// query revoked model
		_, err := queryRevokedModel(setup, vid, pid, softwareVersion, certificationType)
		require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
	}
}

func TestHandler_CertifyRevokedModelForTrackRevocationStrategy(t *testing.T) {
	setup := Setup()

	for _, certificationType := range setup.CertificationTypes {
		// revoke non-existent model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, constants.VID, constants.PID, constants.SoftwareVersion, constants.SoftwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = time.Now().UTC()
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, types.CodeModelDoesNotExist, result.Code)
	}

	// add model
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// revoke model
		revokedModelMsg := msgRevokedModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		revokedModelMsg.RevocationDate = time.Now().UTC()
		result := setup.Handler(setup.Ctx, revokedModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certified model
		receivedComplianceInfo, _ := queryComplianceInfo(setup, vid, pid, softwareVersion, certificationType)
		checkRevokedModel(t, receivedComplianceInfo, revokedModelMsg, certificationType)

		// certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		certifyModelMsg.CertificationDate = revokedModelMsg.RevocationDate.AddDate(0, 0, 1)
		result = setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)
	}
}

func TestHandler_CheckCertificationDone(t *testing.T) {
	setup := Setup()

	// add model add testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, vid, pid, constants.SoftwareVersion, constants.SoftwareVersionString)

	addTestingResult(setup, vid, pid, softwareVersion, softwareVersionString)

	for _, certificationType := range setup.CertificationTypes {
		// certify model
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// create other account certification center
		account := auth.NewAccount(constants.Address3, constants.PubKey3, auth.AccountRoles{auth.CertificationCenter}, 0)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		secondCertifyModelMsg := msgCertifyModel(account.Address, vid, pid, softwareVersion, softwareVersionString, certificationType)
		result = setup.Handler(setup.Ctx, secondCertifyModelMsg)
		require.Equal(t, types.CodeAlreadyCertifyed, result.Code)
	}
}

func queryComplianceInfo(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, certificationType CertificationType) (types.ComplianceInfo, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{
			keeper.QueryComplianceInfo, fmt.Sprintf("%v", vid),
			fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion), fmt.Sprintf("%v", certificationType),
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

func queryCertifiedModel(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, certificationType CertificationType) (bool, sdk.Error) {
	return queryComplianceInfoInState(setup, vid, pid, softwareVersion, keeper.QueryCertifiedModel, certificationType)
}

func queryRevokedModel(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, certificationType CertificationType) (bool, sdk.Error) {
	return queryComplianceInfoInState(setup, vid, pid, softwareVersion, keeper.QueryRevokedModel, certificationType)
}

func queryComplianceInfoInState(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32,
	state string, certificationType CertificationType) (bool, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{state, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion), fmt.Sprintf("%v", certificationType)},
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
	model := model.Model{
		VID:          vid,
		PID:          pid,
		DeviceTypeID: constants.DeviceTypeID,
		ProductName:  constants.ProductName,
		ProductLabel: constants.ProductLabel,
		PartNumber:   constants.PartNumber,
	}

	setup.ModelKeeper.SetModel(setup.Ctx, model)

	return vid, pid
}

func addModelVersion(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, softwareVersionString string) (uint16, uint16, uint32, string) {
	modelVersion := model.ModelVersion{
		VID:                          vid,
		PID:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CDVersionNumber:              constants.CDVersionNumber,
		MinApplicableSoftwareVersion: constants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: constants.MaxApplicableSoftwareVersion,
	}

	setup.ModelKeeper.SetModelVersion(setup.Ctx, modelVersion)

	return vid, pid, softwareVersion, softwareVersionString
}

func addTestingResult(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, softwareVersionString string) (uint16, uint16) {
	testingResult := compliancetest.TestingResult{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		TestResult:            constants.TestResult,
		Owner:                 constants.Owner,
	}

	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	return vid, pid
}

func msgCertifyModel(signer sdk.AccAddress, vid uint16, pid uint16, softwareVersion uint32, softwareVersionString string, certificationType CertificationType) MsgCertifyModel {
	return MsgCertifyModel{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CertificationDate:     constants.CertificationDate,
		CertificationType:     certificationType,
		Signer:                signer,
	}
}

func msgRevokedModel(signer sdk.AccAddress, vid uint16, pid uint16, softwareVersion uint32, softwareVersionString string, certificationType CertificationType) MsgRevokeModel {
	return MsgRevokeModel{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		RevocationDate:        constants.RevocationDate,
		Reason:                constants.RevocationReason,
		CertificationType:     certificationType,
		Signer:                signer,
	}
}

func checkCertifiedModel(t *testing.T, receivedComplianceInfo ComplianceInfo, certifyModelMsg MsgCertifyModel, certificationType CertificationType) {
	require.Equal(t, receivedComplianceInfo.VID, certifyModelMsg.VID)
	require.Equal(t, receivedComplianceInfo.PID, certifyModelMsg.PID)
	require.Equal(t, receivedComplianceInfo.SoftwareVersionCertificationStatus, types.CodeCertified)
	require.Equal(t, receivedComplianceInfo.Date, certifyModelMsg.CertificationDate)
	require.Equal(t, receivedComplianceInfo.CertificationType, certificationType)
}

func checkRevokedModel(t *testing.T, receivedComplianceInfo ComplianceInfo, revokeModelMsg MsgRevokeModel, certificationType CertificationType) {
	require.Equal(t, receivedComplianceInfo.VID, revokeModelMsg.VID)
	require.Equal(t, receivedComplianceInfo.PID, revokeModelMsg.PID)
	require.Equal(t, receivedComplianceInfo.SoftwareVersionCertificationStatus, types.CodeRevoked)
	require.Equal(t, receivedComplianceInfo.Date, revokeModelMsg.RevocationDate)
	require.Equal(t, receivedComplianceInfo.Reason, revokeModelMsg.Reason)
	require.Equal(t, receivedComplianceInfo.CertificationType, certificationType)
}
