// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

func NewRootCertificate(pemCert string, subject string, subjectAsText string, subjectKeyID string,
	serialNumber string, owner string, approvals []*Grant, rejects []*Grant, vid int32, schemaVersion uint32,
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
		SchemaVersion: schemaVersion,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectAsText string, subjectKeyID string, serialNumber string,
	issuer string, authorityKeyID string,
	rootSubject string, rootSubjectKeyID string,
	owner string, vid int32,
	schemaVersion uint32,
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
		Vid:              vid,
		SchemaVersion:    schemaVersion,
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
	schemaVersion uint32,
	certificateType CertificateType,
) Certificate {
	return Certificate{
		PemCert:         pemCert,
		Subject:         subject,
		SubjectAsText:   subjectAsText,
		SubjectKeyId:    subjectKeyID,
		SerialNumber:    serialNumber,
		IsRoot:          true,
		Owner:           owner,
		Vid:             vid,
		CertificateType: certificateType,
		SchemaVersion:   schemaVersion,
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
	schemaVersion uint32,
	certificateType CertificateType,
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
		CertificateType:  certificateType,
		SchemaVersion:    schemaVersion,
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
