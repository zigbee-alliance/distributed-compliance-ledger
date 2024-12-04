package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AllCertificatesBySubjectKeyIDKeyPrefix is the prefix to retrieve all AllCertificatesBySubjectKeyId
	AllCertificatesBySubjectKeyIDKeyPrefix = "AllCertificatesBySubjectKeyID/value/"
)

// AllCertificatesBySubjectKeyIDKey returns the store key to retrieve a AllCertificatesBySubjectKeyId from the index fields
func AllCertificatesBySubjectKeyIDKey(
	subjectKeyID string,
) []byte {
	var key []byte

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
