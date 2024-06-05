package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) ApproveDisableValidator(goCtx context.Context, msg *types.MsgApproveDisableValidator) (*types.MsgApproveDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address: (%s)", err)
	}

	validatorAddr, err := sdk.ValAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid validator address: (%s)", err)
	}

	// check if message creator has enough rights to approve disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.VoteForDisableValidatorRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveDisableValidator transaction should be signed by an account with the %s role",
			types.VoteForDisableValidatorRole,
		)
	}

	// check if proposed disable validator exists
	proposedDisableValidator, isFound := k.GetProposedDisableValidator(ctx, validatorAddr.String())
	if !isFound {
		return nil, types.NewErrProposedDisableValidatorDoesNotExist(msg.Address)
	}

	// check if disable validator already has approval form message creator
	if proposedDisableValidator.HasApprovalFrom(creatorAddr) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disabled validator with address=%v already has approval from=%v",
			msg.Address, msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if disable validator already has reject from message creator
	if proposedDisableValidator.HasRejectDisableFrom(creatorAddr) {
		for i, other := range proposedDisableValidator.Rejects {
			if other.Address == grant.Address {
				proposedDisableValidator.Rejects = append(proposedDisableValidator.Rejects[:i], proposedDisableValidator.Rejects[i+1:]...)

				break
			}
		}
	}
	proposedDisableValidator.Approvals = append(proposedDisableValidator.Approvals, &grant)

	// check if proposed disable validator has enough approvals
	if len(proposedDisableValidator.Approvals) >= k.DisableValidatorApprovalsCount(ctx) {
		// remove disable validator
		k.RemoveProposedDisableValidator(ctx, proposedDisableValidator.Address)

		approvedDisableValidator := types.DisabledValidator{
			Address:             proposedDisableValidator.Address,
			Creator:             proposedDisableValidator.Creator,
			Approvals:           proposedDisableValidator.Approvals,
			Rejects:             proposedDisableValidator.Rejects,
			DisabledByNodeAdmin: false,
		}

		// Disable validator
		validator, _ := k.GetValidator(ctx, validatorAddr)
		k.Jail(ctx, validator, proposedDisableValidator.Approvals[0].Info)

		k.SetDisabledValidator(ctx, approvedDisableValidator)
	} else {
		// update proposed disable validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	}

	return &types.MsgApproveDisableValidatorResponse{}, nil
}
