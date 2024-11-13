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

func CmdListCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-certs",
		Short: "Gets all certificates. This query returns all types of certificates (PAA, PAI, RCAC, ICAC).",
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

			params := &types.QueryAllCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CertificatesAll(context.Background(), params)
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

func CmdShowCertificates() *cobra.Command {
	var (
		subject      string
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use: "cert",
		Short: "Gets certificate by the given combination of subject and subject-key-id or just subject-key-id. " +
			"This query works for all types of certificates (PAA, PAI, RCAC, ICAC).",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if subject != "" {
				var res types.AllCertificates

				return cli.QueryWithProof(
					clientCtx,
					pkitypes.StoreKey,
					types.AllCertificatesKeyPrefix,
					types.AllCertificatesKey(subject, subjectKeyID),
					&res,
				)
			}
			var res types.AllCertificatesBySubjectKeyId

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.AllCertificatesBySubjectKeyIDKeyPrefix,
				types.AllCertificatesBySubjectKeyIDKey(subjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject - required")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex) - required")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
