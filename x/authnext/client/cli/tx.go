package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

	authnextTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdCreateAccount(cdc),
	)...)...)

	return authnextTxCmd
}

func GetCmdCreateAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-account",
		Short: "Create new account for specified address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			_, err = sdk.GetAccPubKeyBech32(viper.GetString(FlagPubkey))
			if err != nil {
				return err
			}

			authtypes.NewBaseAccountWithAddress(addr)

			msg := types.NewMsgAddAccount(addr, viper.GetString(FlagPubkey), cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubkey, "", "Bench32 encoded public key")

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubkey)

	return cmd
}
