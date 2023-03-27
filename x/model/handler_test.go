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

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/testdata"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	vid int32,
) bool {
	args := m.Called(ctx, addr, vid)

	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type ComplianceKeeperMock struct {
	mock.Mock
}

func (m *ComplianceKeeperMock) GetComplianceInfo(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (val dclcompltypes.ComplianceInfo, found bool) {
	args := m.Called(ctx, vid, pid, softwareVersion, certificationType)

	return val, args.Bool(0)
}

var _ types.ComplianceKeeper = &ComplianceKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx              sdk.Context
	Wctx             context.Context
	Keeper           *keeper.Keeper
	DclauthKeeper    *DclauthKeeperMock
	ComplianceKeeper *ComplianceKeeperMock
	Handler          sdk.Handler
	// Querier     sdk.Querier
	Vendor   sdk.AccAddress
	VendorID int32
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vendorID int32,
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
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	complianceKeeper := &ComplianceKeeperMock{}
	keeper, ctx := testkeeper.ModelKeeper(t, dclauthKeeper, complianceKeeper)

	vendor := testdata.GenerateAccAddress()
	vendorID := testconstants.VendorID1

	setup := &TestSetup{
		T:                t,
		Ctx:              ctx,
		Wctx:             sdk.WrapSDKContext(ctx),
		Keeper:           keeper,
		DclauthKeeper:    dclauthKeeper,
		ComplianceKeeper: complianceKeeper,
		Handler:          NewHandler(*keeper),
		Vendor:           vendor,
		VendorID:         vendorID,
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

func TestHandler_OnlyOwnerAndVendorWithSameVidCanUpdateModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// update existing model by user without Vendor role
		msgUpdateModel := NewMsgUpdateModel(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModel)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	// update existing model by vendor with another VendorID
	msgUpdateModel := NewMsgUpdateModel(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model by owner
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID)

	// update existing model by vendor with the same VendorID as owner's one
	msgUpdateModel = NewMsgUpdateModel(vendorWithSameVid)
	msgUpdateModel.ProductLabel += "-updated-once-more"
	msgUpdateModel.LsfRevision++
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)
}

func TestHandler_LsfUpdateValidations(t *testing.T) {
	setup := Setup(t)

	// add new model without lsfURL
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	msgCreateModel.LsfUrl = ""
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
	require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
	require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
	require.Equal(t, msgCreateModel.LsfUrl, "")
	require.Equal(t, testconstants.EmptyLsfRevision, receivedModel.LsfRevision)

	// Update model with lsfRevision, keep the LsfUrl empty
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.LsfRevision = 1
	msgUpdateModel.LsfUrl = ""
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	// Update fails as LsfUrl is empty
	require.Error(t, err)
	require.True(t, types.ErrLsfRevisionIsNotValid.Is(err))

	// Update model with valid LsfUrl, but higher LsfRevision
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.LsfUrl = "https://example.com/lsf.json"
	msgUpdateModel.LsfRevision = 5
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	// Update fails as LsfRevision is not increased monotonically by just 1
	require.Error(t, err)
	require.True(t, types.ErrLsfRevisionIsNotValid.Is(err))

	// Update model with valid LsfUrl and LsfRevision set to 1
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.LsfUrl = "https://example.com/lsf.json"
	msgUpdateModel.LsfRevision = testconstants.LsfRevision
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	// Update succeeds
	require.NoError(t, err)

	// query model
	receivedModel, err = queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, msgUpdateModel.LsfUrl, receivedModel.LsfUrl)
	require.Equal(t, msgUpdateModel.LsfRevision, receivedModel.LsfRevision)

	// Increase LsfRevision by 1
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.LsfUrl = ""
	msgUpdateModel.LsfRevision = testconstants.LsfRevision + 1
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	// Update succeeds
	require.NoError(t, err)

	// query model
	receivedModel, err = queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, "https://example.com/lsf.json", receivedModel.LsfUrl)
	require.Equal(t, msgUpdateModel.LsfRevision, receivedModel.LsfRevision)

	// Increase LsfRevision by more then 1
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.LsfUrl = ""
	msgUpdateModel.LsfRevision = testconstants.LsfRevision + 3
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	// Update fails as LsfRevision is not increased monotonically by just 1
	require.Error(t, err)
	require.True(t, types.ErrLsfRevisionIsNotValid.Is(err))
}

func TestHandler_LsfAddValidation_DefaultValue(t *testing.T) {
	setup := Setup(t)

	// add new model without lsfURL
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
	require.Equal(t, msgCreateModel.LsfUrl, receivedModel.LsfUrl)
	require.Equal(t, testconstants.LsfRevision, receivedModel.LsfRevision)
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
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
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

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

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

func TestHandler_DeleteModel(t *testing.T) {
	setup := Setup(t)

	// try delete not present model
	msgDeleteModel := NewMsgDeleteModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgDeleteModel)
	require.Error(t, err)
	require.True(t, types.ErrModelDoesNotExist.Is(err))

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// delete existing model
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// query deleted model
	_, err = queryModel(setup, msgDeleteModel.Vid, msgDeleteModel.Pid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_DeleteModelWithAssociatedModelVersionsNotCertified(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// mock model versions not to be certified
	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion1.Vid, msgCreateModelVersion1.Pid, msgCreateModelVersion1.SoftwareVersion, mock.Anything).Return(false)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion2.Vid, msgCreateModelVersion2.Pid, msgCreateModelVersion2.SoftwareVersion, mock.Anything).Return(false)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(true)

	msgDeleteModel := NewMsgDeleteModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// query first deleted model version
	_, err = queryModelVersion(
		setup,
		msgCreateModelVersion1.Vid,
		msgCreateModelVersion1.Pid,
		msgCreateModelVersion1.SoftwareVersion,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query second deleted model version
	_, err = queryModelVersion(
		setup,
		msgCreateModelVersion2.Vid,
		msgCreateModelVersion2.Pid,
		msgCreateModelVersion2.SoftwareVersion,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query deleted model
	_, err = queryModel(setup, msgDeleteModel.Vid, msgDeleteModel.Pid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_DeleteModelWithAssociatedModelVersionsCertified(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// mock one model version to be certified
	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion1.Vid, msgCreateModelVersion1.Pid, msgCreateModelVersion1.SoftwareVersion, mock.Anything).Return(false)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(true)

	msgDeleteModel := NewMsgDeleteModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.Error(t, err)

	// query first deleted model version - should not be deleted
	_, err = queryModelVersion(
		setup,
		msgCreateModelVersion1.Vid,
		msgCreateModelVersion1.Pid,
		msgCreateModelVersion1.SoftwareVersion,
	)
	require.NoError(t, err)

	// query second deleted model version - should not be deleted
	_, err = queryModelVersion(
		setup,
		msgCreateModelVersion2.Vid,
		msgCreateModelVersion2.Pid,
		msgCreateModelVersion2.SoftwareVersion,
	)
	require.NoError(t, err)

	// query deleted model - should not be deleted
	_, err = queryModel(setup, msgDeleteModel.Vid, msgDeleteModel.Pid)
	require.NoError(t, err)
}

func TestHandler_OnlyOwnerAndVendorWithSameVidCanDeleteModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// delete existing model by user without Vendor role
		msgDeleteModel := NewMsgDeleteModel(accAddress)
		_, err = setup.Handler(setup.Ctx, msgDeleteModel)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	// delete existing model by vendor with another VendorID
	msgDeleteModel := NewMsgDeleteModel(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// delete existing model by owner
	msgDeleteModel = NewMsgDeleteModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// add new model
	msgCreateModel = NewMsgCreateModel(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID)

	// delete existing model by vendor with the same VendorID as owner's one
	msgDeleteModel = NewMsgDeleteModel(vendorWithSameVid)
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)
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
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
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

	// query model versions
	receivedModelVersions, err := queryAllModelVersions(
		setup,
		msgCreateModelVersion.Vid,
		msgCreateModelVersion.Pid,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion.Vid, receivedModelVersions.Vid)
	require.Equal(t, msgCreateModelVersion.Pid, receivedModelVersions.Pid)
	require.Equal(t, []uint32{msgCreateModelVersion.SoftwareVersion}, receivedModelVersions.SoftwareVersions)
}

func TestHandler_AddMultipleModelVersions(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version 1
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, uint32(1))
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	// add new model version 2
	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, uint32(2))
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// query model version 1
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgCreateModelVersion1.Vid,
		msgCreateModelVersion1.Pid,
		msgCreateModelVersion1.SoftwareVersion,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion1.Vid, receivedModelVersion.Vid)
	require.Equal(t, msgCreateModelVersion1.Pid, receivedModelVersion.Pid)
	require.Equal(t, msgCreateModelVersion1.SoftwareVersion, receivedModelVersion.SoftwareVersion)
	require.Equal(t, msgCreateModelVersion1.SoftwareVersionString, receivedModelVersion.SoftwareVersionString)

	// query model version 2
	receivedModelVersion, err = queryModelVersion(
		setup,
		msgCreateModelVersion2.Vid,
		msgCreateModelVersion2.Pid,
		msgCreateModelVersion2.SoftwareVersion,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion2.Vid, receivedModelVersion.Vid)
	require.Equal(t, msgCreateModelVersion2.Pid, receivedModelVersion.Pid)
	require.Equal(t, msgCreateModelVersion2.SoftwareVersion, receivedModelVersion.SoftwareVersion)
	require.Equal(t, msgCreateModelVersion2.SoftwareVersionString, receivedModelVersion.SoftwareVersionString)

	// query model versions
	receivedModelVersions, err := queryAllModelVersions(
		setup,
		msgCreateModelVersion1.Vid,
		msgCreateModelVersion1.Pid,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion1.Vid, receivedModelVersions.Vid)
	require.Equal(t, msgCreateModelVersion1.Pid, receivedModelVersions.Pid)
	require.Equal(t, []uint32{msgCreateModelVersion1.SoftwareVersion, msgCreateModelVersion2.SoftwareVersion}, receivedModelVersions.SoftwareVersions)
}

func TestHandler_UpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// try update not present model version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, types.ErrModelVersionDoesNotExist.Is(err))

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update existing model version
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
	require.Equal(t, receivedModelVersion.FirmwareInformation, msgCreateModelVersion.FirmwareInformation)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgCreateModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.OtaChecksum, msgCreateModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgCreateModelVersion.OtaChecksumType)

	// query model versions
	receivedModelVersions, err := queryAllModelVersions(
		setup,
		msgCreateModelVersion.Vid,
		msgCreateModelVersion.Pid,
	)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModelVersion.Vid, receivedModelVersions.Vid)
	require.Equal(t, msgCreateModelVersion.Pid, receivedModelVersions.Pid)
	require.Equal(t, []uint32{msgCreateModelVersion.SoftwareVersion}, receivedModelVersions.SoftwareVersions)
}

func TestHandler_PartiallyUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
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

func TestHandler_UpdateOnlyMinApplicableSoftwareVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	msgCreateModelVersion.MinApplicableSoftwareVersion = 5
	msgCreateModelVersion.MaxApplicableSoftwareVersion = 10
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update only min version to a value greater than stored max version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 11
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 0

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, types.ErrMaxSVLessThanMinSV.Is(err))

	// try to update only min version to a value less than stored max version
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 7
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 0

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query updated model version
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check that min version has been updated
	require.Equal(t, uint32(7), receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, uint32(10), receivedModelVersion.MaxApplicableSoftwareVersion)

	// try to update only min version to a value equal to stored max version
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 10
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 0

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query updated model version
	receivedModelVersion, err = queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check that min version has been updated
	require.Equal(t, uint32(10), receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, uint32(10), receivedModelVersion.MaxApplicableSoftwareVersion)
}

func TestHandler_UpdateOnlyMaxApplicableSoftwareVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	msgCreateModelVersion.MinApplicableSoftwareVersion = 5
	msgCreateModelVersion.MaxApplicableSoftwareVersion = 10
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update only max version to a value less than stored min version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 0
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 4

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, types.ErrMaxSVLessThanMinSV.Is(err))

	// try to update only max version to a value greater than stored min version
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 0
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 7

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query updated model version
	receivedModelVersion, err := queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check that max version has been updated
	require.Equal(t, uint32(5), receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, uint32(7), receivedModelVersion.MaxApplicableSoftwareVersion)

	// try to update only max version to a value equal to stored min version
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 0
	msgUpdateModelVersion.MaxApplicableSoftwareVersion = 5

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query updated model version
	receivedModelVersion, err = queryModelVersion(
		setup,
		msgUpdateModelVersion.Vid,
		msgUpdateModelVersion.Pid,
		msgUpdateModelVersion.SoftwareVersion,
	)
	require.NoError(t, err)

	// check that max version has been updated
	require.Equal(t, uint32(5), receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, uint32(5), receivedModelVersion.MaxApplicableSoftwareVersion)
}

func TestHandler_OnlyOwnerAndVendorWithSameVidCanUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreteModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreteModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// update existing model version by user without Vendor role
		msgUpdateModelVersion := NewMsgUpdateModelVersion(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	// update existing model by vendor with another VendorID
	msgUpdateModelVersion := NewMsgUpdateModelVersion(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model version by owner
	msgUpdateModelVersion = NewMsgUpdateModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID)

	// update existing model by vendor with the same VendorID as owner's one
	msgUpdateModelVersion = NewMsgUpdateModelVersion(vendorWithSameVid)
	msgUpdateModelVersion.ReleaseNotesUrl += "/updated-once-more"
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)
}

func TestHandler_DeleteModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgDeleteModelVersion.Vid, msgDeleteModelVersion.Pid, msgDeleteModelVersion.SoftwareVersion, mock.Anything).Return(false)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(true)

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.NoError(t, err)

	// query model version
	_, err = queryModelVersion(
		setup,
		msgDeleteModelVersion.Vid,
		msgDeleteModelVersion.Pid,
		msgDeleteModelVersion.SoftwareVersion,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_DeleteModelVersionNotByVendor(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	setup.AddAccount(testconstants.Address1, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, testconstants.VendorID2)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(testconstants.Address1)

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DeleteModelVersionDifferentVid(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)
	msgDeleteModelVersion.Vid = 55

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DeleteModelVersionDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)
	msgDeleteModelVersion.SoftwareVersion = 3

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, types.ErrModelVersionDoesNotExist)
}

func TestHandler_DeleteModelVersionNotByCreator(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	setup.AddAccount(testconstants.Address1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(testconstants.Address1)

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DeleteModelVersionCertified(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgDeleteModelVersion.Vid, msgDeleteModelVersion.Pid, msgDeleteModelVersion.SoftwareVersion, mock.Anything).Return(true)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, types.ErrModelVersionCertified)
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

func queryAllModelVersions(
	setup *TestSetup,
	vid int32,
	pid int32,
) (*types.ModelVersions, error) {
	req := &types.QueryGetModelVersionsRequest{
		Vid: vid,
		Pid: pid,
	}

	resp, err := setup.Keeper.ModelVersions(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ModelVersions, nil
}

func NewMsgCreateModel(signer sdk.AccAddress) *types.MsgCreateModel {
	return &types.MsgCreateModel{
		Creator:                                  signer.String(),
		Vid:                                      testconstants.VendorID1,
		Pid:                                      testconstants.Pid,
		DeviceTypeId:                             testconstants.DeviceTypeID,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
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

func NewMsgUpdateModel(signer sdk.AccAddress) *types.MsgUpdateModel {
	return &types.MsgUpdateModel{
		Creator:                                  signer.String(),
		Vid:                                      testconstants.VendorID1,
		Pid:                                      testconstants.Pid,
		ProductName:                              testconstants.ProductName + "-updated",
		ProductLabel:                             testconstants.ProductLabel + "-updated",
		PartNumber:                               testconstants.PartNumber + "-updated",
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowURL + "/updated",
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction + "-updated",
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction + "-updated",
		UserManualUrl: testconstants.UserManualURL + "/updated",
		SupportUrl:    testconstants.SupportURL + "/updated",
		ProductUrl:    testconstants.ProductURL + "/updated",
		LsfUrl:        testconstants.LsfURL + "/updated",
		LsfRevision:   testconstants.LsfRevision + 1,
	}
}

func NewMsgDeleteModel(signer sdk.AccAddress) *types.MsgDeleteModel {
	return &types.MsgDeleteModel{
		Creator: signer.String(),
		Vid:     testconstants.VendorID1,
		Pid:     testconstants.Pid,
	}
}

func NewMsgCreateModelVersion(signer sdk.AccAddress, softwareVersion uint32) *types.MsgCreateModelVersion {
	return &types.MsgCreateModelVersion{
		Creator:                      signer.String(),
		Vid:                          testconstants.VendorID1,
		Pid:                          testconstants.Pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareInformation:          testconstants.FirmwareInformation,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL,
	}
}

func NewMsgUpdateModelVersion(signer sdk.AccAddress) *types.MsgUpdateModelVersion {
	return &types.MsgUpdateModelVersion{
		Creator:                      signer.String(),
		Vid:                          testconstants.VendorID1,
		Pid:                          testconstants.Pid,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionValid:         !testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaURL + "/updated",
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion + 1,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion + 1,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL + "/updated",
	}
}

func NewMsgDeleteModelVersion(signer sdk.AccAddress) *types.MsgDeleteModelVersion {
	return &types.MsgDeleteModelVersion{
		Creator:         signer.String(),
		Vid:             testconstants.VendorID1,
		Pid:             testconstants.Pid,
		SoftwareVersion: testconstants.SoftwareVersion,
	}
}
