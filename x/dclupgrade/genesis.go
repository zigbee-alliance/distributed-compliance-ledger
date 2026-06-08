package dclupgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the proposedUpgrade
	for _, elem := range genState.ProposedUpgradeList {
		k.SetProposedUpgrade(ctx, elem)
	}
	// Set all the approvedUpgrade
	for _, elem := range genState.ApprovedUpgradeList {
		k.SetApprovedUpgrade(ctx, elem)
	}
	// Set all the rejectedUpgrade
	for _, elem := range genState.RejectedUpgradeList {
		k.SetRejectedUpgrade(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ProposedUpgradeList = k.GetAllProposedUpgrade(ctx)
	genesis.ApprovedUpgradeList = k.GetAllApprovedUpgrade(ctx)
	genesis.RejectedUpgradeList = k.GetAllRejectedUpgrade(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
