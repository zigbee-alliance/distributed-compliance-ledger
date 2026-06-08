package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesBySubjectKeyIDKeyPrefix is the prefix to retrieve all NocCertificatesBySubjectKeyID
	NocCertificatesBySubjectKeyIDKeyPrefix = "NocCertificatesBySubjectKeyID/value/"
)

// NocCertificatesBySubjectKeyIDKey returns the store key to retrieve a NocCertificatesBySubjectKeyID from the index fields
func NocCertificatesBySubjectKeyIDKey(
	subjectKeyID string,
) []byte {
	var key []byte

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
