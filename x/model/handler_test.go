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

//nolint:testpackage,lll
package model

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)
	return args.Bool(0)
}

func (m *DclauthKeeperMock) HasVendorID(
	ctx sdk.Context,
	addr sdk.AccAddress,
	vid uint64,
) bool {
	args := m.Called(ctx, addr, vid)
	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	Vendor   sdk.AccAddress
	VendorID uint64
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vendorID uint64,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)

	dclauthKeeper.On("HasVendorID", mock.Anything, accAddress, vendorID).Return(true)
	dclauthKeeper.On("HasVendorID", mock.Anything, accAddress, mock.Anything).Return(false)
}

func Setup(t *testing.T) *TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := testkeeper.ModelKeeper(t, dclauthKeeper)

	vendor := GenerateAccAddress()
	vendorID := uint64(testconstants.VendorID1)

	setup := &TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		Handler:       NewHandler(*keeper),
		Vendor:        vendor,
		VendorID:      vendorID,
	}

	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vendorID)

	return setup
}

func TestHandler_AddModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
	require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
	require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
}

func TestHandler_UpdateModel(t *testing.T) {
	setup := Setup(t)

	// try update not present model
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.True(t, types.ErrModelDoesNotExist.Is(err))

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// update existing model
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query updated model
	receivedModel, err := queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)

	// check
	// Mutable Fields ProductName,ProductLable,PartNumber,CommissioningCustomFlowUrl,
	// CommissioningModeInitialStepsInstruction,CommissioningModeSecondaryStepsInstruction,UserManualUrl,SupportUrl,SupportUrl
	require.Equal(t, msgUpdateModel.Vid, receivedModel.Vid)
	require.Equal(t, msgUpdateModel.Pid, receivedModel.Pid)
	require.Equal(t, msgUpdateModel.ProductLabel, receivedModel.ProductLabel)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// update existing model by user without Vendor role
		msgUpdateModel := NewMsgUpdateModel(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModel)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, uint64(testconstants.VendorID2))

	// update existing model by vendor with another VendorID
	msgUpdateModel := NewMsgUpdateModel(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model by owner
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)
}

func TestHandler_AddModelWithEmptyOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add new msgCreateModel
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	msgCreateModel.DeviceTypeId = 0 // Set empty CID

	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, int32(0), receivedModel.DeviceTypeId)
}

func TestHandler_AddModelByNonVendor(t *testing.T) {
	setup := Setup(t)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// add new model
		model := NewMsgCreateModel(accAddress)
		_, err := setup.Handler(setup.Ctx, model)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_AddModelByVendorWithAnotherVendorId(t *testing.T) {
	setup := Setup(t)

	anotherVendor := GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, uint64(testconstants.VendorID2))

	// add new model
	model := NewMsgCreateModel(anotherVendor)
	_, err := setup.Handler(setup.Ctx, model)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_PartiallyUpdateModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgAddModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgAddModel)
	require.NoError(t, err)

	// update Description of existing model
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.ProductName = ""
	msgUpdateModel.ProductLabel = "New Description"

	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)

	// check
	// Mutable Fields ProductName,ProductLable,PartNumber,CommissioningCustomFlowUrl,
	// CommissioningModeInitialStepsInstruction,CommissioningModeSecondaryStepsInstruction,UserManualUrl,SupportUrl,SupportUrl
	require.Equal(t, msgAddModel.ProductName, receivedModel.ProductName)
	require.Equal(t, msgUpdateModel.ProductLabel, receivedModel.ProductLabel)
}

// ----------------------------------------------------------------------------
// Model Version Tests --------------------------------------------------------
// ----------------------------------------------------------------------------

func TestHandler_AddModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgCreateModelVersion.Vid,
		msgCreateModelVersion.Pid,
		msgCreateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion.Vid, receivedModelVersion.Vid)
	require.Equal(t, msgCreateModelVersion.Pid, receivedModelVersion.Pid)
	require.Equal(t, msgCreateModelVersion.SoftwareVersion, receivedModelVersion.SoftwareVersion)
	require.Equal(t, msgCreateModelVersion.SoftwareVersionString, receivedModelVersion.SoftwareVersionString)
}

func TestHandler_UpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update existing model version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check
	// Mutable Fields SoftwareVersionValid,OtaUrl,MinApplicableSoftwareVersion,MaxApplicableSoftwareVersion,ReleaseNotesUrl
	require.Equal(t, receivedModelVersion.Vid, msgUpdateModelVersion.Vid)
	require.Equal(t, receivedModelVersion.Pid, msgUpdateModelVersion.Pid)
	require.Equal(t, receivedModelVersion.SoftwareVersion, msgUpdateModelVersion.SoftwareVersion)

	require.Equal(t, receivedModelVersion.SoftwareVersionValid, msgUpdateModelVersion.SoftwareVersionValid)
	require.Equal(t, receivedModelVersion.OtaUrl, msgUpdateModelVersion.OtaUrl)
	require.Equal(t, receivedModelVersion.MinApplicableSoftwareVersion, msgUpdateModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.MaxApplicableSoftwareVersion, msgUpdateModelVersion.MaxApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.ReleaseNotesUrl, msgUpdateModelVersion.ReleaseNotesUrl)

	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgCreateModelVersion.SoftwareVersionString)
	require.Equal(t, receivedModelVersion.CdVersionNumber, msgCreateModelVersion.CdVersionNumber)
	require.Equal(t, receivedModelVersion.FirmwareDigests, msgCreateModelVersion.FirmwareDigests)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgCreateModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.OtaChecksum, msgCreateModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgCreateModelVersion.OtaChecksumType)
}

func TestHandler_PartiallyUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	// Update only SoftwareVersionValid and ReleaseNotesUrl
	msgUpdateModelVersion.SoftwareVersionValid = !msgCreateModelVersion.SoftwareVersionValid
	msgUpdateModelVersion.OtaUrl = ""
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 0
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 0
	msgUpdateModelVersion.ReleaseNotesUrl = "https://new.releasenotes.url"

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// Mutable Fields SoftwareVersionValid,OtaUrl,MinApplicableSoftwareVersion,MaxApplicableSoftwareVersion,ReleaseNotesUrl
	require.Equal(t, msgUpdateModelVersion.Vid, receivedModelVersion.Vid)
	require.Equal(t, msgUpdateModelVersion.Pid, receivedModelVersion.Pid)
	require.Equal(t, msgUpdateModelVersion.SoftwareVersion, receivedModelVersion.SoftwareVersion)

	require.Equal(t, msgUpdateModelVersion.SoftwareVersionValid, receivedModelVersion.SoftwareVersionValid)
	require.NotEqual(t, msgCreateModelVersion.SoftwareVersionValid, receivedModelVersion.SoftwareVersionValid)

	require.Equal(t, msgUpdateModelVersion.ReleaseNotesUrl, receivedModelVersion.ReleaseNotesUrl)
	require.NotEqual(t, msgCreateModelVersion.ReleaseNotesUrl, receivedModelVersion.ReleaseNotesUrl)

	require.Equal(t, msgCreateModelVersion.OtaUrl, receivedModelVersion.OtaUrl)
	require.Equal(t, msgCreateModelVersion.MinApplicableSoftwareVersion, receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, msgCreateModelVersion.MaxApplicableSoftwareVersion, receivedModelVersion.MaxApplicableSoftwareVersion)
}

func TestHandler_OnlyOwnerCanUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreteModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreteModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// update existing model version by user without Vendor role
		msgUpdateModelVersion := NewMsgUpdateModelVersion(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, uint64(testconstants.VendorID2))

	// update existing model by vendor with another VendorID
	msgUpdateModelVersion := NewMsgUpdateModelVersion(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model version by owner
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)
}

func queryModel(
	setup *TestSetup,
	vid int32,
	pid int32,
) (*types.Model, error) {

	req := &types.QueryGetModelRequest{
		Vid: vid,
		Pid: pid,
	}

	resp, err := setup.Keeper.Model(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return &resp.Model, nil
}

func queryModelVersion(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
) (*types.ModelVersion, error) {

	req := &types.QueryGetModelVersionRequest{
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
	}

	resp, err := setup.Keeper.ModelVersion(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return &resp.ModelVersion, nil
}

func NewMsgCreateModel(signer sdk.AccAddress) *types.MsgCreateModel {
	return &types.MsgCreateModel{
		Creator:                                  signer.String(),
		Vid:                                      int32(testconstants.VendorID1),
		Pid:                                      int32(testconstants.Pid),
		DeviceTypeId:                             testconstants.DeviceTypeId,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
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

func NewMsgUpdateModel(signer sdk.AccAddress) *types.MsgUpdateModel {
	return &types.MsgUpdateModel{
		Creator:                                  signer.String(),
		Vid:                                      int32(testconstants.VendorID1),
		Pid:                                      testconstants.Pid,
		ProductName:                              testconstants.ProductName + "-updated",
		ProductLabel:                             testconstants.ProductLabel + "-updated",
		PartNumber:                               testconstants.PartNumber + "-updated",
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl + "/updated",
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction + "-updated",
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction + "-updated",
		UserManualUrl: testconstants.UserManualUrl + "/updated",
		SupportUrl:    testconstants.SupportUrl + "/updated",
		ProductUrl:    testconstants.ProductUrl + "/updated",
	}
}

func NewMsgCreateModelVersion(signer sdk.AccAddress) *types.MsgCreateModelVersion {
	return &types.MsgCreateModelVersion{
		Creator:                      signer.String(),
		Vid:                          int32(testconstants.VendorID1),
		Pid:                          testconstants.Pid,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl,
	}
}

func NewMsgUpdateModelVersion(signer sdk.AccAddress) *types.MsgUpdateModelVersion {
	return &types.MsgUpdateModelVersion{
		Creator:                      signer.String(),
		Vid:                          int32(testconstants.VendorID1),
		Pid:                          testconstants.Pid,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionValid:         !testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl + "/updated",
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion + 1,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion + 1,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl + "/updated",
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}
