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
			fmt.Sprintf("MsgProposeAddAccount transaction should be signed by an account with the %s role", types.Trustee)).Result()
	}

	// check if account already exists.
	if keeper.IsAccountPresent(ctx, msg.Address) {
		return types.ErrAccountAlreadyExistExist(msg.Address).Result()
	}

	// parse the key.
	pubKey, err := sdk.GetAccPubKeyBech32(msg.PublicKey)
	if err != nil {
		return sdk.ErrInvalidPubKey(err.Error()).Result()
	}

	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if CountAccountApprovals(ctx, keeper) > 1 {
		// create and store pending account.
		account := types.NewPendingAccount(msg.Address, pubKey, msg.Roles, msg.Signer)
		keeper.SetProposedAccount(ctx, account)
	} else {
		// create account, assign account number and store it
		account := types.NewAccount(msg.Address, pubKey, msg.Roles)
		keeper.SetAccount(ctx, keeper.NewAccountWithNumber(ctx, account))
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
	if !keeper.IsProposedAccountPresent(ctx, msg.Address) {
		return types.ErrAccountDoesNotExist(msg.Address).Result()
	}

	// get account
	account := keeper.GetProposedAccount(ctx, msg.Address)

	// check if account already has approval from signer
	if account.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Account associated with the address=%v already has approval from=%v",
				msg.Address, msg.Signer)).Result()
	}

	// append approval
	account.Approvals = append(account.Approvals, msg.Signer)

	// check if account has enough approvals
	if len(account.Approvals) == CountAccountApprovals(ctx, keeper) {
		// create approved account, assign account number and store it
		account := types.NewAccount(account.Address, account.PubKey, account.Roles)
		keeper.SetAccount(ctx, keeper.NewAccountWithNumber(ctx, account))

		// delete pending account record
		keeper.DeleteProposedAccount(ctx, msg.Address)
	} else {
		// update pending account record
		keeper.SetProposedAccount(ctx, account)
	}

	return sdk.Result{}
}

func CountAccountApprovals(ctx sdk.Context, keeper keeper.Keeper) int {
	return int(math.Round(types.DefaultApproveAddAccountPercent * float64(keeper.CountAccountsWithRole(ctx, Trustee))))
}
