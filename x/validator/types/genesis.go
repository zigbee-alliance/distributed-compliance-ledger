package types

/* FIXME issue 99 */

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValidatorList:          []Validator{},
		LastValidatorPowerList: []LastValidatorPower{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorList {
		owner, err := sdk.ValAddressFromBech32(elem.Owner)
		if err != nil {
			return err
		}
		index := string(ValidatorKey(owner))
		if _, ok := validatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validator")
		}
		validatorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in lastValidatorPower
	lastValidatorPowerIndexMap := make(map[string]struct{})

	for _, elem := range gs.LastValidatorPowerList {
		owner, err := sdk.ValAddressFromBech32(elem.Owner)
		if err != nil {
			return err
		}
		index := string(LastValidatorPowerKey(owner))
		if _, ok := lastValidatorPowerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lastValidatorPower")
		}
		lastValidatorPowerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
