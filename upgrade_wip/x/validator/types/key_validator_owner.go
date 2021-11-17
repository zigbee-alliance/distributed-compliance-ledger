package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorOwnerKeyPrefix is the prefix to retrieve all ValidatorOwner
	// ValidatorOwnerKeyPrefix = "ValidatorOwner/value/"
	ValidatorOwnerKeyPrefix = []byte{0x05} // prefix for validator owner
)

// ValidatorOwnerKey returns the store key to retrieve a ValidatorOwner from the index fields
func ValidatorOwnerKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
