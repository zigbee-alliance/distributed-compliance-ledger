// Copyright 2022 DSR Corporation
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
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdCreateVendorInfo() *cobra.Command {
	var (
		vid                  int32
		vendorName           string
		companyLegalName     string
		companyPreferredName string
		vendorLandingPageURL string
	)

	cmd := &cobra.Command{
		Use:   "add-vendor",
		Short: "Add a new VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVendorInfo(
				clientCtx.GetFromAddress().String(),
				vid,
				vendorName,
				companyLegalName,
				companyPreferredName,
				vendorLandingPageURL,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Vendor ID")
	cmd.Flags().StringVarP(&vendorName, FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringVarP(&companyLegalName, FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringVarP(&companyPreferredName, FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringVarP(&vendorLandingPageURL, FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagVendorName)
	_ = cmd.MarkFlagRequired(FlagCompanyLegalName)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdUpdateVendorInfo() *cobra.Command {
	var (
		vid                  int32
		vendorName           string
		companyLegalName     string
		companyPreferredName string
		vendorLandingPageURL string
	)

	cmd := &cobra.Command{
		Use:   "update-vendor",
		Short: "Update a VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVendorInfo(
				clientCtx.GetFromAddress().String(),
				vid,
				vendorName,
				companyLegalName,
				companyPreferredName,
				vendorLandingPageURL,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Vendor ID")
	cmd.Flags().StringVarP(&vendorName, FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringVarP(&companyLegalName, FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringVarP(&companyPreferredName, FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringVarP(&vendorLandingPageURL, FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
