package types

const (
	// ModuleName is the name of the module
	ModuleName = "pki"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	ProposedCertificatePrefix             = []byte{0x1} // prefix for each key to a proposed certificate
	ApprovedCertificatePrefix             = []byte{0x2} // prefix for each key to an approved certificate
	ChildCertificatesPrefix               = []byte{0x3} // prefix for a helper index containing the list of child certificates
	CertificateByIssuerSerialNumberPrefix = []byte{0x6} // prefix for a helper index containing existence flags for certificate issuer/serial number
)

// Key builder for Approved Certificate
func GetApprovedCertificateKey(subject string, subjectKeyId string) []byte {
	return append(ApprovedCertificatePrefix, append([]byte(subject), []byte(subjectKeyId)...)...)
}

// Key builder for Proposed Certificate
func GetProposedCertificateKey(subject string, subjectKeyId string) []byte {
	return append(ProposedCertificatePrefix, append([]byte(subject), []byte(subjectKeyId)...)...)
}

// Key builder for the list of Child Certificates
func GetChildCertificatesKey(subject string, subjectKeyId string) []byte {
	return append(ChildCertificatesPrefix, append([]byte(subject), []byte(subjectKeyId)...)...)
}

// Key builder for Existence flag
func GetCertificateByIssuerSerialNumberKey(issuer string, serialNumber string) []byte {
	return append(CertificateByIssuerSerialNumberPrefix, append([]byte(issuer), []byte(serialNumber)...)...)
}