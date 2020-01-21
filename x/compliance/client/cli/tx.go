package cli

import (
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
		Use:   "add-model-info [id] [family] [cert] [owner]",
		Short: "add new ModelInfo",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			family := args[1]
			cert := args[2]

			owner, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddModelInfo(id, family, cert, owner, cliCtx.GetFromAddress())
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
		Use:   "update-model-info [id] [new-family] [new-cert] [new-owner]",
		Short: "update existing ModelInfo",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			newID := args[0]
			newFamily := args[1]
			newCert := args[2]

			newOwner, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModelInfo(newID, newFamily, newCert, newOwner, cliCtx.GetFromAddress())
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
