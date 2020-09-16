// Copyright 2020 DSR Corporation
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

package genutil

import (
	cosmosgenutil "github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil/types"
)

const (
	ModuleName = types.ModuleName
)

var (
	// functions aliases.
	InitializeNodeValidatorFiles = cosmosgenutil.InitializeNodeValidatorFiles
	ExportGenesisFile            = cosmosgenutil.ExportGenesisFile
	NewInitConfig                = cosmosgenutil.NewInitConfig
	GenesisStateFromGenDoc       = types.GenesisStateFromGenDoc
	GenesisStateFromGenFile      = types.GenesisStateFromGenFile

	// variable aliases.
	ModuleCdc = types.ModuleCdc
)

type (
	GenesisState = types.GenesisState
	InitConfig   = cosmosgenutil.InitConfig
)
