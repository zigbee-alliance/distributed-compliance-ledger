package types

import "fmt"

const (
	// ModuleName is the name of the module
	ModuleName = "modelinfo"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	ModelInfoPrefix      = []byte{0x01} // prefix for each key to a model info
	VendorProductsPrefix = []byte{0x02} // prefix for each key to a vendor products
)

// Key builder for Model Info
func GetModelInfoKey(vid uint16, pid uint16) []byte {
	return append(ModelInfoPrefix, []byte(fmt.Sprintf("%v:%v", vid, pid))...)
}

// Key builder for Vendor Products
func GetVendorProductsKey(vid uint16) []byte {
	return append(VendorProductsPrefix, []byte(fmt.Sprintf("%v", vid))...)
}
