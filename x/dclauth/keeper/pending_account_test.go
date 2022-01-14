package keeper_test

/*

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPendingAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingAccount {
	items := make([]types.PendingAccount, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetPendingAccount(ctx, items[i])
	}
	return items
}

func TestPendingAccountGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNPendingAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPendingAccount(ctx,
			item.Address,
		)
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
		keeper.RemovePendingAccount(ctx,
			item.Address,
		)
		_, found := keeper.GetPendingAccount(ctx,
			item.Address,
		)
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
*/
