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
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func CmdListRejectedUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-rejected-upgrades",
		Short: "Query the list of all rejected upgrades",
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

			params := &types.QueryAllRejectedUpgradeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RejectedUpgradeAll(context.Background(), params)
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

func CmdShowRejectedUpgrade() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "rejected-upgrade --name [name]",
		Short: "Query rejected upgrade by name",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var res types.RejectedUpgrade

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.RejectedUpgradeKeyPrefix,
				types.RejectedUpgradeKey(name),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&name, FlagName, "", "Upgrade name")

	_ = cmd.MarkFlagRequired(FlagName)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
