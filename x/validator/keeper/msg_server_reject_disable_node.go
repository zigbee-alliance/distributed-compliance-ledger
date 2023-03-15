package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) RejectDisableValidator(goCtx context.Context, msg *types.MsgRejectDisableValidator) (*types.MsgRejectDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address: (%s)", err)
	}

	validatorAddr, err := sdk.ValAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid validator address: (%s)", err)
	}

	// check if message creator has enough rights to reject disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.VoteForDisableValidatorRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgRejectDisableValidator transaction should be signed by an account with the %s role",
			types.VoteForDisableValidatorRole,
		)
	}

	// check if proposed disable validator exists
	proposedDisableValidator, isFound := k.GetProposedDisableValidator(ctx, validatorAddr.String())
	if !isFound {
		return nil, types.NewErrProposedDisableValidatorDoesNotExist(msg.Address)
	}

	// check if disable validator already has reject from message creator
	if proposedDisableValidator.HasRejectDisableFrom(creatorAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disabled validator with address=%v already has reject from=%v",
			msg.Address,
			msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if disable validator has approval from message creator
	if proposedDisableValidator.HasApprovalFrom(creatorAddr) {
		// Remove propose disable validator if there are no rejects and other approvals
		if len(proposedDisableValidator.Approvals) == 1 && len(proposedDisableValidator.Rejects) == 0 {
			k.RemoveProposedDisableValidator(ctx, proposedDisableValidator.Address)

			return &types.MsgRejectDisableValidatorResponse{}, nil
		}

		// Remove approval from the list of approvals
		for i, other := range proposedDisableValidator.Approvals {
			if other.Address == grant.Address {
				proposedDisableValidator.Approvals = append(proposedDisableValidator.Approvals[:i], proposedDisableValidator.Approvals[i+1:]...)

				break
			}
		}
	}
	proposedDisableValidator.Rejects = append(proposedDisableValidator.Rejects, &grant)

	// check if proposed disable validator has enough reject approvals
	if len(proposedDisableValidator.Rejects) >= k.DisableValidatorRejectApprovalsCount(ctx) {
		k.RemoveProposedDisableValidator(ctx, proposedDisableValidator.Address)
		rejectedDisableValidator := types.RejectedDisableValidator(proposedDisableValidator)
		k.SetRejectedNode(ctx, rejectedDisableValidator)
	} else {
		// update proposed disable validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	}

	return &types.MsgRejectDisableValidatorResponse{}, nil
}
