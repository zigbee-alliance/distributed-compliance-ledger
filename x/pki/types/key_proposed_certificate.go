package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposedCertificateKeyPrefix is the prefix to retrieve all ProposedCertificate
	ProposedCertificateKeyPrefix = "ProposedCertificate/value/"
)

// ProposedCertificateKey returns the store key to retrieve a ProposedCertificate from the index fields
func ProposedCertificateKey(
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
