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

func CmdListRevokedNocRootCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-noc-x509-root-certs",
		Short: "Gets all revoked noc root certificates (RCACs)",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRevokedNocRootCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RevokedNocRootCertificatesAll(context.Background(), params)
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

func CmdShowRevokedNocRootCertificates() *cobra.Command {
	var (
		subject      string
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use: "revoked-noc-x509-root-cert",
		Short: "Gets revoked noc root certificates (RCACs)" +
			"by the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.RevokedNocRootCertificates

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.RevokedNocRootCertificatesKeyPrefix,
				types.RevokedNocRootCertificatesKey(subject, subjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
