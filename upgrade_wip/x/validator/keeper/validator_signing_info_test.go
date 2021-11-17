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

func createNValidatorSigningInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorSigningInfo {
	items := make([]types.ValidatorSigningInfo, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetValidatorSigningInfo(ctx, items[i])
	}
	return items
}

func TestValidatorSigningInfoGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorSigningInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorSigningInfo(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorSigningInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorSigningInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorSigningInfo(ctx,
			item.Address,
		)
		_, found := keeper.GetValidatorSigningInfo(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestValidatorSigningInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNValidatorSigningInfo(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidatorSigningInfo(ctx))
}
