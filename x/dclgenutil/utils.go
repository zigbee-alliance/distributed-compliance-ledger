// Copyright 2022 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
