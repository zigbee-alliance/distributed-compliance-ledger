package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApprovedRootCertificatesKeyPrefix is the prefix to retrieve all ApprovedRootCertificates
	ApprovedRootCertificatesKeyPrefix = "ApprovedRootCertificates/value/"
)
