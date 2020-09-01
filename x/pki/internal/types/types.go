package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CertificateType string

/*
	Approved Root / Intermediate / Leaf certificates stored in KVStore and matching to the same key
*/
type Certificates struct {
	Items []Certificate `json:"items"`
}

func NewCertificates(items []Certificate) Certificates {
	return Certificates{
		Items: items,
	}
}

func (d Certificates) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	Single Approved Root / Intermediate / Leaf certificate
*/
type Certificate struct {
	PemCert          string         `json:"pem_cert"`
	Subject          string         `json:"subject"`
	SubjectKeyID     string         `json:"subject_key_id"`
	SerialNumber     string         `json:"serial_number"`
	Issuer           string         `json:"issuer,omitempty"`
	AuthorityKeyID   string         `json:"authority_key_id,omitempty"`
	RootSubject      string         `json:"root_subject,omitempty"`
	RootSubjectKeyID string         `json:"root_subject_key_id,omitempty"`
	IsRoot           bool           `json:"is_root"`
	Owner            sdk.AccAddress `json:"owner"`
}

func NewRootCertificate(pemCert string, subject string, subjectKeyID string,
	serialNumber string, owner sdk.AccAddress) Certificate {
	return Certificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		SerialNumber: serialNumber,
		IsRoot:       true,
		Owner:        owner,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectKeyID string, serialNumber string,
	issuer string, authorityKeyID string,
	rootSubject string, rootSubjectKeyID string,
	owner sdk.AccAddress) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectKeyID:     subjectKeyID,
		SerialNumber:     serialNumber,
		Issuer:           issuer,
		AuthorityKeyID:   authorityKeyID,
		RootSubject:      rootSubject,
		RootSubjectKeyID: rootSubjectKeyID,
		IsRoot:           false,
		Owner:            owner,
	}
}

func (d Certificate) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	Proposed (but not Approved yet) Root certificate stored in KVStore
*/
type ProposedCertificate struct {
	PemCert      string           `json:"pem_cert"`
	Subject      string           `json:"subject"`
	SubjectKeyID string           `json:"subject_key_id"`
	SerialNumber string           `json:"serial_number"`
	Owner        sdk.AccAddress   `json:"owner"`
	Approvals    []sdk.AccAddress `json:"approvals"`
}

func NewProposedCertificate(pemCert string, subject string, subjectKeyID string,
	serialNumber string, owner sdk.AccAddress) ProposedCertificate {
	return ProposedCertificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		SerialNumber: serialNumber,
		Owner:        owner,
		Approvals:    []sdk.AccAddress{},
	}
}

func (d ProposedCertificate) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (d ProposedCertificate) HasApprovalFrom(address sdk.Address) bool {
	for _, approval := range d.Approvals {
		if approval.Equals(address) {
			return true
		}
	}

	return false
}

/*
	The list of certificates issued by a given issuer
*/
type ChildCertificates struct {
	Issuer          string                  `json:"issues"`
	AuthorityKeyID  string                  `json:"authority_key_id"`
	CertIdentifiers []CertificateIdentifier `json:"cert_identifiers"`
}

func NewChildCertificates(issuer string, authorityKeyID string) ChildCertificates {
	return ChildCertificates{
		Issuer:          issuer,
		AuthorityKeyID:  authorityKeyID,
		CertIdentifiers: []CertificateIdentifier{},
	}
}

func (d ChildCertificates) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	Composed identifier for certificates
*/

type CertificateIdentifier struct {
	Subject      string `json:"subject"`
	SubjectKeyID string `json:"subject_key_id"`
}

func NewCertificateIdentifier(subject string, subjectKeyID string) CertificateIdentifier {
	return CertificateIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
}

/*
	Proposed (but not Approved yet) Revocation of Root certificate stored in KVStore
*/
type ProposedCertificateRevocation struct {
	Subject      string           `json:"subject"`
	SubjectKeyID string           `json:"subject_key_id"`
	Approvals    []sdk.AccAddress `json:"approvals"`
}

func NewProposedCertificateRevocation(subject string, subjectKeyID string,
	approval sdk.AccAddress) ProposedCertificateRevocation {
	return ProposedCertificateRevocation{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		Approvals:    []sdk.AccAddress{approval},
	}
}

func (d ProposedCertificateRevocation) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (d ProposedCertificateRevocation) HasApprovalFrom(address sdk.Address) bool {
	for _, approval := range d.Approvals {
		if approval.Equals(address) {
			return true
		}
	}

	return false
}
