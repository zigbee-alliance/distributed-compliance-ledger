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
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil/types"
)

const (
	FlagAddress = "address"
	FlagPubKey  = "pubkey"
	FlagRoles   = "roles"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
//nolint:funlen
func AddGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec,
	defaultNodeHome, defaultClientHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account",
		Short: "Add genesis account to genesis.json",
		Args:  cobra.ExactArgs(0),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				kb, err := keys.NewKeyBaseFromDir(viper.GetString(flagClientHome))
				if err != nil {
					return err
				}

				info, err := kb.Get(args[0])
				if err != nil {
					return err
				}

				addr = info.GetAddress()
			}

			pubkey, err := sdk.GetAccPubKeyBech32(viper.GetString(FlagPubKey))
			if err != nil {
				return err
			}

			var roles auth.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, auth.AccountRole(role))
				}
			}
			// Passing the VendorId as zero for Genesis accounts
			account := auth.NewAccount(addr, pubkey, roles, 0)
			if err := account.Validate(); err != nil {
				return err
			}

			// retrieve the app state
			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return err
			}

			// add genesis account to the app state
			var genesisState types.GenesisState

			cdc.MustUnmarshalJSON(appState[genutil.ModuleName], &genesisState)

			if genesisState.Accounts.Contains(addr) {
				return sdk.ErrUnknownRequest(fmt.Sprintf("cannot add account at existing address %v", addr))
			}

			genesisState.Accounts = append(genesisState.Accounts, account)

			genesisStateBz := cdc.MustMarshalJSON(genesisState)
			appState[genutil.ModuleName] = genesisStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return err
			}

			// export app state
			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "Bench32 encoded account public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles (split by comma) to assign to account (supported roles: %v)", auth.Roles))
	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
