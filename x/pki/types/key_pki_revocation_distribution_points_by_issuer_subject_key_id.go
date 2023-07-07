package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix is the prefix to retrieve all PkiRevocationDistributionPointsByIssuerSubjectKeyID.
	PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix = "PkiRevocationDistributionPointsByIssuerSubjectKeyID/value/"
)

// PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey returns the store key to retrieve a PkiRevocationDistributionPointsByIssuerSubjectKeyID from the index fields.
func PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(
	issuerSubjectKeyID string,
) []byte {
	var key []byte

	issuerSubjectKeyIDBytes := []byte(issuerSubjectKeyID)
	key = append(key, issuerSubjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
