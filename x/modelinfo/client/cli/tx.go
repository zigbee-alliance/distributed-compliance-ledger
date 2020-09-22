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
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

			version := viper.GetString(FlagVersion)

			name := viper.GetString(FlagName)

			description, err_ := cliCtx.ReadFromFile(viper.GetString(FlagDescription))
			if err_ != nil {
				return err_
			}

			sku := viper.GetString(FlagSKU)

			hardwareVersion := viper.GetString(FlagHardwareVersion)

			firmwareVersion := viper.GetString(FlagFirmwareVersion)
			otaURL := viper.GetString(FlagOtaURL)
			otaChecksum := viper.GetString(FlagOtaChecksum)
			otaChecksumType := viper.GetString(FlagOtaChecksumType)

			var custom string
			if customFilename := viper.GetString(FlagCustom); len(customFilename) != 0 {
				custom, err_ = cliCtx.ReadFromFile(customFilename)
				if err_ != nil {
					return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid custom:\"%v\"", err_))
				}
			}

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(viper.GetString(FlagTisOrTrpTestingCompleted))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Tis-or-trp-testing-completed: "+
					"Parsing Error: \"%v\" must be boolean", viper.GetString(FlagTisOrTrpTestingCompleted)))
			}

			msg := types.NewMsgAddModelInfo(vid, pid, cid, version, name, description, sku,
				hardwareVersion, firmwareVersion, otaURL, otaChecksum, otaChecksumType,
				custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().StringP(FlagVersion, FlagVersionShortcut, "",
		"Version of model info format")
	cmd.Flags().StringP(FlagName, FlagNameShortcut, "", "Model name")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().String(FlagSKU, "", "Model stock keeping unit")
	cmd.Flags().StringP(FlagHardwareVersion, FlagHardwareVersionShortcut, "",
		"Version of model hardware")
	cmd.Flags().StringP(FlagFirmwareVersion, FlagFirmwareVersionShortcut, "",
		"Version of model firmware")
	cmd.Flags().String(FlagOtaURL, "", "URL of the OTA")
	cmd.Flags().String(FlagOtaChecksum, "", "Checksum of the OTA")
	cmd.Flags().String(FlagOtaChecksumType, "", "Type of the OTA checksum")
	cmd.Flags().StringP(FlagCustom, FlagCustomShortcut, "",
		"Custom information (string or path to file containing data)")
	cmd.Flags().StringP(FlagTisOrTrpTestingCompleted, FlagTisOrTrpTestingCompletedShortcut, "",
		"Whether model has successfully completed TIS/TRP testing")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagDescription)
	_ = cmd.MarkFlagRequired(FlagSKU)
	_ = cmd.MarkFlagRequired(FlagHardwareVersion)
	_ = cmd.MarkFlagRequired(FlagFirmwareVersion)
	_ = cmd.MarkFlagRequired(FlagTisOrTrpTestingCompleted)

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

			var description string
			var err_ error
			if descriptionFilename := viper.GetString(FlagDescription); len(descriptionFilename) != 0 {
				description, err_ = cliCtx.ReadFromFile(descriptionFilename)
				if err_ != nil {
					return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid description:\"%v\"", err_))
				}
			}

			otaURL := viper.GetString(FlagOtaURL)

			var custom string
			if customFilename := viper.GetString(FlagCustom); len(customFilename) != 0 {
				custom, err_ = cliCtx.ReadFromFile(customFilename)
				if err_ != nil {
					return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid custom:\"%v\"", err_))
				}
			}

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(viper.GetString(FlagTisOrTrpTestingCompleted))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid tis-or-trp-testing-completed: "+
					"Parsing Error: \"%v\" must be boolean", viper.GetString(FlagTisOrTrpTestingCompleted)))
			}

			msg := types.NewMsgUpdateModelInfo(vid, pid, cid, description,
				otaURL, custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().StringP(FlagCustom, FlagCustomShortcut, "",
		"Custom information (string or path to file containing data)")
	cmd.Flags().StringP(FlagTisOrTrpTestingCompleted, FlagTisOrTrpTestingCompletedShortcut, "",
		"Whether model has successfully completed TIS/TRP testing")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagTisOrTrpTestingCompleted)

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
