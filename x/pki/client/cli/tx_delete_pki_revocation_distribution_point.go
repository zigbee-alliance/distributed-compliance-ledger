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

func CmdDeletePkiRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-pki-revocation-distribution-point [vid] [label] [issuer-subject-key-id]",
		Short: "Broadcast message delete-pki-revocation-distribution-point",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argLabel := args[1]
			argIssuerSubjectKeyID := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePkiRevocationDistributionPoint(
				clientCtx.GetFromAddress().String(),
				argVid,
				argLabel,
				argIssuerSubjectKeyID,
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
