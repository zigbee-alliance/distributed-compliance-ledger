package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) ApproveUpgrade(goCtx context.Context, msg *types.MsgApproveUpgrade) (*types.MsgApproveUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to approve upgrade
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

	// check if proposed upgrade already has approval form message creator
	if proposedUpgrade.HasApprovalFrom(creatorAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Proposed upgrade with name=%v already has approval from=%v",
			msg.Name, msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	proposedUpgrade.Approvals = append(proposedUpgrade.Approvals, &grant)

	// check if proposed upgrade has enough approvals
	if len(proposedUpgrade.Approvals) == k.UpgradeApprovalsCount(ctx) {
		// remove proposed upgrade
		k.RemoveProposedUpgrade(ctx, proposedUpgrade.Plan.Name)

		// schedule upgrade
		err = k.upgradeKeeper.ScheduleUpgrade(ctx, proposedUpgrade.Plan)
		if err != nil {
			return nil, err
		}

		approvedUpgrage := types.ApprovedUpgrade(proposedUpgrade)
		k.SetApprovedUpgrade(ctx, approvedUpgrage)
	} else {
		// update proposed upgrade
		k.SetProposedUpgrade(ctx, proposedUpgrade)
	}

	return &types.MsgApproveUpgradeResponse{}, nil
}
