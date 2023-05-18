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

func CmdAddPkiRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-revocation-point",
		Short: "Broadcast message add-pki-revocation-distribution-point",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argIsPAA, err := cast.ToBoolE(args[2])
			if err != nil {
				return err
			}
			argLabel := args[3]
			argCrlSignerCertificate := args[4]
			argIssuerSubjectKeyID := args[5]
			argDataUrl := args[6]
			argDataFileSize, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argDataDigest := args[8]
			argDataDigestType, err := cast.ToUint32E(args[9])
			if err != nil {
				return err
			}
			argRevocationType, err := cast.ToUint64E(args[10])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddPkiRevocationDistributionPoint(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argIsPAA,
				argLabel,
				argCrlSignerCertificate,
				argIssuerSubjectKeyID,
				argDataUrl,
				argDataFileSize,
				argDataDigest,
				argDataDigestType,
				argRevocationType,
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
