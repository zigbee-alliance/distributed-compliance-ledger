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

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ComplianceInfoList:   []ComplianceInfo{},
		CertifiedModelList:   []CertifiedModel{},
		RevokedModelList:     []RevokedModel{},
		ProvisionalModelList: []ProvisionalModel{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in complianceInfo
	complianceInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.ComplianceInfoList {
		index := string(ComplianceInfoKey(elem.Vid, elem.Pid, elem.SoftwareVersion, elem.CertificationType))
		if _, ok := complianceInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for complianceInfo")
		}
		complianceInfoIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in certifiedModel
	certifiedModelIndexMap := make(map[string]struct{})

	for _, elem := range gs.CertifiedModelList {
		index := string(CertifiedModelKey(elem.Vid, elem.Pid, elem.SoftwareVersion, elem.CertificationType))
		if _, ok := certifiedModelIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for certifiedModel")
		}
		certifiedModelIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revokedModel
	revokedModelIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevokedModelList {
		index := string(RevokedModelKey(elem.Vid, elem.Pid, elem.SoftwareVersion, elem.CertificationType))
		if _, ok := revokedModelIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revokedModel")
		}
		revokedModelIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in provisionalModel
	provisionalModelIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProvisionalModelList {
		index := string(ProvisionalModelKey(elem.Vid, elem.Pid, elem.SoftwareVersion, elem.CertificationType))
		if _, ok := provisionalModelIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for provisionalModel")
		}
		provisionalModelIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
