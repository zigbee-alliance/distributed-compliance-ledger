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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func NewMsgCreateModel(vid int32, pid int32, signer string) *modeltypes.MsgCreateModel {
	return &modeltypes.MsgCreateModel{
		Creator:                                  signer,
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             testconstants.DeviceTypeID,
		ProductName:                              utils.RandString(),
		ProductLabel:                             utils.RandString(),
		PartNumber:                               utils.RandString(),
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: testconstants.UserManualURL,
		SupportUrl:    testconstants.SupportURL,
		ProductUrl:    testconstants.ProductURL,
		LsfUrl:        testconstants.LsfURL,
	}
}

func NewMsgUpdateModel(vid int32, pid int32, signer string) *modeltypes.MsgUpdateModel {
	return &modeltypes.MsgUpdateModel{
		Creator:                    signer,
		Vid:                        vid,
		Pid:                        pid,
		ProductLabel:               utils.RandString(),
		CommissioningCustomFlowUrl: testconstants.CommissioningCustomFlowURL + "/new",
		UserManualUrl:              testconstants.UserManualURL + "/new",
		SupportUrl:                 testconstants.SupportURL + "/new",
		ProductUrl:                 testconstants.ProductURL + "/new",
		LsfUrl:                     testconstants.LsfURL + "/new",
		LsfRevision:                testconstants.LsfRevision + 1,
	}
}

func NewMsgDeleteModel(vid int32, pid int32, signer string) *modeltypes.MsgDeleteModel {
	return &modeltypes.MsgDeleteModel{
		Creator: signer,
		Vid:     vid,
		Pid:     pid,
	}
}

func NewMsgCreateModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer string,
) *modeltypes.MsgCreateModelVersion {
	return &modeltypes.MsgCreateModelVersion{
		Creator:                      signer,
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareInformation:          testconstants.FirmwareInformation,
		SoftwareVersionValid:         true,
		OtaUrl:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL,
	}
}

func NewMsgUpdateModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	signer string,
) *modeltypes.MsgUpdateModelVersion {
	return &modeltypes.MsgUpdateModelVersion{
		Creator:         signer,
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
		OtaUrl:          testconstants.OtaURL + "/new",
		ReleaseNotesUrl: testconstants.ReleaseNotesURL + "/new",
	}
}

func NewMsgCertifyModelVersion(
	signer string,
	vid int32,
	pid int32,
) *dclcompltypes.MsgCertifyModel {
	return &dclcompltypes.MsgCertifyModel{
		Signer:            signer,
		Vid:               vid,
		Pid:               pid,
		CertificationDate: testconstants.CertificationDate,
		CDCertificateId:   testconstants.CDCertificateID,
		CertificationType: testconstants.CertificationType,
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

func GetModelByHexVidPid(
	suite *utils.TestSuite,
	vid string,
	pid string,
) (*modeltypes.Model, error) {
	var res modeltypes.Model

	if suite.Rest {
		var resp modeltypes.QueryGetModelResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/model/models/%s/%s", vid, pid), &resp)
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

func GetVendorModelsByHexVid(
	suite *utils.TestSuite,
	vid string,
) (*modeltypes.VendorProducts, error) {
	var res modeltypes.VendorProducts

	if suite.Rest {
		var resp modeltypes.QueryGetVendorProductsResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/model/models/%s", vid), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetVendorProducts()
	}

	return &res, nil
}

func DeleteModelWithAssociatedModelVersions(suite *utils.TestSuite) {
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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// New vendor adds a model
	pid := int32(tmrand.Uint16())
	createModelMsg := NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	createModelVersionMsg1 := NewMsgCreateModelVersion(vid, pid, 1, "1", vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelVersionMsg1}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	createModelVersionMsg2 := NewMsgCreateModelVersion(vid, pid, 2, "2", vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelVersionMsg2}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	deleteModelMsg := NewMsgDeleteModel(vid, pid, vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{deleteModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// check if model is deleted
	model, err := GetModel(suite, deleteModelMsg.Vid, deleteModelMsg.Pid)
	require.Error(suite.T, err)
	require.Nil(suite.T, model)

	// check if model version 1 is deleted
	modelVersion1, err := GetModelVersion(suite, createModelVersionMsg1.Vid, createModelVersionMsg1.Pid, createModelVersionMsg1.SoftwareVersion)
	require.Error(suite.T, err)
	require.Nil(suite.T, modelVersion1)

	// check if model version 2 is deleted
	modelVersion2, err := GetModelVersion(suite, createModelVersionMsg2.Vid, createModelVersionMsg2.Pid, createModelVersionMsg2.SoftwareVersion)
	require.Error(suite.T, err)
	require.Nil(suite.T, modelVersion2)
}

func DeleteModelWithAssociatedModelVersionsCertified(suite *utils.TestSuite) {
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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// New vendor adds a model
	pid := int32(tmrand.Uint16())
	createModelMsg := NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	createModelVersionMsg1 := NewMsgCreateModelVersion(vid, pid, 1, "1", vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelVersionMsg1}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	createModelVersionMsg2 := NewMsgCreateModelVersion(vid, pid, 2, "2", vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelVersionMsg2}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// certify model version
	certifyModelVersionMsg := NewMsgCertifyModelVersion(aliceName, vid, pid)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{certifyModelVersionMsg}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	deleteModelMsg := NewMsgDeleteModel(vid, pid, vendorName)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{deleteModelMsg}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// check if model is not deleted
	model, err := GetModel(suite, deleteModelMsg.Vid, deleteModelMsg.Pid)
	require.NoError(suite.T, err)
	require.NotNil(suite.T, model)

	// check if model version 1 is deleted
	modelVersion1, err := GetModelVersion(suite, createModelVersionMsg1.Vid, createModelVersionMsg1.Pid, createModelVersionMsg1.SoftwareVersion)
	require.NoError(suite.T, err)
	require.NotNil(suite.T, modelVersion1)

	// check if model version 2 is deleted
	modelVersion2, err := GetModelVersion(suite, createModelVersionMsg2.Vid, createModelVersionMsg2.Pid, createModelVersionMsg2.SoftwareVersion)
	require.NoError(suite.T, err)
	require.NotNil(suite.T, modelVersion2)
}

func Demo(suite *utils.TestSuite) {
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
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Get all models
	inputModels, err := GetModels(suite)
	require.NoError(suite.T, err)

	// New vendor adds first model
	pid1 := int32(tmrand.Uint16())
	createFirstModelMsg := NewMsgCreateModel(vid, pid1, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createFirstModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check first model is added
	receivedModel, err := GetModel(suite, createFirstModelMsg.Vid, createFirstModelMsg.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, createFirstModelMsg.Vid, receivedModel.Vid)
	require.Equal(suite.T, createFirstModelMsg.Pid, receivedModel.Pid)
	require.Equal(suite.T, createFirstModelMsg.ProductName, receivedModel.ProductName)
	require.Equal(suite.T, createFirstModelMsg.ProductLabel, receivedModel.ProductLabel)

	// Add second model
	pid2 := int32(tmrand.Uint16())
	createSecondModelMsg := NewMsgCreateModel(vid, pid2, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createSecondModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check second model is added
	receivedModel, err = GetModel(suite, createSecondModelMsg.Vid, createSecondModelMsg.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, createSecondModelMsg.Vid, receivedModel.Vid)
	require.Equal(suite.T, createSecondModelMsg.Pid, receivedModel.Pid)
	require.Equal(suite.T, createSecondModelMsg.ProductName, receivedModel.ProductName)
	require.Equal(suite.T, createSecondModelMsg.ProductLabel, receivedModel.ProductLabel)

	// Get all models
	receivedModels, err := GetModels(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputModels)+2, len(receivedModels))

	// Get models of new vendor
	vendorModels, err := GetVendorModels(suite, vid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 2, len(vendorModels.Products))
	require.Equal(suite.T, createFirstModelMsg.Pid, vendorModels.Products[0].Pid)
	require.Equal(suite.T, createSecondModelMsg.Pid, vendorModels.Products[1].Pid)

	// Update second model
	updateSecondModelMsg := NewMsgUpdateModel(createSecondModelMsg.Vid, createSecondModelMsg.Pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{updateSecondModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check second model is updated
	receivedModel, err = GetModel(suite, createSecondModelMsg.Vid, createSecondModelMsg.Pid)
	require.NoError(suite.T, err)
	require.Equal(suite.T, updateSecondModelMsg.ProductLabel, receivedModel.ProductLabel)
}

/* Error cases */

func AddModelByNonVendor(suite *utils.TestSuite) {
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

	// register new account without Vendor role
	nonVendorAccountName := utils.RandString()
	vid := int32(tmrand.Uint16())
	nonVendorAccount := test_dclauth.CreateAccount(
		suite,
		nonVendorAccountName,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	require.NotContains(suite.T, nonVendorAccount.Roles, dclauthtypes.Vendor)

	// try to add createModelMsg
	pid := int32(tmrand.Uint16())
	createModelMsg := NewMsgCreateModel(vid, pid, nonVendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelMsg}, nonVendorAccountName, nonVendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddModelByDifferentVendor(suite *utils.TestSuite) {
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

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid+1,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// try to add createModelMsg
	pid := int32(tmrand.Uint16())
	createModelMsg := NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelMsg}, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, sdkerrors.ErrUnauthorized.Is(err))
}

func AddModelTwice(suite *utils.TestSuite) {
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

	// register new Vendor account
	vendorName := utils.RandString()
	vid := int32(tmrand.Uint16())
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)

	// add model
	pid := int32(tmrand.Uint16())
	createModelMsg := NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// add the same model second time
	_, err = AddModel(suite, createModelMsg, vendorName, vendorAccount)
	require.Error(suite.T, err)
	require.True(suite.T, modeltypes.ErrModelAlreadyExists.Is(err))
}

func GetModelForUnknown(suite *utils.TestSuite) {
	_, err := GetModel(suite, int32(tmrand.Uint16()), int32(tmrand.Uint16()))
	require.Error(suite.T, err)
	suite.AssertNotFound(err)

	_, err = GetModelVersion(suite, int32(tmrand.Uint16()), int32(tmrand.Uint16()), tmrand.Uint32())
	require.Error(suite.T, err)
	suite.AssertNotFound(err)

	_, err = GetVendorModels(suite, int32(tmrand.Uint16()))
	require.Error(suite.T, err)
	suite.AssertNotFound(err)
}

func GetModelForInvalidVidPid(suite *utils.TestSuite) {
	// negative vid
	_, err := GetModel(suite, -1, int32(tmrand.Uint16()))
	require.Error(suite.T, err)
	// FIXME: Consider adding validation for queries.
	// require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
	suite.AssertNotFound(err)

	// negative pid
	_, err = GetModel(suite, int32(tmrand.Uint16()), -1)
	require.Error(suite.T, err)
	// FIXME: Consider adding validation for queries.
	// require.True(suite.T, sdkerrors.ErrInvalidRequest.Is(err))
	suite.AssertNotFound(err)
}

func DemoWithHexVidAndPid(suite *utils.TestSuite) {
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
	vendorName := utils.RandString()
	var vid int32 = 0xA13
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vid,
		aliceName,
		aliceAccount,
		bobName,
		bobAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Get all models
	inputModels, err := GetModels(suite)
	require.NoError(suite.T, err)

	var pid int32 = 0xA11

	// New vendor adds first model
	createFirstModelMsg := NewMsgCreateModel(vid, pid, vendorAccount.Address)
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{createFirstModelMsg}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check first model is added
	receivedModel, err := GetModelByHexVidPid(suite, testconstants.TestVID1String, testconstants.TestPID1String)
	require.NoError(suite.T, err)
	require.Equal(suite.T, createFirstModelMsg.Vid, receivedModel.Vid)
	require.Equal(suite.T, createFirstModelMsg.Pid, receivedModel.Pid)
	require.Equal(suite.T, createFirstModelMsg.ProductName, receivedModel.ProductName)
	require.Equal(suite.T, createFirstModelMsg.ProductLabel, receivedModel.ProductLabel)

	// Get all models
	receivedModels, err := GetModels(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputModels)+1, len(receivedModels))

	// Get models of new vendor
	vendorModels, err := GetVendorModelsByHexVid(suite, testconstants.TestVID1String)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 1, len(vendorModels.Products))
	require.Equal(suite.T, createFirstModelMsg.Pid, vendorModels.Products[0].Pid)
}
