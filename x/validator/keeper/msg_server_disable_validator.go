package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) DisableValidator(goCtx context.Context, msg *types.MsgDisableValidator) (*types.MsgDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, dclauthtypes.NodeAdmin) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disable validator transaction should be signed by an account with the %s role",
			dclauthtypes.NodeAdmin,
		)
	}

	// check if disabled validator exists
	_, isFound := k.GetDisabledValidator(ctx, msg.Address)
	if isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	disabledValidator := types.DisabledValidator{
		Address:             msg.Address,
		Approvals:           []string{msg.Creator},
		DisabledByNodeAdmin: true,
	}

	// store disabled validator
	k.SetDisabledValidator(ctx, disabledValidator)

	return &types.MsgDisableValidatorResponse{}, nil
}
