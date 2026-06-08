package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RevokedCertificatesKeyPrefix is the prefix to retrieve all RevokedCertificates.
	RevokedCertificatesKeyPrefix = "RevokedCertificates/value/"
)

// RevokedCertificatesKey returns the store key to retrieve a RevokedCertificates from the index fields.
func RevokedCertificatesKey(
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
