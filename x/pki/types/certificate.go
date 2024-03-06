package types

func NewRootCertificate(pemCert string, subject string, subjectAsText string, subjectKeyID string,
	serialNumber string, owner string, approvals []*Grant, rejects []*Grant, vid int32,
) Certificate {
	return Certificate{
		PemCert:       pemCert,
		Subject:       subject,
		SubjectAsText: subjectAsText,
		SubjectKeyId:  subjectKeyID,
		SerialNumber:  serialNumber,
		IsRoot:        true,
		Owner:         owner,
		Approvals:     approvals,
		Rejects:       rejects,
		Vid:           vid,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectAsText string, subjectKeyID string, serialNumber string,
	issuer string, authorityKeyID string,
	rootSubject string, rootSubjectKeyID string,
	owner string,
) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectAsText:    subjectAsText,
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

func NewNocRootCertificate(
	pemCert string,
	subject string,
	subjectAsText string,
	subjectKeyID string,
	serialNumber string,
	owner string,
	vid int32,
) Certificate {
	return Certificate{
		PemCert:       pemCert,
		Subject:       subject,
		SubjectAsText: subjectAsText,
		SubjectKeyId:  subjectKeyID,
		SerialNumber:  serialNumber,
		IsRoot:        true,
		Owner:         owner,
		Vid:           vid,
		IsNoc:         true,
	}
}

func NewNocCertificate(
	pemCert string,
	subject string,
	subjectAsText string,
	subjectKeyID string,
	serialNumber string,
	issuer string,
	authorityKeyID string,
	rootSubject string,
	rootSubjectKeyID string,
	owner string,
	vid int32,
) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectAsText:    subjectAsText,
		SubjectKeyId:     subjectKeyID,
		SerialNumber:     serialNumber,
		Issuer:           issuer,
		AuthorityKeyId:   authorityKeyID,
		RootSubject:      rootSubject,
		RootSubjectKeyId: rootSubjectKeyID,
		Vid:              vid,
		Owner:            owner,
		IsRoot:           false,
		IsNoc:            true,
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

func (cert ProposedCertificate) HasRejectFrom(address string) bool {
	for _, rejectApproval := range cert.Rejects {
		if rejectApproval.Address == address {
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

func (cert Certificate) HasApprovalFrom(address string) bool {
	for _, approval := range cert.Approvals {
		if approval.Address == address {
			return true
		}
	}

	return false
}
