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

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
)

// ValidateAccountInGenesis checks that the provided key has sufficient
// coins in the genesis accounts.
func ValidateAccountInGenesis(appGenesisState map[string]json.RawMessage,
	key sdk.Address, cdc *codec.Codec) error {
	accountIsInGenesis := false

	validatorDataBz := appGenesisState[validator.ModuleName]

	var validatorData validator.GenesisState

	cdc.MustUnmarshalJSON(validatorDataBz, &validatorData)

	genUtilDataBz := appGenesisState[validator.ModuleName]

	var genesisState GenesisState

	cdc.MustUnmarshalJSON(genUtilDataBz, &genesisState)

	IterateGenesisAccounts(cdc, appGenesisState, func(acc auth.Account) (stop bool) {
		if acc.Address.Equals(key) {
			accountIsInGenesis = true

			return true
		}

		return false
	})

	if !accountIsInGenesis {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Error account %s in not in the app_state.accounts array of genesis.json", key))
	}

	return nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// DeliverGenTxs - deliver a genesis transaction.
func DeliverGenTxs(ctx sdk.Context, cdc *codec.Codec, genTxs []json.RawMessage,
	validatorKeeper types.ValidatorKeeper, deliverTx deliverTxfn) []abci.ValidatorUpdate {
	for _, genTx := range genTxs {
		var tx authtypes.StdTx

		cdc.MustUnmarshalJSON(genTx, &tx)

		bz := cdc.MustMarshalBinaryLengthPrefixed(tx)
		res := deliverTx(abci.RequestDeliverTx{Tx: bz})

		if !res.IsOK() {
			panic(res.Log)
		}
	}

	return validatorKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
}
