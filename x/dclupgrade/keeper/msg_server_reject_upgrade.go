package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) RejectUpgrade(goCtx context.Context, msg *types.MsgRejectUpgrade) (*types.MsgRejectUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to reject upgrade
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.UpgradeApprovalRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveUpgrade transaction should be signed by an account with the %s role",
			types.UpgradeApprovalRole,
		)
	}

	// check if proposed upgrade exists
	proposedUpgrade, isFound := k.GetProposedUpgrade(ctx, msg.Name)
	if !isFound {
		return nil, types.NewErrProposedUpgradeDoesNotExist(msg.Name)
	}

	// check if proposed upgrade already has reject from message creator
	if proposedUpgrade.HasRejectFrom(creatorAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Proposed upgrade with name=%v already has reject from=%v",
			msg.Name, msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if proposed upgrade already has approval from message creator
	if proposedUpgrade.HasApprovalFrom(creatorAddr) {
		for i, other := range proposedUpgrade.Approvals {
			if other.Address == grant.Address {
				proposedUpgrade.Approvals = append(proposedUpgrade.Approvals[:i], proposedUpgrade.Approvals[i+1:]...)
			}
		}
	}
	proposedUpgrade.Rejects = append(proposedUpgrade.Rejects, &grant)

	// check if proposed upgrade has enough rejects
	if len(proposedUpgrade.Rejects) >= k.UpgradeRejectsCount(ctx) {
		// schedule upgrade
		err = k.upgradeKeeper.ScheduleUpgrade(ctx, proposedUpgrade.Plan)
		if err != nil {
			return nil, err
		}

		// remove proposed upgrade
		k.RemoveProposedUpgrade(ctx, proposedUpgrade.Plan.Name)

		rejectedUpgrade := types.RejectedUpgrade(proposedUpgrade)
		k.SetRejectedUpgrade(ctx, rejectedUpgrade)
	} else {
		// update proposed upgrade
		k.SetProposedUpgrade(ctx, proposedUpgrade)
	}

	return &types.MsgRejectUpgradeResponse{}, nil
}
