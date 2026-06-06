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

import (
	"fmt"
)

// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProposedUpgradeList: []ProposedUpgrade{},
		ApprovedUpgradeList: []ApprovedUpgrade{},
		RejectedUpgradeList: []RejectedUpgrade{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in proposedUpgrade
	proposedUpgradeIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedUpgradeList {
		index := string(ProposedUpgradeKey(elem.Plan.Name))
		if _, ok := proposedUpgradeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedUpgrade")
		}

		proposedUpgradeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in approvedUpgrade
	approvedUpgradeIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApprovedUpgradeList {
		index := string(ApprovedUpgradeKey(elem.Plan.Name))
		if _, ok := approvedUpgradeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for approvedUpgrade")
		}

		approvedUpgradeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in rejectedUpgrade
	rejectedUpgradeIndexMap := make(map[string]struct{})

	for _, elem := range gs.RejectedUpgradeList {
		index := string(RejectedUpgradeKey(elem.Plan.Name))
		if _, ok := rejectedUpgradeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for rejectedUpgrade")
		}
		rejectedUpgradeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
