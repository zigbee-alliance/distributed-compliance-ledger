package types

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "modelinfo"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var (
	ModelInfoPrefix      = []byte{0x01} // prefix for each key to a model info
	VendorProductsPrefix = []byte{0x02} // prefix for each key to a vendor products
)

// Key builder for Model Info.
func GetModelInfoKey(vid uint16, pid uint16) []byte {
	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)

	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)

	return append(ModelInfoPrefix, append(v, p...)...)
}

// Key builder for Vendor Products.
func GetVendorProductsKey(vid uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, vid)

	return append(VendorProductsPrefix, b...)
}
