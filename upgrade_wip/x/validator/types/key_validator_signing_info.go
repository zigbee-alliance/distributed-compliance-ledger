package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorSigningInfoKeyPrefix is the prefix to retrieve all ValidatorSigningInfo
	ValidatorSigningInfoKeyPrefix = "ValidatorSigningInfo/value/"
)

// ValidatorSigningInfoKey returns the store key to retrieve a ValidatorSigningInfo from the index fields
func ValidatorSigningInfoKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
