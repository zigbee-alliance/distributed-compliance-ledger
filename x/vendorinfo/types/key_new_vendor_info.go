package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewVendorInfoKeyPrefix is the prefix to retrieve all NewVendorInfo
	NewVendorInfoKeyPrefix = "NewVendorInfo/value/"
)

// NewVendorInfoKey returns the store key to retrieve a NewVendorInfo from the index fields
func NewVendorInfoKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
