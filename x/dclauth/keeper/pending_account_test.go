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

func createNPendingAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingAccount {
	items := make([]types.PendingAccount, n)
	for i := range items {
		acc := newTestAccount(sample.AccAddress())
		items[i] = types.PendingAccount{Account: &acc}

		keeper.SetPendingAccount(ctx, items[i])
	}

	return items
}

func TestPendingAccountGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPendingAccount(ctx, item.GetAddress())
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestPendingAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePendingAccount(ctx, item.GetAddress())
		_, found := keeper.GetPendingAccount(ctx, item.GetAddress())
		require.False(t, found)
	}
}

func TestPendingAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPendingAccount(ctx)),
	)
}
