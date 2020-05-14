package genaccounts

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genaccounts/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes accounts and deliver genesis transactions
func InitGenesis(ctx sdk.Context, _ *codec.Codec, authKeeper types.AuthKeeper, genesisState GenesisState) {
	genesisState.Sanitize()

	// load the accounts
	for _, acc := range genesisState {
		acc = authKeeper.NewAccount(ctx, acc) // set account number
		authKeeper.SetAccount(ctx, acc)
	}
}

// ExportGenesis exports genesis for all accounts
func ExportGenesis(ctx sdk.Context, authKeeper types.AuthKeeper) GenesisState {

	// iterate to get the accounts
	var accounts []GenesisAccount

	authKeeper.IterateAccounts(ctx,
		func(account auth.Account) (stop bool) {
			accounts = append(accounts, account)
			return false
		},
	)

	return accounts
}
