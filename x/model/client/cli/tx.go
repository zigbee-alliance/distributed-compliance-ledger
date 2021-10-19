// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Model transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	modelTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddModel(cdc),
		GetCmdUpdateModel(cdc),
		GetCmdAddModelVersion(cdc),
		GetCmdUpdateModelVersion(cdc),
		// GetCmdDeleteModel(cdc), Disable deletion
	)...)...)

	return modelTxCmd
}

//nolint:funlen,gocognit
func GetCmdAddModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-model",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			var deviceTypeID uint16
			if deviceTypeIDStr := viper.GetString(FlagDeviceTypeID); len(deviceTypeIDStr) != 0 {
				deviceTypeID, err = conversions.ParseUInt16FromString(FlagDeviceTypeID, deviceTypeIDStr)
				if err != nil {
					return err
				}
			}

			productName := viper.GetString(FlagProductName)

			productLabel, err_ := cliCtx.ReadFromFile(viper.GetString(FlagProductLabel))
			if err_ != nil {
				return err_
			}

			partNumber := viper.GetString(FlagPartNumber)

			var commissioningCustomFlow uint8
			if commissioningCustomFlowStr := viper.GetString(FlagCommissioningCustomFlow); len(commissioningCustomFlowStr) != 0 {
				commissioningCustomFlow, err = conversions.ParseUInt8FromString(commissioningCustomFlowStr)
				if err != nil {
					return err
				}
			}
			commissioningCustomFlowURL := viper.GetString(FlagCommissioningCustomFlowURL)

			var commissioningModeInitialStepsHint uint32
			commissioningModeInitialStepsHintStr := viper.GetString(FlagCommissioningModeInitialStepsHint)
			if len(commissioningModeInitialStepsHintStr) != 0 {
				commissioningModeInitialStepsHint, err = conversions.ParseUInt32FromString(FlagCommissioningModeInitialStepsHint, commissioningModeInitialStepsHintStr)
				if err != nil {
					return err
				}
			}
			commissioningModeInitialStepsInstruction := viper.GetString(FlagCommissioningModeInitialStepsInstruction)

			var commissioningModeSecondaryStepsHint uint32
			commissioningModeSecondaryStepsHintStr := viper.GetString(FlagCommissioningModeSecondaryStepsHint)
			if len(commissioningModeSecondaryStepsHintStr) != 0 {
				commissioningModeSecondaryStepsHint, err = conversions.ParseUInt32FromString(FlagCommissioningModeSecondaryStepsHint, commissioningModeSecondaryStepsHintStr)
				if err != nil {
					return err
				}
			}
			commissioningModeSecondaryStepsInstruction := viper.GetString(FlagCommissioningModeSecondaryStepsInstruction)
			userManualURL := viper.GetString(FlagUserManualURL)
			supportURL := viper.GetString(FlagSupportURL)
			productURL := viper.GetString(FlagProductURL)

			model := types.Model{
				VID:                                      vid,
				PID:                                      pid,
				DeviceTypeID:                             deviceTypeID,
				ProductName:                              productName,
				ProductLabel:                             productLabel,
				PartNumber:                               partNumber,
				CommissioningCustomFlow:                  commissioningCustomFlow,
				CommissioningCustomFlowURL:               commissioningCustomFlowURL,
				CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
				CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
				CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
				CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
				UserManualURL: userManualURL,
				SupportURL:    supportURL,
				ProductURL:    productURL,
			}

			msg := types.NewMsgAddModel(model, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "",
		"Model vendor ID")
	cmd.Flags().String(FlagPID, "",
		"Model product ID")
	cmd.Flags().String(FlagDeviceTypeID, "",
		"Model category ID")
	cmd.Flags().StringP(FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringP(FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().String(FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().String(FlagCommissioningCustomFlow, "",
		`A value of 1 indicates that user interaction with the device (pressing a button, for example) is 
	required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, 
	the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with 
	the necessary details for how to configure the product for initial commissioning.`)
	cmd.Flags().String(FlagCommissioningCustomFlowURL, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
	device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().String(FlagCommissioningModeInitialStepsHint, "",
		`commissioningModeInitialStepsHint SHALL 
	identify a hint for the steps that can be used to put into commissioning mode a device that 
	has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. 
	For example, a value of 1 (bit 0 is set) indicates 
	that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.`)
	cmd.Flags().String(FlagCommissioningModeInitialStepsInstruction, "",
		`commissioningModeInitialStepsInstruction SHALL contain text which relates to specific 
	values of commissioningModeSecondaryStepsHint. Certain values of CommissioningModeInitialStepsHint, 
	as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these 
	values the commissioningModeInitialStepsInstruction SHALL be set`)
	cmd.Flags().String(FlagCommissioningModeSecondaryStepsHint, "",
		`commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can 
	be used to put into commissioning mode a device that has already been commissioned. 
	This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) 
	indicates that a device that has already been commissioned will require the user to visit a 
	current CHIP Administrator to put the device into commissioning mode.`)
	cmd.Flags().String(FlagCommissioningModeSecondaryStepsInstruction, "",
		`commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values 
	of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, 
	as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, 
	and for these values the commissioningModeSecondaryStepInstruction SHALL be set`)
	cmd.Flags().String(FlagUserManualURL, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().String(FlagSupportURL, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().String(FlagProductURL, "",
		"URL that contains product specific web page that contains details for the device model.")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagDeviceTypeID)
	_ = cmd.MarkFlagRequired(FlagProductName)
	_ = cmd.MarkFlagRequired(FlagProductLabel)
	_ = cmd.MarkFlagRequired(FlagPartNumber)

	return cmd
}

//nolint:funlen
func GetCmdUpdateModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			productName := viper.GetString(FlagProductName)

			productLabel, err_ := cliCtx.ReadFromFile(viper.GetString(FlagProductLabel))
			if err_ != nil {
				return err_
			}

			partNumber := viper.GetString(FlagPartNumber)

			commissioningCustomFlowURL := viper.GetString(FlagCommissioningCustomFlowURL)

			userManualURL := viper.GetString(FlagUserManualURL)
			supportURL := viper.GetString(FlagSupportURL)
			productURL := viper.GetString(FlagProductURL)

			model := types.Model{
				VID:                        vid,
				PID:                        pid,
				ProductName:                productName,
				ProductLabel:               productLabel,
				PartNumber:                 partNumber,
				CommissioningCustomFlowURL: commissioningCustomFlowURL,
				UserManualURL:              userManualURL,
				SupportURL:                 supportURL,
				ProductURL:                 productURL,
			}
			msg := types.NewMsgUpdateModel(model, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
	cmd.Flags().String(FlagVID, "",
		"Model vendor ID")
	cmd.Flags().String(FlagPID, "",
		"Model product ID")
	cmd.Flags().StringP(FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringP(FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().String(FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().String(FlagCommissioningCustomFlowURL, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
	device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().String(FlagUserManualURL, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().String(FlagSupportURL, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().String(FlagProductURL, "",
		"URL that contains product specific web page that contains details for the device model.")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	return cmd
}

func GetCmdDeleteModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model",
		Short: "Delete existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModel(vid, pid, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)

	return cmd
}

//nolint:funlen,gocognit
func GetCmdAddModelVersion(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-model-version",
		Short: "Add new Model Version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			// by default the SoftwareVersion is valid, unless passed by user explicitly
			viper.SetDefault(FlagSoftwareVersionValid, true)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			softwareVersion, err := conversions.ParseUInt32FromString(FlagSoftwareVersion, viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}

			softwareVersionString := viper.GetString(FlagSoftwareVersionString)
			if len(softwareVersionString) == 0 {
				return types.ErrSoftwareVersionStringInvalid(softwareVersionString)
			}

			var cdVersionNumber uint16
			if len(viper.GetString(FlagCDVersionNumber)) != 0 {
				cdVersionNumber, err = conversions.ParseUInt16FromString(FlagCDVersionNumber, viper.GetString(FlagCDVersionNumber))
				if err != nil {
					return types.ErrCDVersionNumberInvalid(cdVersionNumber)
				}
			}

			firmwareDigests := viper.GetString(FlagFirmwareDigests)
			if len(firmwareDigests) > 512 {
				return types.ErrFirmwareDigestsInvalid(firmwareDigests)
			}

			softwareVersionValid := viper.GetBool(FlagSoftwareVersionValid)

			otaURL := viper.GetString(FlagOtaURL)
			if len(otaURL) > 256 {
				types.ErrOtaURLInvalid(otaURL)
			}

			var otaFileSize uint64
			if len(viper.GetString(FlagOtaFileSize)) != 0 {
				otaFileSize, err = conversions.ParseUInt64FromString(FlagOtaFileSize, viper.GetString(FlagOtaFileSize))
				if err != nil {
					return err
				}
			}

			otaChecksum := viper.GetString(FlagOtaChecksum)

			var otaChecksumType uint16
			if len(viper.GetString(FlagOtaChecksumType)) > 0 {
				otaChecksumType, err = conversions.ParseUInt16FromString(FlagOtaChecksumType, viper.GetString(FlagOtaChecksumType))
				if err != nil {
					return err
				}
			}

			if len(otaURL) > 1 {
				if otaFileSize == 0 || len(otaChecksum) == 0 || otaChecksumType == 0 {
					return types.ErrMissingOtaInformation()
				}
			}

			minApplicableSoftwareVersion, err := conversions.ParseUInt32FromString(FlagMinApplicableSoftwareVersion, viper.GetString(FlagMinApplicableSoftwareVersion))
			if err != nil {
				return err
			}

			maxApplicableSofwareVersion, err := conversions.ParseUInt32FromString(FlagMaxApplicableSoftwareVersion, viper.GetString(FlagMaxApplicableSoftwareVersion))
			if err != nil {
				return err
			}

			releaseNotesURL := viper.GetString(FlagReleaseNotesURL)
			if len(releaseNotesURL) > 256 {
				return types.ErrReleaseNotesURLInvalid(releaseNotesURL)
			}

			modelVersion := types.ModelVersion{
				VID:                          vid,
				PID:                          pid,
				SoftwareVersion:              softwareVersion,
				SoftwareVersionString:        softwareVersionString,
				CDVersionNumber:              cdVersionNumber,
				FirmwareDigests:              firmwareDigests,
				SoftwareVersionValid:         softwareVersionValid,
				OtaURL:                       otaURL,
				OtaFileSize:                  otaFileSize,
				OtaChecksum:                  otaChecksum,
				OtaChecksumType:              otaChecksumType,
				MinApplicableSoftwareVersion: minApplicableSoftwareVersion,
				MaxApplicableSoftwareVersion: maxApplicableSofwareVersion,
				ReleaseNotesURL:              releaseNotesURL,
			}

			msg := types.MsgAddModelVersion{
				ModelVersion: modelVersion,
				Signer:       cliCtx.FromAddress(),
			}

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "",
		"Model vendor ID")
	cmd.Flags().String(FlagPID, "",
		"Model product ID")
	cmd.Flags().StringP(FlagSoftwareVersion, FlagSoftwareVersionShortcut, "",
		"Software Version of model (uint32)")
	cmd.Flags().String(FlagSoftwareVersionString, "",
		"Software Version String of model")
	cmd.Flags().String(FlagCDVersionNumber, "",
		"CD Version Number of the certification")
	cmd.Flags().String(FlagFirmwareDigests, "",
		`FirmwareDigests field included in the Device Attestation response when this Software Image boots on the device`)
	cmd.Flags().String(FlagSoftwareVersionValid, "",
		"boolean flag to revoke the software version model")
	cmd.Flags().String(FlagOtaURL, "", "URL where to obtain the OTA image")
	cmd.Flags().String(FlagOtaFileSize, "", "OtaFileSize is the total size of the OTA software image in bytes")
	cmd.Flags().String(FlagOtaChecksum, "",
		`Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, 
	encoded in base64 string representation. The digest SHALL have been computed using 
	the algorithm specified in OtaChecksumType`)
	cmd.Flags().String(FlagOtaChecksumType, "", `Numberic identifier as defined in IANA Named Information Hash Algorithm Registry for the type of otaChecksum.
	 For example, a value of 1 would match the sha-256 identifier, which maps to the SHA-256 digest algorithm`)
	cmd.Flags().String(FlagMinApplicableSoftwareVersion, "",
		`MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied`)
	cmd.Flags().String(FlagMaxApplicableSoftwareVersion, "",
		`MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied`)
	cmd.Flags().String(FlagReleaseNotesURL, "",
		`URL that contains product specific web page that contains release notes for the device model.`)
	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCDVersionNumber)
	_ = cmd.MarkFlagRequired(FlagMinApplicableSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagMaxApplicableSoftwareVersion)
	return cmd
}

//nolint:funlen
func GetCmdUpdateModelVersion(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model-version",
		Short: "Update existing Model Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// by default the SoftwareVersion is valid, unless passed by user explicitly
			viper.SetDefault(FlagSoftwareVersionValid, true)
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			softwareVersion, err := conversions.ParseUInt32FromString(FlagSoftwareVersion, viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}

			softwareVersionValid := viper.GetBool(FlagSoftwareVersionValid)

			firmwareDigests := viper.GetString(FlagFirmwareDigests)

			otaURL := viper.GetString(FlagOtaURL)
			if len(otaURL) > 256 {
				types.ErrOtaURLInvalid(otaURL)
			}

			var minApplicableSoftwareVersion uint32
			if len(viper.GetString(FlagMinApplicableSoftwareVersion)) > 0 {
				minApplicableSoftwareVersion, err = conversions.ParseUInt32FromString(FlagMinApplicableSoftwareVersion, viper.GetString(FlagMinApplicableSoftwareVersion))
				if err != nil {
					return err
				}
			}

			var maxApplicableSofwareVersion uint32
			if len(viper.GetString(FlagMaxApplicableSoftwareVersion)) > 0 {
				maxApplicableSofwareVersion, err = conversions.ParseUInt32FromString(FlagMaxApplicableSoftwareVersion, viper.GetString(FlagMaxApplicableSoftwareVersion))
				if err != nil {
					return err
				}
			}

			releaseNotesURL := viper.GetString(FlagReleaseNotesURL)
			if len(releaseNotesURL) > 256 {
				return types.ErrReleaseNotesURLInvalid(releaseNotesURL)
			}

			modelVersion := types.ModelVersion{
				VID:                          vid,
				PID:                          pid,
				SoftwareVersion:              softwareVersion,
				SoftwareVersionValid:         softwareVersionValid,
				FirmwareDigests:              firmwareDigests,
				OtaURL:                       otaURL,
				MinApplicableSoftwareVersion: minApplicableSoftwareVersion,
				MaxApplicableSoftwareVersion: maxApplicableSofwareVersion,
				ReleaseNotesURL:              releaseNotesURL,
			}
			msg := types.MsgUpdateModelVersion{
				ModelVersion: modelVersion,
				Signer:       cliCtx.FromAddress(),
			}

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "",
		"Model vendor ID")
	cmd.Flags().String(FlagPID, "",
		"Model product ID")
	cmd.Flags().StringP(FlagSoftwareVersion, FlagSoftwareVersionShortcut, "",
		"Software Version of model (uint32)")
	cmd.Flags().String(FlagSoftwareVersionValid, "",
		"boolean flag to revoke the software version model")
	cmd.Flags().String(FlagFirmwareDigests, "",
		`FirmwareDigests field included in the Device Attestation response when this Software Image boots on the device`)
	cmd.Flags().String(FlagOtaURL, "", "URL where to obtain the OTA image")
	cmd.Flags().String(FlagMinApplicableSoftwareVersion, "",
		`MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied`)
	cmd.Flags().String(FlagMaxApplicableSoftwareVersion, "",
		`MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied`)
	cmd.Flags().String(FlagReleaseNotesURL, "",
		`URL that contains product specific web page that contains release notes for the device model.`)
	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	return cmd
}
