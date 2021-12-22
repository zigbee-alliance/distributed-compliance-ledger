package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChildCertificatesKeyPrefix is the prefix to retrieve all ChildCertificates
	ChildCertificatesKeyPrefix = "ChildCertificates/value/"
)

// ChildCertificatesKey returns the store key to retrieve a ChildCertificates from the index fields
func ChildCertificatesKey(
	issuer string,
	authorityKeyId string,
) []byte {
	var key []byte

	issuerBytes := []byte(issuer)
	key = append(key, issuerBytes...)
	key = append(key, []byte("/")...)

	authorityKeyIdBytes := []byte(authorityKeyId)
	key = append(key, authorityKeyIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
