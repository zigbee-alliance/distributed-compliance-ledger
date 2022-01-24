package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNVendorProducts(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VendorProducts {
	items := make([]types.VendorProducts, n)
	for i := range items {
		items[i].Vid = int32(i)

		keeper.SetVendorProducts(ctx, items[i])
	}
	return items
}

func TestVendorProductsGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil)
	items := createNVendorProducts(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVendorProducts(ctx,
			item.Vid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVendorProductsRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil)
	items := createNVendorProducts(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVendorProducts(ctx,
			item.Vid,
		)
		_, found := keeper.GetVendorProducts(ctx,
			item.Vid,
		)
		require.False(t, found)
	}
}

func TestVendorProductsGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil)
	items := createNVendorProducts(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVendorProducts(ctx)),
	)
}
