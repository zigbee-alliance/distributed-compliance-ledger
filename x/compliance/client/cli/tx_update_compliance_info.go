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

func CmdUpdateComplianceInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-compliance-info [vid] [pid] [software-version] [software-version-string] [certification-type] [c-d-version-number] [software-version-certification-status] [date] [reason] [owner] [c-d-certificate-id] [certification-route] [program-type] [program-type-version] [compliant-platform-used] [compliant-platform-version] [transport] [family-id] [supported-clusters] [os-version] [parent-child] [certification-id-of-software-component]",
		Short: "Broadcast message UpdateComplianceInfo",
		Args:  cobra.ExactArgs(22),
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
			argCertificationType := args[4]
			argCDVersionNumber, err := cast.ToUint32E(args[5])
			if err != nil {
				return err
			}
			argSoftwareVersionCertificationStatus, err := cast.ToUint32E(args[6])
			if err != nil {
				return err
			}
			argDate := args[7]
			argReason := args[8]
			argOwner := args[9]
			argCDCertificateId := args[10]
			argCertificationRoute := args[11]
			argProgramType := args[12]
			argProgramTypeVersion := args[13]
			argCompliantPlatformUsed := args[14]
			argCompliantPlatformVersion := args[15]
			argTransport := args[16]
			argFamilyId := args[17]
			argSupportedClusters := args[18]
			argOSVersion := args[19]
			argParentChild := args[20]
			argCertificationIdOfSoftwareComponent := args[21]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateComplianceInfo(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
				argSoftwareVersionString,
				argCertificationType,
				argCDVersionNumber,
				argSoftwareVersionCertificationStatus,
				argDate,
				argReason,
				argOwner,
				argCDCertificateId,
				argCertificationRoute,
				argProgramType,
				argProgramTypeVersion,
				argCompliantPlatformUsed,
				argCompliantPlatformVersion,
				argTransport,
				argFamilyId,
				argSupportedClusters,
				argOSVersion,
				argParentChild,
				argCertificationIdOfSoftwareComponent,
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
