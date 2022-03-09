package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) ProposeUpgrade(goCtx context.Context, msg *types.MsgProposeUpgrade) (*types.MsgProposeUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose upgrade
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.UpgradeApprovalRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeUpgrade transaction should be signed by an account with the %s role",
			types.UpgradeApprovalRole,
		)
	}

	// check if proposed upgrade with the same name already exists
	_, isFound := k.GetProposedUpgrade(ctx, msg.Plan.Name)
	if isFound {
		return nil, types.NewErrProposedUpgradeAlreadyExists(msg.Plan.Name)
	}

	// check if approved upgrade with the same name already exists
	_, isFound = k.GetApprovedUpgrade(ctx, msg.Plan.Name)
	if isFound {
		return nil, types.NewErrApprovedUpgradeAlreadyExists(msg.Plan.Name)
	}

	// Execute scheduling upgrade in a new context branch (with branched store)
	// to validate msg.Plan before the proposal proceeds through the approval process.
	// State is not persisted.
	cacheCtx, writeCache := ctx.CacheContext()
	err = k.upgradeKeeper.ScheduleUpgrade(cacheCtx, msg.Plan)
	if err != nil {
		return nil, err
	}

	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	if k.UpgradeApprovalsCount(ctx) > 1 {
		proposedUpgrade := types.ProposedUpgrade{
			Plan:      msg.Plan,
			Creator:   msg.Creator,
			Approvals: []*types.Grant{&grant},
		}

		// store proposed upgrade
		k.SetProposedUpgrade(ctx, proposedUpgrade)
	} else {
		// sufficient approvals count has been reached, so persist cached state (i.e. schedule upgrade)
		writeCache()

		approvedUpgrage := types.ApprovedUpgrade{
			Plan:      msg.Plan,
			Creator:   msg.Creator,
			Approvals: []*types.Grant{&grant},
		}

		k.SetApprovedUpgrade(ctx, approvedUpgrage)
	}

	return &types.MsgProposeUpgradeResponse{}, nil
}
