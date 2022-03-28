package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedCertificatesKeyPrefix is the prefix to retrieve all ApprovedCertificates.
	ApprovedCertificatesKeyPrefix = "ApprovedCertificates/value/"
)

// ApprovedCertificatesKey returns the store key to retrieve a ApprovedCertificates from the index fields.
func ApprovedCertificatesKey(
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
