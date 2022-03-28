package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowApprovedCertificatesBySubject() *cobra.Command {
	var subject string

	cmd := &cobra.Command{
		Use:   "all-subject-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf) associated with subject",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.ApprovedCertificatesBySubject

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ApprovedCertificatesBySubjectKeyPrefix,
				types.ApprovedCertificatesBySubjectKey(subject),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)

	return cmd
}
