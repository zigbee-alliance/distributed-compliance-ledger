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
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	compliancesimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = compliancesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCertifyModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCertifyModel int = 100

	opWeightMsgRevokeModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevokeModel int = 100

	opWeightMsgProvisionModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProvisionModel int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	complianceGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&complianceGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCertifyModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCertifyModel, &weightMsgCertifyModel, nil,
		func(_ *rand.Rand) {
			weightMsgCertifyModel = defaultWeightMsgCertifyModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCertifyModel,
		compliancesimulation.SimulateMsgCertifyModel(am.keeper),
	))

	var weightMsgRevokeModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeModel, &weightMsgRevokeModel, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeModel = defaultWeightMsgRevokeModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeModel,
		compliancesimulation.SimulateMsgRevokeModel(am.keeper),
	))

	var weightMsgProvisionModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProvisionModel, &weightMsgProvisionModel, nil,
		func(_ *rand.Rand) {
			weightMsgProvisionModel = defaultWeightMsgProvisionModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProvisionModel,
		compliancesimulation.SimulateMsgProvisionModel(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
