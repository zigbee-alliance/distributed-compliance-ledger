package types

import "fmt"

const (
	// ModuleName is the name of the module
	ModuleName = "compliance"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	ComplianceInfoPrefix = []byte{0x01} // prefix for each key to a compliance info
)

// Key builder for Compliance Info
func GetComplianceInfoKey(certificationType CertificationType, vid uint16, pid uint16) []byte {
	return append(ComplianceInfoPrefix, []byte(fmt.Sprintf("%v:%v:%v", certificationType, vid, pid))...)
}

// Key builder for Compliance Info
func GetCertificationPrefix(certificationType CertificationType) []byte {
	return append(ComplianceInfoPrefix, []byte(certificationType)...)
}
