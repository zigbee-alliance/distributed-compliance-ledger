package types

import "fmt"

const (
	// ModuleName is the name of the module
	ModuleName = "compliancetest"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	TestingResultsPrefix = []byte{0x1} // prefix for each key to a testing results
)

// Key builder for Testing Results
func GetTestingResultsKey(vid uint16, pid uint16) []byte {
	return append(TestingResultsPrefix, []byte(fmt.Sprintf("%v:%v", vid, pid))...)
}
