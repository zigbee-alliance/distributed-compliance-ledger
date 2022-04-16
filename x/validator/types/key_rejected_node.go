package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RejectedNodeKeyPrefix is the prefix to retrieve all RejectedNode.
	RejectedNodeKeyPrefix = "RejectedNode/value/"
)

// RejectedNodeKey returns the store key to retrieve a RejectedNode from the index fields.
func RejectedNodeKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
