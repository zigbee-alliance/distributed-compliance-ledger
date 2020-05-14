package auth

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddAccount:
			return handleMsgAddAccount(ctx, keeper, msg)
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

func handleMsgAddAccount(ctx sdk.Context, keeper Keeper, msg types.MsgAddAccount) sdk.Result {
	// check if sender has enough rights to create account
	if !keeper.HasRole(ctx, msg.Signer, types.Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgAddAccount transaction should be signed by an account with the %s role", types.Trustee)).Result()
	}

	// check if account already exists
	if keeper.IsAccountPresent(ctx, msg.Address) {
		return types.ErrAccountAlreadyExistExist(msg.Address).Result()
	}

	pubKey, err := sdk.GetAccPubKeyBech32(msg.PublicKey)
	if err != nil {
		return sdk.ErrInvalidPubKey(err.Error()).Result()
	}

	// create and store account
	account := keeper.NewAccount(ctx, types.NewAccount(msg.Address, pubKey, msg.Roles))
	keeper.SetAccount(ctx, account)

	return sdk.Result{}
}

func handleMsgAssignRole(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAssignRole) sdk.Result {
	// check if sender has enough rights to assign role
	if !keeper.HasRole(ctx, msg.Signer, Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgAssignRole transaction should be signed by an account with the %s role", Trustee)).Result()
	}

	// assign role to account
	keeper.AssignRole(ctx, msg.Address, msg.Role)

	return sdk.Result{}
}

func handleMsgRevokeRole(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgRevokeRole) sdk.Result {
	// check if sender has enough rights to revoke role
	if !keeper.HasRole(ctx, msg.Signer, Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgRevokeRole transaction should be signed by an account with the %s role", Trustee)).Result()
	}

	// check if target account has role to revoke
	if !keeper.HasRole(ctx, msg.Address, msg.Role) {
		return sdk.ErrUnauthorized(fmt.Sprintf("Account %s doesn't have role %s to revoke",
			msg.Address.String(), msg.Role)).Result()
	}

	// at least one trustee must be on the ledger
	if msg.Role == Trustee && keeper.CountAccountsWithRole(ctx, Trustee) < 2 {
		return sdk.ErrUnauthorized("there must be at least one Trustee").Result()
	}

	// remove role
	keeper.RevokeRole(ctx, msg.Address, msg.Role)

	return sdk.Result{}
}
