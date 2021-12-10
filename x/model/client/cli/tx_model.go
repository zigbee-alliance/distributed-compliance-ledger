package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-model [vid] [pid] [device-type-id] [product-name] [product-label] [part-number] [commissioning-custom-flow] [commissioning-custom-flow-url] [commissioning-mode-initial-steps-hint] [commissioning-mode-initial-steps-instruction] [commissioning-mode-secondary-steps-hint] [commissioning-mode-secondary-steps-instruction] [user-manual-url] [support-url] [product-url]",
		Short: "Create a new Model",
		Args:  cobra.ExactArgs(15),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			// Get value arguments
			argDeviceTypeId, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argProductName := args[3]
			argProductLabel := args[4]
			argPartNumber := args[5]
			argCommissioningCustomFlow, err := cast.ToInt32E(args[6])
			if err != nil {
				return err
			}
			argCommissioningCustomFlowUrl := args[7]
			argCommissioningModeInitialStepsHint, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}
			argCommissioningModeInitialStepsInstruction := args[9]
			argCommissioningModeSecondaryStepsHint, err := cast.ToUint64E(args[10])
			if err != nil {
				return err
			}
			argCommissioningModeSecondaryStepsInstruction := args[11]
			argUserManualUrl := args[12]
			argSupportUrl := args[13]
			argProductUrl := args[14]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModel(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				argDeviceTypeId,
				argProductName,
				argProductLabel,
				argPartNumber,
				argCommissioningCustomFlow,
				argCommissioningCustomFlowUrl,
				argCommissioningModeInitialStepsHint,
				argCommissioningModeInitialStepsInstruction,
				argCommissioningModeSecondaryStepsHint,
				argCommissioningModeSecondaryStepsInstruction,
				argUserManualUrl,
				argSupportUrl,
				argProductUrl,
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

func CmdUpdateModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model [vid] [pid] [product-name] [product-label] [part-number] [commissioning-custom-flow-url] [commissioning-mode-initial-steps-instruction] [commissioning-mode-secondary-steps-instruction] [user-manual-url] [support-url] [product-url]",
		Short: "Update a Model",
		Args:  cobra.ExactArgs(15),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			// Get value arguments
			argProductName := args[2]
			argProductLabel := args[3]
			argPartNumber := args[4]
			argCommissioningCustomFlowUrl := args[5]
			argCommissioningModeInitialStepsInstruction := args[6]
			argCommissioningModeSecondaryStepsInstruction := args[7]
			argUserManualUrl := args[8]
			argSupportUrl := args[9]
			argProductUrl := args[10]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModel(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				argProductName,
				argProductLabel,
				argPartNumber,
				argCommissioningCustomFlowUrl,
				argCommissioningModeInitialStepsInstruction,
				argCommissioningModeSecondaryStepsInstruction,
				argUserManualUrl,
				argSupportUrl,
				argProductUrl,
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

func CmdDeleteModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model [vid] [pid]",
		Short: "Delete a Model",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModel(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
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
