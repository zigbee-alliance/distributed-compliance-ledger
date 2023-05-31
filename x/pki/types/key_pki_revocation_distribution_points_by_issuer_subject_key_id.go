package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix is the prefix to retrieve all PkiRevocationDistributionPointsByIssuerSubjectKeyId
	PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix = "PkiRevocationDistributionPointsByIssuerSubjectKeyId/value/"
)

// PkiRevocationDistributionPointsByIssuerSubjectKeyIdKey returns the store key to retrieve a PkiRevocationDistributionPointsByIssuerSubjectKeyId from the index fields
func PkiRevocationDistributionPointsByIssuerSubjectKeyIdKey(
	issuerSubjectKeyId string,
) []byte {
	var key []byte

	issuerSubjectKeyIdBytes := []byte(issuerSubjectKeyId)
	key = append(key, issuerSubjectKeyIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
