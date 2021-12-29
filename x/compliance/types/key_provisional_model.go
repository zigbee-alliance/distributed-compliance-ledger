package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProvisionalModelKeyPrefix is the prefix to retrieve all ProvisionalModel
	ProvisionalModelKeyPrefix = "ProvisionalModel/value/"
)

// ProvisionalModelKey returns the store key to retrieve a ProvisionalModel from the index fields
func ProvisionalModelKey(
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) []byte {
	var key []byte

	vidBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(vidBytes, uint32(vid))
	key = append(key, vidBytes...)
	key = append(key, []byte("/")...)

	pidBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(pidBytes, uint32(pid))
	key = append(key, pidBytes...)
	key = append(key, []byte("/")...)

	softwareVersionBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(softwareVersionBytes, softwareVersion)
	key = append(key, softwareVersionBytes...)
	key = append(key, []byte("/")...)

	certificationTypeBytes := []byte(certificationType)
	key = append(key, certificationTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}
