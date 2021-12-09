package keeper

import (
	"context"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ProposeAddAccount(goCtx context.Context, msg *types.MsgProposeAddAccount) (*types.MsgProposeAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check if sender has enough rights to create a validator node
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeAddAccount transaction should be signed by an account with the %s role",
			types.Trustee,
		)
	}

	// check if proposed account has vendor role, vendor id should be present.
	if msg.HasRole(types.Vendor) && msg.VendorID <= 0 {
		return nil, types.ErrMissingVendorIDForVendorAccount()
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	// TODO issue 99: good error
	if err != nil {
		return nil, err
	}

	// check if active account already exists.
	if k.IsAccountPresent(ctx, accAddr) {
		return nil, types.ErrAccountAlreadyExists(msg.Address)
	}

	// check if pending account already exists.
	if k.IsPendingAccountPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountAlreadyExists(msg.Address)
	}

	// parse the key.
	pk, ok := msg.PubKey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pk)
	}

	ba := authtypes.NewBaseAccount(accAddr, pk, 0, 0)
	account := types.NewAccount(ba, msg.Roles, msg.VendorID)

	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account.
		account := types.NewPendingAccount(account, signerAddr)
		k.SetPendingAccount(ctx, *account)
	} else {
		// create account, assign account number and store it
		account.AccountNumber = k.GetNextAccountNumber(ctx)
		k.SetAccountO(ctx, *account)
	}

	return &types.MsgProposeAddAccountResponse{}, nil
}
