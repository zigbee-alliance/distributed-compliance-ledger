package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

func CmdListTestingResults() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-test-results",
		Short: "Get all testing results for all Models",
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
	var (
		vid             int32
		pid             int32
		softwareVersion uint32
	)

	cmd := &cobra.Command{
		Use:   "test-result",
		Short: "Query testing results for Model (identified by the `vid`, `pid` and `softwareVersion`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTestingResultsRequest{
				Vid:             vid,
				Pid:             pid,
				SoftwareVersion: softwareVersion,
			}

			res, err := queryClient.TestingResults(context.Background(), params)
			if cli.HandleError(err) != nil {
				return err
			}
			if err != nil {
				// show default (empty) value in CLI
				res = &types.QueryGetTestingResultsResponse{TestingResults: nil}
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID")
	cmd.Flags().Uint32Var(&softwareVersion, FlagSoftwareVersion, 0,
		"Software Version of model (uint32)")

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
