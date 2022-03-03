package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) ApproveDisableValidator(goCtx context.Context, msg *types.MsgApproveDisableValidator) (*types.MsgApproveDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to approve disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, dclauthtypes.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveDisableValidator transaction should be signed by an account with the %s role",
			dclauthtypes.Trustee,
		)
	}

	// check if proposed disable validator exists
	proposedDisableValidator, isFound := k.GetProposedDisableValidator(ctx, msg.Address)
	if !isFound {
		return nil, types.NewErrProposedDisableValidatorAlreadyExists(msg.Address)
	}

	// check if disable validator already has approval form message creator
	proposedDisableValidator.HasApprovalFrom()
	if proposedDisableValidator.HasApprovalFrom(creatorAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disabled validator with address=%v already has approval from=%v",
			msg.Address, msg.Creator,
		)
	}

	// append approval
	proposedDisableValidator.Approvals = append(proposedDisableValidator.Approvals, creatorAddr.String())

	// check if proposed disable validator has enough approvals
	if len(proposedDisableValidator.Approvals) == k.DisableValidatorApprovalsCount(ctx) {
		// remove disable validator
		k.RemoveDisabledValidator(ctx, proposedDisableValidator.Address)

		approvedUpgrage := types.DisabledValidator{
			Address:             proposedDisableValidator.Address,
			Approvals:           proposedDisableValidator.Approvals,
			DisabledByNodeAdmin: false,
		}

		k.SetDisabledValidator(ctx, approvedUpgrage)
	} else {
		// update proposed disable validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	}

	return &types.MsgApproveDisableValidatorResponse{}, nil
}
