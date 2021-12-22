package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniqueCertificateKeyPrefix is the prefix to retrieve all UniqueCertificate
	UniqueCertificateKeyPrefix = "UniqueCertificate/value/"
)

// UniqueCertificateKey returns the store key to retrieve a UniqueCertificate from the index fields
func UniqueCertificateKey(
	issuer string,
	serialNumber string,
) []byte {
	var key []byte

	issuerBytes := []byte(issuer)
	key = append(key, issuerBytes...)
	key = append(key, []byte("/")...)

	serialNumberBytes := []byte(serialNumber)
	key = append(key, serialNumberBytes...)
	key = append(key, []byte("/")...)

	return key
}
