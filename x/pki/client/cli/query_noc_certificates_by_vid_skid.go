package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowNocCertificatesByVidAndSkid() *cobra.Command {
	var (
		vid          int32
		subjectKeyID string
	)

	cmd := &cobra.Command{
		Use:        "noc-x509-certs",
		Short:      "Gets NOC (Root/ICA) certificates (RCAC/ICAC) by VID and Skid",
		Deprecated: "Command is deprecated in favour of 'noc-x509-cert'",
		Args:       cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.NocCertificatesByVidAndSkid

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.NocCertificatesByVidAndSkidKeyPrefix,
				types.NocCertificatesByVidAndSkidKey(vid, subjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Vendor ID (positive non-zero)")
	cmd.Flags().StringVarP(&subjectKeyID, FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
