package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAccountRoles = "account_roles"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAccountRoles:
			return queryAccountRoles(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown authz query endpoint")
		}
	}
}

func queryAccountRoles(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInvalidAddress(path[0])
	}

	if !keeper.IsAccountRolesPresent(ctx, addr) {
		return nil, types.ErrAccountRolesDoesNotExist(types.DefaultCodespace)
	}

	accountRoles := keeper.GetAccountRoles(ctx, addr)

	res, err := codec.MarshalJSONIndent(keeper.cdc, accountRoles)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
