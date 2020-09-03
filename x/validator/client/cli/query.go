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

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	validatorQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the validator module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	validatorQueryCmd.AddCommand(client.GetCommands(
		GetCmdQueryValidator(queryRoute, cdc),
		GetCmdQueryValidators(queryRoute, cdc))...)

	return validatorQueryCmd
}

// GetCmdQueryValidator implements the node query command.
func GetCmdQueryValidator(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Query a validator node",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.ConsAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			res, height, err := cliCtx.QueryStore(types.GetValidatorKey(addr), storeName)
			if err != nil || res == nil {
				return types.ErrValidatorDoesNotExist(addr)
			}

			var certificate types.Validator
			cdc.MustUnmarshalBinaryBare(res, &certificate)

			return cliCtx.EncodeAndPrintWithHeight(certificate, height)
		},
	}

	cmd.Flags().String(FlagAddress, "", "The Bech32 encoded Address of the validator")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

// GetCmdQueryValidators implements the query all nodes command.
func GetCmdQueryValidators(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-nodes",
		Short: "Query for all validators",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			state := viper.GetString(FlagState)

			paginationParams := pagination.ParsePaginationParamsFromFlags()
			params := types.NewListValidatorsParams(paginationParams, types.ValidatorState(state))

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/validators", storeName), params)
		},
	}

	cmd.Flags().String(FlagState, "", "state of a validator (active/jailed)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of validators to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of validators to take")

	return cmd
}
