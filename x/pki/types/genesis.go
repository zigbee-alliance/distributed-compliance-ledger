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

import (
	"fmt"
)

// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ApprovedCertificatesList:          []ApprovedCertificates{},
		ProposedCertificateList:           []ProposedCertificate{},
		ChildCertificatesList:             []ChildCertificates{},
		ProposedCertificateRevocationList: []ProposedCertificateRevocation{},
		RevokedCertificatesList:           []RevokedCertificates{},
		UniqueCertificateList:             []UniqueCertificate{},
		ApprovedRootCertificates:          nil,
		RevokedRootCertificates:           nil,
		ApprovedCertificatesBySubjectList: []ApprovedCertificatesBySubject{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in approvedCertificates
	approvedCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedCertificatesList {
		index := string(ApprovedCertificatesKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := approvedCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedCertificates")
		}
		approvedCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposedCertificate
	proposedCertificateIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedCertificateList {
		index := string(ProposedCertificateKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := proposedCertificateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedCertificate")
		}
		proposedCertificateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in childCertificates
	childCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChildCertificatesList {
		index := string(ChildCertificatesKey(elem.Issuer, elem.AuthorityKeyId))
		if _, ok := childCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for childCertificates")
		}
		childCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposedCertificateRevocation
	proposedCertificateRevocationIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedCertificateRevocationList {
		index := string(ProposedCertificateRevocationKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := proposedCertificateRevocationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedCertificateRevocation")
		}
		proposedCertificateRevocationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revokedCertificates
	revokedCertificatesIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevokedCertificatesList {
		index := string(RevokedCertificatesKey(elem.Subject, elem.SubjectKeyId))
		if _, ok := revokedCertificatesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revokedCertificates")
		}
		revokedCertificatesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in uniqueCertificate
	uniqueCertificateIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniqueCertificateList {
		index := string(UniqueCertificateKey(elem.Issuer, elem.SerialNumber))
		if _, ok := uniqueCertificateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for uniqueCertificate")
		}
		uniqueCertificateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in approvedCertificatesBySubject
	approvedCertificatesBySubjectIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedCertificatesBySubjectList {
		index := string(ApprovedCertificatesBySubjectKey(elem.Subject))
		if _, ok := approvedCertificatesBySubjectIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedCertificatesBySubject")
		}
		approvedCertificatesBySubjectIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
