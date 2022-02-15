package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func CmdListApprovedUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-approved-upgrade",
		Short: "list all ApprovedUpgrade",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllApprovedUpgradeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ApprovedUpgradeAll(context.Background(), params)
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

func CmdShowApprovedUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-approved-upgrade [name]",
		Short: "shows a ApprovedUpgrade",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetApprovedUpgradeRequest{
				Name: argName,
			}

			res, err := queryClient.ApprovedUpgrade(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
