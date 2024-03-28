package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesKeyPrefix is the prefix to retrieve all NocCertificates.
	NocCertificatesKeyPrefix = "NocCertificates/value/"
)

// NocCertificatesKey returns the store key to retrieve a NocCertificates from the index fields.
func NocCertificatesKey(
	vid int32,
) []byte {
	var key []byte

	vidBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	return key
}
