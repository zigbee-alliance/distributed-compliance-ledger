package dclupgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/mock"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
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

func (m *DclauthKeeperMock) CountAccountsWithRole(ctx sdk.Context, roleToCount dclauthtypes.AccountRole) int {
	return 5
}

type UpgradeKeeperMock struct {
	mock.Mock
}

func (m *UpgradeKeeperMock) ScheduleUpgrade(ctx sdk.Context, plan upgradetypes.Plan) error {
	args := m.Called(ctx, plan)
	return args.Error(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}
