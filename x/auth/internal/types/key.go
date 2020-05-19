package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of the module.
	ModuleName = "auth"

	// StoreKey to be used when creating the KVStore.
	StoreKey = "acc" // it differs from ModuleName to be compatible with cosmos transaction builder and processor.
)

var (
	PendingAccountPrefix = []byte{0x01} //  prefix for each key to a pending account
	AccountPrefix = []byte{0x02} //  prefix for each key to an account

	AccountNumberCounterKey = []byte("globalAccountNumber") // key for account number counter
)

// Key builder for Pending Account.
func GetPendingAccountKey(addr sdk.AccAddress) []byte {
	return append(PendingAccountPrefix, addr.Bytes()...)
}

// Key builder for Account.
func GetAccountKey(addr sdk.AccAddress) []byte {
	return append(AccountPrefix, addr.Bytes()...)
}
