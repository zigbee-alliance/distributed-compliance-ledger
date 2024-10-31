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
