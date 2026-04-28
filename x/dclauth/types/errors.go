package types

// DONTCOVER

import (
	"fmt"

	"cosmossdk.io/errors"
)

// x/dclauth module sentinel errors.
var (
	AccountAlreadyExists                  = errors.Register(ModuleName, 101, "account already exists")
	AccountDoesNotExist                   = errors.Register(ModuleName, 102, "account not found")
	PendingAccountAlreadyExists           = errors.Register(ModuleName, 103, "pending account already exists")
	PendingAccountDoesNotExist            = errors.Register(ModuleName, 104, "pending account not found")
	PendingAccountRevocationAlreadyExists = errors.Register(ModuleName, 105, "pending account revocation already exists")
	PendingAccountRevocationDoesNotExist  = errors.Register(ModuleName, 106, "pending account revocation not found")
	MissingVendorIDForVendorAccount       = errors.Register(ModuleName, 107, "no Vendor ID provided")
	MissingRoles                          = errors.Register(ModuleName, 108, "no roles provided")
)

func ErrAccountAlreadyExists(address interface{}) error {
	return errors.Wrapf(AccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) error {
	return errors.Wrapf(AccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) error {
	return errors.Wrapf(PendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) error {
	return errors.Wrapf(PendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) error {
	return errors.Wrapf(PendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) error {
	return errors.Wrapf(PendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}

func ErrMissingVendorIDForVendorAccount() error {
	return errors.Wrapf(MissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrMissingRoles() error {
	return errors.Wrapf(MissingRoles,
		"No roles provided")
}
