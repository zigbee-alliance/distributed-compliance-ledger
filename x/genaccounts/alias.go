package genaccounts

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genaccounts/internal/types"
)

const (
	ModuleName = types.ModuleName
)

var (
	// functions aliases
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	ValidateGenesis             = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	GenesisAccount  = auth.Account
	GenesisAccounts = types.GenesisAccounts
	GenesisState    = types.GenesisState
)
