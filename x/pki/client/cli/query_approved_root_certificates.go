package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowApprovedRootCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-root-certs",
		Short: "Gets all approved root certificates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var res types.ApprovedRootCertificates

			return cli.QueryWithProofList(
				clientCtx,
				pkitypes.StoreKey,
				pkitypes.ApprovedRootCertificatesKeyPrefix,
				pkitypes.ApprovedRootCertificatesKey,
				&res,
			)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
