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

package compliance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the complianceInfo
	for _, elem := range genState.ComplianceInfoList {
		k.SetComplianceInfo(ctx, elem)
	}
	// Set all the certifiedModel
	for _, elem := range genState.CertifiedModelList {
		k.SetCertifiedModel(ctx, elem)
	}
	// Set all the revokedModel
	for _, elem := range genState.RevokedModelList {
		k.SetRevokedModel(ctx, elem)
	}
	// Set all the provisionalModel
	for _, elem := range genState.ProvisionalModelList {
		k.SetProvisionalModel(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ComplianceInfoList = k.GetAllComplianceInfo(ctx)
	genesis.CertifiedModelList = k.GetAllCertifiedModel(ctx)
	genesis.RevokedModelList = k.GetAllRevokedModel(ctx)
	genesis.ProvisionalModelList = k.GetAllProvisionalModel(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
