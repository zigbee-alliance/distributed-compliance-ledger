package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModel() *cobra.Command {
	var (
		vid                                        int32
		pid                                        int32
		deviceTypeId                               int32
		productName                                string
		productLabel                               string
		partNumber                                 string
		commissioningCustomFlow                    int32
		commissioningCustomFlowUrl                 string
		commissioningModeInitialStepsHint          uint32
		commissioningModeInitialStepsInstruction   string
		commissioningModeSecondaryStepsHint        uint32
		commissioningModeSecondaryStepsInstruction string
		userManualUrl                              string
		supportUrl                                 string
		productUrl                                 string
		lsfUrl                                     string
		lsfRevision                                int32
	)

	cmd := &cobra.Command{
		Use:   "add-model",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			productLabel, err := utils.ReadFromFile(productLabel)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				deviceTypeId,
				productName,
				productLabel,
				partNumber,
				commissioningCustomFlow,
				commissioningCustomFlowUrl,
				commissioningModeInitialStepsHint,
				commissioningModeInitialStepsInstruction,
				commissioningModeSecondaryStepsHint,
				commissioningModeSecondaryStepsInstruction,
				userManualUrl,
				supportUrl,
				productUrl,
				lsfUrl,
				lsfRevision,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&deviceTypeId, FlagDeviceTypeId, 0,
		"Model category ID")
	cmd.Flags().StringVarP(&productName, FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringVarP(&productLabel, FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().StringVar(&partNumber, FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().Int32Var(&commissioningCustomFlow, FlagCommissioningCustomFlow, 0,
		`A value of 1 indicates that user interaction with the device (pressing a button, for example) is 
required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, 
the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with 
the necessary details for how to configure the product for initial commissioning.`)
	cmd.Flags().StringVar(&commissioningCustomFlowUrl, FlagCommissioningCustomFlowUrl, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().Uint32Var(&commissioningModeInitialStepsHint, FlagCommissioningModeInitialStepsHint, 0,
		`commissioningModeInitialStepsHint SHALL 
identify a hint for the steps that can be used to put into commissioning mode a device that 
has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. 
For example, a value of 1 (bit 0 is set) indicates 
that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.`)
	cmd.Flags().StringVar(&commissioningModeInitialStepsInstruction, FlagCommissioningModeInitialStepsInstruction, "",
		`commissioningModeInitialStepsInstruction SHALL contain text which relates to specific 
values of commissioningModeSecondaryStepsHint. Certain values of CommissioningModeInitialStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these 
values the commissioningModeInitialStepsInstruction SHALL be set`)
	cmd.Flags().Uint32Var(&commissioningModeSecondaryStepsHint, FlagCommissioningModeSecondaryStepsHint, 0,
		`commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can 
be used to put into commissioning mode a device that has already been commissioned. 
This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) 
indicates that a device that has already been commissioned will require the user to visit a 
current CHIP Administrator to put the device into commissioning mode.`)
	cmd.Flags().StringVar(&commissioningModeSecondaryStepsInstruction, FlagCommissioningModeSecondaryStepsInstruction, "",
		`commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values 
of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, 
and for these values the commissioningModeSecondaryStepInstruction SHALL be set`)
	cmd.Flags().StringVar(&userManualUrl, FlagUserManualUrl, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().StringVar(&supportUrl, FlagSupportUrl, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().StringVar(&productUrl, FlagProductUrl, "",
		"URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().StringVar(&lsfUrl, FlagLsfUrl, "", "URL to the Localized String File of this product")
	cmd.Flags().Int32Var(&lsfRevision, FlagLsfRevision, 0,
		"LsfRevision is a monotonically increasing positive integer indicating the latest available version of Localized String File")
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagDeviceTypeId)
	_ = cmd.MarkFlagRequired(FlagProductName)
	_ = cmd.MarkFlagRequired(FlagProductLabel)
	_ = cmd.MarkFlagRequired(FlagPartNumber)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdUpdateModel() *cobra.Command {
	var (
		vid                                        int32
		pid                                        int32
		productName                                string
		productLabel                               string
		partNumber                                 string
		commissioningCustomFlowUrl                 string
		commissioningModeInitialStepsInstruction   string
		commissioningModeSecondaryStepsInstruction string
		userManualUrl                              string
		supportUrl                                 string
		productUrl                                 string
		lsfUrl                                     string
		lsfRevision                                int32
	)

	cmd := &cobra.Command{
		Use:   "update-model",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			productLabel, err := utils.ReadFromFile(productLabel)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				productName,
				productLabel,
				partNumber,
				commissioningCustomFlowUrl,
				commissioningModeInitialStepsInstruction,
				commissioningModeSecondaryStepsInstruction,
				userManualUrl,
				supportUrl,
				productUrl,
				lsfUrl,
				lsfRevision,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().StringVarP(&productName, FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringVarP(&productLabel, FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().StringVar(&partNumber, FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().StringVar(&commissioningCustomFlowUrl, FlagCommissioningCustomFlowUrl, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().StringVar(&commissioningModeInitialStepsInstruction, FlagCommissioningModeInitialStepsInstruction, "",
		`commissioningModeInitialStepsInstruction SHALL contain text which relates to specific 
values of commissioningModeSecondaryStepsHint. Certain values of CommissioningModeInitialStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these 
values the commissioningModeInitialStepsInstruction SHALL be set`)
	cmd.Flags().StringVar(&commissioningModeSecondaryStepsInstruction, FlagCommissioningModeSecondaryStepsInstruction, "",
		`commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values 
of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, 
and for these values the commissioningModeSecondaryStepInstruction SHALL be set`)
	cmd.Flags().StringVar(&userManualUrl, FlagUserManualUrl, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().StringVar(&supportUrl, FlagSupportUrl, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().StringVar(&productUrl, FlagProductUrl, "",
		"URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().StringVar(&lsfUrl, FlagLsfUrl, "", "URL to the Localized String File of this product")
	cmd.Flags().Int32Var(&lsfRevision, FlagLsfRevision, 0,
		"LsfRevision is a monotonically increasing positive integer indicating the latest available version of Localized String File")
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdDeleteModel() *cobra.Command {
	var (
		vid int32
		pid int32
	)

	cmd := &cobra.Command{
		Use:   "delete-model",
		Short: "Delete existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgDeleteModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0, "Model product ID (positive non-zero uint16)")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
