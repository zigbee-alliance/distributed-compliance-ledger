package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListUniqueCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-unique-certificate",
		Short: "list all UniqueCertificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUniqueCertificateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UniqueCertificateAll(context.Background(), params)
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

func CmdShowUniqueCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-unique-certificate [issuer] [serial-number]",
		Short: "shows a UniqueCertificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIssuer := args[0]
			argSerialNumber := args[1]

			params := &types.QueryGetUniqueCertificateRequest{
				Issuer:       argIssuer,
				SerialNumber: argSerialNumber,
			}

			res, err := queryClient.UniqueCertificate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
