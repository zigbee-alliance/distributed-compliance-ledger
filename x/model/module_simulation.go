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

package model

/* FIXME issue 110

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	modelsimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = modelsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateModel int = 100

	opWeightMsgUpdateModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateModel int = 100

	opWeightMsgDeleteModel = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteModel int = 100

	opWeightMsgCreateModelVersion = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateModelVersion int = 100

	opWeightMsgUpdateModelVersion = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateModelVersion int = 100

	opWeightMsgDeleteModelVersion = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteModelVersion int = 100

	opWeightMsgCreateModelVersions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateModelVersions int = 100

	opWeightMsgUpdateModelVersions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateModelVersions int = 100

	opWeightMsgDeleteModelVersions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteModelVersions int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	modelGenesis := types.GenesisState{
		ModelList: []types.Model{
			{
				Creator: sample.AccAddress(),
				Vid:     0,
				Pid:     0,
			},
			{
				Creator: sample.AccAddress(),
				Vid:     1,
				Pid:     1,
			},
		},
		ModelVersionList: []types.ModelVersion{
			{
				Creator:         sample.AccAddress(),
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			},
			{
				Creator:         sample.AccAddress(),
				Vid:             1,
				Pid:             1,
				SoftwareVersion: 1,
			},
		},
		ModelVersionsList: []types.ModelVersions{
			{
				Creator: sample.AccAddress(),
				Vid:     0,
				Pid:     0,
			},
			{
				Creator: sample.AccAddress(),
				Vid:     1,
				Pid:     1,
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&modelGenesis)
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

	var weightMsgCreateModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateModel, &weightMsgCreateModel, nil,
		func(_ *rand.Rand) {
			weightMsgCreateModel = defaultWeightMsgCreateModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateModel,
		modelsimulation.SimulateMsgCreateModel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateModel, &weightMsgUpdateModel, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateModel = defaultWeightMsgUpdateModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateModel,
		modelsimulation.SimulateMsgUpdateModel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteModel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteModel, &weightMsgDeleteModel, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteModel = defaultWeightMsgDeleteModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteModel,
		modelsimulation.SimulateMsgDeleteModel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateModelVersion int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateModelVersion, &weightMsgCreateModelVersion, nil,
		func(_ *rand.Rand) {
			weightMsgCreateModelVersion = defaultWeightMsgCreateModelVersion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateModelVersion,
		modelsimulation.SimulateMsgCreateModelVersion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateModelVersion int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateModelVersion, &weightMsgUpdateModelVersion, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateModelVersion = defaultWeightMsgUpdateModelVersion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateModelVersion,
		modelsimulation.SimulateMsgUpdateModelVersion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteModelVersion int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteModelVersion, &weightMsgDeleteModelVersion, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteModelVersion = defaultWeightMsgDeleteModelVersion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteModelVersion,
		modelsimulation.SimulateMsgDeleteModelVersion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateModelVersions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateModelVersions, &weightMsgCreateModelVersions, nil,
		func(_ *rand.Rand) {
			weightMsgCreateModelVersions = defaultWeightMsgCreateModelVersions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateModelVersions,
		modelsimulation.SimulateMsgCreateModelVersions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateModelVersions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateModelVersions, &weightMsgUpdateModelVersions, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateModelVersions = defaultWeightMsgUpdateModelVersions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateModelVersions,
		modelsimulation.SimulateMsgUpdateModelVersions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteModelVersions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteModelVersions, &weightMsgDeleteModelVersions, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteModelVersions = defaultWeightMsgDeleteModelVersions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteModelVersions,
		modelsimulation.SimulateMsgDeleteModelVersions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
*/
