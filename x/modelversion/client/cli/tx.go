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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "model version transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	modelTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddModelVersion(cdc),
		GetCmdUpdateModelVersion(cdc),
		// GetCmdDeleteModelVersion(cdc), Disable deletion
	)...)...)

	return modelTxCmd
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
