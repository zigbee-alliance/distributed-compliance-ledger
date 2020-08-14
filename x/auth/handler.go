package auth

import (
	"fmt"
	"math"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgProposeAddAccount:
			return handleMsgProposeAddAccount(ctx, keeper, msg)
		case types.MsgApproveAddAccount:
			return handleMsgApproveAddAccount(ctx, keeper, msg)
		case types.MsgProposeRevokeAccount:
			return handleMsgProposeRevokeAccount(ctx, keeper, msg)
		case types.MsgApproveRevokeAccount:
			return handleMsgApproveRevokeAccount(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized auth Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgProposeAddAccount(ctx sdk.Context, keeper Keeper, msg types.MsgProposeAddAccount) sdk.Result {
	// check if sender has enough rights to propose account.
	if !keeper.HasRole(ctx, msg.Signer, types.Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgProposeAddAccount transaction should be signed by an account with the %s role",
				types.Trustee)).Result()
	}

	// check if active account already exists.
	if keeper.IsAccountPresent(ctx, msg.Address) {
		return types.ErrAccountAlreadyExist(msg.Address).Result()
	}

	// check if pending account already exists.
	if keeper.IsPendingAccountPresent(ctx, msg.Address) {
		return types.ErrPendingAccountAlreadyExist(msg.Address).Result()
	}

	// parse the key.
	pubKey, err := sdk.GetAccPubKeyBech32(msg.PublicKey)
	if err != nil {
		return sdk.ErrInvalidPubKey(err.Error()).Result()
	}

	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if AddAccountApprovalsCount(ctx, keeper) > 1 {
		// create and store pending account.
		account := types.NewPendingAccount(msg.Address, pubKey, msg.Roles, msg.Signer)
		keeper.SetPendingAccount(ctx, account)
	} else {
		// create account, assign account number and store it
		account := types.NewAccount(msg.Address, pubKey, msg.Roles)
		account.AccountNumber = keeper.GetNextAccountNumber(ctx)
		keeper.SetAccount(ctx, account)
	}

	return sdk.Result{}
}

func handleMsgApproveAddAccount(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgApproveAddAccount) sdk.Result {
	// check if sender has enough rights to approve account
	if !keeper.HasRole(ctx, msg.Signer, Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveAddAccount transaction should be signed by an account with the %s role", Trustee)).Result()
	}

	// check if pending account exists
	if !keeper.IsPendingAccountPresent(ctx, msg.Address) {
		return types.ErrPendingAccountDoesNotExist(msg.Address).Result()
	}

	// get pending account
	pendAcc := keeper.GetPendingAccount(ctx, msg.Address)

	// check if pending account already has approval from signer
	if pendAcc.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Pending account associated with the address=%v already has approval from=%v",
				msg.Address, msg.Signer)).Result()
	}

	// append approval
	pendAcc.Approvals = append(pendAcc.Approvals, msg.Signer)

	// check if pending account has enough approvals
	if len(pendAcc.Approvals) == AddAccountApprovalsCount(ctx, keeper) {
		// create approved account, assign account number and store it
		account := types.NewAccount(pendAcc.Address, pendAcc.PubKey, pendAcc.Roles)
		account.AccountNumber = keeper.GetNextAccountNumber(ctx)
		keeper.SetAccount(ctx, account)

		// delete pending account record
		keeper.DeletePendingAccount(ctx, msg.Address)
	} else {
		// update pending account record
		keeper.SetPendingAccount(ctx, pendAcc)
	}

	return sdk.Result{}
}

func handleMsgProposeRevokeAccount(ctx sdk.Context, keeper keeper.Keeper,
	msg types.MsgProposeRevokeAccount) sdk.Result {
	// check that sender has enough rights to propose account revocation
	if !keeper.HasRole(ctx, msg.Signer, types.Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgProposeRevokeAccount transaction should be signed by an account with the %s role",
				types.Trustee)).Result()
	}

	// check that account exists
	if !keeper.IsAccountPresent(ctx, msg.Address) {
		return types.ErrAccountDoesNotExist(msg.Address).Result()
	}

	// check that pending account revocation does not exist yet
	if keeper.IsPendingAccountRevocationPresent(ctx, msg.Address) {
		return types.ErrPendingAccountRevocationAlreadyExist(msg.Address).Result()
	}

	// if more than 1 trustee's approval is needed, create pending account revocation else delete the account.
	if RevokeAccountApprovalsCount(ctx, keeper) > 1 {
		// create and store pending account revocation record
		revoc := types.NewPendingAccountRevocation(msg.Address, msg.Signer)
		keeper.SetPendingAccountRevocation(ctx, revoc)
	} else {
		// delete account record
		keeper.DeleteAccount(ctx, msg.Address)
	}

	return sdk.Result{}
}

func handleMsgApproveRevokeAccount(ctx sdk.Context, keeper keeper.Keeper,
	msg types.MsgApproveRevokeAccount) sdk.Result {
	// check that sender has enough rights to approve account revocation
	if !keeper.HasRole(ctx, msg.Signer, types.Trustee) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveRevokeAccount transaction should be signed by an account with the %s role",
				types.Trustee)).Result()
	}

	// check that pending account revocation exists
	if !keeper.IsPendingAccountRevocationPresent(ctx, msg.Address) {
		return types.ErrPendingAccountRevocationDoesNotExist(msg.Address).Result()
	}

	// get pending account revocation
	revoc := keeper.GetPendingAccountRevocation(ctx, msg.Address)

	// check if pending account revocation already has approval from signer
	if revoc.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Pending account revocation associated with the address=%v already has approval from=%v",
				msg.Address, msg.Signer)).Result()
	}

	// append approval
	revoc.Approvals = append(revoc.Approvals, msg.Signer)

	// check if pending account revocation has enough approvals
	if len(revoc.Approvals) == AddAccountApprovalsCount(ctx, keeper) {
		// delete account record
		keeper.DeleteAccount(ctx, msg.Address)

		// delete pending account revocation record
		keeper.DeletePendingAccountRevocation(ctx, msg.Address)
	} else {
		// update pending account revocation record
		keeper.SetPendingAccountRevocation(ctx, revoc)
	}

	return sdk.Result{}
}

func AddAccountApprovalsCount(ctx sdk.Context, keeper keeper.Keeper) int {
	return int(math.Round(types.DefaultApproveAddAccountPercent * float64(keeper.CountAccountsWithRole(ctx, Trustee))))
}

func RevokeAccountApprovalsCount(ctx sdk.Context, keeper keeper.Keeper) int {
	return int(math.Round(types.DefaultApproveRevokeAccountPercent * float64(keeper.CountAccountsWithRole(ctx, Trustee))))
}
