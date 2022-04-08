package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO issue 99: do we need that
// DefaultIndex is the default capability global index.
const DefaultIndex uint64 = 1

// TODO issue 99: review - do we need pack/unpack/sanitize for accounts
//	              data in genesis as it is implemented in cosmos now

// DefaultGenesis returns the default Capability genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountList:                  []Account{},
		PendingAccountList:           []PendingAccount{},
		PendingAccountRevocationList: []PendingAccountRevocation{},
		AccountStat:                  nil,
		RevokedAccountList:           []RevokedAccount{},
		RejectedAccountList:          []RejectedAccount{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in account
	accountIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountList {
		addr, err := sdk.AccAddressFromBech32(elem.Address)
		if err != nil {
			return err
		}
		index := string(AccountKey(addr))
		if _, ok := accountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for account")
		}
		accountIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingAccount
	pendingAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingAccountList {
		addr, err := sdk.AccAddressFromBech32(elem.Address)
		if err != nil {
			return err
		}
		index := string(PendingAccountKey(addr))
		if _, ok := pendingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingAccount")
		}
		pendingAccountIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingAccountRevocation
	pendingAccountRevocationIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingAccountRevocationList {
		addr, err := sdk.AccAddressFromBech32(elem.Address)
		if err != nil {
			return err
		}
		index := string(PendingAccountRevocationKey(addr))
		if _, ok := pendingAccountRevocationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingAccountRevocation")
		}
		pendingAccountRevocationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in revokedAccount
	revokedAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.RevokedAccountList {
		index := string(RevokedAccountKey(elem.GetAddress()))
		if _, ok := revokedAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for revokedAccount")
		}
		revokedAccountIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in rejectedAccount
	rejectedAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.RejectedAccountList {
		index := string(RejectedAccountKey(elem.Address))
		if _, ok := rejectedAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for rejectedAccount")
		}
		rejectedAccountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

// GetGenesisStateFromAppState returns x/bank GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// GenesisAccountsIterator implements genesis account iteration.
type GenesisAccountsIterator struct{}

// IterateGenesisAccounts iterates over all the genesis accounts found in
// appGenesis and invokes a callback on each genesis account. If any call
// returns true, iteration stops.
func (GenesisAccountsIterator) IterateGenesisAccounts(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, cb func(GenesisAccount) (stop bool),
) {
	for _, account := range GetGenesisStateFromAppState(cdc, appState).AccountList {
		if cb(account) {
			break
		}
	}
}
