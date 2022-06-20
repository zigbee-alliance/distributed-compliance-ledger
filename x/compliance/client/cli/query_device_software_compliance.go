package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func CmdListDeviceSoftwareCompliance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-device-software-compliance",
		Short: "list all DeviceSoftwareCompliance",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDeviceSoftwareComplianceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DeviceSoftwareComplianceAll(context.Background(), params)
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

func CmdShowDeviceSoftwareCompliance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-device-software-compliance [cd-certificate-id]",
		Short: "shows a DeviceSoftwareCompliance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCdCertificateId := args[0]

			params := &types.QueryGetDeviceSoftwareComplianceRequest{
				CDCertificateId: argCdCertificateId,
			}

			res, err := queryClient.DeviceSoftwareCompliance(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
