package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesKeyPrefix is the prefix to retrieve all NocCertificates
	NocCertificatesKeyPrefix = "NocCertificates/value/"
)

// NocCertificatesKey returns the store key to retrieve a NocCertificates from the index fields
func NocCertificatesKey(
	subject string,
	subjectKeyId string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	subjectKeyIdBytes := []byte(subjectKeyId)
	key = append(key, subjectKeyIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
