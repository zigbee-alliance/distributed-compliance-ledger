package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedCertificatesKeyPrefix is the prefix to retrieve all ApprovedCertificates
	ApprovedCertificatesKeyPrefix = "ApprovedCertificates/value/"
)

// ApprovedCertificatesKey returns the store key to retrieve a ApprovedCertificates from the index fields
func ApprovedCertificatesKey(
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
