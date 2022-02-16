package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNApprovedUpgrade(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedUpgrade {
	items := make([]types.ApprovedUpgrade, n)
	for i := range items {
		items[i].Plan.Name = strconv.Itoa(i)

		keeper.SetApprovedUpgrade(ctx, items[i])
	}
	return items
}

func TestApprovedUpgradeGet(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNApprovedUpgrade(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApprovedUpgrade(ctx,
			item.Plan.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestApprovedUpgradeRemove(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNApprovedUpgrade(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApprovedUpgrade(ctx,
			item.Plan.Name,
		)
		_, found := keeper.GetApprovedUpgrade(ctx,
			item.Plan.Name,
		)
		require.False(t, found)
	}
}

func TestApprovedUpgradeGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNApprovedUpgrade(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApprovedUpgrade(ctx)),
	)
}
