package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ApproveRevokeAccount(goCtx context.Context, msg *types.MsgApproveRevokeAccount) (*types.MsgApproveRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check that sender has enough rights to approve account revocation
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveRevokeAccount transaction should be signed by an account with the %s role",
				types.Trustee))
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check that pending account revocation exists
	if !k.IsPendingAccountRevocationPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountRevocationDoesNotExist(msg.Address)
	}

	// get pending account revocation
	revoc := k.GetPendingAccountRevocation(ctx, accAddr)

	// check if pending account revocation already has approval from signer
	if revoc.HasApprovalFrom(signerAddr) {
		return nil, sdk.ErrUnauthorized(
			fmt.Sprintf("Pending account revocation associated with the address=%v already has approval from=%v",
				msg.Address, msg.Signer))
	}

	// append approval
	revoc.Approvals = append(revoc.Approvals, signerAddr)

	// check if pending account revocation has enough approvals
	if len(revoc.Approvals) == AccountApprovalsCount(ctx, k) {
		// delete account record
		k.DeleteAccount(ctx, accAddr)

		// delete pending account revocation record
		k.DeletePendingAccountRevocation(ctx, accAddr)
	} else {
		// update pending account revocation record
		k.SetPendingAccountRevocation(ctx, revoc)
	}

	return &types.MsgApproveRevokeAccountResponse{}, nil
}
