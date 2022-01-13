// Copyright 2022 DSR Corporation
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

package dclauth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the account
	for _, elem := range genState.AccountList {
		// TODO issue 99: error processing
		_ = elem.SetAccountNumber(k.GetNextAccountNumber(ctx))
		k.SetAccountO(ctx, elem)
	}
	// Set all the pendingAccount
	for _, elem := range genState.PendingAccountList {
		k.SetPendingAccount(ctx, elem)
	}
	// Set all the pendingAccountRevocation
	for _, elem := range genState.PendingAccountRevocationList {
		k.SetPendingAccountRevocation(ctx, elem)
	}
	// Set if defined
	if genState.AccountStat != nil {
		k.SetAccountStat(ctx, *genState.AccountStat)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.AccountList = k.GetAllAccount(ctx)
	genesis.PendingAccountList = k.GetAllPendingAccount(ctx)
	genesis.PendingAccountRevocationList = k.GetAllPendingAccountRevocation(ctx)
	// Get all accountStat
	accountStat, found := k.GetAccountStat(ctx)
	if found {
		genesis.AccountStat = &accountStat
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
