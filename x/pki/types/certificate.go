package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewRootCertificate(pemCert string, subject string, subjectKeyId string,
	serialNumber string, owner string) Certificate {
	return Certificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		SerialNumber: serialNumber,
		IsRoot:       true,
		Owner:        owner,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectKeyId string, serialNumber string,
	issuer string, authorityKeyId string,
	rootSubject string, rootSubjectKeyId string,
	owner string) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectKeyId:     subjectKeyId,
		SerialNumber:     serialNumber,
		Issuer:           issuer,
		AuthorityKeyId:   authorityKeyId,
		RootSubject:      rootSubject,
		RootSubjectKeyId: rootSubjectKeyId,
		IsRoot:           false,
		Owner:            owner,
	}
}

//nolint:interfacer
func (cert ProposedCertificate) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range cert.Approvals {
		if approval == addrStr {
			return true
		}
	}

	return false
}
