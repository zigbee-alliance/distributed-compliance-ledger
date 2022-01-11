package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
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

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNVendorInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VendorInfo {
	items := make([]types.VendorInfo, n)
	for i := range items {
		items[i].VendorID = int32(i)

		keeper.SetVendorInfo(ctx, items[i])
	}
	return items
}

func TestVendorInfoGet(t *testing.T) {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	items := createNVendorInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVendorInfo(ctx,
			item.VendorID,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestVendorInfoRemove(t *testing.T) {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	items := createNVendorInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVendorInfo(ctx,
			item.VendorID,
		)
		_, found := keeper.GetVendorInfo(ctx,
			item.VendorID,
		)
		require.False(t, found)
	}
}

func TestVendorInfoGetAll(t *testing.T) {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	items := createNVendorInfo(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllVendorInfo(ctx))
}
