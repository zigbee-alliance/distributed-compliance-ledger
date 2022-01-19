package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdListVendorInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-vendors",
		Short: "Get information about all vendors",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVendorInfoRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VendorInfoAll(context.Background(), params)
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

func CmdShowVendorInfo() *cobra.Command {
	var vid int32

	cmd := &cobra.Command{
		Use:   "vendor",
		Short: "Get vendor details for the given vendorID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			// node, err := clientCtx.GetNode()
			// if err != nil {
			// 	return err
			// }
			// status, err := node.Status(context.Background())
			// if err != nil {
			// 	return err
			// }
			// height := status.SyncInfo.LatestBlockHeight - 1
			// clientCtx = clientCtx.WithHeight(height)
			// println(clientCtx.Height)

			key := append(types.KeyPrefix(types.VendorInfoKeyPrefix), types.VendorInfoKey(vid)...)
			resBytes, _, err := clientCtx.QueryStore(key, types.StoreKey)

			if err != nil {
				return err
			}
			if resBytes == nil {
				return clientCtx.PrintString(cli.NotFoundOutput)
			}

			var res types.VendorInfo
			clientCtx.Codec.MustUnmarshal(resBytes, &res)
			return clientCtx.PrintProto(&res)

			// queryClient := types.NewQueryClient(clientCtx)

			// params := &types.QueryGetVendorInfoRequest{
			// 	VendorID: vid,
			// }

			// res, err := queryClient.VendorInfo(context.Background(), params)
			// if cli.IsNotFound(err) {
			// 	return clientCtx.PrintString(cli.NotFoundOutput)
			// }
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0, "Unique ID assigned to the vendor")

	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}
