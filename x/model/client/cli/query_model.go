package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdListModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-models",
		Short: "Query the list of all Models",
		Args:  cobra.ExactArgs(0),
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

			params := &types.QueryAllModelRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ModelAll(context.Background(), params)
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

func CmdShowModel() *cobra.Command {
	var (
		vid int32
		pid int32
	)

	cmd := &cobra.Command{
		Use:   "get-model",
		Short: "Query Model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.Model
			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ModelKeyPrefix,
				types.ModelKey(vid, pid),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID")

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
