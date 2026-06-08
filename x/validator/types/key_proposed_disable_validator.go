package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposedDisableValidatorKeyPrefix is the prefix to retrieve all ProposedDisableValidator.
	ProposedDisableValidatorKeyPrefix = "ProposedDisableValidator/value/"
)

// ProposedDisableValidatorKey returns the store key to retrieve a ProposedDisableValidator from the index fields.
func ProposedDisableValidatorKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
