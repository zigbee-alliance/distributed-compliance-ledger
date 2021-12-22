package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListProposedCertificateRevocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs-to-revoke",
		Short: "Gets all proposed but not approved root certificates to be revoked",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposedCertificateRevocationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedCertificateRevocationAll(context.Background(), params)
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

func CmdShowProposedCertificateRevocation() *cobra.Command {
	cmd := &cobra.Command{
		Use: "proposed-x509-root-cert-to-revoke",
		Short: "Gets a proposed but not approved root certificate to be revoked " +
			"with the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			params := &types.QueryGetProposedCertificateRevocationRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			}

			res, err := queryClient.ProposedCertificateRevocation(context.Background(), params)
			if err != nil {
				return err
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
