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
		vid          int32
		subject      string
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use: "noc-x509-cert",
		Short: "Gets NOC certificates (either root or ica) by one of property combinations: " +
			"'subject + subject-key-id' or 'VID and subject-key-id' or just 'subject-key-id'",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if vid != 0 && subject != "" && subjectKeyID != "" {
				return clientCtx.PrintString("Incorrect parameters combination. " +
					"You must provide '--subject, --subject-key-id' or '--vid , --subject-key-id' or " +
					"just '--subject-key-id'")
			}

			if subject != "" && subjectKeyID != "" {
				var res types.NocCertificates

				return cli.QueryWithProof(
					clientCtx,
					pkitypes.StoreKey,
					types.NocCertificatesKeyPrefix,
					types.NocCertificatesKey(subject, subjectKeyID),
					&res,
				)
			}

			if vid != 0 && subjectKeyID != "" {
				var res types.NocCertificatesByVidAndSkid

				return cli.QueryWithProof(
					clientCtx,
					pkitypes.StoreKey,
					types.NocCertificatesByVidAndSkidKeyPrefix,
					types.NocCertificatesByVidAndSkidKey(vid, subjectKeyID),
					&res,
				)
			}

			var res types.NocCertificatesBySubjectKeyID

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.NocCertificatesBySubjectKeyIDKeyPrefix,
				types.NocCertificatesBySubjectKeyIDKey(subjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Vendor ID (positive non-zero) - optional")
	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject - optional")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex) - required")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
