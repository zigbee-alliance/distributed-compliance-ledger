// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package genutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	cfg "github.com/tendermint/tendermint/config"
	tmtypes "github.com/tendermint/tendermint/types"
)

// GenAppStateFromConfig gets the genesis app state from the config.
func GenAppStateFromConfig(cdc *codec.Codec, config *cfg.Config,
	initCfg InitConfig, genDoc tmtypes.GenesisDoc,
) (appState json.RawMessage, err error) {
	// process genesis transactions, else create default genesis.json.
	appGenTxs, persistentPeers, err := CollectStdTxs(
		cdc, config.Moniker, initCfg.GenTxsDir, genDoc)
	if err != nil {
		return appState, err
	}

	config.P2P.PersistentPeers = persistentPeers
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	// if there are no gen txs to be processed, return the default empty state.
	if len(appGenTxs) == 0 {
		return appState, sdk.ErrUnknownRequest("there must be at least one genesis tx")
	}

	// create the app state.
	appGenesisState, err := GenesisStateFromGenDoc(cdc, genDoc)
	if err != nil {
		return appState, err
	}

	appGenesisState, err = SetGenTxsInAppGenesisState(cdc, appGenesisState, appGenTxs)
	if err != nil {
		return appState, err
	}

	appState, err = codec.MarshalJSONIndent(cdc, appGenesisState)

	if err != nil {
		return appState, err
	}

	genDoc.AppState = appState
	err = ExportGenesisFile(&genDoc, config.GenesisFile())

	return appState, err
}

// CollectStdTxs processes and validates application's genesis StdTxs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
//nolint:funlen
func CollectStdTxs(cdc *codec.Codec, name, genTxsDir string,
	genDoc tmtypes.GenesisDoc,
) (appGenTxs []authtypes.StdTx, persistentPeers string, err error) {
	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)

	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	// prepare a map of all accounts in genesis state to then validate
	// against the validators addresses.
	var appState map[string]json.RawMessage
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]auth.Account)

	IterateGenesisAccounts(cdc, appState,
		func(acc auth.Account) (stop bool) {
			addrMap[acc.Address.String()] = acc

			return false
		},
	)

	// addresses and IPs (and port) validator server info.
	var addressesIPs []string

	for _, fo := range fos {
		filename := filepath.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (filepath.Ext(filename) != ".json") {
			continue
		}

		// get the genStdTx.
		var jsonRawTx []byte

		if jsonRawTx, err = ioutil.ReadFile(filename); err != nil {
			return appGenTxs, persistentPeers, err
		}

		var genStdTx authtypes.StdTx

		if err = cdc.UnmarshalJSON(jsonRawTx, &genStdTx); err != nil {
			return appGenTxs, persistentPeers, err
		}

		appGenTxs = append(appGenTxs, genStdTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656".
		nodeAddrIP := genStdTx.GetMemo()
		if len(nodeAddrIP) == 0 {
			return appGenTxs, persistentPeers, sdk.ErrUnknownRequest(
				fmt.Sprintf("couldn't find node's address and IP in %s", fo.Name()))
		}

		// genesis transactions must be single-message.
		msgs := genStdTx.GetMsgs()
		if len(msgs) != 1 {
			return appGenTxs, persistentPeers, sdk.ErrUnknownRequest(
				"each genesis transaction must provide a single genesis message")
		}

		msg := msgs[0].(validator.MsgCreateValidator)
		account := msg.Signer.String()

		_, valOk := addrMap[account]
		if !valOk {
			return appGenTxs, persistentPeers, sdk.ErrUnknownRequest(
				fmt.Sprintf("Error account %v not in genesis.json: %+v", account, addrMap))
		}

		// exclude itself from persistent peers.
		if msg.Description.Name != name {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}

// SetGenTxsInAppGenesisState - sets the genesis transactions in the app genesis state.
func SetGenTxsInAppGenesisState(cdc *codec.Codec, appGenesisState map[string]json.RawMessage,
	genTxs []authtypes.StdTx) (map[string]json.RawMessage, error) {
	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesisState)
	// convert all the GenTxs to JSON
	genTxsBz := make([]json.RawMessage, len(genTxs))

	for i, genTx := range genTxs {
		txBz, err := cdc.MarshalJSON(genTx)
		if err != nil {
			return appGenesisState, err
		}

		genTxsBz[i] = txBz
	}

	genesisState.GenTxs = genTxsBz

	return SetGenesisStateInAppState(cdc, appGenesisState, genesisState), nil
}

// SetGenesisStateInAppState sets the genutil genesis state within the expected app state.
func SetGenesisStateInAppState(cdc *codec.Codec,
	appState map[string]json.RawMessage, genesisState GenesisState) map[string]json.RawMessage {
	genesisStateBz := cdc.MustMarshalJSON(genesisState)
	appState[ModuleName] = genesisStateBz

	return appState
}
