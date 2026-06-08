package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposedCertificateKeyPrefix is the prefix to retrieve all ProposedCertificate.
	ProposedCertificateKeyPrefix = "ProposedCertificate/value/"
)

// ProposedCertificateKey returns the store key to retrieve a ProposedCertificate from the index fields.
func ProposedCertificateKey(
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
