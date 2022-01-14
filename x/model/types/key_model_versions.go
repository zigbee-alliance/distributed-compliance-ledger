package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModelVersionsKeyPrefix is the prefix to retrieve all ModelVersions.
	ModelVersionsKeyPrefix = "ModelVersions/value/"
)

// ModelVersionsKey returns the store key to retrieve a ModelVersions from the index fields.
func ModelVersionsKey(
	vid int32,
	pid int32,
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

	return key
}
