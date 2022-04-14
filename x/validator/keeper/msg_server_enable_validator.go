package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) EnableValidator(goCtx context.Context, msg *types.MsgEnableValidator) (*types.MsgEnableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, sdk.AccAddress(creatorAddr), types.EnableDisableValidatorRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Enable validator transaction should be signed by an account with the %s role",
			types.EnableDisableValidatorRole,
		)
	}

	// check if disabled validator exists
	_, isFound := k.GetDisabledValidator(ctx, msg.Creator)
	if !isFound {
		return nil, types.NewErrDisabledValidatorDoesNotExist(msg.Creator)
	}

	// Enable validator
	validator, _ := k.GetValidator(ctx, creatorAddr)
	k.Unjail(ctx, validator)

	// remove disabled validator from store
	k.RemoveDisabledValidator(ctx, msg.Creator)

	return &types.MsgEnableValidatorResponse{}, nil
}
