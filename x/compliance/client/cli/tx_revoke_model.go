package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ = strconv.Itoa(0)

func CmdRevokeModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-model",
		Short: "Revoke an existing model. Note that the corresponding model version and test results must be present on ledger.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint32E(args[2])
			if err != nil {
				return err
			}
			argSoftwareVersionString := args[3]
			argRevocationDate := args[4]
			argCertificationType := args[5]
			argReason := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeModel(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
				argSoftwareVersionString,
				argRevocationDate,
				argCertificationType,
				argReason,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().String(FlagSoftwareVersionString, "", "Model software version string")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().StringP(FlagCertificationDate, FlagCertificationDateShortcut, "",
		"The date of model certification (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "",
		"Optional comment describing the reason of certification")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagCertificationDate)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
