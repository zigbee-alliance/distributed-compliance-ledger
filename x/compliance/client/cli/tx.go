package cli

import (
	"strconv"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliance transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(client.PostCommands(
		GetCmdAddModelInfo(cdc),
		GetCmdUpdateModelInfo(cdc),
		GetCmdDeleteModelInfo(cdc),
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdAddModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "add-model-info [id] [name] [owner] [description] [sku] [firmware-version] [hardware-version] " +
			"[certificate-id] [certified-date] [tis-or-trp-testing-completed]",
		Short: "add new ModelInfo",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			name := args[1]

			owner, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			description := args[3]
			sku := args[4]
			firmwareVersion := args[5]
			hardwareVersion := args[6]
			certificateID := args[7]

			certifiedDate, err := time.Parse(time.RFC3339, args[8])
			if err != nil {
				return err
			}

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[9])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddModelInfo(id, name, owner, description, sku, firmwareVersion, hardwareVersion,
				certificateID, certifiedDate, tisOrTrpTestingCompleted, cliCtx.GetFromAddress())

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

//nolint dupl
func GetCmdUpdateModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "update-model-info [id] [new-name] [new-owner] [new-description] [new-sku] [new-firmware-version] " +
			"[new-hardware-version] [new-certificate-id] [new-certified-date] [new-tis-or-trp-testing-completed]",
		Short: "update existing ModelInfo",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			newName := args[1]

			newOwner, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			newDescription := args[3]
			newSku := args[4]
			newFirmwareVersion := args[5]
			newHardwareVersion := args[6]
			newCertificateID := args[7]

			newCertifiedDate, err := time.Parse(time.RFC3339, args[8])
			if err != nil {
				return err
			}

			newTisOrTrpTestingCompleted, err := strconv.ParseBool(args[9])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModelInfo(id, newName, newOwner, newDescription, newSku, newFirmwareVersion,
				newHardwareVersion, newCertificateID, newCertifiedDate, newTisOrTrpTestingCompleted, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-model-info [id]",
		Short: "delete existing ModelInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]

			msg := types.NewMsgDeleteModelInfo(id, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
