package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func CmdListProposedUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-upgrades",
		Short: "Query the list of all proposed upgrades",
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

			params := &types.QueryAllProposedUpgradeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedUpgradeAll(context.Background(), params)
			if cli.IsKeyNotFoundRpcError(err) {
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

func CmdShowProposedUpgrade() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "proposed-upgrade --name [name]",
		Short: "Query proposed upgrade by name",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.ProposedUpgrade

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ProposedUpgradeKeyPrefix,
				types.ProposedUpgradeKey(name),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&name, FlagName, "", "Upgrade name")

	_ = cmd.MarkFlagRequired(FlagName)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
