package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TestingResultsList: []TestingResults{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in testingResults
	testingResultsIndexMap := make(map[string]struct{})

	for _, elem := range gs.TestingResultsList {
		index := string(TestingResultsKey(elem.Vid, elem.Pid, elem.SoftwareVersion))
		if _, ok := testingResultsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for testingResults")
		}
		testingResultsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
