package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModelVersionKeyPrefix is the prefix to retrieve all ModelVersion.
	ModelVersionKeyPrefix = "ModelVersion/value/"
)

// ModelVersionKey returns the store key to retrieve a ModelVersion from the index fields.
func ModelVersionKey(
	vid int32,
	pid int32,
	softwareVersion uint32,
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

	return key
}
