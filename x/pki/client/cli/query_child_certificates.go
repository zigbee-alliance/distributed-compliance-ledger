package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListChildCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-child-certificates",
		Short: "list all ChildCertificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllChildCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ChildCertificatesAll(context.Background(), params)
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

func CmdShowChildCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-child-certificates [issuer] [authority-key-id]",
		Short: "shows a ChildCertificates",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIssuer := args[0]
			argAuthorityKeyId := args[1]

			params := &types.QueryGetChildCertificatesRequest{
				Issuer:         argIssuer,
				AuthorityKeyId: argAuthorityKeyId,
			}

			res, err := queryClient.ChildCertificates(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
