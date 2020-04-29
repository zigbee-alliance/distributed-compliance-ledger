package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authnextTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authentication extensions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authnextTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateAccount(cdc),
	)...)

	return authnextTxCmd
}

func GetCmdCreateAccount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-account [addr] [pub-key]",
		Short: "Create new account for specified address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			authtypes.NewBaseAccountWithAddress(addr)

			msg := types.NewMsgAddAccount(addr, args[1], cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}
