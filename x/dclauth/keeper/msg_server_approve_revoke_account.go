package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ApproveRevokeAccount(goCtx context.Context, msg *types.MsgApproveRevokeAccount) (*types.MsgApproveRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: (%s)", err)
	}

	// check that sender has enough rights to approve account revocation
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveRevokeAccount transaction should be signed by an account with the %s role",
			types.Trustee,
		)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check that pending account revocation exists
	if !k.IsPendingAccountRevocationPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountRevocationDoesNotExist(msg.Address)
	}

	// get pending account revocation
	revoc, _ := k.GetPendingAccountRevocation(ctx, accAddr)

	// check if pending account revocation already has approval from signer
	if revoc.HasRevocationFrom(signerAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Pending account revocation associated with the address=%v already has approval from=%v",
			msg.Address,
			msg.Signer,
		)
	}
	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}
	// append approval
	revoc.Approvals = append(revoc.Approvals, &grant)

	// check if pending account revocation has enough approvals
	if len(revoc.Approvals) == k.AccountApprovalsCount(ctx) {
		// delete account record
		k.RemoveAccount(ctx, accAddr)

		// delete pending account revocation record
		k.RemovePendingAccountRevocation(ctx, accAddr)
	} else {
		// update pending account revocation record
		k.SetPendingAccountRevocation(ctx, revoc)
	}

	return &types.MsgApproveRevokeAccountResponse{}, nil
}
