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
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	sdktestdata "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/testdata"
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
	Plan types.Plan
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

func (setup *TestSetup) ScheduleUpgradeBehavior(
	plan types.Plan,
	err error,
) {
	setup.UpgradeKeeper.On("ScheduleUpgrade", mock.Anything, plan).Return(err)
}

func Setup(t *testing.T) TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	upgradeKeeper := &UpgradeKeeperMock{}
	keeper, ctx := testkeeper.DclupgradeKeeper(t, dclauthKeeper, upgradeKeeper)

	trustee := sdktestdata.GenerateAccAddress()
	plan := testconstants.Plan
	setup := TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		UpgradeKeeper: upgradeKeeper,
		Handler:       NewHandler(*keeper),
		Plan:          plan,
	}

	setup.AddAccount(trustee, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.ScheduleUpgradeBehavior(plan, nil)
	return setup
}

func TestHandler_ProposedUpgradeExists(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// propose the same upgrade
	_, err = setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err)
	require.True(t, types.ErrProposedUpgradeAlreadyExists.Is(err))
}

func TestHandler_OnlyTrusteeCanProposeUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := sdktestdata.GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// propose upgrade by user without Trustee role
		msgProposeUpgrade := NewMsgProposeUpgrade(accAddress)
		_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
		setup.UpgradeKeeper.AssertNotCalled(t, "ScheduleUpgrade", mock.Anything, msgProposeUpgrade.Plan)
	}

	// propose and approve upgrade by Trustees (3 Trustees, >=2/3 approvals needed)

	// propose upgrade by user with Trustee role
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
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
}

func TestHandler_ApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose and approve upgrade by Trustees (3 Trustees, >=2/3 approvals needed)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
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
}

func TestHandler_OnlyTrusteeCanApproveUpgrade(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.TestHouse,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Vendor,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := sdktestdata.GenerateAccAddress()
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

	setup.UpgradeKeeper.AssertCalled(
		t,
		"ScheduleUpgrade",
		mock.MatchedBy(isContextWithNonCachedMultiStore),
		msgProposeUpgrade.Plan,
	)
}

func TestHandler_ProposedUpgradeDoesNotExist(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
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

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress2 := sdktestdata.GenerateAccAddress()
	trusteeAccAddress3 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress2, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.AddAccount(trusteeAccAddress3, []dclauthtypes.AccountRole{dclauthtypes.Trustee})

	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(3)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.NoError(t, err)

	// approve upgrade
	msgApproveUpgrade := NewMsgApproveUpgrade(trusteeAccAddress1)
	_, err = setup.Handler(setup.Ctx, msgApproveUpgrade)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func TestHandler_PlanHeightLessBlockHeight(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(1)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Height = 1
	setup.Ctx = setup.Ctx.WithBlockHeight(100)
	blockHeight := setup.Ctx.BlockHeight()

	require.Less(t, msgProposeUpgrade.Plan.Height, blockHeight)

	setup.ScheduleUpgradeBehavior(msgProposeUpgrade.Plan, sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func TestHandler_TimeIsNotZero(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(1)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Time = time.Now()

	require.False(t, msgProposeUpgrade.Plan.Time.IsZero())

	setup.ScheduleUpgradeBehavior(msgProposeUpgrade.Plan, sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func TestHandler_UpgradedClientStateNotNil(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(1)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.UpgradedClientState = &codectypes.Any{TypeUrl: "123"}

	require.True(t, msgProposeUpgrade.Plan.UpgradedClientState != nil)

	setup.ScheduleUpgradeBehavior(msgProposeUpgrade.Plan, sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func TestHandler_PlanNameLenIs0(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(1)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Name = ""

	require.Equal(t, len(msgProposeUpgrade.Plan.Name), 0)

	setup.ScheduleUpgradeBehavior(msgProposeUpgrade.Plan, sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func TestHandler_PlanHeightLessOrEqual0(t *testing.T) {
	setup := Setup(t)

	trusteeAccAddress1 := sdktestdata.GenerateAccAddress()
	setup.AddAccount(trusteeAccAddress1, []dclauthtypes.AccountRole{dclauthtypes.Trustee})
	setup.DclauthKeeper.On("CountAccountsWithRole", mock.Anything, dclauthtypes.Trustee).Return(1)

	// propose new upgrade
	msgProposeUpgrade := NewMsgProposeUpgrade(trusteeAccAddress1)
	msgProposeUpgrade.Plan.Height = -1

	require.LessOrEqual(t, msgProposeUpgrade.Plan.Height, int64(0))

	setup.ScheduleUpgradeBehavior(msgProposeUpgrade.Plan, sdkerrors.ErrInvalidRequest)
	_, err := setup.Handler(setup.Ctx, msgProposeUpgrade)
	require.Error(t, err, sdkerrors.ErrInvalidRequest)
}

func isContextWithNonCachedMultiStore(ctx sdk.Context) bool {
	_, ok := ctx.MultiStore().(storetypes.CacheMultiStore)
	return !ok // not CacheMultiStore
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
