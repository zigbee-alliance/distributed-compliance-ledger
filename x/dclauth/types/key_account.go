package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountKeyPrefix is the prefix to retrieve all Account
	AccountKeyPrefix = "Account/value/"
)

// AccountKey returns the store key to retrieve a Account from the index fields
func AccountKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
