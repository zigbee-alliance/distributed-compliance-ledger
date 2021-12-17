package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniqueCertificateKeyPrefix is the prefix to retrieve all UniqueCertificate
	UniqueCertificateKeyPrefix = "UniqueCertificate/value/"
)
