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

package rest_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/common"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `dclcli rest-server --chain-id dclchain`

	TODO: provide tests for error cases
*/

//nolint:funlen
func TestComplianceDemo_KeepTrackCompliance(t *testing.T) {
	// Register new Vendor account
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Register new TestHouse account
	testHouse := utils.CreateNewAccount(auth.AccountRoles{auth.TestHouse})

	// Register new ZBCertificationCenter account
	zb := utils.CreateNewAccount(auth.AccountRoles{auth.ZBCertificationCenter})

	// Get all compliance infos
	inputComplianceInfos, _ := utils.GetComplianceInfos()

	// Get all certified models
	inputCertifiedModels, _ := utils.GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := utils.GetAllRevokedModels()

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	_, _ = utils.AddModelInfo(modelInfo, vendor)

	// Check if model either certified or revoked before Compliance record was created
	modelIsCertified, _ := utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, compliance.ZbCertificationType)
	require.False(t, modelIsCertified.Value)

	modelIsRevoked, _ := utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, compliance.ZbCertificationType)
	require.False(t, modelIsRevoked.Value)

	// Publish testing result
	testingResult := utils.NewMsgAddTestingResult(modelInfo.VID, modelInfo.PID, testHouse.Address)
	_, _ = utils.PublishTestingResult(testingResult, testHouse)

	// Certify model
	certifyModelMsg := compliance.NewMsgCertifyModel(modelInfo.VID, modelInfo.PID, time.Now().UTC(),
		compliance.CertificationType(testconstants.CertificationType), testconstants.EmptyString, zb.Address)
	_, _ = utils.PublishCertifiedModel(certifyModelMsg, zb)

	// Check model is certified
	modelIsCertified, _ = utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.True(t, modelIsCertified.Value)

	// Register other ZBCertificationCenter account
	secondZb := utils.CreateNewAccount(auth.AccountRoles{auth.ZBCertificationCenter})

	// Certify model by other ZBCertificationCenter account
	secondCertifyModelMsg := compliance.NewMsgCertifyModel(modelInfo.VID, modelInfo.PID, time.Now().UTC(),
		compliance.CertificationType(testconstants.CertificationType), testconstants.EmptyString, secondZb.Address)
	secondCertifyResult, _ := utils.PublishCertifiedModel(secondCertifyModelMsg, secondZb)

	require.Equal(t, compliance.CodeAlreadyCertifyed, sdk.CodeType(secondCertifyResult.Code))

	modelIsRevoked, _ = utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.False(t, modelIsRevoked.Value)

	// Get all certified models
	certifiedModels, _ := utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total)+1, utils.ParseUint(certifiedModels.Total))

	// Revoke model certification
	revocationTime := certifyModelMsg.CertificationDate.AddDate(0, 0, 1)
	revokeModelMsg := compliance.NewMsgRevokeModel(modelInfo.VID, modelInfo.PID, revocationTime,
		compliance.CertificationType(testconstants.CertificationType), testconstants.RevocationReason, zb.Address)
	_, _ = utils.PublishRevokedModel(revokeModelMsg, zb)

	// Check model is revoked
	modelIsCertified, _ = utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, revokeModelMsg.CertificationType)
	require.False(t, modelIsCertified.Value)

	modelIsRevoked, _ = utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, revokeModelMsg.CertificationType)
	require.True(t, modelIsRevoked.Value)

	// Get all revoked models
	revokedModels, _ := utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total)+1, utils.ParseUint(revokedModels.Total))

	// Get all certified models
	certifiedModels, _ = utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total), utils.ParseUint(certifiedModels.Total))

	// Get all compliance infos
	complianceInfos, _ := utils.GetComplianceInfos()
	require.Equal(t, utils.ParseUint(inputComplianceInfos.Total)+1, utils.ParseUint(complianceInfos.Total))

	// Get compliance info
	complianceInfo, _ := utils.GetComplianceInfo(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.Equal(t, complianceInfo.State, compliance.RevokedState)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, complianceInfo.History[0].State, compliance.CertifiedState)
}

func TestComplianceDemo_KeepTrackRevocation(t *testing.T) {
	// Register new Vendor account
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	_, _ = utils.AddModelInfo(modelInfo, vendor)

	// Get all certified models
	inputCertifiedModels, _ := utils.GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := utils.GetAllRevokedModels()

	// Register new ZBCertificationCenter account
	zb := utils.CreateNewAccount(auth.AccountRoles{auth.ZBCertificationCenter})

	vid, pid := modelInfo.VID, modelInfo.PID

	// Revoke non-existent model
	revocationTime := time.Now().UTC()
	revokeModelMsg := compliance.NewMsgRevokeModel(common.RandUint16(), common.RandUint16(), revocationTime,
		compliance.CertificationType(testconstants.CertificationType), testconstants.RevocationReason, zb.Address)
	_, _ = utils.PublishRevokedModel(revokeModelMsg, zb)

	// Check non-existent model is revoked
	modelIsRevoked, _ := utils.GetRevokedModel(revokeModelMsg.VID,
		revokeModelMsg.PID, revokeModelMsg.CertificationType)
	require.False(t, modelIsRevoked.Value)

	// Revoke model
	revocationTime = time.Now().UTC()
	revokeModelMsg = compliance.NewMsgRevokeModel(vid, pid, revocationTime,
		compliance.CertificationType(testconstants.CertificationType), testconstants.RevocationReason, zb.Address)
	_, _ = utils.PublishRevokedModel(revokeModelMsg, zb)

	// Check model is revoked
	modelIsRevoked, _ = utils.GetRevokedModel(revokeModelMsg.VID,
		revokeModelMsg.PID, revokeModelMsg.CertificationType)
	require.True(t, modelIsRevoked.Value)

	modelIsCertified, _ := utils.GetCertifiedModel(revokeModelMsg.VID,
		revokeModelMsg.PID, revokeModelMsg.CertificationType)
	require.False(t, modelIsCertified.Value)

	// Get all revoked models
	revokedModels, _ := utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total)+1, utils.ParseUint(revokedModels.Total))

	// Certify model
	certificationTime := revocationTime.AddDate(0, 0, 1)
	certifyModelMsg := compliance.NewMsgCertifyModel(vid, pid, certificationTime,
		compliance.CertificationType(testconstants.CertificationType), testconstants.EmptyString, zb.Address)
	_, _ = utils.PublishCertifiedModel(certifyModelMsg, zb)

	// Check model is certified
	modelIsRevoked, _ = utils.GetRevokedModel(certifyModelMsg.VID,
		certifyModelMsg.PID, certifyModelMsg.CertificationType)
	require.False(t, modelIsRevoked.Value)

	modelIsCertified, _ = utils.GetCertifiedModel(certifyModelMsg.VID,
		certifyModelMsg.PID, certifyModelMsg.CertificationType)
	require.True(t, modelIsCertified.Value)

	// Get all certified models
	certifiedModels, _ := utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total)+1, utils.ParseUint(certifiedModels.Total))

	// Get all revoked models
	revokedModels, _ = utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total), utils.ParseUint(revokedModels.Total))
}
