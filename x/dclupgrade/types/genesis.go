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
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
