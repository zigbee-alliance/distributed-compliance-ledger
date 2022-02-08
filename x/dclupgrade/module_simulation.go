package dclupgrade

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
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
	_ = simappparams.StakePerAccount
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
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator.
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
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

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
