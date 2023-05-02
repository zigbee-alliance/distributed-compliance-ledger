package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdUpdatePkiRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pki-revocation-distribution-point [vid] [label] [crl-signer-certificate] [issuer-subject-key-id] [data-url] [data-file-size] [data-digest] [data-digest-type]",
		Short: "Broadcast message update-pki-revocation-distribution-point",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argLabel := args[1]
			argCrlSignerCertificate := args[2]
			argIssuerSubjectKeyID := args[3]
			argDataUrl := args[4]
			argDataFileSize, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argDataDigest := args[6]
			argDataDigestType, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePkiRevocationDistributionPoint(
				clientCtx.GetFromAddress().String(),
				argVid,
				argLabel,
				argCrlSignerCertificate,
				argIssuerSubjectKeyID,
				argDataUrl,
				argDataFileSize,
				argDataDigest,
				argDataDigestType,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
