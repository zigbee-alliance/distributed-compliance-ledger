package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

// PendingAccountKey returns the store key to retrieve a PendingAccount from the index fields
func PendingAccountKey(
	address sdk.AccAddress,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
