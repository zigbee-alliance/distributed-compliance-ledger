package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorKeyPrefix is the prefix to retrieve all Validator
	//ValidatorKeyPrefix = "Validator/value/"
	ValidatorKeyPrefix = []byte{0x01} // prefix for each key to a validator
)

// ValidatorKey returns the store key to retrieve a Validator from the index fields
func ValidatorKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
