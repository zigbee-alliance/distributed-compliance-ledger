// Copyright 2022 DSR Corporation
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

func (d ProposedCertificateRevocation) HasApprovalFrom(address sdk.Address) bool {
	addrStr := address.String()
	for _, approval := range d.Approvals {
		if approval == addrStr {
			return true
		}
	}
	return false
}
