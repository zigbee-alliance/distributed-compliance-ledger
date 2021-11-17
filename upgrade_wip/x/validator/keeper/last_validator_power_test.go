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

func createNLastValidatorPower(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LastValidatorPower {
	items := make([]types.LastValidatorPower, n)
	for i := range items {
		items[i].ConsensusAddress = strconv.Itoa(i)

		keeper.SetLastValidatorPower(ctx, items[i])
	}
	return items
}

func TestLastValidatorPowerGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLastValidatorPower(ctx,
			item.ConsensusAddress,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestLastValidatorPowerRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLastValidatorPower(ctx,
			item.ConsensusAddress,
		)
		_, found := keeper.GetLastValidatorPower(ctx,
			item.ConsensusAddress,
		)
		require.False(t, found)
	}
}

func TestLastValidatorPowerGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllLastValidatorPower(ctx))
}
