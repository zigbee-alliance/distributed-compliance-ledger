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
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRightsToChange(ctx sdk.Context, addr sdk.AccAddress, pid int32) bool {
	args := m.Called(ctx, addr, pid)

	return args.Bool(0)
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

	complianceInfo, ok := args.Get(0).(dclcompltypes.ComplianceInfo)
	if ok {
		val = complianceInfo
	}

	return val, args.Bool(len(args) - 1)
}

var _ keeper.ComplianceKeeper = &ComplianceKeeperMock{}

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
	Vendor     sdk.AccAddress
	VendorID   int32
	ProductIDs []*commontypes.Uint16Range
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vendorID int32,
	productIDs []*commontypes.Uint16Range,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)

	dclauthKeeper.On("HasVendorID", mock.Anything, accAddress, vendorID).Return(true)
	dclauthKeeper.On("HasVendorID", mock.Anything, accAddress, mock.Anything).Return(false)

	if len(productIDs) == 0 {
		dclauthKeeper.On("HasRightsToChange", mock.Anything, accAddress, mock.Anything).Return(true)
	}
	for _, productIDRange := range productIDs {
		for productID := productIDRange.Min; productID <= productIDRange.Max; productID++ {
			dclauthKeeper.On("HasRightsToChange", mock.Anything, accAddress, productID).Return(true)
		}
	}
	dclauthKeeper.On("HasRightsToChange", mock.Anything, accAddress, mock.Anything).Return(false)
}

func Setup(t *testing.T) *TestSetup {
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	complianceKeeper := &ComplianceKeeperMock{}
	keeper, ctx := testkeeper.ModelKeeper(t, dclauthKeeper, complianceKeeper)

	vendor := testdata.GenerateAccAddress()
	vendorID := testconstants.VendorID1
	productIDs := testconstants.ProductIDsEmpty
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
		ProductIDs:       productIDs,
	}

	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vendorID, productIDs)

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
	require.Equal(t, msgCreateModel.CommissioningFallbackUrl, receivedModel.CommissioningFallbackUrl)
}

func TestHandler_CreateModelByVendorAdminUpdateDeleteByVendor(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add a new model by VendorAdmin
	msgCreateModel := NewMsgCreateModel(vendorAdmin)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// update model by Vendor
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.Vid = msgCreateModel.Vid
	msgUpdateModel.Pid = msgCreateModel.Pid
	msgUpdateModel.ProductName = "Updated by Vendor"
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, msgUpdateModel.ProductName, receivedModel.ProductName)

	// delete model by Vendor
	msgDeleteModel := NewMsgDeleteModel(setup.Vendor)
	msgDeleteModel.Vid = msgCreateModel.Vid
	msgDeleteModel.Pid = msgCreateModel.Pid
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// query model
	_, err = queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_CreateModelByVendorUpdateDeleteByVendorAdmin(t *testing.T) {
	setup := Setup(t)

	// add new model by Vendor
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// update model by VendorAdmin
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	msgUpdateModel := NewMsgUpdateModel(vendorAdmin)
	msgUpdateModel.Vid = msgCreateModel.Vid
	msgUpdateModel.Pid = msgCreateModel.Pid
	msgUpdateModel.ProductName = "Updated by VendorAdmin"
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, msgUpdateModel.ProductName, receivedModel.ProductName)

	// delete model by VendorAdmin
	msgDeleteModel := NewMsgDeleteModel(vendorAdmin)
	msgDeleteModel.Vid = msgCreateModel.Vid
	msgDeleteModel.Pid = msgCreateModel.Pid
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// query model
	_, err = queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_AddModelByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(vendorAdmin)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
	require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
}

func TestHandler_UpdateModelByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// update model
	msgUpdateModel := NewMsgUpdateModel(vendorAdmin)
	msgUpdateModel.Vid = msgCreateModel.Vid
	msgUpdateModel.Pid = msgCreateModel.Pid
	msgUpdateModel.ProductName = "Updated Product Name"
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgUpdateModel.ProductName, receivedModel.ProductName)
}

func TestHandler_DeleteModelByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// delete model
	msgDeleteModel := NewMsgDeleteModel(vendorAdmin)
	msgDeleteModel.Vid = msgCreateModel.Vid
	msgDeleteModel.Pid = msgCreateModel.Pid
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// query model
	_, err = queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_AddModelVersionByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add model version
	msgCreateModelVersion := NewMsgCreateModelVersion(vendorAdmin, 1)
	msgCreateModelVersion.Vid = msgCreateModel.Vid
	msgCreateModelVersion.Pid = msgCreateModel.Pid
	// mock getting compliance record
	setup.ComplianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion.Vid, msgCreateModelVersion.Pid, msgCreateModelVersion.SoftwareVersion, mock.Anything).Return(false)

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.NoError(t, err)
	require.Equal(t, uint32(1), receivedModelVersion.SoftwareVersion)
}

func TestHandler_UpdateModelVersionByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, 1)
	msgCreateModelVersion.Vid = msgCreateModel.Vid
	msgCreateModelVersion.Pid = msgCreateModel.Pid
	// mock getting compliance record
	setup.ComplianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion.Vid, msgCreateModelVersion.Pid, msgCreateModelVersion.SoftwareVersion, mock.Anything).Return(false)

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update model version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(vendorAdmin)
	msgUpdateModelVersion.Vid = msgCreateModel.Vid
	msgUpdateModelVersion.Pid = msgCreateModel.Pid
	msgUpdateModelVersion.SoftwareVersion = 1
	msgUpdateModelVersion.OtaUrl = "https://updated-ota-url.com"
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.NoError(t, err)
	require.Equal(t, "https://updated-ota-url.com", receivedModelVersion.OtaUrl)
}

func TestHandler_DeleteModelVersionByVendorAdmin(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, 1)
	msgCreateModelVersion.Vid = msgCreateModel.Vid
	msgCreateModelVersion.Pid = msgCreateModel.Pid
	// mock getting compliance record
	setup.ComplianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion.Vid, msgCreateModelVersion.Pid, msgCreateModelVersion.SoftwareVersion, mock.Anything).Return(false)

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// delete model version
	msgDeleteModelVersion := NewMsgDeleteModelVersion(vendorAdmin)
	msgDeleteModelVersion.Vid = msgCreateModel.Vid
	msgDeleteModelVersion.Pid = msgCreateModel.Pid
	msgDeleteModelVersion.SoftwareVersion = 1
	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.NoError(t, err)

	// query model version
	_, err = queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_CreateModelVersionByVendorAdminUpdateDeleteByVendor(t *testing.T) {
	setup := Setup(t)
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add model version by VendorAdmin
	msgCreateModelVersion := NewMsgCreateModelVersion(vendorAdmin, 1)
	msgCreateModelVersion.Vid = msgCreateModel.Vid
	msgCreateModelVersion.Pid = msgCreateModel.Pid
	setup.ComplianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion.Vid, msgCreateModelVersion.Pid, msgCreateModelVersion.SoftwareVersion, mock.Anything).Return(false)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update model version by Vendor
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.Vid = msgCreateModel.Vid
	msgUpdateModelVersion.Pid = msgCreateModel.Pid
	msgUpdateModelVersion.SoftwareVersion = 1
	msgUpdateModelVersion.OtaUrl = "https://vendor-updated-ota-url.com"
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.NoError(t, err)
	require.Equal(t, "https://vendor-updated-ota-url.com", receivedModelVersion.OtaUrl)

	// delete model version by Vendor
	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)
	msgDeleteModelVersion.Vid = msgCreateModel.Vid
	msgDeleteModelVersion.Pid = msgCreateModel.Pid
	msgDeleteModelVersion.SoftwareVersion = 1
	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.NoError(t, err)

	// query model version
	_, err = queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_CreateModelVersionByVendorUpdateDeleteByVendorAdmin(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	// add model version by Vendor
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, 1)
	msgCreateModelVersion.Vid = msgCreateModel.Vid
	msgCreateModelVersion.Pid = msgCreateModel.Pid
	setup.ComplianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion.Vid, msgCreateModelVersion.Pid, msgCreateModelVersion.SoftwareVersion, mock.Anything).Return(false)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update model version by VendorAdmin
	vendorAdmin := testdata.GenerateAccAddress()
	setup.AddAccount(vendorAdmin, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0, nil)

	msgUpdateModelVersion := NewMsgUpdateModelVersion(vendorAdmin)
	msgUpdateModelVersion.Vid = msgCreateModel.Vid
	msgUpdateModelVersion.Pid = msgCreateModel.Pid
	msgUpdateModelVersion.SoftwareVersion = 1
	msgUpdateModelVersion.OtaUrl = "https://vendoradmin-updated-ota-url.com"
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	// query model version
	receivedModelVersion, err := queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.NoError(t, err)
	require.Equal(t, "https://vendoradmin-updated-ota-url.com", receivedModelVersion.OtaUrl)

	// delete model version by VendorAdmin
	msgDeleteModelVersion := NewMsgDeleteModelVersion(vendorAdmin)
	msgDeleteModelVersion.Vid = msgCreateModel.Vid
	msgDeleteModelVersion.Pid = msgCreateModel.Pid
	msgDeleteModelVersion.SoftwareVersion = 1
	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.NoError(t, err)

	// query model version
	_, err = queryModelVersion(setup, msgCreateModel.Vid, msgCreateModel.Pid, 1)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_AddModel_CheckCommissioningModeInitialStepsHintHandling(t *testing.T) {
	cases := []struct {
		name                                      string
		commissioningModeInitialStepsHint         uint32
		expectedCommissioningModeInitialStepsHint uint32
	}{
		{
			name:                              "CommissioningModeInitialStepsHint=0 Sets Default 1",
			commissioningModeInitialStepsHint: 0,
			expectedCommissioningModeInitialStepsHint: 1,
		},
		{
			name:                              "CommissioningModeInitialStepsHint=2 Remains 2",
			commissioningModeInitialStepsHint: 2,
			expectedCommissioningModeInitialStepsHint: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			// add new model
			msgCreateModel := NewMsgCreateModel(setup.Vendor)
			msgCreateModel.CommissioningModeInitialStepsHint = tc.commissioningModeInitialStepsHint
			_, err := setup.Handler(setup.Ctx, msgCreateModel)
			require.NoError(t, err)

			// query model
			receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
			require.NoError(t, err)

			// check
			require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
			require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
			require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
			require.Equal(t, tc.expectedCommissioningModeInitialStepsHint, receivedModel.CommissioningModeInitialStepsHint)
		})
	}
}

func TestHandler_AddModel_CheckCommissioningModeSecondaryStepsHintHandling(t *testing.T) {
	cases := []struct {
		name                                        string
		commissioningModeSecondaryStepsHint         uint32
		expectedCommissioningModeSecondaryStepsHint uint32
	}{
		{
			name:                                "CommissioningModeSecondaryStepsHint=0 Sets Default 4",
			commissioningModeSecondaryStepsHint: 0,
			expectedCommissioningModeSecondaryStepsHint: 4,
		},
		{
			name:                                "CommissioningModeSecondaryStepsHint=3 Remains 3",
			commissioningModeSecondaryStepsHint: 3,
			expectedCommissioningModeSecondaryStepsHint: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			// add new model
			msgCreateModel := NewMsgCreateModel(setup.Vendor)
			msgCreateModel.CommissioningModeSecondaryStepsHint = tc.commissioningModeSecondaryStepsHint
			_, err := setup.Handler(setup.Ctx, msgCreateModel)
			require.NoError(t, err)

			// query model
			receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
			require.NoError(t, err)

			// check
			require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
			require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
			require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
			require.Equal(t, tc.expectedCommissioningModeSecondaryStepsHint, receivedModel.CommissioningModeSecondaryStepsHint)
		})
	}
}

func TestHandler_AddModel_CheckIcdUserActiveModeTriggerHintHandling(t *testing.T) {
	cases := []struct {
		name                                 string
		icdUserActiveModeTriggerHint         uint32
		expectedIcdUserActiveModeTriggerHint uint32
	}{
		{
			name:                                 "icdUserActiveModeTriggerHint=0 Sets Default 1",
			icdUserActiveModeTriggerHint:         0,
			expectedIcdUserActiveModeTriggerHint: 1,
		},
		{
			name:                                 "icdUserActiveModeTriggerHint=3 Remains 3",
			icdUserActiveModeTriggerHint:         3,
			expectedIcdUserActiveModeTriggerHint: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			// add new model
			msgCreateModel := NewMsgCreateModel(setup.Vendor)
			msgCreateModel.IcdUserActiveModeTriggerHint = tc.icdUserActiveModeTriggerHint
			_, err := setup.Handler(setup.Ctx, msgCreateModel)
			require.NoError(t, err)

			// query model
			receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
			require.NoError(t, err)

			// check
			require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
			require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
			require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
			require.Equal(t, tc.expectedIcdUserActiveModeTriggerHint, receivedModel.IcdUserActiveModeTriggerHint)
		})
	}
}

func TestHandler_AddModel_CheckFactoryResetStepsHintHandling(t *testing.T) {
	cases := []struct {
		name                          string
		factoryResetStepsHint         uint32
		expectedFactoryResetStepsHint uint32
	}{
		{
			name:                          "factoryResetStepsHint=0 Sets Default 1",
			factoryResetStepsHint:         0,
			expectedFactoryResetStepsHint: 1,
		},
		{
			name:                          "factoryResetStepsHint=3 Remains 3",
			factoryResetStepsHint:         3,
			expectedFactoryResetStepsHint: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := Setup(t)

			// add new model
			msgCreateModel := NewMsgCreateModel(setup.Vendor)
			msgCreateModel.FactoryResetStepsHint = tc.factoryResetStepsHint
			_, err := setup.Handler(setup.Ctx, msgCreateModel)
			require.NoError(t, err)

			// query model
			receivedModel, err := queryModel(setup, msgCreateModel.Vid, msgCreateModel.Pid)
			require.NoError(t, err)

			// check
			require.Equal(t, msgCreateModel.Vid, receivedModel.Vid)
			require.Equal(t, msgCreateModel.Pid, receivedModel.Pid)
			require.Equal(t, msgCreateModel.DeviceTypeId, receivedModel.DeviceTypeId)
			require.Equal(t, tc.expectedFactoryResetStepsHint, receivedModel.FactoryResetStepsHint)
		})
	}
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

	// query model
	receivedModel, err := queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, testconstants.SchemaVersion, receivedModel.SchemaVersion)

	// update existing model
	var newSchemaVersion uint32 = 2
	var newCommissioningModeInitialStepsHint uint32 = 8
	var newCommissioningModeSecondaryStepsHint uint32 = 9
	var newIcdUserActiveModeTriggerHint uint32 = 6
	var newFactoryResetStepsHint uint32 = 7
	msgUpdateModel.SchemaVersion = newSchemaVersion
	msgUpdateModel.CommissioningModeInitialStepsHint = newCommissioningModeInitialStepsHint
	msgUpdateModel.CommissioningModeSecondaryStepsHint = newCommissioningModeSecondaryStepsHint
	msgUpdateModel.IcdUserActiveModeTriggerHint = newIcdUserActiveModeTriggerHint
	msgUpdateModel.FactoryResetStepsHint = newFactoryResetStepsHint
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query updated model
	receivedModel, err = queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)

	// check
	// Mutable Fields ProductName,ProductLable,PartNumber,CommissioningCustomFlowUrl,
	// CommissioningModeInitialStepsInstruction,CommissioningModeSecondaryStepsInstruction,UserManualUrl,SupportUrl,SupportUrl
	require.Equal(t, msgUpdateModel.Vid, receivedModel.Vid)
	require.Equal(t, msgUpdateModel.Pid, receivedModel.Pid)
	require.Equal(t, msgUpdateModel.ProductLabel, receivedModel.ProductLabel)
	require.Equal(t, newCommissioningModeInitialStepsHint, receivedModel.CommissioningModeInitialStepsHint)
	require.Equal(t, newCommissioningModeSecondaryStepsHint, receivedModel.CommissioningModeSecondaryStepsHint)
	require.Equal(t, newSchemaVersion, receivedModel.SchemaVersion)
	require.Equal(t, msgUpdateModel.CommissioningFallbackUrl, receivedModel.CommissioningFallbackUrl)
}

func TestHandler_UpdateModelByVendorWithProductIds(t *testing.T) {
	setup := Setup(t)

	owner := testdata.GenerateAccAddress()
	setup.AddAccount(owner, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1, testconstants.ProductIDs200)

	// add new model
	msgCreateModel := NewMsgCreateModel(owner)
	msgCreateModel.Pid = 200
	enhancedSetupFlowTCRevision := msgCreateModel.EnhancedSetupFlowTCRevision
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDs100)

	// update existing model by vendor with another VendorID
	msgUpdateModel := NewMsgUpdateModel(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model by owner
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(owner)
	msgUpdateModel.Pid = 200
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	vendorWithoutProductIDs := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithoutProductIDs, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDsEmpty)

	// update existing model by vendor with the same VendorID as owner's one
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(vendorWithoutProductIDs)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
	msgUpdateModel.Pid = 200
	msgUpdateModel.ProductLabel += "-updated-once-more"
	msgUpdateModel.LsfRevision++
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)
}

func TestHandler_OnlyOwnerAndVendorWithSameVidCanUpdateModel(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)
	enhancedSetupFlowTCRevision := msgCreateModel.EnhancedSetupFlowTCRevision

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID, setup.ProductIDs)

		// update existing model by user without Vendor role
		msgUpdateModel := NewMsgUpdateModel(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModel)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, setup.ProductIDs)

	// update existing model by vendor with another VendorID
	msgUpdateModel := NewMsgUpdateModel(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model by owner
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, setup.ProductIDs)

	// update existing model by vendor with the same VendorID as owner's one
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(vendorWithSameVid)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
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
	enhancedSetupFlowTCRevision := msgCreateModel.EnhancedSetupFlowTCRevision

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
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
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
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
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
	enhancedSetupFlowTCRevision++
	msgUpdateModel = NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.EnhancedSetupFlowTCRevision = enhancedSetupFlowTCRevision
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
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID, setup.ProductIDs)

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
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDsEmpty)

	// add new model
	model := NewMsgCreateModel(anotherVendor)
	_, err := setup.Handler(setup.Ctx, model)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_AddModelByVendorWithProductIds(t *testing.T) {
	setup := Setup(t)

	owner := testdata.GenerateAccAddress()
	setup.AddAccount(owner, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDs200)

	model := NewMsgCreateModel(owner)
	model.Pid = 200
	_, err := setup.Handler(setup.Ctx, model)
	require.NoError(t, err)

	vendor := testdata.GenerateAccAddress()
	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDs100)
	model = NewMsgCreateModel(vendor)
	model.Pid = 101
	_, err = setup.Handler(setup.Ctx, model)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDsEmpty)

	// add model by vendor with non-assigned PIDs
	model = NewMsgCreateModel(vendorWithSameVid)
	model.Pid = 201
	_, err = setup.Handler(setup.Ctx, model)
	require.NoError(t, err)
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

func TestHandler_UpdateModelEnhancedSetupFlowTCRevisionUnsetIncrement(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgAddModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgAddModel)
	require.NoError(t, err)

	// update EnhancedSetupFlowTCRevision of existing model
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)

	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgUpdateModel.EnhancedSetupFlowTCRevision, receivedModel.EnhancedSetupFlowTCRevision)
}

func TestHandler_UpdateModelEnhancedSetupFlowTCRevisionIncrement(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgAddModel := NewMsgCreateModel(setup.Vendor)
	msgAddModel.EnhancedSetupFlowOptions = testconstants.EnhancedSetupFlowOptions
	msgAddModel.EnhancedSetupFlowTCUrl = testconstants.EnhancedSetupFlowTCURL
	msgAddModel.EnhancedSetupFlowTCRevision = int32(testconstants.EnhancedSetupFlowTCRevision)
	msgAddModel.EnhancedSetupFlowTCDigest = testconstants.EnhancedSetupFlowTCDigest
	msgAddModel.EnhancedSetupFlowTCFileSize = uint32(testconstants.EnhancedSetupFlowTCFileSize)
	msgAddModel.MaintenanceUrl = testconstants.MaintenanceURL
	_, err := setup.Handler(setup.Ctx, msgAddModel)
	require.NoError(t, err)

	// update EnhancedSetupFlowTCRevision of existing model
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)

	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.NoError(t, err)

	// query model
	receivedModel, err := queryModel(setup, msgUpdateModel.Vid, msgUpdateModel.Pid)
	require.NoError(t, err)

	// check
	require.Equal(t, msgAddModel.EnhancedSetupFlowTCRevision+1, msgUpdateModel.EnhancedSetupFlowTCRevision)
	require.Equal(t, msgUpdateModel.EnhancedSetupFlowTCRevision, receivedModel.EnhancedSetupFlowTCRevision)
}

func TestHandler_UpdateModelEnhancedSetupFlowTCRevisionIncorrectIncrement(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgAddModel := NewMsgCreateModel(setup.Vendor)
	msgAddModel.EnhancedSetupFlowOptions = testconstants.EnhancedSetupFlowOptions
	msgAddModel.EnhancedSetupFlowTCUrl = testconstants.EnhancedSetupFlowTCURL
	msgAddModel.EnhancedSetupFlowTCRevision = int32(testconstants.EnhancedSetupFlowTCRevision)
	msgAddModel.EnhancedSetupFlowTCDigest = testconstants.EnhancedSetupFlowTCDigest
	msgAddModel.EnhancedSetupFlowTCFileSize = uint32(testconstants.EnhancedSetupFlowTCFileSize)
	msgAddModel.MaintenanceUrl = testconstants.MaintenanceURL
	_, err := setup.Handler(setup.Ctx, msgAddModel)
	require.NoError(t, err)

	// update EnhancedSetupFlowTCRevision of existing model
	msgUpdateModel := NewMsgUpdateModel(setup.Vendor)
	msgUpdateModel.EnhancedSetupFlowTCRevision = msgAddModel.EnhancedSetupFlowTCRevision + 2
	_, err = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrEnhancedSetupFlowTCRevisionInvalid)
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

func TestHandler_DeleteModelAfterDeletingModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	complianceKeeper = setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	msgDeleteModelVersion1 := NewMsgDeleteModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion1)
	require.NoError(t, err)

	msgDeleteModel := NewMsgDeleteModel(setup.Vendor)
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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// mock model versions not to be certified
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

	complianceKeeper := setup.ComplianceKeeper

	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion1.Vid, msgCreateModelVersion1.Pid, msgCreateModelVersion1.SoftwareVersion, mock.Anything).Times(len(dclcompltypes.CertificationTypesList)).Return(false)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgCreateModelVersion2.Vid, msgCreateModelVersion2.Pid, msgCreateModelVersion2.SoftwareVersion, mock.Anything).Times(len(dclcompltypes.CertificationTypesList)).Return(false)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// mock one model version to be certified
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

func TestHandler_DeleteModelByVendorWitProductIds(t *testing.T) {
	setup := Setup(t)

	owner := testdata.GenerateAccAddress()
	setup.AddAccount(owner, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1, testconstants.ProductIDs200)

	// add new model
	msgCreateModel := NewMsgCreateModel(owner)
	msgCreateModel.Pid = 200
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	vendor := testdata.GenerateAccAddress()
	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDs100)

	// delete existing model by vendor with another VendorID
	msgDeleteModel := NewMsgDeleteModel(vendor)
	msgDeleteModel.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// delete existing model by owner
	msgDeleteModel = NewMsgDeleteModel(owner)
	msgDeleteModel.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
	require.NoError(t, err)

	// add new model
	msgCreateModel = NewMsgCreateModel(owner)
	msgCreateModel.Pid = 199
	_, err = setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	vendorWithSameVid := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDsEmpty)

	// delete existing model by vendor with non-assigned PIDs
	msgDeleteModel = NewMsgDeleteModel(vendorWithSameVid)
	msgDeleteModel.Pid = 199
	_, err = setup.Handler(setup.Ctx, msgDeleteModel)
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
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID, setup.ProductIDs)

		// delete existing model by user without Vendor role
		msgDeleteModel := NewMsgDeleteModel(accAddress)
		_, err = setup.Handler(setup.Ctx, msgDeleteModel)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDsEmpty)

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
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, setup.ProductIDs)

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

	softwareVersion := testconstants.SoftwareVersion

	positiveTests := []struct {
		name string
		msg  *types.MsgCreateModelVersion
	}{
		{
			name: "when compliance info does not exist",
			msg: func(msg *types.MsgCreateModelVersion) *types.MsgCreateModelVersion {
				complianceKeeper := setup.ComplianceKeeper
				complianceKeeper.On("GetComplianceInfo", mock.Anything, msg.Vid, msg.Pid, msg.SoftwareVersion, mock.Anything).Return(false)

				return msg
			}(NewMsgCreateModelVersion(setup.Vendor, softwareVersion)),
		},
		{
			name: "when compliance info already exists",
			msg: func(msg *types.MsgCreateModelVersion) *types.MsgCreateModelVersion {
				complianceKeeper := setup.ComplianceKeeper
				complianceInfo := dclcompltypes.ComplianceInfo{
					Vid:                   msg.Vid,
					Pid:                   msg.Pid,
					SoftwareVersion:       msg.SoftwareVersion,
					SoftwareVersionString: msg.SoftwareVersionString,
					CDVersionNumber:       uint32(msg.CdVersionNumber),
				}

				complianceKeeper.On("GetComplianceInfo", mock.Anything, msg.Vid, msg.Pid, msg.SoftwareVersion, mock.Anything).Return(complianceInfo, true)

				return msg
			}(NewMsgCreateModelVersion(setup.Vendor, softwareVersion+1)),
		},
	}

	negativeTests := []struct {
		name string
		msg  *types.MsgCreateModelVersion
		err  error
	}{
		{
			name: "when compliance info software version string does not match",
			msg: func(msg *types.MsgCreateModelVersion) *types.MsgCreateModelVersion {
				complianceKeeper := setup.ComplianceKeeper

				complianceInfo := dclcompltypes.ComplianceInfo{
					Vid:                   msg.Vid,
					Pid:                   msg.Pid,
					SoftwareVersion:       msg.SoftwareVersion,
					SoftwareVersionString: "4.0",
					CDVersionNumber:       uint32(msg.CdVersionNumber),
				}
				complianceKeeper.On("GetComplianceInfo", mock.Anything, msg.Vid, msg.Pid, msg.SoftwareVersion, mock.Anything).Return(complianceInfo, true)

				return msg
			}(NewMsgCreateModelVersion(setup.Vendor, softwareVersion+2)),
			err: types.ErrComplianceInfoSoftwareVersionStringDoesNotMatch,
		},
		{
			name: "when compliance info CD version does not match",
			msg: func(msg *types.MsgCreateModelVersion) *types.MsgCreateModelVersion {
				complianceKeeper := setup.ComplianceKeeper
				complianceInfo := dclcompltypes.ComplianceInfo{
					Vid:                   msg.Vid,
					Pid:                   msg.Pid,
					SoftwareVersion:       msg.SoftwareVersion,
					SoftwareVersionString: msg.SoftwareVersionString,
					CDVersionNumber:       uint32(msg.CdVersionNumber + 1),
				}

				complianceKeeper.On("GetComplianceInfo", mock.Anything, msg.Vid, msg.Pid, msg.SoftwareVersion, mock.Anything).Return(complianceInfo, true)

				return msg
			}(NewMsgCreateModelVersion(setup.Vendor, softwareVersion+3)),
			err: types.ErrComplianceInfoCDVersionNumberDoesNotMatch,
		},
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			msgCreateModelVersion := tt.msg
			currentModelVersions, _ := queryAllModelVersions(
				setup,
				msgCreateModelVersion.Vid,
				msgCreateModelVersion.Pid,
			)

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
			if currentModelVersions == nil {
				require.Equal(t, msgCreateModelVersion.Vid, receivedModelVersions.Vid)
				require.Equal(t, msgCreateModelVersion.Pid, receivedModelVersions.Pid)
				require.Equal(t, []uint32{msgCreateModelVersion.SoftwareVersion}, receivedModelVersions.SoftwareVersions)
				require.Equal(t, msgCreateModelVersion.SchemaVersion, receivedModelVersion.SchemaVersion)
			} else {
				currentModelVersions.SoftwareVersions = append(currentModelVersions.SoftwareVersions, msgCreateModelVersion.SoftwareVersion)
				require.Equal(t, currentModelVersions, receivedModelVersions)
			}
		})

		for _, tt := range negativeTests {
			t.Run(tt.name, func(t *testing.T) {
				_, err = setup.Handler(setup.Ctx, tt.msg)
				require.ErrorIs(t, err, tt.err)
			})
		}
	}
}

func TestHandler_AddMultipleModelVersions(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// update existing model version
	var newSchemaVersion uint32 = 2
	msgUpdateModelVersion.SchemaVersion = newSchemaVersion
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
	require.Equal(t, receivedModelVersion.OtaUrl, msgCreateModelVersion.OtaUrl+"/updated")
	require.Equal(t, receivedModelVersion.OtaChecksum, msgCreateModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgCreateModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.MinApplicableSoftwareVersion, msgUpdateModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.MaxApplicableSoftwareVersion, msgUpdateModelVersion.MaxApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.ReleaseNotesUrl, msgUpdateModelVersion.ReleaseNotesUrl)

	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgCreateModelVersion.SoftwareVersionString)
	require.Equal(t, receivedModelVersion.CdVersionNumber, msgCreateModelVersion.CdVersionNumber)
	require.Equal(t, receivedModelVersion.FirmwareInformation, msgCreateModelVersion.FirmwareInformation)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgCreateModelVersion.OtaChecksumType)
	require.Equal(t, receivedModelVersion.SpecificationVersion, msgCreateModelVersion.SpecificationVersion)
	require.Equal(t, newSchemaVersion, receivedModelVersion.SchemaVersion)

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

func TestHandler_UpdateModelVersionByVendorWithProductIds(t *testing.T) {
	setup := Setup(t)

	owner := testdata.GenerateAccAddress()
	setup.AddAccount(owner, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1, testconstants.ProductIDs200)

	// add new model
	msgCreteModel := NewMsgCreateModel(owner)
	msgCreteModel.Pid = 200
	_, err := setup.Handler(setup.Ctx, msgCreteModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(owner, testconstants.SoftwareVersion)
	msgCreateModelVersion.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	vendor := testdata.GenerateAccAddress()
	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDs100)

	// update existing model by vendor with another productIDs
	msgUpdateModelVersion := NewMsgUpdateModelVersion(vendor)
	msgUpdateModelVersion.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing model version by owner
	msgUpdateModelVersion = NewMsgUpdateModelVersion(owner)
	msgUpdateModelVersion.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)

	vendorWithoutProductIDs := testdata.GenerateAccAddress()
	setup.AddAccount(vendorWithoutProductIDs, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, testconstants.ProductIDsEmpty)

	msgUpdateModelVersion = NewMsgUpdateModelVersion(vendorWithoutProductIDs)
	msgUpdateModelVersion.Pid = 200

	msgUpdateModelVersion.ReleaseNotesUrl += "/updated-once-more"
	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.NoError(t, err)
}

func TestHandler_PartiallyUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	// Update only SoftwareVersionValid and ReleaseNotesUrl
	msgUpdateModelVersion.SoftwareVersionValid = !msgCreateModelVersion.SoftwareVersionValid

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

	require.Equal(t, msgCreateModelVersion.OtaUrl+"/updated", receivedModelVersion.OtaUrl)
	require.Equal(t, msgCreateModelVersion.MinApplicableSoftwareVersion, receivedModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, msgCreateModelVersion.MaxApplicableSoftwareVersion, receivedModelVersion.MaxApplicableSoftwareVersion)
}

func TestHandler_UpdateOnlyMinApplicableSoftwareVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
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

func TestHandler_UpdateOTAFieldsInitiallyNotSet(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	msgCreateModelVersion.OtaUrl = ""
	msgCreateModelVersion.OtaFileSize = 0
	msgCreateModelVersion.OtaChecksum = ""

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update only max version to a value less than stored min version
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.OtaUrl = "https://123.com"
	msgUpdateModelVersion.OtaFileSize = 4
	msgUpdateModelVersion.OtaChecksum = "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk="
	msgUpdateModelVersion.OtaChecksumType = 1

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

	// check that OTA fields has been updated
	require.Equal(t, msgUpdateModelVersion.OtaUrl, receivedModelVersion.OtaUrl)
	require.Equal(t, msgUpdateModelVersion.OtaFileSize, receivedModelVersion.OtaFileSize)
	require.Equal(t, msgUpdateModelVersion.OtaChecksum, receivedModelVersion.OtaChecksum)
	require.Equal(t, msgUpdateModelVersion.OtaChecksumType, receivedModelVersion.OtaChecksumType)
}

func TestHandler_UpdateOTAFieldsInitiallyNotSet_OtherMissing(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	msgCreateModelVersion.OtaUrl = ""
	msgCreateModelVersion.OtaFileSize = 0
	msgCreateModelVersion.OtaChecksum = ""
	msgCreateModelVersion.OtaChecksumType = 0

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update OtaUrl but missing other fields
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.OtaUrl = "https://123.com"

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.ErrorIs(t, err, types.ErrOtaFieldsMissProvided)
}

func TestHandler_UpdateOTAFieldsInitiallySet(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update OTA fields
	newOTAUrl := "https://123.com"

	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.OtaUrl = newOTAUrl
	msgUpdateModelVersion.OtaFileSize = 0
	msgUpdateModelVersion.OtaChecksum = ""
	msgUpdateModelVersion.OtaChecksumType = 0

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

	// check that only OTA URL has been updated
	require.Equal(t, newOTAUrl, receivedModelVersion.OtaUrl)
	require.Equal(t, msgCreateModelVersion.OtaChecksum, receivedModelVersion.OtaChecksum)
	require.Equal(t, msgCreateModelVersion.OtaFileSize, receivedModelVersion.OtaFileSize)
	require.Equal(t, msgCreateModelVersion.OtaChecksumType, receivedModelVersion.OtaChecksumType)
}

func TestHandler_UpdateOTAFieldsInitiallySet_OtherPresent(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)

	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	// try to update OTA fields but other fields are present
	msgUpdateModelVersion := NewMsgUpdateModelVersion(setup.Vendor)
	msgUpdateModelVersion.OtaUrl = "https://123.com"
	msgUpdateModelVersion.OtaFileSize = 4

	_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.ErrorIs(t, err, types.ErrOtaFieldsCannotBeUpdated)
}
func TestHandler_OnlyOwnerAndVendorWithSameVidCanUpdateModelVersion(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreteModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreteModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
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
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID, setup.ProductIDs)

		// update existing model version by user without Vendor role
		msgUpdateModelVersion := NewMsgUpdateModelVersion(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateModelVersion)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	anotherVendor := testdata.GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDsEmpty)

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
	setup.AddAccount(vendorWithSameVid, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID, setup.ProductIDs)

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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)

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

	// query model versions
	_, err = queryAllModelVersions(
		setup,
		msgDeleteModelVersion.Vid,
		msgDeleteModelVersion.Pid,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestHandler_DeleteOneOfTwoModelVersions(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	// add two new model versions
	msgCreateModelVersion1 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion1)
	require.NoError(t, err)

	msgCreateModelVersion2 := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion+1)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion2)
	require.NoError(t, err)

	// mock model versions not to be certified
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.NoError(t, err)

	// query deleted model version
	_, err = queryModelVersion(
		setup,
		msgCreateModelVersion1.Vid,
		msgCreateModelVersion1.Pid,
		msgCreateModelVersion1.SoftwareVersion,
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))

	// query not deleted model version
	modelVersion, err := queryModelVersion(
		setup,
		msgCreateModelVersion2.Vid,
		msgCreateModelVersion2.Pid,
		msgCreateModelVersion2.SoftwareVersion,
	)
	require.NoError(t, err)
	require.NotNil(t, modelVersion)

	modelVersions, err := queryAllModelVersions(setup, msgCreateModel.Vid, msgCreateModel.Pid)
	require.NoError(t, err)
	require.Equal(t, []uint32{msgCreateModelVersion2.SoftwareVersion}, modelVersions.SoftwareVersions)
}

func TestHandler_DeleteModelVersionDifferentAccSameVid(t *testing.T) {
	setup := Setup(t)

	// add new model
	msgCreateModel := NewMsgCreateModel(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	secondAcc := testdata.GenerateAccAddress()
	secondAccVid := testconstants.VendorID1

	setup.AddAccount(secondAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, secondAccVid, testconstants.ProductIDsEmpty)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(secondAcc)

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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	setup.AddAccount(testconstants.Address1, []dclauthtypes.AccountRole{dclauthtypes.Trustee}, testconstants.VendorID2, testconstants.ProductIDsEmpty)

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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	setup.AddAccount(testconstants.Address1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2, testconstants.ProductIDsEmpty)

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

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Times(len(dclcompltypes.CertificationTypesList)).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(setup.Vendor, testconstants.SoftwareVersion)
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(setup.Vendor)

	complianceKeeper.On("GetComplianceInfo", mock.Anything, msgDeleteModelVersion.Vid, msgDeleteModelVersion.Pid, msgDeleteModelVersion.SoftwareVersion, mock.Anything).Return(true)
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, types.ErrModelVersionDeletionCertified)
}

func TestHandler_DeleteModelVersionByVendorWithProductIds(t *testing.T) {
	setup := Setup(t)

	owner := testdata.GenerateAccAddress()
	setup.AddAccount(owner, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1, testconstants.ProductIDs200)

	// add new model
	msgCreateModel := NewMsgCreateModel(owner)
	msgCreateModel.Pid = 200
	_, err := setup.Handler(setup.Ctx, msgCreateModel)
	require.NoError(t, err)

	complianceKeeper := setup.ComplianceKeeper
	complianceKeeper.On("GetComplianceInfo", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(false)
	// add new model version
	msgCreateModelVersion := NewMsgCreateModelVersion(owner, testconstants.SoftwareVersion)
	msgCreateModelVersion.Pid = 200
	_, err = setup.Handler(setup.Ctx, msgCreateModelVersion)
	require.NoError(t, err)

	setup.AddAccount(testconstants.Address1, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1, testconstants.ProductIDs100)

	msgDeleteModelVersion := NewMsgDeleteModelVersion(testconstants.Address1)
	msgDeleteModelVersion.Pid = 200

	_, err = setup.Handler(setup.Ctx, msgDeleteModelVersion)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	msgDeleteModelVersion = NewMsgDeleteModelVersion(owner)
	msgDeleteModelVersion.Pid = 200

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
		FactoryResetStepsHint:                      testconstants.FactoryResetStepsHint,
		FactoryResetStepsInstruction:               testconstants.FactoryResetStepsInstruction,
		IcdUserActiveModeTriggerHint:               testconstants.IcdUserActiveModeTriggerHint,
		IcdUserActiveModeTriggerInstruction:        testconstants.IcdUserActiveModeTriggerInstruction,
		UserManualUrl:                              testconstants.UserManualURL,
		SupportUrl:                                 testconstants.SupportURL,
		ProductUrl:                                 testconstants.ProductURL,
		LsfUrl:                                     testconstants.LsfURL,
		CommissioningFallbackUrl:                   testconstants.CommissioningFallbackURL,
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
		FactoryResetStepsInstruction:               testconstants.FactoryResetStepsInstruction + "-updated",
		IcdUserActiveModeTriggerInstruction:        testconstants.IcdUserActiveModeTriggerInstruction + "-updated",
		UserManualUrl:                              testconstants.UserManualURL + "/updated",
		SupportUrl:                                 testconstants.SupportURL + "/updated",
		ProductUrl:                                 testconstants.ProductURL + "/updated",
		LsfUrl:                                     testconstants.LsfURL + "/updated",
		LsfRevision:                                testconstants.LsfRevision + 1,
		EnhancedSetupFlowOptions:                   testconstants.EnhancedSetupFlowOptions + 2,
		EnhancedSetupFlowTCUrl:                     testconstants.EnhancedSetupFlowTCURL + "/updated",
		EnhancedSetupFlowTCRevision:                int32(testconstants.EnhancedSetupFlowTCRevision + 1),
		EnhancedSetupFlowTCDigest:                  testconstants.EnhancedSetupFlowTCDigest,
		EnhancedSetupFlowTCFileSize:                uint32(testconstants.EnhancedSetupFlowTCFileSize + 1),
		MaintenanceUrl:                             testconstants.MaintenanceURL + "/updated",
		CommissioningFallbackUrl:                   testconstants.CommissioningFallbackURL + "/updated",
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
		SchemaVersion:                testconstants.SchemaVersion,
		SpecificationVersion:         testconstants.SpecificationVersion,
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
		OtaFileSize:                  0,
		OtaChecksum:                  "",
		OtaChecksumType:              0,
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
