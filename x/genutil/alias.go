package genutil

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	cosmosgenutil "github.com/cosmos/cosmos-sdk/x/genutil"
)

const (
	ModuleName = types.ModuleName
)

var (
	// function aliases
	InitializeNodeValidatorFiles = cosmosgenutil.InitializeNodeValidatorFiles
	ExportGenesisFile            = cosmosgenutil.ExportGenesisFile
	NewInitConfig                = cosmosgenutil.NewInitConfig
	GenesisStateFromGenDoc       = types.GenesisStateFromGenDoc
	GenesisStateFromGenFile      = types.GenesisStateFromGenFile

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	GenesisState = types.GenesisState
	InitConfig   = cosmosgenutil.InitConfig
)
