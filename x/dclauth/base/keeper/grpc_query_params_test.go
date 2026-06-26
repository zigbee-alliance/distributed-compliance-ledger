package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery_Success(t *testing.T) {
	keeper, ctx := setupKeeperWithParams(t)
	wctx := sdk.WrapSDKContext(ctx)

	expected := authtypes.DefaultGenesisState().Params
	keeper.SetParams(ctx, expected)

	resp, err := keeper.Params(wctx, &authtypes.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, expected, resp.Params)
}
