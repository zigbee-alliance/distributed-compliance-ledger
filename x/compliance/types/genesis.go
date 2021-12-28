package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ComplianceInfoList: []ComplianceInfo{},
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
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
