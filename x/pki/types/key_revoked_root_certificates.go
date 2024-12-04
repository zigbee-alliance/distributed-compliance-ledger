package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RevokedRootCertificatesKeyPrefix is the prefix to retrieve all RevokedRootCertificates
	RevokedRootCertificatesKeyPrefix = "RevokedRootCertificates/value/"
)
