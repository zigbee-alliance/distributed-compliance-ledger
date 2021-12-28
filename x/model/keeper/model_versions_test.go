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

func createNModelVersions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ModelVersions {
	items := make([]types.ModelVersions, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)

		keeper.SetModelVersions(ctx, items[i])
	}
	return items
}

func TestModelVersionsGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersions(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetModelVersions(ctx,
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

func TestModelVersionsRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersions(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveModelVersions(ctx,
			item.Vid,
			item.Pid,
		)
		_, found := keeper.GetModelVersions(ctx,
			item.Vid,
			item.Pid,
		)
		require.False(t, found)
	}
}

func TestModelVersionsGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllModelVersions(ctx)),
	)
}
