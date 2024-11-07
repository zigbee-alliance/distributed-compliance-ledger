package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AllCertificatesKeyPrefix is the prefix to retrieve all Certificates
	AllCertificatesKeyPrefix = "AllCertificates/value/"
)

// AllCertificatesKey returns the store key to retrieve a Certificates from the index fields
func AllCertificatesKey(
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
