package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorKeyPrefix is the prefix to retrieve all Validator
	ValidatorKeyPrefix = "Validator/value/"
)

// ValidatorKey returns the store key to retrieve a Validator from the index fields
func ValidatorKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
