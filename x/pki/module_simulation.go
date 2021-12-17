package pki

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkisimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = pkisimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeAddX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeAddX509RootCert int = 100

	opWeightMsgApproveAddX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveAddX509RootCert int = 100

	opWeightMsgAddX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddX509Cert int = 100

	opWeightMsgProposeRevokeX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeRevokeX509RootCert int = 100

	opWeightMsgApproveRevokeX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveRevokeX509RootCert int = 100

	opWeightMsgRevokeX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevokeX509Cert int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	pkiGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&pkiGenesis)
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

	var weightMsgProposeAddX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeAddX509RootCert, &weightMsgProposeAddX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgProposeAddX509RootCert = defaultWeightMsgProposeAddX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeAddX509RootCert,
		pkisimulation.SimulateMsgProposeAddX509RootCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveAddX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveAddX509RootCert, &weightMsgApproveAddX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgApproveAddX509RootCert = defaultWeightMsgApproveAddX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveAddX509RootCert,
		pkisimulation.SimulateMsgApproveAddX509RootCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddX509Cert, &weightMsgAddX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgAddX509Cert = defaultWeightMsgAddX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddX509Cert,
		pkisimulation.SimulateMsgAddX509Cert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProposeRevokeX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeRevokeX509RootCert, &weightMsgProposeRevokeX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgProposeRevokeX509RootCert = defaultWeightMsgProposeRevokeX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeRevokeX509RootCert,
		pkisimulation.SimulateMsgProposeRevokeX509RootCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveRevokeX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveRevokeX509RootCert, &weightMsgApproveRevokeX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRevokeX509RootCert = defaultWeightMsgApproveRevokeX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRevokeX509RootCert,
		pkisimulation.SimulateMsgApproveRevokeX509RootCert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevokeX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeX509Cert, &weightMsgRevokeX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeX509Cert = defaultWeightMsgRevokeX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeX509Cert,
		pkisimulation.SimulateMsgRevokeX509Cert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
