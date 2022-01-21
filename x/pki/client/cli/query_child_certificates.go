package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowChildCertificates() *cobra.Command {
	var (
		subject      string
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use:   "all-child-x509-certs",
		Short: "Gets all child certificates for the given certificate",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.ChildCertificates
			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ChildCertificatesKeyPrefix,
				types.ChildCertificatesKey(subject, subjectKeyID),
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
