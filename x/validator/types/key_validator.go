package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// ValidatorKeyPrefix is the prefix to retrieve all Validator.
	ValidatorKeyPrefix = "Validator/value/"
	// ValidatorByConsAddrKeyPrefix is the prefix to retrieve all Validator by consensus address.
	ValidatorByConsAddrKeyPrefix = "ValidatorByConsAddr/value/"
)

// ValidatorKey returns the store key to retrieve a Validator from the index fields.
func ValidatorKey(
	owner sdk.ValAddress,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ValidatorByConsAddrKey(
	addr sdk.ConsAddress,
) []byte {
	var key []byte

	ownerBytes := []byte(addr)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
