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
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	dclupgradesimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// avoid unused import issue.
var (
	_ = sample.AccAddress
	_ = dclupgradesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeUpgrade = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgProposeUpgrade int = 100

	opWeightMsgApproveUpgrade = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgApproveUpgrade int = 100

	opWeightMsgRejectUpgrade = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgRejectUpgrade int = 100

	// this line is used by starport scaffolding # simapp/module/const.
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dclupgradeGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState.
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dclupgradeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent { //nolint:staticcheck
	return nil
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgProposeUpgrade int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeUpgrade, &weightMsgProposeUpgrade, nil,
		func(_ *rand.Rand) {
			weightMsgProposeUpgrade = defaultWeightMsgProposeUpgrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeUpgrade,
		dclupgradesimulation.SimulateMsgProposeUpgrade(am.keeper),
	))

	var weightMsgApproveUpgrade int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveUpgrade, &weightMsgApproveUpgrade, nil,
		func(_ *rand.Rand) {
			weightMsgApproveUpgrade = defaultWeightMsgApproveUpgrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveUpgrade,
		dclupgradesimulation.SimulateMsgApproveUpgrade(am.keeper),
	))

	var weightMsgRejectUpgrade int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRejectUpgrade, &weightMsgRejectUpgrade, nil,
		func(_ *rand.Rand) {
			weightMsgRejectUpgrade = defaultWeightMsgRejectUpgrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRejectUpgrade,
		dclupgradesimulation.SimulateMsgRejectUpgrade(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
