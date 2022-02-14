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

package compliance

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	test_model "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAllComplianceInfo(suite *utils.TestSuite) (res []compliancetypes.ComplianceInfo, err error) {
	if suite.Rest {
		var resp compliancetypes.QueryAllComplianceInfoResponse
		err := suite.QueryREST("/dcl/compliance/compliance-info", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetComplianceInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.ComplianceInfoAll(
			context.Background(),
			&compliancetypes.QueryAllComplianceInfoRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetComplianceInfo()
	}

	return res, nil
}

func GetComplianceInfo(suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string) (*compliancetypes.ComplianceInfo, error) {
	var res compliancetypes.ComplianceInfo

	if suite.Rest {
		var resp compliancetypes.QueryGetComplianceInfoResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/compliance-info/%d/%d/%d/%s",
				vid,
				pid,
				sv,
				ct,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetComplianceInfo()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.ComplianceInfo(
			context.Background(),
			&compliancetypes.QueryGetComplianceInfoRequest{
				Vid:               vid,
				Pid:               pid,
				SoftwareVersion:   sv,
				CertificationType: ct,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetComplianceInfo()
	}

	return &res, nil
}

func GetAllCertifiedModels(suite *utils.TestSuite) (res []compliancetypes.CertifiedModel, err error) {
	if suite.Rest {
		var resp compliancetypes.QueryAllCertifiedModelResponse
		err := suite.QueryREST("/dcl/compliance/certified-models", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertifiedModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.CertifiedModelAll(
			context.Background(),
			&compliancetypes.QueryAllCertifiedModelRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertifiedModel()
	}

	return res, nil
}

func GetCertifiedModel(suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string) (*compliancetypes.CertifiedModel, error) {
	var res compliancetypes.CertifiedModel

	if suite.Rest {
		var resp compliancetypes.QueryGetCertifiedModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/certified-models/%d/%d/%d/%s",
				vid,
				pid,
				sv,
				ct,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertifiedModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.CertifiedModel(
			context.Background(),
			&compliancetypes.QueryGetCertifiedModelRequest{
				Vid:               vid,
				Pid:               pid,
				SoftwareVersion:   sv,
				CertificationType: ct,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertifiedModel()
	}

	return &res, nil
}

func GetAllRevokedModels(suite *utils.TestSuite) (res []compliancetypes.RevokedModel, err error) {
	if suite.Rest {
		var resp compliancetypes.QueryAllRevokedModelResponse
		err := suite.QueryREST("/dcl/compliance/revoked-models", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.RevokedModelAll(
			context.Background(),
			&compliancetypes.QueryAllRevokedModelRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedModel()
	}

	return res, nil
}

func GetRevokedModel(suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string) (*compliancetypes.RevokedModel, error) {
	var res compliancetypes.RevokedModel

	if suite.Rest {
		var resp compliancetypes.QueryGetRevokedModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/revoked-models/%d/%d/%d/%s",
				vid,
				pid,
				sv,
				ct,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.RevokedModel(
			context.Background(),
			&compliancetypes.QueryGetRevokedModelRequest{
				Vid:               vid,
				Pid:               pid,
				SoftwareVersion:   sv,
				CertificationType: ct,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedModel()
	}

	return &res, nil
}

func GetAllProvisionalModels(suite *utils.TestSuite) (res []compliancetypes.ProvisionalModel, err error) {
	if suite.Rest {
		var resp compliancetypes.QueryAllProvisionalModelResponse
		err := suite.QueryREST("/dcl/compliance/provisional-models", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProvisionalModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.ProvisionalModelAll(
			context.Background(),
			&compliancetypes.QueryAllProvisionalModelRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProvisionalModel()
	}

	return res, nil
}

func GetProvisionalModel(suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string) (*compliancetypes.ProvisionalModel, error) {
	var res compliancetypes.ProvisionalModel

	if suite.Rest {
		var resp compliancetypes.QueryGetProvisionalModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/provisional-models/%d/%d/%d/%s",
				vid,
				pid,
				sv,
				ct,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProvisionalModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.ProvisionalModel(
			context.Background(),
			&compliancetypes.QueryGetProvisionalModelRequest{
				Vid:               vid,
				Pid:               pid,
				SoftwareVersion:   sv,
				CertificationType: ct,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProvisionalModel()
	}

	return &res, nil
}

func ComplianceDemoTrackCompliance(suite *utils.TestSuite) {
	// Query for unknown
	_, err := GetComplianceInfo(suite, testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.CertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.CertificationType)
	suite.AssertNotFound(err)
	_, err = GetCertifiedModel(suite, testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.CertificationType)
	suite.AssertNotFound(err)
	_, err = GetProvisionalModel(suite, testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.CertificationType)
	suite.AssertNotFound(err)

	// Query for empty test results
	inputAllComplianceInfo, _ := GetAllComplianceInfo(suite)
	require.Equal(suite.T, 0, len(inputAllComplianceInfo))
	inputAllCertifiedModels, _ := GetAllCertifiedModels(suite)
	require.Equal(suite.T, 0, len(inputAllCertifiedModels))
	inputAllRevokedModels, _ := GetAllRevokedModels(suite)
	require.Equal(suite.T, 0, len(inputAllRevokedModels))
	inputAllProvisionalModels, _ := GetAllProvisionalModels(suite)
	require.Equal(suite.T, 0, len(inputAllProvisionalModels))

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register new CertificationCenter account
	certCenter := utils.RandString()
	certCenterAccount := test_dclauth.CreateAccount(
		suite,
		certCenter,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter},
		1,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, certCenterAccount)

	// Publish model info
	pid := int32(tmrand.Uint16())
	firstModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	sv := tmrand.Uint32()
	svs := utils.RandString()
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check if model either certified or revoked before Compliance record was created
	_, err = GetComplianceInfo(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetCertifiedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetProvisionalModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)

	// Certify model
	certReason := "some reason"
	certDate := "2020-01-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Certify model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))

	// Check model is certified
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.Equal(suite.T, compliancetypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	modelIsCertified, _ := GetCertifiedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ := GetRevokedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ := GetProvisionalModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)

	// Get all models
	complianceInfos, _ := GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ := GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ := GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ := GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))

	// Revoke model certification
	revocReason := "some reason 2"
	revocDate := "2020-02-01T00:00:01Z"
	revocModelMsg := compliancetypes.MsgRevokeModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		RevocationDate:        revocDate,
		CertificationType:     "zigbee",
		Reason:                revocReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is revoked
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.Equal(suite.T, compliancetypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(3), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, revocReason, complianceInfo.Reason)
	require.Equal(suite.T, revocDate, complianceInfo.Date)
	modelIsCertified, _ = GetCertifiedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ = GetRevokedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ = GetProvisionalModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels), len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels)+1, len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
}

func ComplianceDemoTrackRevocation(suite *utils.TestSuite) {
	inputAllComplianceInfo, _ := GetAllComplianceInfo(suite)
	inputAllCertifiedModels, _ := GetAllCertifiedModels(suite)
	inputAllRevokedModels, _ := GetAllRevokedModels(suite)
	inputAllProvisionalModels, _ := GetAllProvisionalModels(suite)

	// TODO: simplify initialization

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register new CertificationCenter account
	certCenter := utils.RandString()
	certCenterAccount := test_dclauth.CreateAccount(
		suite,
		certCenter,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter},
		1,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, certCenterAccount)

	// Publish model info
	pid := int32(tmrand.Uint16())
	firstModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	sv := tmrand.Uint32()
	svs := utils.RandString()
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Revoke non-certified model
	revocReason := "some reason 3"
	revocDate := "2020-03-01T00:00:01Z"
	revocModelMsg := compliancetypes.MsgRevokeModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		RevocationDate:        revocDate,
		CertificationType:     "zigbee",
		Reason:                revocReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Revoke model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyRevoked.Is(err))

	// Check non-certified model is revoked
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.Equal(suite.T, compliancetypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(3), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, revocReason, complianceInfo.Reason)
	require.Equal(suite.T, revocDate, complianceInfo.Date)
	_, err = GetCertifiedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	modelIsRevoked, _ := GetRevokedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsRevoked.Value)
	_, err = GetProvisionalModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)

	// Get all
	complianceInfos, _ := GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ := GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels), len(certifiedModels))
	revokedModels, _ := GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels)+1, len(revokedModels))
	provisionalModels, _ := GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))

	// Certify model
	certReason := "some reason 4"
	certDate := "2020-05-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.Equal(suite.T, compliancetypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModel(suite, vid, pid, sv, compliancetypes.ZigbeeCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
}

func ComplianceDemoTrackProvision(suite *utils.TestSuite) {
	inputAllComplianceInfo, _ := GetAllComplianceInfo(suite)
	inputAllCertifiedModels, _ := GetAllCertifiedModels(suite)
	inputAllRevokedModels, _ := GetAllRevokedModels(suite)
	inputAllProvisionalModels, _ := GetAllProvisionalModels(suite)

	// TODO: simplify initialization

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register new CertificationCenter account
	certCenter := utils.RandString()
	certCenterAccount := test_dclauth.CreateAccount(
		suite,
		certCenter,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter},
		1,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
	)
	require.NotNil(suite.T, certCenterAccount)

	pid := int32(tmrand.Uint16())
	sv := tmrand.Uint32()
	svs := utils.RandString()

	// Provision non-existent model
	provReason := "some reason 10"
	provDate := "2021-03-01T00:00:01Z"
	provModelMsg := compliancetypes.MsgProvisionModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		ProvisionalDate:       provDate,
		CertificationType:     "matter",
		Reason:                provReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Provision model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyProvisional.Is(err))

	// Check non-existent model is provisioned
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.Equal(suite.T, compliancetypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(1), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, provReason, complianceInfo.Reason)
	require.Equal(suite.T, provDate, complianceInfo.Date)
	_, err = GetCertifiedModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	suite.AssertNotFound(err)
	provisionModel, _ := GetProvisionalModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.True(suite.T, provisionModel.Value)

	// Get all
	complianceInfos, _ := GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ := GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels), len(certifiedModels))
	revokedModels, _ := GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ := GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels)+1, len(provisionalModels))

	// Publish model info
	firstModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Certify model
	certReason := "some reason 44"
	certDate := "2021-10-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "matter",
		Reason:                certReason,
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.Equal(suite.T, compliancetypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModel(suite, vid, pid, sv, compliancetypes.MatterCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))

	// Can not provision certified model
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))
}
