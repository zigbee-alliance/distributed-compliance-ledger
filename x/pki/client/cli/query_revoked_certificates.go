package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListRevokedCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-x509-certs",
		Short: "Gets all revoked certificates (root, intermediate and leaf)",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRevokedCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RevokedCertificatesAll(context.Background(), params)
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

func CmdShowRevokedCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoked-x509-cert",
		Short: "Gets revoked certificates (either root, intermediate or leaf) " +
			"by the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			params := &types.QueryGetRevokedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			}

			res, err := queryClient.RevokedCertificates(context.Background(), params)
			if HandleError(err) != nil {
				return err
			}
			if err != nil {
				// show default (empty) value in CLI
				res = &types.QueryGetRevokedCertificatesResponse{RevokedCertificates: nil}
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
