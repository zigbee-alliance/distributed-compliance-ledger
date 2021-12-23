package cli

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdCreateVendorInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-vendor",
		Short: "Add a new VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVendorID, err := cast.ToInt32E(FlagVID)
			if err != nil {
				return err
			}

			// Get value arguments
			argVendorName := FlagVendorName
			argCompanyLegalName := FlagCompanyLegalName
			argCompanyPrefferedName := FlagCompanyPreferredName
			argVendorLandingPageURL := FlagVendorLandingPageURL

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVendorInfo(
				clientCtx.GetFromAddress().String(),
				indexVendorID,
				argVendorName,
				argCompanyLegalName,
				argCompanyPrefferedName,
				argVendorLandingPageURL,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagVID,
		"", "Vendor ID")
	cmd.Flags().StringP(FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringP(FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringP(FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringP(FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagVendorName)
	_ = cmd.MarkFlagRequired(FlagCompanyLegalName)

	return cmd
}

func CmdUpdateVendorInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vendor [vendor-id] [vendor-name] [company-legal-name] [company-preffered-name] [vendor-landing-page-url]",
		Short: "Update a VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVendorID, err := cast.ToInt32E(FlagVID)
			if err != nil {
				return err
			}

			// Get value arguments
			argVendorName := FlagVendorName
			argCompanyLegalName := FlagCompanyLegalName
			argCompanyPrefferedName := FlagCompanyPreferredName
			argVendorLandingPageURL := FlagVendorLandingPageURL

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVendorInfo(
				clientCtx.GetFromAddress().String(),
				indexVendorID,
				argVendorName,
				argCompanyLegalName,
				argCompanyPrefferedName,
				argVendorLandingPageURL,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagVID,
		"", "Vendor ID")
	cmd.Flags().StringP(FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringP(FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringP(FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringP(FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}
