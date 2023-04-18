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

func createNModel(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Model {
	items := make([]types.Model, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)

		keeper.SetModel(ctx, items[i])
	}

	return items
}

func TestModelGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil, nil)
	items := createNModel(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetModel(ctx,
			item.Vid,
			item.Pid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestModelRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil, nil)
	items := createNModel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveModel(ctx,
			item.Vid,
			item.Pid,
		)
		_, found := keeper.GetModel(ctx,
			item.Vid,
			item.Pid,
		)
		require.False(t, found)
	}
}

func TestModelGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil, nil)
	items := createNModel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllModel(ctx)),
	)
}
