package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RevokedNocIcaCertificatesKeyPrefix is the prefix to retrieve all RevokedNocIcaCertificates
	RevokedNocIcaCertificatesKeyPrefix = "RevokedNocIcaCertificates/value/"
)

// RevokedNocIcaCertificatesKey returns the store key to retrieve a RevokedNocIcaCertificates from the index fields
func RevokedNocIcaCertificatesKey(
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
