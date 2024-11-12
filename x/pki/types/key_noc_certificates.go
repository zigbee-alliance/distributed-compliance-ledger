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
	subjectKeyID string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
