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

package dclupgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the proposedUpgrade
	for _, elem := range genState.ProposedUpgradeList {
		k.SetProposedUpgrade(ctx, elem)
	}
	// Set all the approvedUpgrade
	for _, elem := range genState.ApprovedUpgradeList {
		k.SetApprovedUpgrade(ctx, elem)
	}
	// Set all the rejectedUpgrade
	for _, elem := range genState.RejectedUpgradeList {
		k.SetRejectedUpgrade(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ProposedUpgradeList = k.GetAllProposedUpgrade(ctx)
	genesis.ApprovedUpgradeList = k.GetAllApprovedUpgrade(ctx)
	genesis.RejectedUpgradeList = k.GetAllRejectedUpgrade(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
