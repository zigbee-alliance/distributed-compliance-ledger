package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CertificateType string

const (
	RootCertificate         CertificateType = "root"
	IntermediateCertificate CertificateType = "intermediate"
)

/*
	Approved Root / Intermediate / Leaf certificate stored in KVStore
*/
type Certificate struct {
	PemCert       string          `json:"pem_cert"`
	Subject       string          `json:"subject"`
	SubjectKeyId  string          `json:"subject_key_id"`
	SerialNumber  string          `json:"serial_number"`
	Type          CertificateType `json:"type"`
	Owner         sdk.AccAddress  `json:"owner"`
	RootSubjectId string          `json:"root_subject_key_id"`
}

func NewRootCertificate(pemCert string, subject string, subjectKeyId string, serialNumber string, owner sdk.AccAddress) Certificate {
	return Certificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		SerialNumber: serialNumber,
		Type:         RootCertificate,
		Owner:        owner,
	}
}

func NewIntermediateCertificate(pemCert string, subject string, subjectKeyId string, serialNumber string, rootSubjectId string, owner sdk.AccAddress) Certificate {
	return Certificate{
		PemCert:       pemCert,
		Subject:       subject,
		SubjectKeyId:  subjectKeyId,
		SerialNumber:  serialNumber,
		Type:          IntermediateCertificate,
		Owner:         owner,
		RootSubjectId: rootSubjectId,
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
	SubjectKeyId string           `json:"subject_key_id"`
	SerialNumber string           `json:"serial_number"`
	Approvals    []sdk.AccAddress `json:"approvals"`
	Owner        sdk.AccAddress   `json:"owner"`
}

func NewProposedCertificate(pemCert string, subject string, subjectKeyId string, serialNumber string, owner sdk.AccAddress) ProposedCertificate {
	return ProposedCertificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		SerialNumber: serialNumber,
		Approvals:    []sdk.AccAddress{},
		Owner:        owner,
	}
}

func (d ProposedCertificate) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (d ProposedCertificate) HasApprovalFrom(address sdk.AccAddress) bool {
	for _, approval := range d.Approvals {
		if approval.Equals(address) {
			return true
		}
	}
	return false
}

/*
	The list of direct child certificates (depending of Subject/SubjectKeyId parent certificate ) stored in KVStore
*/
type ChildCertificates struct {
	Subject           string                  `json:"subject"`
	SubjectKeyId      string                  `json:"subject_key_id"`
	ChildCertificates []CertificateIdentified `json:"child_certificates"`
}

func NewChildCertificates(subject string, subjectKeyId string) ChildCertificates {
	return ChildCertificates{
		Subject:           subject,
		SubjectKeyId:      subjectKeyId,
		ChildCertificates: []CertificateIdentified{},
	}
}

func (d *ChildCertificates) AddChildCertificate(keyId CertificateIdentified) {
	d.ChildCertificates = append(d.ChildCertificates, keyId)
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

type CertificateIdentified struct {
	Subject      string `json:"subject"`
	SubjectKeyId string `json:"subject_key_id"`
}

func NewCertificateIdentifier(subject string, subjectKeyId string) CertificateIdentified {
	return CertificateIdentified{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}
}
