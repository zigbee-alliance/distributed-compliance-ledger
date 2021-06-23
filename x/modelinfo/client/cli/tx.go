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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Modelinfo transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	modelinfoTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddModel(cdc),
		GetCmdUpdateModel(cdc),
		// GetCmdDeleteModel(cdc), Disable deletion
	)...)...)

	return modelinfoTxCmd
}

//nolint:funlen
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

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			name := viper.GetString(FlagName)

			description, err_ := cliCtx.ReadFromFile(viper.GetString(FlagDescription))
			if err_ != nil {
				return err_
			}

			sku := viper.GetString(FlagSKU)

			var softwareVersion uint32
			if softwareVersionStr := viper.GetString(FlagSoftwareVersion); len(softwareVersionStr) != 0 {
				softwareVersion, err = conversions.ParseUInt32FromString(softwareVersionStr)
				if err != nil {
					return err
				}
			}

			softwareVersionString := viper.GetString(FlagSoftwareVersionString)

			var hardwareVersion uint32
			if hardwareVersionStr := viper.GetString(FlagHardwareVersion); len(hardwareVersionStr) != 0 {
				hardwareVersion, err = conversions.ParseUInt32FromString(hardwareVersionStr)
				if err != nil {
					return err
				}
			}

			hardwareVersionString := viper.GetString(FlagHardwareVersionString)

			var cdVersionNumber uint16
			if cdVersionNumberStr := viper.GetString(FlagCDVersionNumber); len(cdVersionNumberStr) != 0 {
				cdVersionNumber, err = conversions.ParseUInt16FromString(cdVersionNumberStr)
				if err != nil {
					return err
				}
			}

			firmwareDigests := viper.GetString(FlagFirmwareDigests)
			// bool
			revoked := viper.GetBool(FlagRevoked)
			otaURL := viper.GetString(FlagOtaURL)
			otaChecksum := viper.GetString(FlagOtaChecksum)
			otaChecksumType := viper.GetString(FlagOtaChecksumType)
			otaBlob := viper.GetString(FlagOtaBlob)
			var commissioningCustomFlow uint8
			if commissioningCustomFlowStr := viper.GetString(FlagCommissioningCustomFlow); len(commissioningCustomFlowStr) != 0 {
				commissioningCustomFlow, err = conversions.ParseUInt8FromString(commissioningCustomFlowStr)
				if err != nil {
					return err
				}
			}
			commissioningCustomFlowUrl := viper.GetString(FlagCommissioningCustomFlowUrl)

			var commissioningModeInitialStepsHint uint32
			if commissioningModeInitialStepsHintStr := viper.GetString(FlagCommissioningModeInitialStepsHint); len(commissioningModeInitialStepsHintStr) != 0 {
				commissioningModeInitialStepsHint, err = conversions.ParseUInt32FromString(commissioningModeInitialStepsHintStr)
				if err != nil {
					return err
				}
			}
			commissioningModeInitialStepsInstruction := viper.GetString(FlagCommissioningModeInitialStepsInstruction)

			var commissioningModeSecondaryStepsHint uint32
			if commissioningModeSecondaryStepsHintStr := viper.GetString(FlagCommissioningModeSecondaryStepsHint); len(commissioningModeSecondaryStepsHintStr) != 0 {
				commissioningModeSecondaryStepsHint, err = conversions.ParseUInt32FromString(commissioningModeSecondaryStepsHintStr)
				if err != nil {
					return err
				}
			}
			commissioningModeSecondaryStepsInstruction := viper.GetString(FlagCommissioningModeSecondaryStepsInstruction)
			releaseNotesUrl := viper.GetString(FlagReleaseNotesUrl)
			userManualUrl := viper.GetString(FlagUserManualUrl)
			supportUrl := viper.GetString(FlagSupportUrl)
			productURL := viper.GetString(FlagProductURL)
			chipBlob := viper.GetString(FlagChipBlob)
			vendorBlob := viper.GetString(FlagVendorBlob)

			model := types.Model{

				VID:                                      vid,
				PID:                                      pid,
				CID:                                      cid,
				Name:                                     name,
				Description:                              description,
				SKU:                                      sku,
				SoftwareVersion:                          softwareVersion,
				SoftwareVersionString:                    softwareVersionString,
				HardwareVersion:                          hardwareVersion,
				HardwareVersionString:                    hardwareVersionString,
				CDVersionNumber:                          cdVersionNumber,
				FirmwareDigests:                          firmwareDigests,
				Revoked:                                  revoked,
				OtaURL:                                   otaURL,
				OtaChecksum:                              otaChecksum,
				OtaChecksumType:                          otaChecksumType,
				OtaBlob:                                  otaBlob,
				CommissioningCustomFlow:                  commissioningCustomFlow,
				CommissioningCustomFlowUrl:               commissioningCustomFlowUrl,
				CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
				CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
				CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
				CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
				ReleaseNotesUrl: releaseNotesUrl,
				UserManualUrl:   userManualUrl,
				SupportUrl:      supportUrl,
				ProductURL:      productURL,
				ChipBlob:        chipBlob,
				VendorBlob:      vendorBlob,
			}

			msg := types.NewMsgAddModelInfo(model, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().StringP(FlagName, FlagNameShortcut, "", "Model name")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().String(FlagSKU, "", "Model stock keeping unit")
	cmd.Flags().StringP(FlagSoftwareVersion, FlagSoftwareVersionShortcut, "",
		"Software Version of model (uint32)")
	cmd.Flags().String(FlagSoftwareVersionString, "", "Software Version String of model")
	cmd.Flags().StringP(FlagHardwareVersion, FlagHardwareVersionShortcut, "",
		"Version of model hardware")
	cmd.Flags().String(FlagHardwareVersionString, "", "Hardware Version String of model")
	cmd.Flags().String(FlagCDVersionNumber, "", "CD Version Number of the Certification")
	cmd.Flags().String(FlagFirmwareDigests, "", "FirmwareDigests field included in the Device Attestation response when this Software Image boots on the device")
	cmd.Flags().String(FlagRevoked, "", "boolean flag to revoke the model")
	cmd.Flags().String(FlagOtaURL, "", "Url for OTA")
	cmd.Flags().String(FlagOtaChecksum, "", "Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType")
	cmd.Flags().String(FlagOtaChecksumType, "", "Legal values for OtaChecksumType are : SHA-256")
	cmd.Flags().String(FlagOtaBlob, "", "Metadata about OTA")
	cmd.Flags().String(FlagCommissioningCustomFlow, "", "A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with the necessary details for how to configure the product for initial commissioning.")
	cmd.Flags().String(FlagCommissioningCustomFlowUrl, "", "commisioning-custom-flow-url SHALL identify a vendor specific commissioning URL for the device model when the commisioning-custom-flow field is set to '2'")
	cmd.Flags().String(FlagCommissioningModeInitialStepsHint, "", "commissioning-mode-initial-steps-hint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.")
	cmd.Flags().String(FlagCommissioningModeInitialStepsInstruction, "", "commissioning-mode-initial-steps-instruction SHALL contain text which relates to specific values of commissioning-mode-secondary-steps-hint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioning-mode-initial-steps-instruction SHALL be set")
	cmd.Flags().String(FlagCommissioningModeSecondaryStepsHint, "", "commissioning-mode-secondary-steps-hint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.")
	cmd.Flags().String(FlagCommissioningModeSecondaryStepsInstruction, "", "commissioning-mode-secondary-step-instruction SHALL contain text which relates to specific values of commissioning-mode-secondary-steps-hint. Certain values of commissioning-mode-secondary-steps-hint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioning-mode-secondary-step-instruction SHALL be set")
	cmd.Flags().String(FlagReleaseNotesUrl, "", "URL that contains product specific web page that contains release notes for the device model.")
	cmd.Flags().String(FlagUserManualUrl, "", "URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().String(FlagSupportUrl, "", "URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().String(FlagProductURL, "", "URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().String(FlagChipBlob, "", "chip-blob SHALL identify CHIP specific configurations")
	cmd.Flags().String(FlagVendorBlob, "", "field for vendors to provide any additional metadata about the device model using a string, blob, or URL.")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagDescription)
	_ = cmd.MarkFlagRequired(FlagSKU)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagHardwareVersion)
	_ = cmd.MarkFlagRequired(FlagHardwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCDVersionNumber)
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

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			description, err_ := cliCtx.ReadFromFile(viper.GetString(FlagDescription))
			if err_ != nil {
				return err_
			}

			var cdVersionNumber uint16
			if cdVersionNumberStr := viper.GetString(FlagCDVersionNumber); len(cdVersionNumberStr) != 0 {
				cdVersionNumber, err = conversions.ParseUInt16FromString(cdVersionNumberStr)
				if err != nil {
					return err
				}
			}

			// bool
			revoked := viper.GetBool(FlagRevoked)
			otaURL := viper.GetString(FlagOtaURL)
			otaChecksum := viper.GetString(FlagOtaChecksum)
			otaChecksumType := viper.GetString(FlagOtaChecksumType)
			otaBlob := viper.GetString(FlagOtaBlob)

			commissioningCustomFlowUrl := viper.GetString(FlagCommissioningCustomFlowUrl)

			releaseNotesUrl := viper.GetString(FlagReleaseNotesUrl)
			userManualUrl := viper.GetString(FlagUserManualUrl)
			supportUrl := viper.GetString(FlagSupportUrl)
			productURL := viper.GetString(FlagProductURL)
			chipBlob := viper.GetString(FlagChipBlob)
			vendorBlob := viper.GetString(FlagVendorBlob)

			model := types.Model{

				VID:                        vid,
				PID:                        pid,
				CID:                        cid,
				Description:                description,
				CDVersionNumber:            cdVersionNumber,
				Revoked:                    revoked,
				OtaURL:                     otaURL,
				OtaChecksum:                otaChecksum,
				OtaChecksumType:            otaChecksumType,
				OtaBlob:                    otaBlob,
				CommissioningCustomFlowUrl: commissioningCustomFlowUrl,
				ReleaseNotesUrl:            releaseNotesUrl,
				UserManualUrl:              userManualUrl,
				SupportUrl:                 supportUrl,
				ProductURL:                 productURL,
				ChipBlob:                   chipBlob,
				VendorBlob:                 vendorBlob,
			}
			msg := types.NewMsgUpdateModelInfo(model, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().String(FlagCDVersionNumber, "", "CD Version Number of the Certification")
	cmd.Flags().String(FlagRevoked, "", "boolean flag to revoke the model")
	cmd.Flags().String(FlagOtaURL, "", "Url for OTA")
	cmd.Flags().String(FlagOtaChecksum, "", "Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType")
	cmd.Flags().String(FlagOtaChecksumType, "", "Legal values for OtaChecksumType are : SHA-256")
	cmd.Flags().String(FlagOtaBlob, "", "Metadata about OTA")
	cmd.Flags().String(FlagCommissioningCustomFlowUrl, "", "CommissioningCustomFlowUrl SHALL identify a vendor specific commissioning URL for the device model when the CommissioningCustomFlow field is set to '2'")
	cmd.Flags().String(FlagReleaseNotesUrl, "", "URL that contains product specific web page that contains release notes for the device model.")
	cmd.Flags().String(FlagUserManualUrl, "", "URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().String(FlagSupportUrl, "", "URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().String(FlagProductURL, "", "URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().String(FlagChipBlob, "", "ChipBlob SHALL identify CHIP specific configurations")
	cmd.Flags().String(FlagVendorBlob, "", "VendorBlob is a optional field for vendors to provide any additional metadata about the device model using a string, blob, or URL.")

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

			msg := types.NewMsgDeleteModelInfo(vid, pid, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)

	return cmd
}
