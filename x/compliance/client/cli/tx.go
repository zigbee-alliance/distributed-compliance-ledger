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
		//GetCmdDeleteModelInfo(cdc), Disable deletion
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdAddModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "add-model-info <id> <name> <description> <sku> <firmware-version> <hardware-version> " +
			"<tis-or-trp-testing-completed> [certificate-id] [certified-date]",
		Short: "add new ModelInfo",
		Args:  cobra.RangeArgs(7, 9),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			name := args[1]

			description := args[2]
			sku := args[3]
			firmwareVersion := args[4]
			hardwareVersion := args[5]

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[6])
			if err != nil {
				return err
			}

			var certificateID string
			var certifiedDate time.Time

			switch len(args) {
			case 8: // or error?
				{
					certificateID = args[7]
				}
			case 9:
				{
					certificateID = args[7]
					certifiedDate, err = time.Parse(time.RFC3339, args[8])
					if err != nil {
						return err
					}
				}
			default:
			}

			msg := types.NewMsgAddModelInfo(id, name, description, sku, firmwareVersion, hardwareVersion,
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
		Use: "update-model-info <id> <new-name> <new-owner> <new-description> <new-sku> <new-firmware-version> " +
			"<new-hardware-version> <new-tis-or-trp-testing-completed> [new-certificate-id] [new-certified-date]",
		Short: "update existing ModelInfo",
		Args:  cobra.RangeArgs(9, 10),
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

			newTisOrTrpTestingCompleted, err := strconv.ParseBool(args[7])
			if err != nil {
				return err
			}

			var newCertificateID string
			var newCertifiedDate time.Time

			switch len(args) {
			case 9: // or error?
				{
					newCertificateID = args[8]
				}
			case 10:
				{
					newCertificateID = args[8]
					newCertifiedDate, err = time.Parse(time.RFC3339, args[9])
					if err != nil {
						return err
					}
				}
			default:
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
		Use:   "delete-model-info <id>",
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
