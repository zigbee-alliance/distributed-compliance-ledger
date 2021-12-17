package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModelVersion() *cobra.Command {
	var (
		vid                          int32
		pid                          int32
		softwareVersion              uint64
		softwareVersionString        string
		cdVersionNumber              int32
		firmwareDigests              string
		softwareVersionValid         bool
		otaUrl                       string
		otaFileSize                  uint64
		otaChecksum                  string
		otaChecksumType              int32
		minApplicableSoftwareVersion uint64
		maxApplicableSoftwareVersion uint64
		releaseNotesUrl              string
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
				firmwareDigests,
				softwareVersionValid,
				otaUrl,
				otaFileSize,
				otaChecksum,
				otaChecksumType,
				minApplicableSoftwareVersion,
				maxApplicableSoftwareVersion,
				releaseNotesUrl,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID")
	cmd.Flags().Uint64VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0,
		"Software Version of model (uint32)")
	cmd.Flags().StringVar(&softwareVersionString, FlagSoftwareVersionString, "",
		"Software Version String of model")
	cmd.Flags().Int32Var(&cdVersionNumber, FlagCdVersionNumber, 0,
		"CD Version Number of the certification")
	cmd.Flags().StringVar(&firmwareDigests, FlagFirmwareDigests, "",
		`FirmwareDigests field included in the Device Attestation response
		 when this Software Image boots on the device`)
	// by default the Software Version is valid, unless --softwareVersionValid is passed by user explicitly
	cmd.Flags().BoolVar(&softwareVersionValid, FlagSoftwareVersionValid, true,
		"boolean flag to revoke the software version model")
	cmd.Flags().StringVar(&otaUrl, FlagOtaUrl, "",
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
	cmd.Flags().Uint64Var(&minApplicableSoftwareVersion, FlagMinApplicableSoftwareVersion, 0,
		`MinApplicableSoftwareVersion should specify the lowest 
SoftwareVersion for which this image can be applied`)
	cmd.Flags().Uint64Var(&maxApplicableSoftwareVersion, FlagMaxApplicableSoftwareVersion, 0,
		`MaxApplicableSoftwareVersion should specify the highest 
SoftwareVersion for which this image can be applied`)
	cmd.Flags().StringVar(&releaseNotesUrl, FlagReleaseNotesUrl, "",
		`URL that contains product specific web page that contains 
release notes for the device model.`)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCdVersionNumber)
	_ = cmd.MarkFlagRequired(FlagMinApplicableSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagMaxApplicableSoftwareVersion)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateModelVersion() *cobra.Command {
	var (
		vid                          int32
		pid                          int32
		softwareVersion              uint64
		softwareVersionValid         bool
		otaUrl                       string
		minApplicableSoftwareVersion uint64
		maxApplicableSoftwareVersion uint64
		releaseNotesUrl              string
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
				otaUrl,
				minApplicableSoftwareVersion,
				maxApplicableSoftwareVersion,
				releaseNotesUrl,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID")
	cmd.Flags().Uint64VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0,
		"Software Version of model (uint32)")
	// by default the Software Version is valid, unless --softwareVersionValid is passed by user explicitly
	// FIXME: This behavior looks erroneous because the user can implicitly change invalid model version to valid
	cmd.Flags().BoolVar(&softwareVersionValid, FlagSoftwareVersionValid, true,
		"boolean flag to revoke the software version model")
	cmd.Flags().StringVar(&otaUrl, FlagOtaUrl, "",
		"URL where to obtain the OTA image")
	cmd.Flags().Uint64Var(&minApplicableSoftwareVersion, FlagMinApplicableSoftwareVersion, 0,
		`MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied`)
	cmd.Flags().Uint64Var(&maxApplicableSoftwareVersion, FlagMaxApplicableSoftwareVersion, 0,
		`MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied`)
	cmd.Flags().StringVar(&releaseNotesUrl, FlagReleaseNotesUrl, "",
		`URL that contains product specific web page that contains release notes for the device model.`)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
