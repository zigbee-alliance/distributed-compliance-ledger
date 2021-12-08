package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VendorProductsList: []VendorProducts{},
		ModelList:          []Model{},
		ModelVersionList:   []ModelVersion{},
		ModelVersionsList:  []ModelVersions{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in vendorProducts
	vendorProductsIndexMap := make(map[string]struct{})

	for _, elem := range gs.VendorProductsList {
		index := string(VendorProductsKey(elem.Vid))
		if _, ok := vendorProductsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vendorProducts")
		}
		vendorProductsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in model
	modelIndexMap := make(map[string]struct{})

	for _, elem := range gs.ModelList {
		index := string(ModelKey(elem.Vid, elem.Pid))
		if _, ok := modelIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for model")
		}
		modelIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in modelVersion
	modelVersionIndexMap := make(map[string]struct{})

	for _, elem := range gs.ModelVersionList {
		index := string(ModelVersionKey(elem.Vid, elem.Pid, elem.SoftwareVersion))
		if _, ok := modelVersionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for modelVersion")
		}
		modelVersionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in modelVersions
	modelVersionsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ModelVersionsList {
		index := string(ModelVersionsKey(elem.Vid, elem.Pid))
		if _, ok := modelVersionsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for modelVersions")
		}
		modelVersionsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
