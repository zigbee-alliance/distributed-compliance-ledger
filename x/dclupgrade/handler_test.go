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
package dclupgrade

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	UpgradeKeeper *UpgradeKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	Vendor   sdk.AccAddress
	VendorID int32
	Plan     types.Plan
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vendorID int32,
) {
	for _, role := range roles {
		setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)

	setup.DclauthKeeper.On("HasVendorID", mock.Anything, accAddress, vendorID).Return(true)
	setup.DclauthKeeper.On("HasVendorID", mock.Anything, accAddress, mock.Anything).Return(false)
}

func (setup *TestSetup) ScheduleUpgrade(
	plan types.Plan,
) {
	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, plan).Return(nil)
}

func Setup(t *testing.T) TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	upgradeKeeper := &UpgradeKeeperMock{}
	keeper, ctx := testkeeper.DclupgradeKeeper(t, dclauthKeeper, upgradeKeeper)

	vendor := GenerateAccAddress()
	vendorID := testconstants.VendorID1
	plan := testconstants.Plan
	setup := TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		UpgradeKeeper: upgradeKeeper,
		Handler:       NewHandler(*keeper),
		Vendor:        vendor,
		VendorID:      vendorID,
		Plan:          plan,
	}

	setup.AddAccount(vendor, []dclauthtypes.AccountRole{types.UpgradeApprovalRole}, vendorID)
	setup.ScheduleUpgrade(plan)
	return setup
}

func TestHandler_ProposeUpgrade(t *testing.T) {
	setup := Setup(t)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)
}

func TestHandler_ProposedUpgradeExists(t *testing.T) {
	setup := Setup(t)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)

	// propose the save upgrade
	_, err = setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrProposedUpgradeAlreadyExists.Is(err))
}

func TestHandler_ProposeUpgrade_OnlyTrusteeCanProposeUpgrade(t *testing.T) {
	setup := Setup(t)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// propose upgrade by user without Trustee role
		msgProposeUpgrade := NewMsgProposeUpgrade(accAddress)
		_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		setup.UpgradeKeeper.AssertNotCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)
	}

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{types.UpgradeApprovalRole}, setup.VendorID)

	// propose upgrade by user with Trustee role
	msgProposeUpgrade := NewMsgProposeUpgrade(accAddress)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)
	setup.UpgradeKeeper.AssertCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)
}

// ApproveUpgrade tests

func TestHandler_ApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{types.UpgradeApprovalRole}, setup.VendorID)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(accAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)
}

func TestHandler_ApproveUpgrade_OnlyTrusteeCanApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// approve upgrade by user without Trustee role
		msgApproveUpgrade := NewMsgApproveUpgrade(accAddress)
		_, err := setup.Handler(setup.Ctx, msgApproveUpgrade)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{types.UpgradeApprovalRole}, setup.VendorID)

	// propose upgrade by user with Trustee role
	msgApproveUpgrade := NewMsgApproveUpgrade(accAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)
}

func TestHandler_ApproveUpgrade_ProposedUpgradeDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrProposedUpgradeDoesNotExist.Is(err))
}

func TestHandler_ApproveUpgrade_MessageCreatorAlreadyApprovedUpgrade(t *testing.T) {
	setup := Setup(t)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func NewMsgProposeUpgrade(signer sdk.AccAddress) *types.MsgProposeUpgrade {
	return &types.MsgProposeUpgrade{
		Creator: signer.String(),
		Plan:    testconstants.Plan,
	}
}

func NewMsgApproveUpgrade(signer sdk.AccAddress) *types.MsgApproveUpgrade {
	return &types.MsgApproveUpgrade{
		Creator: signer.String(),
		Name:    testconstants.Plan.Name,
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}
