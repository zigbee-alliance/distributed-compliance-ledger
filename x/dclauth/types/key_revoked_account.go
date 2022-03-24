package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// RevokedAccountKeyPrefix is the prefix to retrieve all RevokedAccount.
	RevokedAccountKeyPrefix = "RevokedAccount/value/"
)

// RevokedAccountKey returns the store key to retrieve a RevokedAccount from the index fields.
func RevokedAccountKey(
	address sdk.AccAddress,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
