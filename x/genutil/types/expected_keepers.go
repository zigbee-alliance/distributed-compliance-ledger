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

package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
)

// ValidatorKeeper defines the expected validator keeper.
type ValidatorKeeper interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate)
}

// AccountKeeper defines the expected account keeper.
type AuthKeeper interface {
	GetNextAccountNumber(sdk.Context) uint64
	SetAccount(sdk.Context, auth.Account)
	IterateAccounts(ctx sdk.Context, process func(auth.Account) (stop bool))
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object.
type GenesisAccountsIterator interface {
	IterateGenesisAccounts(
		cdc *codec.Codec,
		appGenesis map[string]json.RawMessage,
		iterateFn func(auth.Account) (stop bool),
	)
}
