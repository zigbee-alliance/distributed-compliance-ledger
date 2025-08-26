package keeper

import (
	"context"

	"cosmossdk.io/errors"
	tmstrings "github.com/cometbft/cometbft/libs/strings"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	// check if sender has enough rights to create a validator node
	if !k.dclauthKeeper.HasRole(ctx, sdk.AccAddress(valAddr), dclauthtypes.NodeAdmin) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"CreateValidator transaction should be signed by an account with the \"%s\" role",
			types.EnableDisableValidatorRole,
		)
	}

	// check if we has not reached the limit of nodes
	if k.CountLastValidators(ctx) == types.MaxNodes {
		return nil, types.ErrPoolIsFull()
	}

	// check to see if the pubkey or sender has been registered before
	if _, found := k.GetValidator(ctx, valAddr); found {
		return nil, sdkstakingtypes.ErrValidatorOwnerExists
	}

	pk, ok := msg.PubKey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pk)
	}

	if _, found := k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk)); found {
		return nil, sdkstakingtypes.ErrValidatorPubKeyExists
	}

	if err := msg.Description.Validate(); err != nil {
		return nil, err
	}

	// check key type
	cp := ctx.ConsensusParams()
	if cp != nil && cp.Validator != nil {
		if pk.Type() != "tendermint/PubKeyEd25519" {
			return nil, errors.Wrapf(
				sdkstakingtypes.ErrValidatorPubKeyTypeNotSupported,
				"consensus pubkey %s is not supported (only ed25519 allowed)",
				pk.Type(),
			)
		}
		if !tmstrings.StringInSlice(pk.Type(), cp.Validator.PubKeyTypes) {
			return nil, errors.Wrapf(
				sdkstakingtypes.ErrValidatorPubKeyTypeNotSupported,
				"got: %s, expected: %s", pk.Type(), cp.Validator.PubKeyTypes,
			)
		}
	}

	validator, err := types.NewValidator(valAddr, pk, msg.Description)
	if err != nil {
		return nil, err
	}

	k.SetValidator(ctx, validator)
	err = k.SetValidatorByConsAddr(ctx, validator)
	if err != nil {
		return nil, err
	}

	// TODO issue 99: val after-hooks if needed
	// ...

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgCreateValidatorResponse{}, nil
}
