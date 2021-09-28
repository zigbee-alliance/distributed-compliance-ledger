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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	vendorinfoTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Vendorinfo transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	vendorinfoTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddVendor(cdc),
		GetCmdUpdateVendor(cdc),
		// GetCmdDeleteModel(cdc), Disable deletion
	)...)...)

	return vendorinfoTxCmd
}

//nolint:funlen,gocognit
func GetCmdAddVendor(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-vendor",
		Short: "Add new vendor info",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vendorId, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			vendorName := viper.GetString(FlagVendorName)
			companyLegalName := viper.GetString(FlagCompanyLegalName)
			companyPreferredName := viper.GetString(FlagCompanyPreferredName)
			vendorLandingPageUrl := viper.GetString(FlagVendorLandingPageUrl)

			vendorInfo := types.VendorInfo{
				VendorId:             vendorId,
				VendorName:           vendorName,
				CompanyLegalName:     companyLegalName,
				CompanyPreferredName: companyPreferredName,
				VendorLandingPageUrl: vendorLandingPageUrl,
			}

			msg := types.NewMsgAddVendorInfo(vendorInfo, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "",
		"Vendor ID")
	cmd.Flags().String(FlagVendorName, "",
		"Vendor Name")
	cmd.Flags().String(FlagCompanyLegalName, "",
		"Company Legal Name")
	cmd.Flags().String(FlagCompanyPreferredName, "",
		"Company Preferred Name")
	cmd.Flags().String(FlagVendorLandingPageUrl, "",
		"Landing Page URL for the Vendor")
	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagVendorName)
	_ = cmd.MarkFlagRequired(FlagCompanyLegalName)
	return cmd
}

//nolint:funlen
func GetCmdUpdateVendor(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vendor",
		Short: "Update existing Vendor info",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vendorId, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			vendorName := viper.GetString(FlagVendorName)
			companyLegalName := viper.GetString(FlagCompanyLegalName)
			companyPreferredName := viper.GetString(FlagCompanyPreferredName)
			vendorLandingPageUrl := viper.GetString(FlagVendorLandingPageUrl)

			vendorInfo := types.VendorInfo{
				VendorId:             vendorId,
				VendorName:           vendorName,
				CompanyLegalName:     companyLegalName,
				CompanyPreferredName: companyPreferredName,
				VendorLandingPageUrl: vendorLandingPageUrl,
			}
			msg := types.NewMsgUpdateVendorInfo(vendorInfo, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
	cmd.Flags().String(FlagVID, "",
		"Vendor ID")
	cmd.Flags().String(FlagVendorName, "",
		"Vendor Name")
	cmd.Flags().String(FlagCompanyLegalName, "",
		"Company Legal Name")
	cmd.Flags().String(FlagCompanyPreferredName, "",
		"Company Preferred Name")
	cmd.Flags().String(FlagVendorLandingPageUrl, "",
		"Landing Page URL for the Vendor")

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}
