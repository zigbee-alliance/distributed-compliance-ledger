package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModelKeyPrefix is the prefix to retrieve all Model.
	ModelKeyPrefix = "Model/value/"
)

// ModelKey returns the store key to retrieve a Model from the index fields.
func ModelKey(
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
