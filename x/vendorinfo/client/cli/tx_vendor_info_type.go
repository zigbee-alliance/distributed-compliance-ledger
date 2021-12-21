package cli

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdCreateVendorInfoType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vendor-info-type [vendor-id] [vendor-name] [company-legal-name] [company-preffered-name] [vendor-landing-page-url]",
		Short: "Create a new VendorInfoType",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVendorID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argVendorName := args[1]
			argCompanyLegalName := args[2]
			argCompanyPrefferedName := args[3]
			argVendorLandingPageURL := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVendorInfoType(
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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateVendorInfoType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vendor-info-type [vendor-id] [vendor-name] [company-legal-name] [company-preffered-name] [vendor-landing-page-url]",
		Short: "Update a VendorInfoType",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVendorID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argVendorName := args[1]
			argCompanyLegalName := args[2]
			argCompanyPrefferedName := args[3]
			argVendorLandingPageURL := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVendorInfoType(
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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteVendorInfoType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-vendor-info-type [vendor-id]",
		Short: "Delete a VendorInfoType",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexVendorID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVendorInfoType(
				clientCtx.GetFromAddress().String(),
				indexVendorID,
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
