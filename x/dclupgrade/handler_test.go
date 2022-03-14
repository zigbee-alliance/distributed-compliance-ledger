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

package dclupgrade

import (
	"context"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/testdata"
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
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	for _, role := range roles {
		setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func Setup(t *testing.T) TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	upgradeKeeper := &UpgradeKeeperMock{}
	keeper, ctx := testkeeper.DclupgradeKeeper(t, dclauthKeeper, upgradeKeeper)

	setup := TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		UpgradeKeeper: upgradeKeeper,
		Handler:       NewHandler(*keeper),
	}

	return setup
}

func TestHandler_ProposedUpgradeExists(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// propose the same upgrade
	msgProposeUpgrade.Creator = trusteeAccAddress2.String()
	_, err = setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrProposedUpgradeAlreadyExists.Is(err))
}

func TestHandler_OnlyTrusteeCanProposeUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// propose upgrade by user without Trustee role
		msgProposeUpgrade := NewMsgProposeUpgrade(accAddress)
		_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	// propose upgrade by user with Trustee role
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)
}

func TestHandler_ProposeUpgradeWhenSeveralVotesNeeded(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(
		t,
		"ScheduleUpgrade",
		mock.MatchedBy(isContextWithCachedMultiStore),
		msgProposeUpgrade.Plan,
	)

	setup.UpgradeKeeper.AssertNotCalled(
		t,
		"ScheduleUpgrade",
		mock.MatchedBy(isContextWithNonCachedMultiStore),
		msgProposeUpgrade.Plan,
	)

	// check proposed upgrade for being created
	proposedUpgrade, isFound := setup.Keeper.GetProposedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.True(t, isFound)

	require.Equal(t, msgProposeUpgrade.Plan, proposedUpgrade.Plan)
	require.Equal(t, msgProposeUpgrade.Creator, proposedUpgrade.Creator)

	require.Equal(t, 1, len(proposedUpgrade.Approvals))

	require.Equal(t, msgProposeUpgrade.Creator, proposedUpgrade.Approvals[0].Address)
	require.Equal(t, msgProposeUpgrade.Time, proposedUpgrade.Approvals[0].Time)
	require.Equal(t, msgProposeUpgrade.Info, proposedUpgrade.Approvals[0].Info)

	// check approved upgrade for not being created
	_, isFound = setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)
}

// TODO Implement this test
// func TestHandler_ProposeUpgradeWhenOneVoteNeeded(t *testing.T) {
// 	setup := Setup(t)

// 	trusteeAccAddress1 := testdata.GenerateAccAddress()
// 	trusteeAccAddress2 := testdata.GenerateAccAddress()
// 	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
// 	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

// 	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(2)

// 	// propose upgrade
// 	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

// 	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
// 	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
// 	require.NoError(t, err)

// 	_, isFound := setup.Keeper.GetProposedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
// 	require.False(t, isFound)
// }

func TestHandler_ApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose and approve upgrade by Trustees (3 Trustees, >=2/3 approvals needed)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress2)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)

	setup.UpgradeKeeper.AssertCalled(
		t,
		"ScheduleUpgrade",
		mock.MatchedBy(isContextWithNonCachedMultiStore),
		msgProposeUpgrade.Plan,
	)

	// check proposed upgrade for being deleted
	_, isFound := setup.Keeper.GetProposedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)

	// check upgrade for being added to ApprovedUpgrade store
	approvedUpgrade, isFound := setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.True(t, isFound)

	// verification of ApprovedUpgrade
	require.Equal(t, msgProposeUpgrade.Plan, approvedUpgrade.Plan)
	require.Equal(t, msgProposeUpgrade.Creator, approvedUpgrade.Creator)

	// one approval is from propose stage, another is from approve stage
	require.Equal(t, 2, len(approvedUpgrade.Approvals))

	require.Equal(t, msgProposeUpgrade.Creator, approvedUpgrade.Approvals[0].Address)
	require.Equal(t, msgProposeUpgrade.Time, approvedUpgrade.Approvals[0].Time)
	require.Equal(t, msgProposeUpgrade.Info, approvedUpgrade.Approvals[0].Info)

	require.Equal(t, msgApproveUpgrade.Creator, approvedUpgrade.Approvals[1].Address)
	require.Equal(t, msgApproveUpgrade.Time, approvedUpgrade.Approvals[1].Time)
	require.Equal(t, msgApproveUpgrade.Info, approvedUpgrade.Approvals[1].Info)
}

func TestHandler_ProposeUpgradeWhenApprovedUpgradeWithSameNameExists(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose and approve upgrade by Trustees (3 Trustees, >=2/3 approvals needed)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress2)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)

	// check proposed upgrade for being deleted
	_, isFound := setup.Keeper.GetProposedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)

	// check upgrade for being added to ApprovedUpgrade store
	_, isFound = setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.True(t, isFound)

	// propose upgrade with the same name
	msgProposeUpgrade.Creator = trusteeAccAddress3.String()
	_, err = setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrApprovedUpgradeAlreadyExists.Is(err))
}

func TestHandler_UpgradeApprovalWhenMoreVotesNeeded(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	trusteeAccAddress4 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress4, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(4)

	// propose and approve upgrade by Trustees (4 Trustees, >=2/3 approvals needed)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress2)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)

	// check proposed proposedUpgrade for not being deleted
	proposedUpgrade, isFound := setup.Keeper.GetProposedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.True(t, isFound)

	require.Equal(t, msgProposeUpgrade.Plan, proposedUpgrade.Plan)
	require.Equal(t, msgProposeUpgrade.Creator, proposedUpgrade.Creator)

	// one approval is from propose stage, another is from approve stage
	require.Equal(t, 2, len(proposedUpgrade.Approvals))

	require.Equal(t, msgProposeUpgrade.Creator, proposedUpgrade.Approvals[0].Address)
	require.Equal(t, msgProposeUpgrade.Time, proposedUpgrade.Approvals[0].Time)
	require.Equal(t, msgProposeUpgrade.Info, proposedUpgrade.Approvals[0].Info)

	require.Equal(t, msgApproveUpgrade.Creator, proposedUpgrade.Approvals[1].Address)
	require.Equal(t, msgApproveUpgrade.Time, proposedUpgrade.Approvals[1].Time)
	require.Equal(t, msgApproveUpgrade.Info, proposedUpgrade.Approvals[1].Info)

	// check upgrade for not being added to ApprovedUpgrade store
	_, isFound = setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)
}

func TestHandler_OnlyTrusteeCanApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := testdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// approve upgrade by user without Trustee role
		msgApproveUpgrade := NewMsgApproveUpgrade(accAddress)
		_, err := setup.Handler(setup.Ctx, msgApproveUpgrade)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}

	// approve upgrade by user with Trustee role
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress2)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.NoError(t, err)
}

func TestHandler_ProposedUpgradeDoesNotExist(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress1)
	_, err := setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrProposedUpgradeDoesNotExist.Is(err))
}

func TestHandler_MessageCreatorAlreadyApprovedUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress1)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// check upgrade for not being added to ApprovedUpgrade store
	_, isFound := setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)
}

func TestHandler_ProposeUpgradePlanHeightLessBlockHeight(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade with plan height < block height
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Height = 1
	setup.Ctx = setup.Ctx.WithBlockHeight(100)

	// error returned because height in plan is less than block height on the propose stage
	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func TestHandler_ApproveUpgradePlanHeightLessBlockHeight(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := testdata.GenerateAccAddress()
	trusteeAccAddress2 := testdata.GenerateAccAddress()
	trusteeAccAddress3 := testdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Height = 2
	setup.Ctx = setup.Ctx.WithBlockHeight(1)

	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(nil).Once()
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// create approve message from trustee2
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress2)

	// approve new upgrade with plan height < block height
	setup.Ctx = setup.Ctx.WithBlockHeight(3)
	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan).Return(sdkerrors.ErrInvalidRequest).Once()
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)

	// check upgrade for not being added to ApprovedUpgrade store
	_, isFound := setup.Keeper.GetApprovedUpgrade(setup.Ctx, msgProposeUpgrade.Plan.Name)
	require.False(t, isFound)
}

func isContextWithCachedMultiStore(ctx sdk.Context) bool {
	_, ok := ctx.MultiStore().(storetypes.CacheMultiStore)
	return ok // CacheMultiStore
}

func isContextWithNonCachedMultiStore(ctx sdk.Context) bool {
	_, ok := ctx.MultiStore().(storetypes.CacheMultiStore)
	return !ok // not CacheMultiStore
}

func NewMsgProposeUpgrade(signer sdk.AccAddress) *types.MsgProposeUpgrade {
	return &types.MsgProposeUpgrade{
		Creator: signer.String(),
		Plan: types.Plan{
			Name:   testconstants.UpgradePlanName,
			Height: testconstants.UpgradePlanHeight,
			Info:   testconstants.UpgradePlanInfo,
		},
		Info: testconstants.Info,
		Time: testconstants.Time,
	}
}

func NewMsgApproveUpgrade(signer sdk.AccAddress) *types.MsgApproveUpgrade {
	return &types.MsgApproveUpgrade{
		Creator: signer.String(),
		Name:    testconstants.UpgradePlanName,
		Info:    testconstants.Info2,
		Time:    testconstants.Time2,
	}
}
