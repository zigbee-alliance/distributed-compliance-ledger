package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dclauth module sentinel errors
var (
	CodeAccountAlreadyExists                  = sdkerrors.Register(ModuleName, 101, "account already exists")
	CodeAccountDoesNotExist                   = sdkerrors.Register(ModuleName, 102, "account not found")
	CodePendingAccountAlreadyExists           = sdkerrors.Register(ModuleName, 103, "pending account already exists")
	CodePendingAccountDoesNotExist            = sdkerrors.Register(ModuleName, 104, "pending account not found")
	CodePendingAccountRevocationAlreadyExists = sdkerrors.Register(ModuleName, 105, "pending account revocation already exists")
	CodePendingAccountRevocationDoesNotExist  = sdkerrors.Register(ModuleName, 106, "pending account revocation not found")
	CodeMissingVendorIDForVendorAccount       = sdkerrors.Register(ModuleName, 107, "no Vendor ID provided")
	CodeMissingRoles                          = sdkerrors.Register(ModuleName, 108, "no roles provided")
)

func ErrAccountAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(CodeAccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(CodePendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(CodePendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(CodePendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}

func ErrMissingVendorIDForVendorAccount() error {
	return sdkerrors.Wrapf(CodeMissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrMissingRoles() error {
	return sdkerrors.Wrapf(CodeMissingRoles,
		"No roles provided")
}
