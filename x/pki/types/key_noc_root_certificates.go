package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocRootCertificatesKeyPrefix is the prefix to retrieve all NocRootCertificates
	NocRootCertificatesKeyPrefix = "NocRootCertificates/value/"
)

// NocRootCertificatesKey returns the store key to retrieve a NocRootCertificates from the index fields
func NocRootCertificatesKey(
	vid int32,
) []byte {
	var key []byte

	vidBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	return key
}
