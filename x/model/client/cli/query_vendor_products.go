package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdListVendorProducts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vendor-products",
		Short: "list all VendorProducts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVendorProductsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VendorProductsAll(context.Background(), params)
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

func CmdShowVendorProducts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vendor-products [vid]",
		Short: "shows a VendorProducts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetVendorProductsRequest{
				Vid: argVid,
			}

			res, err := queryClient.VendorProducts(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
