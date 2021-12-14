package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dclauth module sentinel errors
const (
	DefaultCodespace string = ModuleName

	CodeAccountAlreadyExists                  uint32 = 101
	CodeAccountDoesNotExist                   uint32 = 102
	CodePendingAccountAlreadyExists           uint32 = 103
	CodePendingAccountDoesNotExist            uint32 = 104
	CodePendingAccountRevocationAlreadyExists uint32 = 105
	CodePendingAccountRevocationDoesNotExist  uint32 = 106
	CodeMissingVendorIDForVendorAccount       uint32 = 107
	CodeMissingRoles                          uint32 = 108
)

func ErrAccountAlreadyExists(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeAccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodePendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodePendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodePendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}

func ErrMissingVendorIDForVendorAccount() *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeMissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrMissingRoles() *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeMissingRoles,
		"No roles provided")
}
