package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposedUpgradeKeyPrefix is the prefix to retrieve all ProposedUpgrade
	ProposedUpgradeKeyPrefix = "ProposedUpgrade/value/"
)

// ProposedUpgradeKey returns the store key to retrieve a ProposedUpgrade from the index fields
func ProposedUpgradeKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
