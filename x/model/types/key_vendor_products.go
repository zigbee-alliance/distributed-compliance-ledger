package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VendorProductsKeyPrefix is the prefix to retrieve all VendorProducts
	VendorProductsKeyPrefix = "VendorProducts/value/"
)

// VendorProductsKey returns the store key to retrieve a VendorProducts from the index fields
func VendorProductsKey(
	vid int32,
) []byte {
	var key []byte

	vidBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	return key
}
