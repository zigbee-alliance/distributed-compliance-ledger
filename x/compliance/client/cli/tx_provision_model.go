package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ = strconv.Itoa(0)

func CmdProvisionModel() *cobra.Command {
	var (
		vid                   int32
		pid                   int32
		softwareVersion       uint32
		softwareVersionString string
		provisionalDate       string
		certificationType     string
		reason                string
		cdVersionNumber       uint32
		programTypeVersion    string
		certificationID       string
	)

	cmd := &cobra.Command{
		Use:   "provision-model",
		Short: "Set provisional state for the model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProvisionModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				softwareVersion,
				softwareVersionString,
				cdVersionNumber,
				provisionalDate,
				certificationType,
				reason,
				programTypeVersion,
				certificationID,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPID, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0,
		"Software Version of model (uint32)")
	cmd.Flags().StringVar(&softwareVersionString, FlagSoftwareVersionString, "",
		"Software Version String of model")
	cmd.Flags().Uint32Var(&cdVersionNumber, FlagCDVersionNumber, 0,
		"CD Version Number of the certification")
	cmd.Flags().StringVarP(&certificationType, FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().StringVarP(&provisionalDate, FlagProvisionalDate, FlagDateShortcut, "",
		"The date of model provisioning (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z")
	cmd.Flags().StringVar(&reason, FlagReason, "",
		"Optional comment describing the reason of provisioning")
	cmd.Flags().StringVar(&programTypeVersion, FlagProgramTypeVersion, "",
		"Program Type Version of the certification")
	cmd.Flags().StringVar(&certificationID, FlagCertificationID, "",
		"Certification ID of the certification")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagProvisionalDate)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	cli.AddTxFlagsToCmd(cmd)

	return cmd
}
