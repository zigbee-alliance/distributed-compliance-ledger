package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PkiRevocationDistributionPointKeyPrefix is the prefix to retrieve all PkiRevocationDistributionPoint.
	PkiRevocationDistributionPointKeyPrefix = "PkiRevocationDistributionPoint/value/"
)

// PkiRevocationDistributionPointKey returns the store key to retrieve a PkiRevocationDistributionPoint from the index fields.
func PkiRevocationDistributionPointKey(
	vid int32,
	label string,
	issuerSubjectKeyID string,
) []byte {
	var key []byte

	vidBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	labelBytes := []byte(label)
	key = append(key, labelBytes...)
	key = append(key, []byte("/")...)

	issuerSubjectKeyIDBytes := []byte(issuerSubjectKeyID)
	key = append(key, issuerSubjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
