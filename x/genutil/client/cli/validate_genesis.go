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

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Validate genesis command takes.
func ValidateGenesisCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager) *cobra.Command {
	return &cobra.Command{
		Use:   "validate-genesis [file]",
		Args:  cobra.RangeArgs(0, 1),
		Short: "validates the genesis file at the default location or at the location passed as an arg",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Load default if passed no args, otherwise load passed file
			var genesis string
			if len(args) == 0 {
				genesis = ctx.Config.GenesisFile()
			} else {
				genesis = args[0]
			}

			fmt.Fprintf(os.Stderr, "Validating genesis file at %s\n", genesis)

			var genDoc *tmtypes.GenesisDoc
			if genDoc, err = tmtypes.GenesisDocFromFile(genesis); err != nil {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Error loading genesis doc from %s: %s", genesis, err.Error()))
			}

			var genState map[string]json.RawMessage
			if err = cdc.UnmarshalJSON(genDoc.AppState, &genState); err != nil {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Error unmarshaling genesis doc %s: %s", genesis, err.Error()))
			}

			if err = mbm.ValidateGenesis(genState); err != nil {
				return sdk.ErrUnknownRequest(
					fmt.Sprintf("Error validating genesis file %s: %s", genesis, err.Error()))
			}

			fmt.Printf("File at %s is a valid genesis file\n", genesis)

			return nil
		},
	}
}
