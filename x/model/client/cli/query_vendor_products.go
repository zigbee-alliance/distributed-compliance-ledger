package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdShowVendorProducts() *cobra.Command {
	var vid int32

	cmd := &cobra.Command{
		Use:   "vendor-models",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetVendorProductsRequest{
				Vid: vid,
			}

			res, err := queryClient.VendorProducts(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")

	_ = cmd.MarkFlagRequired(FlagVid)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
