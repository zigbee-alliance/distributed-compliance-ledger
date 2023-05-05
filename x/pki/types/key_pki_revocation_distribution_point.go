package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // PKIRevocationDistributionPointKeyPrefix is the prefix to retrieve all PKIRevocationDistributionPoint
	PKIRevocationDistributionPointKeyPrefix = "PKIRevocationDistributionPoint/value/"
)

// PKIRevocationDistributionPointKey returns the store key to retrieve a PKIRevocationDistributionPoint from the index fields
func PKIRevocationDistributionPointKey(
vid uint64,
label string,
issuerSubjectKeyID string,
) []byte {
	var key []byte
    
    vidBytes := make([]byte, 8)
  					binary.BigEndian.PutUint64(vidBytes, vid)
    key = append(key, vidBytes...)
    key = append(key, []byte("/")...)
    
    labelBytes := []byte(label)
    key = append(key, labelBytes...)
    key = append(key, []byte("/")...)
    
    issuerSubjectKeyIDBytes := []byte(issuerSubjectKeyID)
    key = append(key, issuerSubjectKeyIDBytes...)
    key = append(key, []byte("/")...)
    
	return key
}