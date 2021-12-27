package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

func CmdListTestingResults() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-testing-results",
		Short: "list all TestingResults",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTestingResultsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TestingResultsAll(context.Background(), params)
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

func CmdShowTestingResults() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-testing-results [vid] [pid] [software-version]",
		Short: "shows a TestingResults",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint32E(args[2])
			if err != nil {
				return err
			}

			params := &types.QueryGetTestingResultsRequest{
				Vid:             argVid,
				Pid:             argPid,
				SoftwareVersion: argSoftwareVersion,
			}

			res, err := queryClient.TestingResults(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
