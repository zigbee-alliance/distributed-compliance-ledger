package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) EnableValidator(goCtx context.Context, msg *types.MsgEnableValidator) (*types.MsgEnableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, dclauthtypes.NodeAdmin) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Enable validator transaction should be signed by an account with the %s role",
			dclauthtypes.NodeAdmin,
		)
	}

	// check if disabled validator exists
	_, isFound := k.GetDisabledValidator(ctx, msg.Address)
	if isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	// store disabled validator
	k.RemoveDisabledValidator(ctx, msg.Address)

	return &types.MsgEnableValidatorResponse{}, nil
}
