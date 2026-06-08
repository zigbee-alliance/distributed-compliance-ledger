package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RejectedUpgradeKeyPrefix is the prefix to retrieve all RejectedUpgrade.
	RejectedUpgradeKeyPrefix = "RejectedUpgrade/value/"
)

// RejectedUpgradeKey returns the store key to retrieve a RejectedUpgrade from the index fields.
func RejectedUpgradeKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
