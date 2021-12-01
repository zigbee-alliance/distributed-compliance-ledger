package dclauth

/* FIXME issue 99

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	dclauthsimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dclauthsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeAddAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeAddAccount int = 100

	opWeightMsgApproveAddAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveAddAccount int = 100

	opWeightMsgProposeRevokeAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeRevokeAccount int = 100

	opWeightMsgApproveRevokeAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveRevokeAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dclauthGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dclauthGenesis)
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

	var weightMsgProposeAddAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeAddAccount, &weightMsgProposeAddAccount, nil,
		func(_ *rand.Rand) {
			weightMsgProposeAddAccount = defaultWeightMsgProposeAddAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeAddAccount,
		dclauthsimulation.SimulateMsgProposeAddAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveAddAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveAddAccount, &weightMsgApproveAddAccount, nil,
		func(_ *rand.Rand) {
			weightMsgApproveAddAccount = defaultWeightMsgApproveAddAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveAddAccount,
		dclauthsimulation.SimulateMsgApproveAddAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProposeRevokeAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeRevokeAccount, &weightMsgProposeRevokeAccount, nil,
		func(_ *rand.Rand) {
			weightMsgProposeRevokeAccount = defaultWeightMsgProposeRevokeAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeRevokeAccount,
		dclauthsimulation.SimulateMsgProposeRevokeAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveRevokeAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveRevokeAccount, &weightMsgApproveRevokeAccount, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRevokeAccount = defaultWeightMsgApproveRevokeAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRevokeAccount,
		dclauthsimulation.SimulateMsgApproveRevokeAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
*/
