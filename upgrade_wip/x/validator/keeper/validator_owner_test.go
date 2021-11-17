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

func createNValidatorOwner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorOwner {
	items := make([]types.ValidatorOwner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetValidatorOwner(ctx, items[i])
	}
	return items
}

func TestValidatorOwnerGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorOwner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorOwner(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorOwnerRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorOwner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorOwner(ctx,
			item.Address,
		)
		_, found := keeper.GetValidatorOwner(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestValidatorOwnerGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorOwner(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidatorOwner(ctx))
}
