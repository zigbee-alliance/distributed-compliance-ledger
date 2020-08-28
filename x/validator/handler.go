package validator

import (
	"fmt"
	"strings"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/string"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

func NewHandler(k Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k, authKeeper)
		default:
			errMsg := fmt.Sprintf("unrecognized validator Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateValidator(ctx sdk.Context, msg types.MsgCreateValidator,
	k Keeper, authKeeper auth.Keeper) sdk.Result {
	// check if sender has enough rights to create a validator node
	if !authKeeper.HasRole(ctx, msg.Signer, auth.NodeAdmin) {
		return sdk.ErrUnauthorized(fmt.Sprintf("CreateValidator transaction should be "+
			"signed by an account with the \"%s\" role", auth.NodeAdmin)).Result()
	}

	if k.AccountHasValidator(ctx, msg.Signer) {
		return types.ErrAccountAlreadyHasNode(msg.Signer).Result()
	}

	// check if we has not reached the limit of nodes
	if k.CountLastValidators(ctx) == types.MaxNodes {
		return types.ErrPoolIsFull().Result()
	}

	// check if a validator with a given address already exists
	if k.IsValidatorPresent(ctx, msg.Address) {
		return types.ErrValidatorExists(msg.Address).Result()
	}

	// check key type
	if ctx.ConsensusParams() != nil {
		tmPubKey := tmtypes.TM2PB.PubKey(msg.GetPubKey())
		if !string.StringInSlice(tmPubKey.Type, ctx.ConsensusParams().Validator.PubKeyTypes) {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Validator pubkey type \"%s\" is not supported. Supported types: [%s]",
					tmPubKey.Type, strings.Join(ctx.ConsensusParams().Validator.PubKeyTypes, ","))).Result()
		}
	}

	// create and store validator
	validator := NewValidator(msg.Address, msg.PubKey, msg.Description, msg.Signer)

	k.SetValidator(ctx, validator)
	k.SetValidatorOwner(ctx, msg.Signer, msg.Address)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Address.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}
