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

func CmdListNocCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-noc-x509-certs",
		Short: "Gets all noc certificates (root and ica)",
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

			params := &types.QueryNocCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NocCertificatesAll(context.Background(), params)
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

func CmdShowNocCertificates() *cobra.Command {
	var (
		subject      string
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use: "noc-x509-cert",
		Short: "Gets certificates (either root or ica) " +
			"by the given combination of subject and subject-key-id or just subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if subject != "" {
				var res types.NocCertificates

				return cli.QueryWithProof(
					clientCtx,
					pkitypes.StoreKey,
					types.NocCertificatesKeyPrefix,
					types.NocCertificatesKey(subject, subjectKeyID),
					&res,
				)
			}
			var res types.NocCertificatesBySubjectKeyId

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.NocCertificatesBySubjectKeyIdKeyPrefix,
				types.NocCertificatesBySubjectKeyIdKey(subjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject - optional")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex) - required")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
