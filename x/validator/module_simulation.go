package validator

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	validatorsimulation "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// avoid unused import issue.
var (
	_ = sample.AccAddress
	_ = validatorsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeDisableValidator = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgProposeDisableValidator int = 100

	opWeightMsgApproveDisableValidator = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgApproveDisableValidator int = 100

	opWeightMsgDisableValidator = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgDisableValidator int = 100

	opWeightMsgEnableValidator = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value.
	defaultWeightMsgEnableValidator int = 100

	// this line is used by starport scaffolding # simapp/module/const.
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	validatorGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&validatorGenesis)
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

	var weightMsgProposeDisableValidator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeDisableValidator, &weightMsgProposeDisableValidator, nil,
		func(_ *rand.Rand) {
			weightMsgProposeDisableValidator = defaultWeightMsgProposeDisableValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeDisableValidator,
		validatorsimulation.SimulateMsgProposeDisableValidator(am.keeper),
	))

	var weightMsgApproveDisableValidator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveDisableValidator, &weightMsgApproveDisableValidator, nil,
		func(_ *rand.Rand) {
			weightMsgApproveDisableValidator = defaultWeightMsgApproveDisableValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveDisableValidator,
		validatorsimulation.SimulateMsgApproveDisableValidator(am.keeper),
	))

	var weightMsgDisableValidator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDisableValidator, &weightMsgDisableValidator, nil,
		func(_ *rand.Rand) {
			weightMsgDisableValidator = defaultWeightMsgDisableValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisableValidator,
		validatorsimulation.SimulateMsgDisableValidator(am.keeper),
	))

	var weightMsgEnableValidator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnableValidator, &weightMsgEnableValidator, nil,
		func(_ *rand.Rand) {
			weightMsgEnableValidator = defaultWeightMsgEnableValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnableValidator,
		validatorsimulation.SimulateMsgEnableValidator(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
