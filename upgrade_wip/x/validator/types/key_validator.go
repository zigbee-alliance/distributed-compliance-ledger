package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

// ValidatorKey returns the store key to retrieve a Validator from the index fields
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
