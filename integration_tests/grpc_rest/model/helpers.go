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

package model

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func NewMsgCreateModel(vid int32, pid int32) *modeltypes.MsgCreateModel {
	return &modeltypes.MsgCreateModel{
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             testconstants.DeviceTypeId,
		ProductName:                              utils.RandString(),
		ProductLabel:                             utils.RandString(),
		PartNumber:                               utils.RandString(),
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: testconstants.UserManualUrl,
		SupportUrl:    testconstants.SupportUrl,
		ProductUrl:    testconstants.ProductUrl,
	}
}

func NewMsgUpdateModel(vid int32, pid int32) *modeltypes.MsgUpdateModel {
	return &modeltypes.MsgUpdateModel{
		Vid:                        vid,
		Pid:                        pid,
		ProductLabel:               utils.RandString(),
		CommissioningCustomFlowUrl: testconstants.CommissioningCustomFlowUrl + "/new",
		UserManualUrl:              testconstants.UserManualUrl + "/new",
		SupportUrl:                 testconstants.SupportUrl + "/new",
		ProductUrl:                 testconstants.ProductUrl + "/new",
	}
}

func NewMsgCreateModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
) *modeltypes.MsgCreateModelVersion {
	return &modeltypes.MsgCreateModelVersion{
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         true,
		OtaUrl:                       testconstants.OtaUrl,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl,
	}
}

func NewMsgUpdateModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
) *modeltypes.MsgUpdateModelVersion {
	return &modeltypes.MsgUpdateModelVersion{
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
		OtaUrl:          testconstants.OtaUrl + "/new",
		ReleaseNotesUrl: testconstants.ReleaseNotesUrl + "/new",
	}
}

func AddModel(
	suite *utils.TestSuite,
	msg *modeltypes.MsgCreateModel,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func AddModelVersion(
	suite *utils.TestSuite,
	msg *modeltypes.MsgCreateModelVersion,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func UpdateModel(
	suite *utils.TestSuite,
	msg *modeltypes.MsgUpdateModel,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func UpdateModelVersion(
	suite *utils.TestSuite,
	msg *modeltypes.MsgUpdateModelVersion,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {
	msg.Creator = suite.GetAddress(signerName).String()
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func GetModel(
	suite *utils.TestSuite,
	vid int32,
	pid int32,
) (*modeltypes.Model, error) {
	var res modeltypes.Model

	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp modeltypes.QueryGetModelResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/model/models/%v/%v", vid, pid), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		modelClient := modeltypes.NewQueryClient(grpcConn)
		resp, err := modelClient.Model(
			context.Background(),
			&modeltypes.QueryGetModelRequest{Vid: vid, Pid: pid},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetModel()
	}

	return &res, nil
}

func GetModelVersion(
	suite *utils.TestSuite,
	vid int32,
	pid int32,
	softwareVersion uint32,
) (*modeltypes.ModelVersion, error) {
	var res modeltypes.ModelVersion

	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp modeltypes.QueryGetModelVersionResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/model/versions/%v/%v/%v", vid, pid, softwareVersion), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetModelVersion()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		modelClient := modeltypes.NewQueryClient(grpcConn)
		resp, err := modelClient.ModelVersion(
			context.Background(),
			&modeltypes.QueryGetModelVersionRequest{Vid: vid, Pid: pid, SoftwareVersion: softwareVersion},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetModelVersion()
	}

	return &res, nil
}

func GetModels(suite *utils.TestSuite) (res []modeltypes.Model, err error) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp modeltypes.QueryAllModelResponse
		err := suite.QueryREST("/dcl/model/models", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetModel()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		modelClient := modeltypes.NewQueryClient(grpcConn)
		resp, err := modelClient.ModelAll(
			context.Background(),
			&modeltypes.QueryAllModelRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetModel()
	}

	return res, nil
}

func GetVendorModels(
	suite *utils.TestSuite,
	vid int32,
) (*modeltypes.VendorProducts, error) {
	var res modeltypes.VendorProducts

	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp modeltypes.QueryGetVendorProductsResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/model/models/%v", vid), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorProducts()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		modelClient := modeltypes.NewQueryClient(grpcConn)
		resp, err := modelClient.VendorProducts(
			context.Background(),
			&modeltypes.QueryGetVendorProductsRequest{Vid: vid},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorProducts()
	}

	return &res, nil
}

func ModelDemo(suite *utils.TestSuite) {
	// Alice and Bob are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := test_dclauth.GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Register new Vendor account
	vid := int32(tmrand.Uint16())
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		uint64(vid),
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
	)
	require.NotNil(suite.T, vendorAccount)

	// Get all models
	inputModels, err := GetModels(suite)
	require.NoError(suite.T, err)

	// New vendor adds first model
	pid1 := int32(tmrand.Uint16())
	firstModel := NewMsgCreateModel(vid, pid1)
	_, err = AddModel(suite, firstModel, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check first model is added
	receivedModel, err := GetModel(suite, firstModel.Vid, firstModel.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, firstModel.Vid, receivedModel.Vid)
	require.Equal(suite.T, firstModel.Pid, receivedModel.Pid)
	require.Equal(suite.T, firstModel.ProductName, receivedModel.ProductName)
	require.Equal(suite.T, firstModel.ProductLabel, receivedModel.ProductLabel)

	// Add second model
	pid2 := int32(tmrand.Uint16())
	secondModel := NewMsgCreateModel(vid, pid2)
	_, err = AddModel(suite, secondModel, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check second model is added
	receivedModel, err = GetModel(suite, secondModel.Vid, secondModel.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, secondModel.Vid, receivedModel.Vid)
	require.Equal(suite.T, secondModel.Pid, receivedModel.Pid)
	require.Equal(suite.T, secondModel.ProductName, receivedModel.ProductName)
	require.Equal(suite.T, secondModel.ProductLabel, receivedModel.ProductLabel)

	// Get all models
	receivedModels, err := GetModels(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputModels)+2, len(receivedModels))

	// Get models of new vendor
	vendorModels, err := GetVendorModels(suite, vid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 2, len(vendorModels.Products))
	require.Equal(suite.T, firstModel.Pid, vendorModels.Products[0].Pid)
	require.Equal(suite.T, secondModel.Pid, vendorModels.Products[1].Pid)

	// Update second model
	secondModelUpdate := NewMsgUpdateModel(secondModel.Vid, secondModel.Pid)
	_, err = UpdateModel(suite, secondModelUpdate, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check second model is updated
	receivedModel, err = GetModel(suite, secondModel.Vid, secondModel.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, secondModelUpdate.ProductLabel, receivedModel.ProductLabel)
}

// /* Error cases */

// func Test_AddModel_ByNonVendor(suite *utils.TestSuite) {
// 	// register new account
// 	vid := common.RandUint16()
// 	testAccount := utils.CreateNewAccount(auth.AccountRoles{}, vid)

// 	// try to publish model info
// 	model := utils.NewMsgAddModel(testAccount.Address, vid)
// 	res, _ := utils.SignAndBroadcastMessage(testAccount, model)
// 	require.Equal(suite.T, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
// }

// func Test_AddModel_ByDifferentVendor(suite *utils.TestSuite) {
// 	// register new account
// 	vid := common.RandUint16()
// 	testAccount := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid+1)

// 	// try to publish model info
// 	model := utils.NewMsgAddModel(testAccount.Address, vid)
// 	res, _ := utils.SignAndBroadcastMessage(testAccount, model)
// 	require.Equal(suite.T, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
// }

// func Test_AddModel_Twice(suite *utils.TestSuite) {
// 	// register new account
// 	vid := common.RandUint16()
// 	testAccount := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid)

// 	// publish modelMsg info
// 	modelMsg := utils.NewMsgAddModel(testAccount.Address, vid)
// 	res, _ := utils.AddModel(modelMsg, testAccount)
// 	require.Equal(suite.T, sdk.CodeOK, sdk.CodeType(res.Code))

// 	// publish second time
// 	res, _ = utils.AddModel(modelMsg, testAccount)
// 	require.Equal(suite.T, model.CodeModelAlreadyExists, sdk.CodeType(res.Code))
// }

// func Test_GetModel_ForUnknown(suite *utils.TestSuite) {
// 	_, code := utils.GetModel(common.RandUint16(), common.RandUint16())
// 	require.Equal(suite.T, http.StatusNotFound, code)
// }

// func Test_GetModel_ForInvalidVidPid(suite *utils.TestSuite) {
// 	// zero vid
// 	_, code := utils.GetModel(0, common.RandUint16())
// 	require.Equal(suite.T, http.StatusBadRequest, code)

// 	// zero pid
// 	_, code = utils.GetModel(common.RandUint16(), 0)
// 	require.Equal(suite.T, http.StatusBadRequest, code)
// }
