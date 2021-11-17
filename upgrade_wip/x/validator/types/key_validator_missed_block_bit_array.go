package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorMissedBlockBitArrayKeyPrefix is the prefix to retrieve all ValidatorMissedBlockBitArray
	ValidatorMissedBlockBitArrayKeyPrefix = "ValidatorMissedBlockBitArray/value/"
)

// ValidatorMissedBlockBitArrayKey returns the store key to retrieve a ValidatorMissedBlockBitArray from the index fields
func ValidatorMissedBlockBitArrayKey(
	address string,
	index uint64,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
