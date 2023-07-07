package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowPkiRevocationDistributionPointsByIssuerSubjectKeyID() *cobra.Command {
	var issuerSubjectKeyID string

	cmd := &cobra.Command{
		Use:   "revocation-points",
		Short: "Gets all revocation points associated with issuer's subject key id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.PkiRevocationDistributionPointsByIssuerSubjectKeyID

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix,
				types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(issuerSubjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&issuerSubjectKeyID, FlagIssuerSubjectKeyID, "", "Issuer's subject key id")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagIssuerSubjectKeyID)

	return cmd
}
