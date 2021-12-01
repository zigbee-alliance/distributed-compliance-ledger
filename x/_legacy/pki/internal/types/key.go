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
