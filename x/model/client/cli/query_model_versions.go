package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdShowModelVersions() *cobra.Command {
	var (
		vid int32
		pid int32
	)

	cmd := &cobra.Command{
		Use:   "all-model-versions",
		Short: "Query the list of all versions for a given Device Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetModelVersionsRequest{
				Vid: vid,
				Pid: pid,
			}

			res, err := queryClient.ModelVersions(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
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
