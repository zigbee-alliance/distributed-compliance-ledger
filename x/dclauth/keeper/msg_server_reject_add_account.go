package keeper

import (
	"context"

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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: (%s)", err)
	}

	// check if sendor has enough rights to create a validator node
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveAddAccount transaction should be signed by an account with the %s role",
			types.Trustee,
		)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if pending account exists
	if !k.IsPendingAccountPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountDoesNotExist(msg.Address)
	}

	// get pending account
	pendAcc, _ := k.GetPendingAccount(ctx, accAddr)

	// check if pending account already has approval from signer
	if pendAcc.HasApprovalFrom(signerAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Pending account associated with the address=%v already has approval from=%v",
			msg.Address,
			msg.Signer,
		)
	}

	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}
	pendAcc.RejectApprovals = append(pendAcc.RejectApprovals, &grant)

	// check if pending account has enough reject approvals
	if len(pendAcc.RejectApprovals) == k.AccountRejectApprovalsCount(ctx) {
		account := types.NewAccount(pendAcc.BaseAccount, pendAcc.Roles, pendAcc.Approvals, pendAcc.VendorID)
		account.RejectApprovals = pendAcc.RejectApprovals

		err = account.SetAccountNumber(k.GetNextAccountNumber(ctx))
		if err != nil {
			return nil, err
		}

		rejectedAccount := types.RejectedAccount{
			Account: account,
		}

		k.SetRejectedAccount(ctx, rejectedAccount)

		// delete revoked account record if we have else don't delete
		k.RemoveRevokedAccount(ctx, accAddr)

		// delete pending account record
		k.RemovePendingAccount(ctx, accAddr)
	} else {
		// update pending account record
		k.SetPendingAccount(ctx, pendAcc)
	}

	return &types.MsgRejectAddAccountResponse{}, nil
}
