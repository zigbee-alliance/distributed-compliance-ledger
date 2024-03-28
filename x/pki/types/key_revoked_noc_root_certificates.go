package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RevokedNocRootCertificatesKeyPrefix is the prefix to retrieve all RevokedNocRootCertificates.
	RevokedNocRootCertificatesKeyPrefix = "RevokedNocRootCertificates/value/"
)

// RevokedNocRootCertificatesKey returns the store key to retrieve a RevokedNocRootCertificates from the index fields.
func RevokedNocRootCertificatesKey(
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
