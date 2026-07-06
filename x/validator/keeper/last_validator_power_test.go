package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func createNLastValidatorPower(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LastValidatorPower {
	items := make([]types.LastValidatorPower, n)
	for i := range items {
		items[i].Owner = sample.ValAddress()
		items[i].Power = int32(i)

		keeper.SetLastValidatorPower(ctx, items[i])
	}

	return items
}

func TestLastValidatorPowerGet(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		owner, err := sdk.ValAddressFromBech32(item.Owner)
		require.NoError(t, err)

		rst, found := keeper.GetLastValidatorPower(ctx, owner)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestLastValidatorPowerRemove(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		owner, err := sdk.ValAddressFromBech32(item.Owner)
		require.NoError(t, err)

		keeper.RemoveLastValidatorPower(ctx, owner)
		_, found := keeper.GetLastValidatorPower(ctx, owner)
		require.False(t, found)
	}
}

func TestLastValidatorPowerGetAll(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	items := createNLastValidatorPower(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllLastValidatorPower(ctx))
}

func TestLastValidatorPowerCount(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	items := createNLastValidatorPower(keeper, ctx, 10)
	require.Equal(t, len(items), keeper.CountLastValidators(ctx))
}
