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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	pkisimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// avoid unused import issue.
var (
	_ = sample.AccAddress
	_ = pkisimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeAddX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgProposeAddX509RootCert int = 100

	opWeightMsgApproveAddX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgApproveAddX509RootCert int = 100

	opWeightMsgAddX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAddX509Cert int = 100

	opWeightMsgProposeRevokeX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgProposeRevokeX509RootCert int = 100

	opWeightMsgApproveRevokeX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgApproveRevokeX509RootCert int = 100

	opWeightMsgRevokeX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgRevokeX509Cert int = 100

	opWeightMsgRejectAddX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgRejectAddX509RootCert int = 100

	opWeightMsgAddPkiRevocationDistributionPoint = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAddPkiRevocationDistributionPoint int = 100

	opWeightMsgUpdatePkiRevocationDistributionPoint = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgUpdatePkiRevocationDistributionPoint int = 100

	opWeightMsgDeletePkiRevocationDistributionPoint = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgDeletePkiRevocationDistributionPoint int = 100

	opWeightMsgAssignVid = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAssignVid int = 100

	opWeightMsgAddNocX509RootCert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAddNocX509RootCert int = 100

	opWeightMsgRemoveX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgRemoveX509Cert int = 100

	opWeightMsgAddNocX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgAddNocX509Cert int = 100

	opWeightMsgRevokeNocRootX509Cert = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgRevokeNocRootX509Cert int = 100

	// this line is used by starport scaffolding # simapp/module/const.
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	pkiGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[pkitypes.ModuleName] = simState.Cdc.MustMarshalJSON(&pkiGenesis)
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

	var weightMsgProposeAddX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeAddX509RootCert, &weightMsgProposeAddX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgProposeAddX509RootCert = defaultWeightMsgProposeAddX509RootCert
		},
	)

	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeAddX509RootCert,
		pkisimulation.SimulateMsgProposeAddX509RootCert(am.keeper),
	))

	var weightMsgApproveAddX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveAddX509RootCert, &weightMsgApproveAddX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgApproveAddX509RootCert = defaultWeightMsgApproveAddX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveAddX509RootCert,
		pkisimulation.SimulateMsgApproveAddX509RootCert(am.keeper),
	))

	var weightMsgAddX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddX509Cert, &weightMsgAddX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgAddX509Cert = defaultWeightMsgAddX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddX509Cert,
		pkisimulation.SimulateMsgAddX509Cert(am.keeper),
	))

	var weightMsgProposeRevokeX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeRevokeX509RootCert, &weightMsgProposeRevokeX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgProposeRevokeX509RootCert = defaultWeightMsgProposeRevokeX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeRevokeX509RootCert,
		pkisimulation.SimulateMsgProposeRevokeX509RootCert(am.keeper),
	))

	var weightMsgApproveRevokeX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveRevokeX509RootCert, &weightMsgApproveRevokeX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRevokeX509RootCert = defaultWeightMsgApproveRevokeX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRevokeX509RootCert,
		pkisimulation.SimulateMsgApproveRevokeX509RootCert(am.keeper),
	))

	var weightMsgRevokeX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeX509Cert, &weightMsgRevokeX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeX509Cert = defaultWeightMsgRevokeX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeX509Cert,
		pkisimulation.SimulateMsgRevokeX509Cert(am.keeper),
	))

	var weightMsgRejectAddX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRejectAddX509RootCert, &weightMsgRejectAddX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgRejectAddX509RootCert = defaultWeightMsgRejectAddX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRejectAddX509RootCert,
		pkisimulation.SimulateMsgRejectAddX509RootCert(am.keeper),
	))

	var weightMsgAddPkiRevocationDistributionPoint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddPkiRevocationDistributionPoint, &weightMsgAddPkiRevocationDistributionPoint, nil,
		func(_ *rand.Rand) {
			weightMsgAddPkiRevocationDistributionPoint = defaultWeightMsgAddPkiRevocationDistributionPoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddPkiRevocationDistributionPoint,
		pkisimulation.SimulateMsgAddPkiRevocationDistributionPoint(am.keeper),
	))

	var weightMsgUpdatePkiRevocationDistributionPoint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePkiRevocationDistributionPoint, &weightMsgUpdatePkiRevocationDistributionPoint, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePkiRevocationDistributionPoint = defaultWeightMsgUpdatePkiRevocationDistributionPoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePkiRevocationDistributionPoint,
		pkisimulation.SimulateMsgUpdatePkiRevocationDistributionPoint(am.keeper),
	))

	var weightMsgDeletePkiRevocationDistributionPoint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePkiRevocationDistributionPoint, &weightMsgDeletePkiRevocationDistributionPoint, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePkiRevocationDistributionPoint = defaultWeightMsgDeletePkiRevocationDistributionPoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePkiRevocationDistributionPoint,
		pkisimulation.SimulateMsgDeletePkiRevocationDistributionPoint(am.keeper),
	))

	var weightMsgAssignVid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAssignVid, &weightMsgAssignVid, nil,
		func(_ *rand.Rand) {
			weightMsgAssignVid = defaultWeightMsgAssignVid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAssignVid,
		pkisimulation.SimulateMsgAssignVid(am.keeper),
	))

	var weightMsgAddNocX509RootCert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddNocX509RootCert, &weightMsgAddNocX509RootCert, nil,
		func(_ *rand.Rand) {
			weightMsgAddNocX509RootCert = defaultWeightMsgAddNocX509RootCert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddNocX509RootCert,
		pkisimulation.SimulateMsgAddNocX509RootCert(am.keeper),
	))

	var weightMsgRemoveX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveX509Cert, &weightMsgRemoveX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveX509Cert = defaultWeightMsgRemoveX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveX509Cert,
		pkisimulation.SimulateMsgRemoveX509Cert(am.keeper),
	))

	var weightMsgAddNocX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddNocX509Cert, &weightMsgAddNocX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgAddNocX509Cert = defaultWeightMsgAddNocX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddNocX509Cert,
		pkisimulation.SimulateMsgAddNocX509Cert(am.keeper),
	))

	var weightMsgRevokeNocRootX509Cert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeNocRootX509Cert, &weightMsgRevokeNocRootX509Cert, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeNocRootX509Cert = defaultWeightMsgRevokeNocRootX509Cert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeNocRootX509Cert,
		pkisimulation.SimulateMsgRevokeNocRootX509Cert(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
