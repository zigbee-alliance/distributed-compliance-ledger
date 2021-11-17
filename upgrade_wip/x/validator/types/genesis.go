package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValidatorList:                    []Validator{},
		LastValidatorPowerList:           []LastValidatorPower{},
		ValidatorSigningInfoList:         []ValidatorSigningInfo{},
		ValidatorMissedBlockBitArrayList: []ValidatorMissedBlockBitArray{},
		ValidatorOwnerList:               []ValidatorOwner{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorList {
		index := string(ValidatorKey(elem.Address))
		if _, ok := validatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validator")
		}
		validatorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in lastValidatorPower
	lastValidatorPowerIndexMap := make(map[string]struct{})

	for _, elem := range gs.LastValidatorPowerList {
		index := string(LastValidatorPowerKey(elem.ConsensusAddress))
		if _, ok := lastValidatorPowerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lastValidatorPower")
		}
		lastValidatorPowerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in validatorSigningInfo
	validatorSigningInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorSigningInfoList {
		index := string(ValidatorSigningInfoKey(elem.Address))
		if _, ok := validatorSigningInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validatorSigningInfo")
		}
		validatorSigningInfoIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in validatorMissedBlockBitArray
	validatorMissedBlockBitArrayIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorMissedBlockBitArrayList {
		index := string(ValidatorMissedBlockBitArrayKey(elem.Address, elem.Index))
		if _, ok := validatorMissedBlockBitArrayIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validatorMissedBlockBitArray")
		}
		validatorMissedBlockBitArrayIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in validatorOwner
	validatorOwnerIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorOwnerList {
		index := string(ValidatorOwnerKey(elem.Address))
		if _, ok := validatorOwnerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validatorOwner")
		}
		validatorOwnerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
