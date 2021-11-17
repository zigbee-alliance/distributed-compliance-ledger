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

func createNValidatorMissedBlockBitArray(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorMissedBlockBitArray {
	items := make([]types.ValidatorMissedBlockBitArray, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Index = uint64(i)

		keeper.SetValidatorMissedBlockBitArray(ctx, items[i])
	}
	return items
}

func TestValidatorMissedBlockBitArrayGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorMissedBlockBitArray(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorMissedBlockBitArray(ctx,
			item.Address,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorMissedBlockBitArrayRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorMissedBlockBitArray(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorMissedBlockBitArray(ctx,
			item.Address,
			item.Index,
		)
		_, found := keeper.GetValidatorMissedBlockBitArray(ctx,
			item.Address,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestValidatorMissedBlockBitArrayGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorMissedBlockBitArray(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidatorMissedBlockBitArray(ctx))
}
