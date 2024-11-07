package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowNocCertificatesBySubject() *cobra.Command {
	var subject string

	cmd := &cobra.Command{
		Use:   "all-noc-subject-x509-certs",
		Short: "Gets all noc certificates (root and ica) associated with subject",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var res types.NocCertificatesBySubject

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.NocCertificatesBySubjectKeyPrefix,
				types.NocCertificatesBySubjectKey(subject),
				&res,
			)
		},
	}

	cmd.Flags().StringVarP(&subject, FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)

	return cmd
}
