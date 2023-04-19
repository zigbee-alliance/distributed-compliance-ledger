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
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAllComplianceInfo(suite *utils.TestSuite) (res []dclcompltypes.ComplianceInfo, err error) {
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

func GetComplianceInfo(
	suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string,
) (*dclcompltypes.ComplianceInfo, error) {
	var res dclcompltypes.ComplianceInfo

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

func GetComplianceInfoByHexVidAndPid(
	suite *utils.TestSuite, vid string, pid string, sv uint32, ct string,
) (*dclcompltypes.ComplianceInfo, error) {
	var res dclcompltypes.ComplianceInfo

	if suite.Rest {
		var resp compliancetypes.QueryGetComplianceInfoResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/compliance-info/%s/%s/%d/%s",
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

func GetCertifiedModel(
	suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string,
) (*compliancetypes.CertifiedModel, error) {
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

func GetCertifiedModelByHexVidAndPid(
	suite *utils.TestSuite, vid string, pid string, sv uint32, ct string,
) (*compliancetypes.CertifiedModel, error) {
	var res compliancetypes.CertifiedModel

	if suite.Rest {
		var resp compliancetypes.QueryGetCertifiedModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/certified-models/%s/%s/%d/%s",
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

func GetRevokedModel(
	suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string,
) (*compliancetypes.RevokedModel, error) {
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

func GetRevokedModelByHexVidAndPid(
	suite *utils.TestSuite, vid string, pid string, sv uint32, ct string,
) (*compliancetypes.RevokedModel, error) {
	var res compliancetypes.RevokedModel

	if suite.Rest {
		var resp compliancetypes.QueryGetRevokedModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/revoked-models/%s/%s/%d/%s",
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

func GetProvisionalModel(
	suite *utils.TestSuite, vid int32, pid int32, sv uint32, ct string,
) (*compliancetypes.ProvisionalModel, error) {
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

func GetProvisionalModelByHexVidAndPid(
	suite *utils.TestSuite, vid string, pid string, sv uint32, ct string,
) (*compliancetypes.ProvisionalModel, error) {
	var res compliancetypes.ProvisionalModel

	if suite.Rest {
		var resp compliancetypes.QueryGetProvisionalModelResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/provisional-models/%s/%s/%d/%s",
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
	}

	return &res, nil
}

func GetAllDeviceSoftwareCompliance(suite *utils.TestSuite) (res []compliancetypes.DeviceSoftwareCompliance, err error) {
	if suite.Rest {
		var resp compliancetypes.QueryAllDeviceSoftwareComplianceResponse
		err := suite.QueryREST("/dcl/compliance/device-software-compliance", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetDeviceSoftwareCompliance()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.DeviceSoftwareComplianceAll(
			context.Background(),
			&compliancetypes.QueryAllDeviceSoftwareComplianceRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetDeviceSoftwareCompliance()
	}

	return res, nil
}

func GetDeviceSoftwareCompliance(
	suite *utils.TestSuite, cDCertificateID string,
) (*compliancetypes.DeviceSoftwareCompliance, error) {
	var res compliancetypes.DeviceSoftwareCompliance

	if suite.Rest {
		var resp compliancetypes.QueryGetDeviceSoftwareComplianceResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/compliance/device-software-compliance/%s",
				cDCertificateID,
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetDeviceSoftwareCompliance()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		client := compliancetypes.NewQueryClient(grpcConn)
		resp, err := client.DeviceSoftwareCompliance(
			context.Background(),
			&compliancetypes.QueryGetDeviceSoftwareComplianceRequest{
				CDCertificateId: cDCertificateID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetDeviceSoftwareCompliance()
	}

	return &res, nil
}

const certDate = "2021-10-01T00:00:01Z"

const provDate = "2021-03-01T00:00:01Z"

func DemoTrackCompliance(suite *utils.TestSuite) {
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
	inputAllDeviceSoftwareCompliance, _ := GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, 0, len(inputAllDeviceSoftwareCompliance))

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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
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
	_, err = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)

	// Certify model
	certReason := "some reason 1"
	certDate := "2020-01-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Certify model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))

	// Check model is certified
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	modelIsCertified, _ := GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ := GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ := GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)

	// Check device software compliance
	deviceSoftwareCompliance, _ := GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.CDCertificateId)
	require.Equal(suite.T, 1, len(deviceSoftwareCompliance.ComplianceInfo))
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, deviceSoftwareCompliance.ComplianceInfo[0].CertificationType)
	require.Equal(suite.T, uint32(2), deviceSoftwareCompliance.ComplianceInfo[0].SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, deviceSoftwareCompliance.ComplianceInfo[0].Vid)
	require.Equal(suite.T, pid, deviceSoftwareCompliance.ComplianceInfo[0].Pid)
	require.Equal(suite.T, sv, deviceSoftwareCompliance.ComplianceInfo[0].SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.ComplianceInfo[0].CDCertificateId)
	require.Equal(suite.T, certReason, deviceSoftwareCompliance.ComplianceInfo[0].Reason)
	require.Equal(suite.T, certDate, deviceSoftwareCompliance.ComplianceInfo[0].Date)

	// Get all models
	complianceInfos, _ := GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ := GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ := GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ := GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ := GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliance)+1, len(deviceSoftwareCompliances))

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
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is revoked
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(3), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, revocReason, complianceInfo.Reason)
	require.Equal(suite.T, revocDate, complianceInfo.Date)
	modelIsCertified, _ = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)

	// Check modek is revoked from the entity Device Software Compliance
	_, err = GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	suite.AssertNotFound(err)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels), len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels)+1, len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliance), len(deviceSoftwareCompliances))

	oldComplianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)

	updateComplianceInfoMsg := compliancetypes.MsgUpdateComplianceInfo{
		Creator:           certCenterAccount.Address,
		Vid:               vid,
		Pid:               pid,
		SoftwareVersion:   sv,
		CertificationType: dclcompltypes.ZigbeeCertificationType,
		ProgramType:       "new program type",
		Reason:            "new reason",
		ParentChild:       "child",
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&updateComplianceInfoMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	updatedComplianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)

	// updated fields
	require.Equal(suite.T, updatedComplianceInfo.ProgramType, updateComplianceInfoMsg.ProgramType)
	require.Equal(suite.T, updatedComplianceInfo.Reason, updateComplianceInfoMsg.Reason)
	require.Equal(suite.T, updatedComplianceInfo.ParentChild, updateComplianceInfoMsg.ParentChild)

	// not updated fields
	require.Equal(suite.T, updatedComplianceInfo.CDCertificateId, oldComplianceInfo.CDCertificateId)
	require.Equal(suite.T, updatedComplianceInfo.CDVersionNumber, oldComplianceInfo.CDVersionNumber)
	require.Equal(suite.T, updatedComplianceInfo.CertificationIdOfSoftwareComponent, oldComplianceInfo.CertificationIdOfSoftwareComponent)
	require.Equal(suite.T, updatedComplianceInfo.Date, oldComplianceInfo.Date)
	require.Equal(suite.T, updatedComplianceInfo.SoftwareVersionCertificationStatus, oldComplianceInfo.SoftwareVersionCertificationStatus)

	// Publish model info
	pid = int32(tmrand.Uint16())
	secondModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish model version
	sv = tmrand.Uint32()
	svs = utils.RandString()
	secondModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Certify model with all optional fields
	certReason = "some reason 3"
	certifyModelMsg = compliancetypes.MsgCertifyModel{
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    sv,
		SoftwareVersionString:              svs,
		CertificationDate:                  certDate,
		CertificationType:                  "zigbee",
		Reason:                             certReason,
		CDCertificateId:                    testconstants.CDCertificateID,
		ProgramTypeVersion:                 testconstants.ProgramTypeVersion,
		FamilyId:                           testconstants.FamilyID,
		SupportedClusters:                  testconstants.SupportedClusters,
		CompliantPlatformUsed:              testconstants.CompliantPlatformUsed,
		CompliantPlatformVersion:           testconstants.CompliantPlatformVersion,
		OSVersion:                          testconstants.OSVersion,
		CertificationRoute:                 testconstants.CertificationRoute,
		ProgramType:                        testconstants.ProgramType,
		Transport:                          testconstants.Transport,
		ParentChild:                        testconstants.ParentChild1,
		CertificationIdOfSoftwareComponent: testconstants.CertificationIDOfSoftwareComponent,
		CDVersionNumber:                    uint32(testconstants.CdVersionNumber),
		Signer:                             certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)

	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	require.Equal(suite.T, testconstants.ProgramTypeVersion, complianceInfo.ProgramTypeVersion)
	require.Equal(suite.T, testconstants.FamilyID, complianceInfo.FamilyId)
	require.Equal(suite.T, testconstants.SupportedClusters, complianceInfo.SupportedClusters)
	require.Equal(suite.T, testconstants.CompliantPlatformUsed, complianceInfo.CompliantPlatformUsed)
	require.Equal(suite.T, testconstants.CompliantPlatformVersion, complianceInfo.CompliantPlatformVersion)
	require.Equal(suite.T, testconstants.OSVersion, complianceInfo.OSVersion)
	require.Equal(suite.T, testconstants.CertificationRoute, complianceInfo.CertificationRoute)
	require.Equal(suite.T, testconstants.Transport, complianceInfo.Transport)
	require.Equal(suite.T, testconstants.ParentChild1, complianceInfo.ParentChild)
	require.Equal(suite.T, testconstants.CertificationIDOfSoftwareComponent, complianceInfo.CertificationIdOfSoftwareComponent)

	modelIsCertified, _ = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)

	// Check Device Software Compliance
	deviceSoftwareCompliance, _ = GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.CDCertificateId)
	require.Equal(suite.T, 1, len(deviceSoftwareCompliance.ComplianceInfo))
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, deviceSoftwareCompliance.ComplianceInfo[0].CertificationType)
	require.Equal(suite.T, uint32(2), deviceSoftwareCompliance.ComplianceInfo[0].SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, deviceSoftwareCompliance.ComplianceInfo[0].Vid)
	require.Equal(suite.T, pid, deviceSoftwareCompliance.ComplianceInfo[0].Pid)
	require.Equal(suite.T, sv, deviceSoftwareCompliance.ComplianceInfo[0].SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.ComplianceInfo[0].CDCertificateId)
	require.Equal(suite.T, certReason, deviceSoftwareCompliance.ComplianceInfo[0].Reason)
	require.Equal(suite.T, certDate, deviceSoftwareCompliance.ComplianceInfo[0].Date)
	require.Equal(suite.T, testconstants.ProgramTypeVersion, deviceSoftwareCompliance.ComplianceInfo[0].ProgramTypeVersion)
	require.Equal(suite.T, testconstants.FamilyID, deviceSoftwareCompliance.ComplianceInfo[0].FamilyId)
	require.Equal(suite.T, testconstants.SupportedClusters, deviceSoftwareCompliance.ComplianceInfo[0].SupportedClusters)
	require.Equal(suite.T, testconstants.CompliantPlatformUsed, deviceSoftwareCompliance.ComplianceInfo[0].CompliantPlatformUsed)
	require.Equal(suite.T, testconstants.CompliantPlatformVersion, deviceSoftwareCompliance.ComplianceInfo[0].CompliantPlatformVersion)
	require.Equal(suite.T, testconstants.OSVersion, deviceSoftwareCompliance.ComplianceInfo[0].OSVersion)
	require.Equal(suite.T, testconstants.CertificationRoute, deviceSoftwareCompliance.ComplianceInfo[0].CertificationRoute)
	require.Equal(suite.T, testconstants.Transport, deviceSoftwareCompliance.ComplianceInfo[0].Transport)
	require.Equal(suite.T, testconstants.ParentChild1, deviceSoftwareCompliance.ComplianceInfo[0].ParentChild)
	require.Equal(suite.T, testconstants.CertificationIDOfSoftwareComponent, deviceSoftwareCompliance.ComplianceInfo[0].CertificationIdOfSoftwareComponent)

	// Get all models
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+2, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels)+1, len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliance)+1, len(deviceSoftwareCompliances))
}

func DemoTrackRevocation(suite *utils.TestSuite) {
	inputAllComplianceInfo, _ := GetAllComplianceInfo(suite)
	inputAllCertifiedModels, _ := GetAllCertifiedModels(suite)
	inputAllRevokedModels, _ := GetAllRevokedModels(suite)
	inputAllProvisionalModels, _ := GetAllProvisionalModels(suite)
	inputAllDeviceSoftwareCompliances, _ := GetAllDeviceSoftwareCompliance(suite)

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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
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
	revocReason := "some reason 4"
	revocDate := "2020-03-01T00:00:01Z"
	revocModelMsg := compliancetypes.MsgRevokeModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		RevocationDate:        revocDate,
		CertificationType:     "zigbee",
		Reason:                revocReason,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Revoke model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyRevoked.Is(err))

	// Check non-certified model is revoked
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(3), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, revocReason, complianceInfo.Reason)
	require.Equal(suite.T, revocDate, complianceInfo.Date)
	_, err = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	modelIsRevoked, _ := GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsRevoked.Value)
	_, err = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
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
	deviceSoftwareCompliances, _ := GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))

	// Certify model
	certReason := "some reason 5"
	certDate := "2020-05-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Check Device Software Compliance
	deviceSoftwareCompliance, _ := GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.CDCertificateId)
	require.Equal(suite.T, 2, len(deviceSoftwareCompliance.ComplianceInfo))
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, deviceSoftwareCompliance.ComplianceInfo[1].CertificationType)
	require.Equal(suite.T, uint32(2), deviceSoftwareCompliance.ComplianceInfo[1].SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, deviceSoftwareCompliance.ComplianceInfo[1].Vid)
	require.Equal(suite.T, pid, deviceSoftwareCompliance.ComplianceInfo[1].Pid)
	require.Equal(suite.T, sv, deviceSoftwareCompliance.ComplianceInfo[1].SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.ComplianceInfo[1].CDCertificateId)
	require.Equal(suite.T, certReason, deviceSoftwareCompliance.ComplianceInfo[1].Reason)
	require.Equal(suite.T, certDate, deviceSoftwareCompliance.ComplianceInfo[1].Date)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))
}

func DemoTrackProvision(suite *utils.TestSuite) {
	inputAllComplianceInfo, _ := GetAllComplianceInfo(suite)
	inputAllCertifiedModels, _ := GetAllCertifiedModels(suite)
	inputAllRevokedModels, _ := GetAllRevokedModels(suite)
	inputAllProvisionalModels, _ := GetAllProvisionalModels(suite)
	inputAllDeviceSoftwareCompliances, _ := GetAllDeviceSoftwareCompliance(suite)

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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
	)
	require.NotNil(suite.T, certCenterAccount)

	pid := int32(tmrand.Uint16())
	sv := tmrand.Uint32()
	svs := utils.RandString()

	firstModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Provision non-existent model
	provReason := "some reason 6"
	provModelMsg := compliancetypes.MsgProvisionModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		ProvisionalDate:       provDate,
		CertificationType:     "matter",
		Reason:                provReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Provision model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyProvisional.Is(err))

	// Check non-existent model is provisioned
	complianceInfo, _ := GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(1), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, provReason, complianceInfo.Reason)
	require.Equal(suite.T, provDate, complianceInfo.Date)
	_, err = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	provisionModel, _ := GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
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
	deviceSoftwareCompliances, _ := GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))

	// Certify model
	certReason := "some reason 7"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "matter",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Check Device Software Compliance
	deviceSoftwareCompliance, _ := GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.CDCertificateId)
	require.Equal(suite.T, 3, len(deviceSoftwareCompliance.ComplianceInfo))
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, deviceSoftwareCompliance.ComplianceInfo[2].CertificationType)
	require.Equal(suite.T, uint32(2), deviceSoftwareCompliance.ComplianceInfo[2].SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, deviceSoftwareCompliance.ComplianceInfo[2].Vid)
	require.Equal(suite.T, pid, deviceSoftwareCompliance.ComplianceInfo[2].Pid)
	require.Equal(suite.T, sv, deviceSoftwareCompliance.ComplianceInfo[2].SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.ComplianceInfo[2].CDCertificateId)
	require.Equal(suite.T, certReason, deviceSoftwareCompliance.ComplianceInfo[2].Reason)
	require.Equal(suite.T, certDate, deviceSoftwareCompliance.ComplianceInfo[2].Date)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+1, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))

	// Can not provision certified model
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))

	pid = int32(tmrand.Uint16())
	sv = tmrand.Uint32()
	svs = utils.RandString()

	secondModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	secondModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{secondModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Provision non-existent model with all optional fields
	provReason = "some reason 8"
	provModelMsg = compliancetypes.MsgProvisionModel{
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    sv,
		SoftwareVersionString:              svs,
		ProvisionalDate:                    provDate,
		CertificationType:                  "matter",
		Reason:                             provReason,
		CDCertificateId:                    testconstants.CDCertificateID,
		ProgramTypeVersion:                 testconstants.ProgramTypeVersion,
		FamilyId:                           testconstants.FamilyID,
		SupportedClusters:                  testconstants.SupportedClusters,
		CompliantPlatformUsed:              testconstants.CompliantPlatformUsed,
		CompliantPlatformVersion:           testconstants.CompliantPlatformVersion,
		OSVersion:                          testconstants.OSVersion,
		CertificationRoute:                 testconstants.CertificationRoute,
		ProgramType:                        testconstants.ProgramType,
		Transport:                          testconstants.Transport,
		ParentChild:                        testconstants.ParentChild1,
		CertificationIdOfSoftwareComponent: testconstants.CertificationIDOfSoftwareComponent,
		CDVersionNumber:                    uint32(testconstants.CdVersionNumber),
		Signer:                             certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check non-existed model is provisioned
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)

	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(1), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, provReason, complianceInfo.Reason)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, provDate, complianceInfo.Date)
	require.Equal(suite.T, testconstants.ProgramTypeVersion, complianceInfo.ProgramTypeVersion)
	require.Equal(suite.T, testconstants.FamilyID, complianceInfo.FamilyId)
	require.Equal(suite.T, testconstants.SupportedClusters, complianceInfo.SupportedClusters)
	require.Equal(suite.T, testconstants.CompliantPlatformUsed, complianceInfo.CompliantPlatformUsed)
	require.Equal(suite.T, testconstants.CompliantPlatformVersion, complianceInfo.CompliantPlatformVersion)
	require.Equal(suite.T, testconstants.OSVersion, complianceInfo.OSVersion)
	require.Equal(suite.T, testconstants.CertificationRoute, complianceInfo.CertificationRoute)
	require.Equal(suite.T, testconstants.Transport, complianceInfo.Transport)
	require.Equal(suite.T, testconstants.ParentChild1, complianceInfo.ParentChild)
	require.Equal(suite.T, testconstants.CertificationIDOfSoftwareComponent, complianceInfo.CertificationIdOfSoftwareComponent)

	_, err = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	provisionModel, _ = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.True(suite.T, provisionModel.Value)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+2, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+1, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels)+1, len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))

	// Certify model with some optional fields
	certReason = "some reason 9"
	certifyModelMsg = compliancetypes.MsgCertifyModel{
		Vid:                                vid,
		Pid:                                pid,
		SoftwareVersion:                    sv,
		SoftwareVersionString:              svs,
		CertificationDate:                  certDate,
		CertificationType:                  "matter",
		Reason:                             certReason,
		CDCertificateId:                    testconstants.CDCertificateID,
		ProgramTypeVersion:                 "pTypeVersion",
		FamilyId:                           "familyID",
		SupportedClusters:                  "sClusters",
		CompliantPlatformUsed:              "WIFI",
		CompliantPlatformVersion:           "V1",
		CertificationIdOfSoftwareComponent: "x5732",
		CDVersionNumber:                    uint32(testconstants.CdVersionNumber),
		Signer:                             certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfo(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)

	// After certify model tx some fields will be update and another fields should be no change
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	require.Equal(suite.T, "pTypeVersion", complianceInfo.ProgramTypeVersion)
	require.Equal(suite.T, "familyID", complianceInfo.FamilyId)
	require.Equal(suite.T, "sClusters", complianceInfo.SupportedClusters)
	require.Equal(suite.T, "WIFI", complianceInfo.CompliantPlatformUsed)
	require.Equal(suite.T, "V1", complianceInfo.CompliantPlatformVersion)
	require.Equal(suite.T, "x5732", complianceInfo.CertificationIdOfSoftwareComponent)
	require.Equal(suite.T, testconstants.OSVersion, complianceInfo.OSVersion)
	require.Equal(suite.T, testconstants.CertificationRoute, complianceInfo.CertificationRoute)
	require.Equal(suite.T, testconstants.Transport, complianceInfo.Transport)
	require.Equal(suite.T, testconstants.ParentChild1, complianceInfo.ParentChild)

	certifiedModel, _ = GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ = GetRevokedModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ = GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Check Device Software Compliance
	deviceSoftwareCompliance, _ = GetDeviceSoftwareCompliance(suite, testconstants.CDCertificateID)
	require.Equal(suite.T, 4, len(deviceSoftwareCompliance.ComplianceInfo))
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.CDCertificateId)
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, deviceSoftwareCompliance.ComplianceInfo[3].CertificationType)
	require.Equal(suite.T, uint32(2), deviceSoftwareCompliance.ComplianceInfo[3].SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, deviceSoftwareCompliance.ComplianceInfo[3].Vid)
	require.Equal(suite.T, pid, deviceSoftwareCompliance.ComplianceInfo[3].Pid)
	require.Equal(suite.T, sv, deviceSoftwareCompliance.ComplianceInfo[3].SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, deviceSoftwareCompliance.ComplianceInfo[3].CDCertificateId)
	require.Equal(suite.T, certReason, deviceSoftwareCompliance.ComplianceInfo[3].Reason)
	require.Equal(suite.T, certDate, deviceSoftwareCompliance.ComplianceInfo[3].Date)
	require.Equal(suite.T, "pTypeVersion", deviceSoftwareCompliance.ComplianceInfo[3].ProgramTypeVersion)
	require.Equal(suite.T, "familyID", deviceSoftwareCompliance.ComplianceInfo[3].FamilyId)
	require.Equal(suite.T, "sClusters", deviceSoftwareCompliance.ComplianceInfo[3].SupportedClusters)
	require.Equal(suite.T, "WIFI", deviceSoftwareCompliance.ComplianceInfo[3].CompliantPlatformUsed)
	require.Equal(suite.T, "V1", deviceSoftwareCompliance.ComplianceInfo[3].CompliantPlatformVersion)
	require.Equal(suite.T, "x5732", deviceSoftwareCompliance.ComplianceInfo[3].CertificationIdOfSoftwareComponent)
	require.Equal(suite.T, testconstants.OSVersion, deviceSoftwareCompliance.ComplianceInfo[3].OSVersion)
	require.Equal(suite.T, testconstants.CertificationRoute, deviceSoftwareCompliance.ComplianceInfo[3].CertificationRoute)
	require.Equal(suite.T, testconstants.Transport, deviceSoftwareCompliance.ComplianceInfo[3].Transport)
	require.Equal(suite.T, testconstants.ParentChild1, deviceSoftwareCompliance.ComplianceInfo[3].ParentChild)

	// Get all
	complianceInfos, _ = GetAllComplianceInfo(suite)
	require.Equal(suite.T, len(inputAllComplianceInfo)+2, len(complianceInfos))
	certifiedModels, _ = GetAllCertifiedModels(suite)
	require.Equal(suite.T, len(inputAllCertifiedModels)+2, len(certifiedModels))
	revokedModels, _ = GetAllRevokedModels(suite)
	require.Equal(suite.T, len(inputAllRevokedModels), len(revokedModels))
	provisionalModels, _ = GetAllProvisionalModels(suite)
	require.Equal(suite.T, len(inputAllProvisionalModels), len(provisionalModels))
	deviceSoftwareCompliances, _ = GetAllDeviceSoftwareCompliance(suite)
	require.Equal(suite.T, len(inputAllDeviceSoftwareCompliances), len(deviceSoftwareCompliances))
}

func DemoTrackComplianceWithHexVidAndPid(suite *utils.TestSuite) {
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
	var vid int32 = 0xA13
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
	)
	require.NotNil(suite.T, certCenterAccount)

	// Publish model info
	var pid int32 = 0xA11
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
	_, err = GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID1String, testconstants.TestPID1String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModelByHexVidAndPid(suite, testconstants.TestVID1String, testconstants.TestPID1String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetCertifiedModelByHexVidAndPid(suite, testconstants.TestVID1String, testconstants.TestPID1String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	_, err = GetProvisionalModelByHexVidAndPid(suite, testconstants.TestVID1String, testconstants.TestPID1String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)

	// Certify model
	certReason := "some reason 10"
	certDate := "2020-01-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Certify model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))

	// Check model is certified
	complianceInfo, _ := GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID1String, testconstants.TestPID1String, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	modelIsCertified, _ := GetCertifiedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsCertified.Value)
	modelIsRevoked, _ := GetRevokedModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsRevoked.Value)
	modelIsProvisional, _ := GetProvisionalModel(suite, vid, pid, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, modelIsProvisional.Value)
}

func DemoTrackRevocationWithHexVidAndPid(suite *utils.TestSuite) {
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
	var vid int32 = 0xA14
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
	)
	require.NotNil(suite.T, certCenterAccount)

	// Publish model info
	var pid int32 = 0xA15
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
	revocReason := "some reason 11"
	revocDate := "2020-03-01T00:00:01Z"
	revocModelMsg := compliancetypes.MsgRevokeModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		RevocationDate:        revocDate,
		CertificationType:     "zigbee",
		Reason:                revocReason,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Revoke model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&revocModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyRevoked.Is(err))

	// Check non-certified model is revoked
	complianceInfo, _ := GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(3), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, revocReason, complianceInfo.Reason)
	require.Equal(suite.T, revocDate, complianceInfo.Date)
	_, err = GetCertifiedModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)
	modelIsRevoked, _ := GetRevokedModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, modelIsRevoked.Value)
	_, err = GetProvisionalModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	suite.AssertNotFound(err)

	// Certify model
	certReason := "some reason 12"
	certDate := "2020-05-01T00:00:01Z"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "zigbee",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.Equal(suite.T, dclcompltypes.ZigbeeCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModelByHexVidAndPid(suite, testconstants.TestVID2String, testconstants.TestPID2String, sv, dclcompltypes.ZigbeeCertificationType)
	require.False(suite.T, provisionalModel.Value)
}

func DemoTrackProvisionByHexVidAndPid(suite *utils.TestSuite) {
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
	var vid int32 = 0xA16
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
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
		testconstants.Info,
	)
	require.NotNil(suite.T, certCenterAccount)

	var pid int32 = 0xA17
	sv := tmrand.Uint32()
	svs := utils.RandString()

	// Publish model info
	firstModel := test_model.NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModel}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Publish modelVersion
	firstModelVersion := test_model.NewMsgCreateModelVersion(vid, pid, sv, svs, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{firstModelVersion}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Provision non-existent model
	provReason := "some reason 13"
	provModelMsg := compliancetypes.MsgProvisionModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		ProvisionalDate:       provDate,
		CertificationType:     "matter",
		Reason:                provReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Provision model again
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyProvisional.Is(err))

	// Check non-existent model is provisioned
	complianceInfo, _ := GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(1), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, provReason, complianceInfo.Reason)
	require.Equal(suite.T, provDate, complianceInfo.Date)
	_, err = GetCertifiedModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	_, err = GetRevokedModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	suite.AssertNotFound(err)
	provisionModel, _ := GetProvisionalModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.True(suite.T, provisionModel.Value)

	// Certify model
	certReason := "some reason 14"
	certifyModelMsg := compliancetypes.MsgCertifyModel{
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       sv,
		SoftwareVersionString: svs,
		CertificationDate:     certDate,
		CertificationType:     "matter",
		Reason:                certReason,
		CDCertificateId:       testconstants.CDCertificateID,
		CDVersionNumber:       uint32(testconstants.CdVersionNumber),
		Signer:                certCenterAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&certifyModelMsg}, certCenter, certCenterAccount)
	require.NoError(suite.T, err)

	// Check model is certified
	complianceInfo, _ = GetComplianceInfoByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.Equal(suite.T, dclcompltypes.MatterCertificationType, complianceInfo.CertificationType)
	require.Equal(suite.T, uint32(2), complianceInfo.SoftwareVersionCertificationStatus)
	require.Equal(suite.T, vid, complianceInfo.Vid)
	require.Equal(suite.T, pid, complianceInfo.Pid)
	require.Equal(suite.T, sv, complianceInfo.SoftwareVersion)
	require.Equal(suite.T, testconstants.CDCertificateID, complianceInfo.CDCertificateId)
	require.Equal(suite.T, certReason, complianceInfo.Reason)
	require.Equal(suite.T, certDate, complianceInfo.Date)
	certifiedModel, _ := GetCertifiedModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.True(suite.T, certifiedModel.Value)
	revokedModel, _ := GetRevokedModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, revokedModel.Value)
	provisionalModel, _ := GetProvisionalModelByHexVidAndPid(suite, testconstants.TestVID3String, testconstants.TestPID3String, sv, dclcompltypes.MatterCertificationType)
	require.False(suite.T, provisionalModel.Value)

	// Can not provision certified model
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&provModelMsg}, certCenter, certCenterAccount)
	require.Error(suite.T, err)
	require.True(suite.T, compliancetypes.ErrAlreadyCertified.Is(err))
}
