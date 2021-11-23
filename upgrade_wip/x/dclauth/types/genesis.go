package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountList:                  []Account{},
		PendingAccountList:           []PendingAccount{},
		PendingAccountRevocationList: []PendingAccountRevocation{},
		AccountStat:                  nil,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in account
	accountIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountList {
		index := string(AccountKey(elem.Address))
		if _, ok := accountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for account")
		}
		accountIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingAccount
	pendingAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingAccountList {
		index := string(PendingAccountKey(elem.Address))
		if _, ok := pendingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingAccount")
		}
		pendingAccountIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingAccountRevocation
	pendingAccountRevocationIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingAccountRevocationList {
		index := string(PendingAccountRevocationKey(elem.Address))
		if _, ok := pendingAccountRevocationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingAccountRevocation")
		}
		pendingAccountRevocationIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
