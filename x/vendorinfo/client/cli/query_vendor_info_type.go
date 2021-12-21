package cli

import (
	"context"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdListVendorInfoType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vendor-info-type",
		Short: "list all VendorInfoType",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVendorInfoTypeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VendorInfoTypeAll(context.Background(), params)
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

func CmdShowVendorInfoType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vendor-info-type [vendor-id]",
		Short: "shows a VendorInfoType",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVendorID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetVendorInfoTypeRequest{
				VendorID: argVendorID,
			}

			res, err := queryClient.VendorInfoType(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
