package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func createNRejectedAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RejectedAccount {
	items := make([]types.RejectedAccount, n)
	for i := range items {
		acc := newTestAccount(sample.AccAddress())
		items[i] = types.RejectedAccount{Account: &acc}

		keeper.SetRejectedAccount(ctx, items[i])
	}

	return items
}

func TestRejectedAccountGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNRejectedAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRejectedAccount(ctx, item.GetAddress())
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRejectedAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNRejectedAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRejectedAccount(ctx, item.GetAddress())
		_, found := keeper.GetRejectedAccount(ctx, item.GetAddress())
		require.False(t, found)
	}
}

func TestRejectedAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNRejectedAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRejectedAccount(ctx)),
	)
}
