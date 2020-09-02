package types

const (
	// ModuleName is the name of the module.
	ModuleName = "pki"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var (
	// prefix for each key to a proposed certificate.
	ProposedCertificatePrefix = []byte{0x01}
	// prefix for each key to an approved certificate.
	ApprovedCertificatePrefix = []byte{0x02}
	// prefix for a helper index containing the list of child certificates.
	ChildCertificatesPrefix = []byte{0x03}
	// prefix for each key to a proposed certificate revocation.
	ProposedCertificateRevocationPrefix = []byte{0x04}
	// prefix for each key to a revoked certificate.
	RevokedCertificatePrefix = []byte{0x05}
	// prefix for each key to a certificate existence flag.
	UniqueCertificateKeyPrefix = []byte{0x06}
)

// Key builder for Proposed Certificate.
func GetProposedCertificateKey(subject string, subjectKeyID string) []byte {
	return append(ProposedCertificatePrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for Approved Certificate.
func GetApprovedCertificateKey(subject string, subjectKeyID string) []byte {
	return append(ApprovedCertificatePrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for the list of Child Certificates.
func GetChildCertificatesKey(issuer string, authorityKeyID string) []byte {
	return append(ChildCertificatesPrefix, append([]byte(issuer), []byte(authorityKeyID)...)...)
}

// Key builder for Proposed Certificate Revocation.
func GetProposedCertificateRevocationKey(subject string, subjectKeyID string) []byte {
	return append(ProposedCertificateRevocationPrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for Revoked Certificate.
func GetRevokedCertificateKey(subject string, subjectKeyID string) []byte {
	return append(RevokedCertificatePrefix, append([]byte(subject), []byte(subjectKeyID)...)...)
}

// Key builder for Certificate Existence Flag.
func GetUniqueCertificateKey(issuer string, serialNumber string) []byte {
	return append(UniqueCertificateKeyPrefix, append([]byte(issuer), []byte(serialNumber)...)...)
}
