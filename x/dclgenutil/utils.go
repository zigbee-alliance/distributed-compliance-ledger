package dclgenutil

import (
	cosmosgenutil "github.com/cosmos/cosmos-sdk/x/genutil"
)

var (
	// functions aliases.
	ExportGenesisFile                        = cosmosgenutil.ExportGenesisFile
	ExportGenesisFileWithTime                = cosmosgenutil.ExportGenesisFileWithTime
	NewInitConfig                            = cosmosgenutil.NewInitConfig
	InitializeNodeValidatorFiles             = cosmosgenutil.InitializeNodeValidatorFiles
	InitializeNodeValidatorFilesFromMnemonic = cosmosgenutil.InitializeNodeValidatorFilesFromMnemonic
)

type (
	InitConfig = cosmosgenutil.InitConfig
)
