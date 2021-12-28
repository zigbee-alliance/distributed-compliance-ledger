package compliancetest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the testingResults
	for _, elem := range genState.TestingResultsList {
		k.SetTestingResults(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.TestingResultsList = k.GetAllTestingResults(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
