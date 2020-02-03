package authz

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAssignRole:
			return handleMsgAssignRole(ctx, keeper, msg)
		case types.MsgRevokeRole:
			return handleMsgRevokeRole(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAssignRole(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAssignRole) sdk.Result {
	if !keeper.HasRole(ctx, msg.Signer, Administrator) {
		return sdk.ErrUnauthorized("you are not authorized to perform this action").Result()
	}

	keeper.AssignRole(ctx, msg.Address, msg.Role)

	return sdk.Result{}
}

func handleMsgRevokeRole(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgRevokeRole) sdk.Result {
	if !keeper.HasRole(ctx, msg.Signer, Administrator) {
		return sdk.ErrUnauthorized("you are not authorized to perform this action").Result()
	}

	if !keeper.HasRole(ctx, msg.Address, msg.Role) {
		return sdk.ErrUnauthorized(fmt.Sprintf("account %s doesn't have role %s", msg.Address.String(), msg.Role)).Result()
	}

	if msg.Role == Administrator && keeper.CountAccounts(ctx, Administrator) < 2 {
		return sdk.ErrUnauthorized("there must be at least one administrator").Result()
	}

	keeper.RevokeRole(ctx, msg.Address, msg.Role)

	return sdk.Result{}
}
