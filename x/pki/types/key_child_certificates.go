package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChildCertificatesKeyPrefix is the prefix to retrieve all ChildCertificates.
	ChildCertificatesKeyPrefix = "ChildCertificates/value/"
)

// ChildCertificatesKey returns the store key to retrieve a ChildCertificates from the index fields.
func ChildCertificatesKey(
	issuer string,
	authorityKeyID string,
) []byte {
	var key []byte

	issuerBytes := []byte(issuer)
	key = append(key, issuerBytes...)
	key = append(key, []byte("/")...)

	authorityKeyIDBytes := []byte(authorityKeyID)
	key = append(key, authorityKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
