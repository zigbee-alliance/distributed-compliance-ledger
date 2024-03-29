package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// PendingAccountKeyPrefix is the prefix to retrieve all PendingAccount.
	PendingAccountKeyPrefix = "PendingAccount/value/"
)

// PendingAccountKey returns the store key to retrieve a PendingAccount from the index fields.
func PendingAccountKey(
	address sdk.AccAddress,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
