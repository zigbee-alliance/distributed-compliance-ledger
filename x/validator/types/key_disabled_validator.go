package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// DisabledValidatorKeyPrefix is the prefix to retrieve all DisabledValidator.
	DisabledValidatorKeyPrefix = "DisabledValidator/value/"
)

// DisabledValidatorKey returns the store key to retrieve a DisabledValidator from the index fields.
func DisabledValidatorKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
