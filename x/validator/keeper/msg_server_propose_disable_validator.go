package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) ProposeDisableValidator(goCtx context.Context, msg *types.MsgProposeDisableValidator) (*types.MsgProposeDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, dclauthtypes.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeDisableValidator transaction should be signed by an account with the %s role",
			types.DisableValidatorRole,
		)
	}

	// check if proposed disable validator exists
	_, isFound := k.GetDisabledValidator(ctx, msg.Address)
	if isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	if k.DisableValidatorApprovalsCount(ctx) > 1 {
		proposedDisableValidator := types.ProposedDisableValidator{
			Address:   msg.Address,
			Approvals: []string{msg.Creator},
		}

		// store disabled validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	} else {
		disabledValidator := types.DisabledValidator{
			Address:             msg.Address,
			Approvals:           []string{msg.Creator},
			DisabledByNodeAdmin: false,
		}

		// store disabled validator
		k.SetDisabledValidator(ctx, disabledValidator)
	}

	return &types.MsgProposeDisableValidatorResponse{}, nil
}
