package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNewVendorInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NewVendorInfo {
	items := make([]types.NewVendorInfo, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNewVendorInfo(ctx, items[i])
	}
	return items
}

func TestNewVendorInfoGet(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNNewVendorInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewVendorInfo(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestNewVendorInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNNewVendorInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewVendorInfo(ctx,
			item.Index,
		)
		_, found := keeper.GetNewVendorInfo(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNewVendorInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.VendorinfoKeeper(t)
	items := createNNewVendorInfo(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllNewVendorInfo(ctx))
}
