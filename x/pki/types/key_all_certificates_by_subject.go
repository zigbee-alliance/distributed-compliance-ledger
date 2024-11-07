package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AllCertificatesBySubjectKeyPrefix is the prefix to retrieve all AllCertificatesBySubject
	AllCertificatesBySubjectKeyPrefix = "AllCertificatesBySubject/value/"
)

// AllCertificatesBySubjectKey returns the store key to retrieve a AllCertificatesBySubject from the index fields
func AllCertificatesBySubjectKey(
	subject string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	return key
}
