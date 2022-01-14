package dclgenutil

import (
	cosmosgenutil "github.com/cosmos/cosmos-sdk/x/genutil"
	cosmosgenutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

var (
	// functions aliases.
	ExportGenesisFile                        = cosmosgenutil.ExportGenesisFile
	ExportGenesisFileWithTime                = cosmosgenutil.ExportGenesisFileWithTime
	NewInitConfig                            = cosmosgenutiltypes.NewInitConfig
	InitializeNodeValidatorFiles             = cosmosgenutil.InitializeNodeValidatorFiles
	InitializeNodeValidatorFilesFromMnemonic = cosmosgenutil.InitializeNodeValidatorFilesFromMnemonic
)

type (
	InitConfig = cosmosgenutiltypes.InitConfig
)
