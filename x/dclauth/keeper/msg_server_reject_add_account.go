package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) RejectAddAccount(
	goCtx context.Context, msg *types.MsgRejectAddAccount,
) (*types.MsgRejectAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: (%s)", err)
	}

	// check if sendor has enough rights to reject a validator node
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgRejectAccount transaction should be signed by an account with the %s role",
			types.Trustee,
		)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if pending account exists
	if !k.IsPendingAccountPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountDoesNotExist(msg.Address)
	}

	// get pending account
	pendAcc, _ := k.GetPendingAccount(ctx, accAddr)

	// check if pending account already has reject approval from signer
	if pendAcc.HasRejectApprovalFrom(signerAddr) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"Pending account associated with the address=%v already has reject from=%v",
			msg.Address,
			msg.Signer,
		)
	}

	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if pending account already has approval from signer
	if pendAcc.HasApprovalFrom(signerAddr) {
		// Remove pending account if there are no rejects and other approvals
		if len(pendAcc.Approvals) == 1 && len(pendAcc.Rejects) == 0 {
			k.RemovePendingAccount(ctx, accAddr)

			return &types.MsgRejectAddAccountResponse{}, nil
		}

		// Remove approval from the list of approvals
		for i, other := range pendAcc.Approvals {
			if other.Address == grant.Address {
				pendAcc.Approvals = append(pendAcc.Approvals[:i], pendAcc.Approvals[i+1:]...)

				break
			}
		}
	}
	pendAcc.Rejects = append(pendAcc.Rejects, &grant)

	var percent float64

	if pendAcc.HasOnlyVendorRole(types.Vendor) {
		percent = types.VendorAccountApprovalsPercent
	} else {
		percent = types.AccountApprovalsPercent
	}

	// check if pending account has enough reject approvals
	if len(pendAcc.Rejects) >= k.AccountRejectApprovalsCount(ctx, percent) {
		account := types.NewAccount(pendAcc.BaseAccount, pendAcc.Roles, pendAcc.Approvals, pendAcc.Rejects, pendAcc.VendorID, pendAcc.ProductIDs)
		err = account.SetAccountNumber(k.GetNextAccountNumber(ctx))
		if err != nil {
			return nil, err
		}

		rejectedAccount := types.RejectedAccount{
			Account: account,
		}

		k.SetRejectedAccount(ctx, rejectedAccount)

		// delete pending account record
		k.RemovePendingAccount(ctx, accAddr)
	} else {
		// update pending account record
		k.SetPendingAccount(ctx, pendAcc)
	}

	return &types.MsgRejectAddAccountResponse{}, nil
}
