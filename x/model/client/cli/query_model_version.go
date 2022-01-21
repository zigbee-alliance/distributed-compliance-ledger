package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdShowModelVersion() *cobra.Command {
	var (
		vid             int32
		pid             int32
		softwareVersion uint32
	)

	cmd := &cobra.Command{
		Use:   "get-model-version",
		Short: "Query Model Version by combination of Vendor ID, Product ID and Software Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.ModelVersion
			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ModelVersionKeyPrefix,
				types.ModelVersionKey(vid, pid, softwareVersion),
				&res,
			)
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
