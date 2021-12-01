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

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions.
func InitGenesis(ctx sdk.Context, cdc *codec.Codec, authKeeper types.AuthKeeper, validatorKeeper types.ValidatorKeeper,
	deliverTx deliverTxfn, genesisState GenesisState) []abci.ValidatorUpdate {
	// load the accounts
	for _, acc := range genesisState.Accounts {
		err := acc.SetAccountNumber(authKeeper.GetNextAccountNumber(ctx))
		if err != nil {
			panic(err)
		}

		authKeeper.SetAccount(ctx, acc)
	}

	// deliver validator transactions
	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, validatorKeeper, deliverTx)
	}

	return validators
}

func IterateGenesisAccounts(cdc *codec.Codec, appGenesis map[string]json.RawMessage,
	iterateFn func(auth.Account) (stop bool)) {
	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesis)
	for _, acc := range genesisState.Accounts {
		if iterateFn(acc) {
			break
		}
	}
}
