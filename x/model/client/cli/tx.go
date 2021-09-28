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
