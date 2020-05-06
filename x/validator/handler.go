package validator

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/functions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k, authzKeeper)

		default:
			errMsg := fmt.Sprintf("unrecognized validator Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateValidator(ctx sdk.Context, msg types.MsgCreateValidator, k Keeper, authzKeeper authz.Keeper) sdk.Result {
	//if !authzKeeper.HasRole(ctx, sdk.AccAddress(msg.ValidatorAddress), authz.NodeAdmin) {
	//	return sdk.ErrUnauthorized(fmt.Sprintf("MsgCreateValidator transaction should be signed by an account with the %s role", authz.NodeAdmin)).Result()
	//}

	if k.IsValidatorPresent(ctx, msg.ValidatorAddress) {
		return types.ErrValidatorOperatorAddressExists(msg.ValidatorAddress).Result()
	}

	if k.IsValidatorByConsAddrPresent(ctx, sdk.GetConsAddress(msg.GetPubKey())) {
		return types.ErrValidatorPubKeyExists(msg.ValidatorAddress).Result()
	}

	if ctx.ConsensusParams() != nil {
		tmPubKey := tmtypes.TM2PB.PubKey(msg.GetPubKey())
		if !functions.StringInSlice(tmPubKey.Type, ctx.ConsensusParams().Validator.PubKeyTypes) {
			return types.ErrValidatorPubKeyTypeNotSupported(tmPubKey.Type, ctx.ConsensusParams().Validator.PubKeyTypes).Result()
		}
	}

	validator := NewValidator(msg.ValidatorAddress, msg.PubKey, msg.Description)

	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddr(ctx, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}
