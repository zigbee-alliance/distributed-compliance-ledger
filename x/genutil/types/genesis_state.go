package types

//nolint:goimports
import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
)

// GenesisState defines the genesis account and validators.
type GenesisState struct {
	Accounts GenesisAccounts   `json:"accounts"`
	GenTxs   []json.RawMessage `json:"gentxs"`
}

// GetGenesisStateFromAppState gets the genutil genesis state from the expected app state.
func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

// SetGenesisStateInAppState sets the genutil genesis state within the expected app state.
func SetGenesisStateInAppState(cdc *codec.Codec,
	appState map[string]json.RawMessage, genesisState GenesisState) map[string]json.RawMessage {
	genesisStateBz := cdc.MustMarshalJSON(genesisState)
	appState[ModuleName] = genesisStateBz

	return appState
}

// GenesisStateFromGenDoc creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenDoc(cdc *codec.Codec, genDoc tmtypes.GenesisDoc,
) (genesisState map[string]json.RawMessage, err error) {
	if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
		return genesisState, err
	}

	return genesisState, nil
}

// GenesisStateFromGenFile creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenFile(cdc *codec.Codec, genFile string,
) (genesisState map[string]json.RawMessage, genDoc *tmtypes.GenesisDoc, err error) {
	if !common.FileExists(genFile) {
		return genesisState, genDoc, sdk.ErrUnknownRequest(
			fmt.Sprintf("%s does not exist, run `init` first", genFile))
	}

	genDoc, err = tmtypes.GenesisDocFromFile(genFile)

	if err != nil {
		return genesisState, genDoc, err
	}

	genesisState, err = GenesisStateFromGenDoc(cdc, *genDoc)

	return genesisState, genDoc, err
}

// ValidateGenesis performs validation of genesis accounts. It
// ensures that there are no duplicate accounts in the genesis state.
func ValidateGenesis(genesisState GenesisState) error {
	addrMap := make(map[string]bool, len(genesisState.Accounts))

	for _, acc := range genesisState.Accounts {
		addrStr := acc.Address.String()

		// disallow any duplicate accounts
		if _, ok := addrMap[addrStr]; ok {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("duplicate account found in genesis state; address: %s", addrStr))
		}

		addrMap[addrStr] = true
	}

	for i, genTx := range genesisState.GenTxs {
		var tx authtypes.StdTx
		if err := ModuleCdc.UnmarshalJSON(genTx, &tx); err != nil {
			return err
		}

		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return sdk.ErrUnknownRequest("must provide genesis StdTx with exactly 1 CreateValidator message")
		}

		if _, ok := msgs[0].(validator.MsgCreateValidator); !ok {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("genesis transaction %v does not contain a MsgCreateValidator", i))
		}
	}

	return nil
}
