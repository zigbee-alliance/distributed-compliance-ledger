package authz

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	AccountRoles []AccountRoles `json:"account_roles"`
}

func NewGenesisState() GenesisState {
	return GenesisState{AccountRoles: []AccountRoles{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.AccountRoles {
		if record.Address == nil {
			return fmt.Errorf("invalid AccountRoles: Value: %s. Error: Missing Address", record.Address)
		}

		if record.Roles == nil {
			return fmt.Errorf("invalid AccountRoles: Value: %s. Error: Missing Roles", record.Roles)
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.AccountRoles {
		keeper.SetAccountRoles(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []AccountRoles

	k.IterateAccountRoles(ctx, func(accountRoles types.AccountRoles) (stop bool) {
		records = append(records, accountRoles)
		return false
	})

	return GenesisState{AccountRoles: records}
}
