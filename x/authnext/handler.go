package authnext

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(accKeeper AccountKeeper, authzKeeper authz.Keeper, cdc *codec.Codec) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddAccount:
			return handleMsgAddAccount(ctx, accKeeper, authzKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddAccount(ctx sdk.Context, keeper AccountKeeper, authzKeeper authz.Keeper, msg types.MsgAddAccount) sdk.Result {
	// check if sender has enough rights to create account
	if !authzKeeper.HasRole(ctx, msg.Signer, authz.Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgAddAccount transaction should be signed by an account with the %s role", authz.Trustee)).Result()
	}

	// check if account already exists
	if account := keeper.GetAccount(ctx, msg.Address); account != nil {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Account associated with the address=%v already exists on the ledger", msg.Address)).Result()
	}

	pubKey, err := sdk.GetAccPubKeyBech32(msg.PublicKey)
	if err != nil {
		return sdk.ErrInvalidPubKey(err.Error()).Result()
	}

	// create account and fill key
	account := keeper.NewAccountWithAddress(ctx, msg.Address)
	err = account.SetPubKey(pubKey)
	if err != nil {
		return sdk.ErrInvalidPubKey(err.Error()).Result()
	}

	// store account
	keeper.SetAccount(ctx, account)

	return sdk.Result{}
}
