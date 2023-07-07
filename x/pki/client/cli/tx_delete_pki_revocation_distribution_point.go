package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdDeletePkiRevocationDistributionPoint() *cobra.Command {
	var (
		vid                int32
		label              string
		issuerSubjectKeyID string
	)

	cmd := &cobra.Command{
		Use:   "delete-revocation-point",
		Short: "Deletes a PKI Revocation distribution endpoint owned by the sender",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePkiRevocationDistributionPoint(
				clientCtx.GetFromAddress().String(),
				vid,
				label,
				issuerSubjectKeyID,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Vendor ID (positive non-zero). Must be the same as Vendor account's VID and vid field in the VID-scoped CRLSignerCertificate")
	cmd.Flags().StringVarP(&label, FlagLabel, FlagLabelShortcut, "", " A label to disambiguate multiple revocation information partitions of a particular issuer")
	cmd.Flags().StringVar(&issuerSubjectKeyID, FlagIssuerSubjectKeyID, "", "Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: 5A880E6C3653D07FB08971A3F473790930E62BDB")
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagLabel)
	_ = cmd.MarkFlagRequired(FlagIssuerSubjectKeyID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
