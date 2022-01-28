package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) ProposeUpgrade(goCtx context.Context, msg *types.MsgProposeUpgrade) (*types.MsgProposeUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if creator has enough rights to create model
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to create a validator node
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, dclauthtypes.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeUpgrade transaction should be signed by an account with the %s role",
			dclauthtypes.Trustee,
		)
	}

	// check if proposed upgrade exists
	_, isFound := k.GetProposedUpgrade(ctx, msg.Plan.Name)
	if isFound {
		return nil, types.NewErrProposedUpgradeAlreadyExists(msg.Plan.Name)
	}

	// schedule upgrade
	err = k.upgradeKeeper.ScheduleUpgrade(ctx, msg.Plan)
	if err != nil {
		return nil, err
	}

	proposedUpgrade := types.ProposedUpgrade{
		Plan:      msg.Plan,
		Creator:   msg.Creator,
		Approvals: []string{msg.Creator},
	}

	// store proposed upgrade
	k.SetProposedUpgrade(ctx, proposedUpgrade)

	return &types.MsgProposeUpgradeResponse{}, nil
}
