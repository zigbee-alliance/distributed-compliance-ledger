package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListProposedCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proposed-certificate",
		Short: "list all ProposedCertificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposedCertificateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedCertificateAll(context.Background(), params)
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

func CmdShowProposedCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proposed-certificate [subject] [subject-key-id]",
		Short: "shows a ProposedCertificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSubject := args[0]
			argSubjectKeyId := args[1]

			params := &types.QueryGetProposedCertificateRequest{
				Subject:      argSubject,
				SubjectKeyId: argSubjectKeyId,
			}

			res, err := queryClient.ProposedCertificate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
