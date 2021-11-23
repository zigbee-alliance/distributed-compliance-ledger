package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)

		keeper.SetValidator(ctx, items[i])
	}
	return items
}

func TestValidatorGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidator(ctx,
			item.Owner,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidator(ctx,
			item.Owner,
		)
		_, found := keeper.GetValidator(ctx,
			item.Owner,
		)
		require.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidator(ctx))
}
