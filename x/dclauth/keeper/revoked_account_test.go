package keeper_test

/*
import (
	"fmt"
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

func createNRevokedAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RevokedAccount {
	items := make([]types.RevokedAccount, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetRevokedAccount(ctx, items[i])
	}
	return items
}

func TestRevokedAccountGet(t *testing.T) {
	fmt.Print("H!")
	keeper, ctx := keepertest.DclauthKeeper(t)
	fmt.Print("He!")
	items := createNRevokedAccount(keeper, ctx, 10)

	fmt.Print("Hello!")
	for _, item := range items {
		rst, found := keeper.GetRevokedAccount(ctx,
			item.GetAddress(),
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRevokedAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNRevokedAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRevokedAccount(ctx,
			item.GetAddress(),
		)
		_, found := keeper.GetRevokedAccount(ctx,
			item.GetAddress(),
		)
		require.False(t, found)
	}
}

func TestRevokedAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNRevokedAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRevokedAccount(ctx)),
	)
}
*/
