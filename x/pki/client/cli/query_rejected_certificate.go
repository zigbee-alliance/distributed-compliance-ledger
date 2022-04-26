package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListRejectedCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-rejected-certificate",
		Short: "list all RejectedCertificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRejectedCertificateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RejectedCertificateAll(context.Background(), params)
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

func CmdShowRejectedCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-rejected-certificate [subject] [subject-key-id]",
		Short: "shows a RejectedCertificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSubject := args[0]
			argSubjectKeyId := args[1]

			params := &types.QueryGetRejectedCertificateRequest{
				Subject:      argSubject,
				SubjectKeyId: argSubjectKeyId,
			}

			res, err := queryClient.RejectedCertificate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
