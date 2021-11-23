package keeper

import (
	"context"
	"fmt"

	// sdkvalerrors "github.com/cosmos/cosmos-sdk/staking/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	// check if sender has enough rights to create a validator node
	if !k.dclauthKeeper.HasRole(ctx, sdk.AccAddress(valAddr), auth.NodeAdmin) {
		return sdk.ErrUnauthorized(fmt.Sprintf("CreateValidator transaction should be "+
			"signed by an account with the \"%s\" role", auth.NodeAdmin)).Result()
	}

	// check if we has not reached the limit of nodes
	if k.CountLastValidators(ctx) == types.MaxNodes {
		return types.ErrPoolIsFull().Result()
	}

	// check to see if the pubkey or sender has been registered before
	if _, found := k.GetValidator(ctx, valAddr); found {
		return nil, types.ErrValidatorExists(msg.Signer).Result()
	}

	pk, ok := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdk.Wrapf(sdk.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pk)
	}

	if _, found := k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk)); found {
		return nil, sdkvalerrors.ErrValidatorPubKeyExists
	}

	if _, err := msg.Description.Validate(); err != nil {
		return nil, err
	}

	// check key type
	cp := ctx.ConsensusParams()
	if cp != nil && cp.Validator != nil {
		if !tmstrings.StringInSlice(pk.Type(), cp.Validator.PubKeyTypes) {
			return nil, sdk.Wrapf(
				sdkvalerrors.ErrValidatorPubKeyTypeNotSupported,
				"got: %s, expected: %s", pk.Type(), cp.Validator.PubKeyTypes,
			)
		}
	}

	validator, err := types.NewValidator(valAddr, pk, msg.Description)
	if err != nil {
		return nil, err
	}

	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddr(ctx, validator)

	// TODO issue 99: vall after- hooks if needed
	// ...

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

	return &types.MsgCreateValidatorResponse{}, nil
}
