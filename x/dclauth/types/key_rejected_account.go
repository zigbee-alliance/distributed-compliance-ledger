package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RejectedAccountKeyPrefix is the prefix to retrieve all RejectedAccount
	RejectedAccountKeyPrefix = "RejectedAccount/value/"
)

// RejectedAccountKey returns the store key to retrieve a RejectedAccount from the index fields
func RejectedAccountKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
