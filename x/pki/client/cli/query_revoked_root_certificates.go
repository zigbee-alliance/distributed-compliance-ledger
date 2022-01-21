package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowRevokedRootCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-x509-root-certs",
		Short: "Gets all revoked root certificates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.RevokedRootCertificates
			return cli.QueryWithProofList(
				clientCtx,
				types.StoreKey,
				types.RevokedRootCertificatesKeyPrefix,
				types.RevokedRootCertificatesKey,
				&res,
			)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
