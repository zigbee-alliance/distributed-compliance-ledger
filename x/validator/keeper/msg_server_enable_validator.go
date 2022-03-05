package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) EnableValidator(goCtx context.Context, msg *types.MsgEnableValidator) (*types.MsgEnableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.EnableDisableValidatorRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Enable validator transaction should be signed by an account with the %s role",
			types.EnableDisableValidatorRole,
		)
	}

	// check if disabled validator exists
	_, isFound := k.GetDisabledValidator(ctx, msg.Address)
	if isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	// Enable validator
	validator, _ := k.GetValidator(ctx, sdk.ValAddress(msg.Address))
	k.Unjail(ctx, validator)

	// store disabled validator
	k.RemoveDisabledValidator(ctx, msg.Address)

	return &types.MsgEnableValidatorResponse{}, nil
}
