package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestAccountKeeper_AccessHelpers(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)

	addr := sdk.MustAccAddressFromBech32(sample.AccAddress())
	missing := sdk.MustAccAddressFromBech32(sample.AccAddress())

	acc := types.NewAccount(
		authtypes.NewBaseAccount(addr, nil, 0, 0),
		types.AccountRoles{types.Vendor}, nil, nil, 7,
		[]*commontypes.Uint16Range{{Min: 10, Max: 20}},
	)
	keeper.SetAccountO(ctx, *acc)

	require.NotNil(t, keeper.GetAccount(ctx, addr))
	require.Nil(t, keeper.GetAccount(ctx, missing))

	require.Nil(t, keeper.GetModuleAddress("any"))

	require.True(t, keeper.HasVendorID(ctx, addr, 7))
	require.False(t, keeper.HasVendorID(ctx, addr, 8))
	require.False(t, keeper.HasVendorID(ctx, missing, 7))

	require.True(t, keeper.HasRightsToChange(ctx, addr, 15))
	require.False(t, keeper.HasRightsToChange(ctx, addr, 99))
	require.False(t, keeper.HasRightsToChange(ctx, missing, 15))

	require.Equal(t, 1, keeper.CountAccountsWithRole(ctx, types.Vendor))
	require.Equal(t, 0, keeper.CountAccountsWithRole(ctx, types.Trustee))
}
