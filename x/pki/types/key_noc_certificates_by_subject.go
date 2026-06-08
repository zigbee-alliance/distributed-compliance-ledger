package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesBySubjectKeyPrefix is the prefix to retrieve all NocCertificatesBySubject
	NocCertificatesBySubjectKeyPrefix = "NocCertificatesBySubject/value/"
)

// NocCertificatesBySubjectKey returns the store key to retrieve a NocCertificatesBySubject from the index fields
func NocCertificatesBySubjectKey(
	subject string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	return key
}
