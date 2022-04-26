package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RejectedCertificateKeyPrefix is the prefix to retrieve all RejectedCertificate
	RejectedCertificateKeyPrefix = "RejectedCertificate/value/"
)

// RejectedCertificateKey returns the store key to retrieve a RejectedCertificate from the index fields
func RejectedCertificateKey(
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
