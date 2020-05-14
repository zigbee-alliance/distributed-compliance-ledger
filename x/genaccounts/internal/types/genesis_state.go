package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"sort"
)

// State to Unmarshal
type GenesisState GenesisAccounts

// get the genesis state from the expected app state
func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

// Sanitize sorts accounts.
func (gs GenesisState) Sanitize() {
	sort.Slice(gs, func(i, j int) bool {
		return gs[i].AccountNumber < gs[j].AccountNumber
	})
}

// ValidateGenesis performs validation of genesis accounts. It
// ensures that there are no duplicate accounts in the genesis state.
func ValidateGenesis(genesisState GenesisState) error {
	addrMap := make(map[string]bool, len(genesisState))
	for _, acc := range genesisState {
		addrStr := acc.Address.String()

		// disallow any duplicate accounts
		if _, ok := addrMap[addrStr]; ok {
			return fmt.Errorf("duplicate account found in genesis state; address: %s", addrStr)
		}

		addrMap[addrStr] = true
	}

	return nil
}
