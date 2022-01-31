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
		Use:   "list-proposed-upgrade",
		Short: "List all proposed upgrades",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposedUpgradeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedUpgradeAll(context.Background(), params)
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
	cmd := &cobra.Command{
		Use:   "show-proposed-upgrade [name]",
		Short: "Show proposed upgrade with given name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			name := args[0]
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

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
