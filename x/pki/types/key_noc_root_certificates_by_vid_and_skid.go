package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocRootCertificatesByVidAndSkidKeyPrefix is the prefix to retrieve all NocRootCertificatesByVidAndSkid
	NocRootCertificatesByVidAndSkidKeyPrefix = "NocRootCertificatesByVidAndSkid/value/"
)

// NocRootCertificatesByVidAndSkidKey returns the store key to retrieve a NocRootCertificatesByVidAndSkid from the index fields
func NocRootCertificatesByVidAndSkidKey(
	vid int32,
	subjectKeyId string,
) []byte {
	var key []byte

	vidBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	subjectKeyIdBytes := []byte(subjectKeyId)
	key = append(key, subjectKeyIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
