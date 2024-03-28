package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListProposedCertificateRevocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs-to-revoke",
		Short: "Gets all proposed but not approved root certificates to be revoked",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposedCertificateRevocationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedCertificateRevocationAll(context.Background(), params)
			if cli.IsKeyNotFoundRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
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
	var (
		subject      string
		subjectKeyID string
		serialNumber string
	)

	cmd := &cobra.Command{
		Use: "proposed-x509-root-cert-to-revoke",
		Short: "Gets a proposed but not approved root certificate to be revoked " +
			"with the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var res types.ProposedCertificateRevocation

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.ProposedCertificateRevocationKeyPrefix,
				types.ProposedCertificateRevocationKey(subject, subjectKeyID, serialNumber),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	cmd.Flags().StringVarP(&serialNumber, FlagSerialNumber, FlagSerialNumberShortcut, "", "Certificate's serial number")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
