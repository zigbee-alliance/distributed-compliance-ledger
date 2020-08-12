package auth

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Accounts        []Account              `json:"accounts"`
	PendingAccounts []types.PendingAccount `json:"pending_accounts"`
}

func NewGenesisState() GenesisState {
	return GenesisState{Accounts: []Account{}}
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Accounts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, record := range data.PendingAccounts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Accounts {
		keeper.SetAccount(ctx, record)
	}

	for _, record := range data.PendingAccounts {
		keeper.SetPendingAccount(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var accounts []Account
	var pendingAccounts []PendingAccount

	k.IterateAccounts(ctx, func(account types.Account) (stop bool) {
		accounts = append(accounts, account)
		return false
	})

	k.IteratePendingAccounts(ctx, func(account types.PendingAccount) (stop bool) {
		pendingAccounts = append(pendingAccounts, account)
		return false
	})

	return GenesisState{Accounts: accounts, PendingAccounts: pendingAccounts}
}
