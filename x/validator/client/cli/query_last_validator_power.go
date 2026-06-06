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
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListLastValidatorPower() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-last-powers",
		Short: "list all LastValidatorPower",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLastValidatorPowerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LastValidatorPowerAll(context.Background(), params)
			if cli.IsKeyNotFoundRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLastValidatorPower() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last-power",
		Short: "shows a LastValidatorPower",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.ValAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				owner, err2 := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
				if err2 != nil {
					return err2
				}
				addr = sdk.ValAddress(owner)
			}

			var res types.LastValidatorPower

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.LastValidatorPowerKeyPrefix,
				types.LastValidatorPowerKey(addr),
				&res,
			)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded validator address or owner account")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
