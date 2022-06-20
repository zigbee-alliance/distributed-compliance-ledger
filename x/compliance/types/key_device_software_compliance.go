package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceSoftwareComplianceKeyPrefix is the prefix to retrieve all DeviceSoftwareCompliance
	DeviceSoftwareComplianceKeyPrefix = "DeviceSoftwareCompliance/value/"
)

// DeviceSoftwareComplianceKey returns the store key to retrieve a DeviceSoftwareCompliance from the index fields
func DeviceSoftwareComplianceKey(
	cdCertificateId string,
) []byte {
	var key []byte

	cdCertificateIdBytes := []byte(cdCertificateId)
	key = append(key, cdCertificateIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
