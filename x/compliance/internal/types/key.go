package types

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "compliance"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var ComplianceInfoPrefix = []byte{0x01} // prefix for each key to a compliance info

// Key builder for Compliance Info.
func GetComplianceInfoKey(certificationType CertificationType, vid uint16, pid uint16) []byte {
	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)

	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)

	return append(ComplianceInfoPrefix, append([]byte(certificationType), append(v, p...)...)...)
}

// Key builder for Compliance Info.
func GetCertificationPrefix(certificationType CertificationType) []byte {
	return append(ComplianceInfoPrefix, []byte(certificationType)...)
}
