package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VendorInfoTypeKeyPrefix is the prefix to retrieve all VendorInfoType
	VendorInfoTypeKeyPrefix = "VendorInfoType/value/"
)

// VendorInfoTypeKey returns the store key to retrieve a VendorInfoType from the index fields
func VendorInfoTypeKey(
	vendorID uint64,
) []byte {
	var key []byte

	vendorIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(vendorIDBytes, vendorID)
	key = append(key, vendorIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
