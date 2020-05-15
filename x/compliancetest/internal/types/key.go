package types

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the module
	ModuleName = "compliancetest"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	TestingResultsPrefix = []byte{0x01} // prefix for each key to a testing results
)

// Key builder for Testing Results
func GetTestingResultsKey(vid uint16, pid uint16) []byte {
	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)
	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)
	return append(TestingResultsPrefix, append(v, p...)...)
}
