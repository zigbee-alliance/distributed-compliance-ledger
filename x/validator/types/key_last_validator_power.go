package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

// LastValidatorPowerKey returns the store key to retrieve a LastValidatorPower from the index fields
func LastValidatorPowerKey(
	owner sdk.ValAddress,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
