package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdShowVendorProducts() *cobra.Command {
	var vid int32

	cmd := &cobra.Command{
		Use:   "vendor-models",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.VendorProducts

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.VendorProductsKeyPrefix,
				types.VendorProductsKey(vid),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")

	_ = cmd.MarkFlagRequired(FlagVid)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
