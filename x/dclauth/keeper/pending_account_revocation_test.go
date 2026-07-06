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

func createNPendingAccountRevocation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingAccountRevocation {
	items := make([]types.PendingAccountRevocation, n)
	for i := range items {
		items[i] = types.PendingAccountRevocation{Address: sample.AccAddress()}

		keeper.SetPendingAccountRevocation(ctx, items[i])
	}

	return items
}

func TestPendingAccountRevocationGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccountRevocation(keeper, ctx, 10)
	for _, item := range items {
		addr := sdk.MustAccAddressFromBech32(item.Address)
		rst, found := keeper.GetPendingAccountRevocation(ctx, addr)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestPendingAccountRevocationRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccountRevocation(keeper, ctx, 10)
	for _, item := range items {
		addr := sdk.MustAccAddressFromBech32(item.Address)
		keeper.RemovePendingAccountRevocation(ctx, addr)
		_, found := keeper.GetPendingAccountRevocation(ctx, addr)
		require.False(t, found)
	}
}

func TestPendingAccountRevocationGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccountRevocation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPendingAccountRevocation(ctx)),
	)
}
