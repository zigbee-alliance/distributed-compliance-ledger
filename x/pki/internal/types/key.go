package types

const (
	// ModuleName is the name of the module
	ModuleName = "pki"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// prefix for each key to a proposed certificate.
	ProposedCertificatePrefix = []byte{0x01}
	// prefix for each key to an approved certificate.
	ApprovedCertificatePrefix = []byte{0x02}
	// prefix for a helper index containing the list of child certificates.
	ChildCertificatesPrefix = []byte{0x03}
	// prefix for a helper index containing existence flags for certificate issuer/serial number.
	CertificateByIssuerSerialNumberPrefix = []byte{0x06}
)

// Key builder for Approved Certificate.
func GetApprovedCertificateKey(subject string, subjectKeyID string) []byte {
	return append(ApprovedCertificatePrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for Proposed Certificate.
func GetProposedCertificateKey(subject string, subjectKeyID string) []byte {
	return append(ProposedCertificatePrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for the list of Child Certificates.
func GetChildCertificatesKey(subject string, subjectKeyID string) []byte {
	return append(ChildCertificatesPrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for Existence flag.
func GetCertificateByIssuerSerialNumberKey(issuer string, serialNumber string) []byte {
	return append(CertificateByIssuerSerialNumberPrefix, append([]byte(issuer), []byte(serialNumber)...)...)
}
