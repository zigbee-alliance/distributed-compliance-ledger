package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesBySubjectKeyIdKeyPrefix is the prefix to retrieve all NocCertificatesBySubjectKeyId
	NocCertificatesBySubjectKeyIdKeyPrefix = "NocCertificatesBySubjectKeyId/value/"
)

// NocCertificatesBySubjectKeyIdKey returns the store key to retrieve a NocCertificatesBySubjectKeyId from the index fields
func NocCertificatesBySubjectKeyIdKey(
	subjectKeyID string,
) []byte {
	var key []byte

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
