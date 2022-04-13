package types

import (
	"encoding/json"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/* FIXME issue 99 */

// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValidatorList:                []Validator{},
		LastValidatorPowerList:       []LastValidatorPower{},
		ProposedDisableValidatorList: []ProposedDisableValidator{},
		DisabledValidatorList:        []DisabledValidator{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// TODO issue 99: review - cosmos checks duplication for consensus addr here

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
	// Check for duplicated index in proposedDisableValidator
	proposedDisableValidatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposedDisableValidatorList {
		index := string(ProposedDisableValidatorKey(elem.Address))
		if _, ok := proposedDisableValidatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposedDisableValidator")
		}
		proposedDisableValidatorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in disabledValidator
	disabledValidatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.DisabledValidatorList {
		index := string(DisabledValidatorKey(elem.Address))
		if _, ok := disabledValidatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for disabledValidator")
		}
		disabledValidatorIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

// GetGenesisStateFromAppState returns x/staking GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
