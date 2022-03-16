package types

func NewRootCertificate(pemCert string, subject string, subjectKeyID string,
	serialNumber string, owner string, approvals []*Grant) Certificate {
	return Certificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
		IsRoot:       true,
		Owner:        owner,
		Approvals:    approvals,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectKeyID string, serialNumber string,
	issuer string, authorityKeyID string,
	rootSubject string, rootSubjectKeyID string,
	owner string) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectKeyId:     subjectKeyID,
		SerialNumber:     serialNumber,
		Issuer:           issuer,
		AuthorityKeyId:   authorityKeyID,
		RootSubject:      rootSubject,
		RootSubjectKeyId: rootSubjectKeyID,
		IsRoot:           false,
		Owner:            owner,
	}
}

func (cert ProposedCertificate) HasApprovalFrom(address string) bool {
	for _, approval := range cert.Approvals {
		if approval.Address == address {
			return true
		}
	}

	return false
}

func (d ProposedCertificateRevocation) HasApprovalFrom(address string) bool {
	for _, approvals := range d.Approvals {
		if approvals.Address == address {
			return true
		}
	}

	return false
}
