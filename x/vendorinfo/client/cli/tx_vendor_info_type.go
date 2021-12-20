package cli

import (
    

    "github.com/spf13/cobra"
     "github.com/spf13/cast" 

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdCreateVendorInfoType() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-vendor-info-type [index] [vendor-id] [vendor-name] [company-legal-name] [company-preffered-name] [vendor-landing-page-url]",
		Short: "Create a new VendorInfoType",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		 argVendorID, err := cast.ToUint64E(args[1])
            if err != nil {
                return err
            }
		 argVendorName := args[2]
		 argCompanyLegalName := args[3]
		 argCompanyPrefferedName := args[4]
		 argVendorLandingPageURL := args[5]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVendorInfoType(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
                argVendorID,
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
		Use:   "update-vendor-info-type [index] [vendor-id] [vendor-name] [company-legal-name] [company-preffered-name] [vendor-landing-page-url]",
		Short: "Update a VendorInfoType",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		 argVendorID, err := cast.ToUint64E(args[1])
            if err != nil {
                return err
            }
		 argVendorName := args[2]
		 argCompanyLegalName := args[3]
		 argCompanyPrefferedName := args[4]
		 argVendorLandingPageURL := args[5]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVendorInfoType(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
                argVendorID,
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
		Use:   "delete-vendor-info-type [index]",
		Short: "Delete a VendorInfoType",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
             indexIndex := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVendorInfoType(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
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