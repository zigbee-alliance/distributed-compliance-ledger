package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModelVersion() *cobra.Command {
	var (
		vid                          int32
		pid                          int32
		softwareVersion              uint32
		softwareVersionString        string
		cdVersionNumber              int32
		firmwareInformation          string
		softwareVersionValid         bool
		otaURL                       string
		otaFileSize                  uint64
		otaChecksum                  string
		otaChecksumType              int32
		minApplicableSoftwareVersion uint32
		maxApplicableSoftwareVersion uint32
		releaseNotesURL              string
	)

	cmd := &cobra.Command{
		Use:   "add-model-version",
		Short: "Add new Model Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateModelVersion(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				softwareVersion,
				softwareVersionString,
				cdVersionNumber,
				firmwareInformation,
				softwareVersionValid,
				otaURL,
				otaFileSize,
				otaChecksum,
				otaChecksumType,
				minApplicableSoftwareVersion,
				maxApplicableSoftwareVersion,
				releaseNotesURL,
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
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0,
		"Software Version of model (uint32)")
	cmd.Flags().StringVar(&softwareVersionString, FlagSoftwareVersionString, "",
		"Software Version String of model")
	cmd.Flags().Int32Var(&cdVersionNumber, FlagCdVersionNumber, 0,
		"CD Version Number of the certification")
	cmd.Flags().StringVar(&firmwareInformation, FlagFirmwareInformation, "",
		`FirmwareInformation field included in the Device Attestation response
		 when this Software Image boots on the device`)
	// by default the Software Version is valid, unless --softwareVersionValid is passed by user explicitly
	cmd.Flags().BoolVar(&softwareVersionValid, FlagSoftwareVersionValid, true,
		"boolean flag to revoke the software version model")
	cmd.Flags().StringVar(&otaURL, FlagOtaURL, "",
		"URL where to obtain the OTA image")
	cmd.Flags().Uint64Var(&otaFileSize, FlagOtaFileSize, 0,
		"OtaFileSize is the total size of the OTA software image in bytes")
	cmd.Flags().StringVar(&otaChecksum, FlagOtaChecksum, "",
		`Digest of the entire contents of the associated OTA 
		Software Update Image under the OtaUrl attribute, 
		encoded in base64 string representation. The digest SHALL have been computed using 
		the algorithm specified in OtaChecksumType`)
	cmd.Flags().Int32Var(&otaChecksumType, FlagOtaChecksumType, 0,
		`Numberic identifier as defined in 
IANA Named Information Hash Algorithm Registry for the type of otaChecksum.
For example, a value of 1 would match the sha-256 identifier, 
which maps to the SHA-256 digest algorithm`)
	cmd.Flags().Uint32Var(&minApplicableSoftwareVersion, FlagMinApplicableSoftwareVersion, 0,
		`MinApplicableSoftwareVersion should specify the lowest 
SoftwareVersion for which this image can be applied`)
	cmd.Flags().Uint32Var(&maxApplicableSoftwareVersion, FlagMaxApplicableSoftwareVersion, 0,
		`MaxApplicableSoftwareVersion should specify the highest 
SoftwareVersion for which this image can be applied`)
	cmd.Flags().StringVar(&releaseNotesURL, FlagReleaseNotesURL, "",
		`URL that contains product specific web page that contains 
release notes for the device model.`)

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCdVersionNumber)
	_ = cmd.MarkFlagRequired(FlagMinApplicableSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagMaxApplicableSoftwareVersion)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdUpdateModelVersion() *cobra.Command {
	var (
		vid                          int32
		pid                          int32
		softwareVersion              uint32
		softwareVersionValid         bool
		otaURL                       string
		minApplicableSoftwareVersion uint32
		maxApplicableSoftwareVersion uint32
		releaseNotesURL              string
	)

	cmd := &cobra.Command{
		Use:   "update-model-version",
		Short: "Update existing Model Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdateModelVersion(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				softwareVersion,
				softwareVersionValid,
				otaURL,
				minApplicableSoftwareVersion,
				maxApplicableSoftwareVersion,
				releaseNotesURL,
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
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0,
		"Software Version of model (uint32)")
	// by default the Software Version is valid, unless --softwareVersionValid is passed by user explicitly
	// FIXME: This behavior looks erroneous because the user can implicitly change invalid model version to valid
	cmd.Flags().BoolVar(&softwareVersionValid, FlagSoftwareVersionValid, true,
		"boolean flag to revoke the software version model")
	cmd.Flags().StringVar(&otaURL, FlagOtaURL, "",
		"URL where to obtain the OTA image")
	cmd.Flags().Uint32Var(&minApplicableSoftwareVersion, FlagMinApplicableSoftwareVersion, 0,
		`MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied`)
	cmd.Flags().Uint32Var(&maxApplicableSoftwareVersion, FlagMaxApplicableSoftwareVersion, 0,
		`MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied`)
	cmd.Flags().StringVar(&releaseNotesURL, FlagReleaseNotesURL, "",
		`URL that contains product specific web page that contains release notes for the device model.`)

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdDeleteModelVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model-version [vid] [pid] [software-version]",
		Short: "Broadcast message DeleteModelVersion",
		Args:  cobra.ExactArgs(3),
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelVersion(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
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
