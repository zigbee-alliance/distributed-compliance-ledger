package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // VendorInfoTypeKeyPrefix is the prefix to retrieve all VendorInfoType
	VendorInfoTypeKeyPrefix = "VendorInfoType/value/"
)

// VendorInfoTypeKey returns the store key to retrieve a VendorInfoType from the index fields
func VendorInfoTypeKey(
index string,
) []byte {
	var key []byte
    
    indexBytes := []byte(index)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}