package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedCertificatesBySubjectKeyIDKeyPrefix is the prefix to retrieve all ApprovedCertificatesBySubjectKeyId.
	ApprovedCertificatesBySubjectKeyIDKeyPrefix = "ApprovedCertificatesBySubjectKeyID/value/"
)

// ApprovedCertificatesBySubjectKeyIDKey returns the store key to retrieve a ApprovedCertificatesBySubjectKeyId from the index fields.
func ApprovedCertificatesBySubjectKeyIDKey(
	subjectKeyID string,
) []byte {
	var key []byte

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
