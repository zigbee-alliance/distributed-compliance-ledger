package types

import "regexp"

var allowedDataDigestTypes = [6]uint32{1, 7, 8, 10, 11, 12}
var allowedRevocationTypes = [1]uint32{1}

func VerifyRevocationPointIssuerSubjectKeyIDFormat(issuerSubjectKeyID string) bool {
	match, _ := regexp.MatchString("^(?:[0-9A-F]{2})+$", issuerSubjectKeyID)

	return match
}
