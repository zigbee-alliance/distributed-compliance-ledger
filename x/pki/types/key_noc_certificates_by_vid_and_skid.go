package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesByVidAndSkidKeyPrefix is the prefix to retrieve all NocCertificatesByVidAndSkid.
	NocCertificatesByVidAndSkidKeyPrefix = "NocCertificatesByVidAndSkid/value/"
)

// NocCertificatesByVidAndSkidKey returns the store key to retrieve a NocCertificatesByVidAndSkid from the index fields.
func NocCertificatesByVidAndSkidKey(
	vid int32,
	subjectKeyID string,
) []byte {
	var key []byte

	vidBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
