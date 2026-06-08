package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedUpgradeKeyPrefix is the prefix to retrieve all ApprovedUpgrade.
	ApprovedUpgradeKeyPrefix = "ApprovedUpgrade/value/"
)

// ApprovedUpgradeKey returns the store key to retrieve a ApprovedUpgrade from the index fields.
func ApprovedUpgradeKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
