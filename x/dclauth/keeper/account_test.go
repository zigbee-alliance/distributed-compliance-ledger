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

func createNAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccount(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx,
			item.Address,
		)
		_, found := keeper.GetAccount(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccount(ctx)),
	)
}
*/
