package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ApproveAddAccount(
	goCtx context.Context,
	msg *types.MsgApproveAddAccount,
) (*types.MsgApproveAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check if sender has enough rights to create a validator node
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveAddAccount transaction should be signed by an account with the %s role", types.Trustee))
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check if pending account exists
	if !k.IsPendingAccountPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountDoesNotExist(msg.Address)
	}

	// get pending account
	pendAcc := k.GetPendingAccount(ctx, accAddr)

	// check if pending account already has approval from signer
	if pendAcc.HasApprovalFrom(signerAddr) {
		return nil, sdk.ErrUnauthorized(
			fmt.Sprintf("Pending account associated with the address=%v already has approval from=%v",
				msg.Address, msg.Signer))
	}

	// append approval
	pendAcc.Approvals = append(pendAcc.Approvals, signerAddr)

	// check if pending account has enough approvals
	if len(pendAcc.Approvals) == AccountApprovalsCount(ctx, k) {
		// create approved account, assign account number and store it
		// TODO issue 99: create a separate instance of BaseAccount with
		//		AccountNumber and Sequence set to zero
		account := types.NewAccount(pendAcc.BaseAccount, pendAcc.Roles, pendAcc.VendorID)
		account.AccountNumber = k.GetNextAccountNumber(ctx)
		k.SetAccount(ctx, account)

		// delete pending account record
		k.DeletePendingAccount(ctx, accAddr)
	} else {
		// update pending account record
		k.SetPendingAccount(ctx, pendAcc)
	}

	return &types.MsgApproveAddAccountResponse{}, nil
}
