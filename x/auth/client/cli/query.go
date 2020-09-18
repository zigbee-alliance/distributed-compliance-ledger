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

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authorization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	authQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccount(storeKey, cdc),
		GetCmdAccounts(storeKey, cdc),
		GetCmdProposedAccounts(storeKey, cdc),
		GetCmdAccountsWithProof(storeKey, cdc),
		GetCmdProposedAccountsToRevoke(storeKey, cdc),
	)...)

	return authQueryCmd
}

func GetCmdAccount(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Get account associated with the address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			res, height, err := cliCtx.QueryStore(types.GetAccountKey(address), queryRoute)
			if err != nil || res == nil {
				return types.ErrAccountDoesNotExist(address)
			}

			var account types.Account
			cdc.MustUnmarshalBinaryBare(res, &account)

			// the trick to prevent appending of `type` field by cdc
			out := codec.Cdc.MustMarshalJSON(account)

			return cliCtx.PrintWithHeight(out, height)
		},
	}

	cmd.Flags().String(FlagAddress, "", FlagAddressUsage)
	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

func GetCmdProposedAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-accounts",
		Short: "Get all proposed but not approved accounts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/%s", queryRoute, keeper.QueryAllPendingAccounts), params)
		},
	}

	pagination.AddPaginationParams(cmd)

	return cmd
}

func GetCmdAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-accounts",
		Short: "Get all accounts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/%s", queryRoute, keeper.QueryAllAccounts), params)
		},
	}

	pagination.AddPaginationParams(cmd)

	return cmd
}

func GetCmdAccountsWithProof(storeKey string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-accounts-with-proof",
		Short: "Get all accounts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			valueUnmarshaler := func(bytes []byte) json.RawMessage {
				value := types.Account{}
				cdc.MustUnmarshalBinaryBare(bytes, &value)

				// the trick to prevent appending of `type` field by cdc
				return codec.Cdc.MustMarshalJSON(value)
			}

			return cliCtx.QueryAllWithProof(storeKey, types.AccountPrefix, types.AccountsTotalKey, valueUnmarshaler)
		},
	}

	return cmd
}

func GetCmdProposedAccountsToRevoke(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-accounts-to-revoke",
		Short: "Get all proposed but not approved accounts to be revoked",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/%s", queryRoute, keeper.QueryAllPendingAccountRevocations), params)
		},
	}

	pagination.AddPaginationParams(cmd)

	return cmd
}
