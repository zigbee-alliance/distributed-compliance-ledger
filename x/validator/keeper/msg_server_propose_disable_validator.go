package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) ProposeDisableValidator(goCtx context.Context, msg *types.MsgProposeDisableValidator) (*types.MsgProposeDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address: (%s)", err)
	}

	validatorAddr, err := sdk.ValAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid validator address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.VoteForDisableValidatorRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeDisableValidator transaction should be signed by an account with the %s role",
			types.VoteForDisableValidatorRole,
		)
	}

	// check if validator exists
	isFound := k.Keeper.IsValidatorPresent(ctx, validatorAddr)
	if !isFound {
		return nil, sdkstakingtypes.ErrNoValidatorFound
	}

	// check if disabled validator exists
	_, isFound = k.GetDisabledValidator(ctx, validatorAddr.String())
	if isFound {
		return nil, types.NewErrDisabledValidatorAlreadyExists(msg.Address)
	}

	// check if proposed disable validator exists
	_, isFound = k.GetProposedDisableValidator(ctx, validatorAddr.String())
	if isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	if k.DisableValidatorApprovalsCount(ctx) > 1 {
		proposedDisableValidator := types.ProposedDisableValidator{
			Address: msg.Address,
			Creator: msg.Creator,
			Approvals: []*types.Grant{
				{
					Address: msg.Creator,
					Time:    msg.Time,
					Info:    msg.Info,
				},
			},
		}

		// store proposed disable validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	} else {
		disabledValidator := types.DisabledValidator{
			Address: msg.Address,
			Creator: msg.Creator,
			Approvals: []*types.Grant{
				{
					Address: msg.Creator,
					Time:    msg.Time,
					Info:    msg.Info,
				},
			},
			DisabledByNodeAdmin: false,
		}

		// Disable validator
		validator, _ := k.GetValidator(ctx, validatorAddr)
		k.Jail(ctx, validator, msg.Info)

		// store disabled validator
		k.SetDisabledValidator(ctx, disabledValidator)
	}

	return &types.MsgProposeDisableValidatorResponse{}, nil
}
