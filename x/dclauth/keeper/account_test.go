package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// newTestAccount builds a valid DCL Account (with a base account) for the given
// bech32 address. Shared by the dclauth keeper test files in this package.
func newTestAccount(addr string) types.Account {
	accAddr := sdk.MustAccAddressFromBech32(addr)
	ba := authtypes.NewBaseAccount(accAddr, nil, 0, 0)

	return *types.NewAccount(ba, nil, nil, nil, 0, nil)
}

func createNAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i] = newTestAccount(sample.AccAddress())

		keeper.SetAccountO(ctx, items[i])
	}

	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountO(ctx, item.GetAddress())
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
		keeper.RemoveAccount(ctx, item.GetAddress())
		_, found := keeper.GetAccountO(ctx, item.GetAddress())
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
