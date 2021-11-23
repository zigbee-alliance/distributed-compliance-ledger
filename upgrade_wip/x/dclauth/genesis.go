package dclauth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, elem)
	}
	// Set all the pendingAccount
	for _, elem := range genState.PendingAccountList {
		k.SetPendingAccount(ctx, elem)
	}
	// Set all the pendingAccountRevocation
	for _, elem := range genState.PendingAccountRevocationList {
		k.SetPendingAccountRevocation(ctx, elem)
	}
	// Set if defined
	if genState.AccountStat != nil {
		k.SetAccountStat(ctx, *genState.AccountStat)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.AccountList = k.GetAllAccount(ctx)
	genesis.PendingAccountList = k.GetAllPendingAccount(ctx)
	genesis.PendingAccountRevocationList = k.GetAllPendingAccountRevocation(ctx)
	// Get all accountStat
	accountStat, found := k.GetAccountStat(ctx)
	if found {
		genesis.AccountStat = &accountStat
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
