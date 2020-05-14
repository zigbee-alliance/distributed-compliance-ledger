package auth

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Accounts []Account `json:"accounts"`
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

	return nil
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Accounts {
		keeper.SetAccount(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Account

	k.IterateAccounts(ctx, func(account types.Account) (stop bool) {
		records = append(records, account)
		return false
	})

	return GenesisState{Accounts: records}
}
