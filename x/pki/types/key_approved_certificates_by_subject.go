package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedCertificatesBySubjectKeyPrefix is the prefix to retrieve all ApprovedCertificatesBySubject.
	ApprovedCertificatesBySubjectKeyPrefix = "ApprovedCertificatesBySubject/value/"
)

// ApprovedCertificatesBySubjectKey returns the store key to retrieve a ApprovedCertificatesBySubject from the index fields.
func ApprovedCertificatesBySubjectKey(
	subject string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	return key
}
