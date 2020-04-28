package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authzTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authorization subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authzTxCmd.AddCommand(client.PostCommands(
		GetCmdAddAssignRole(cdc),
		GetCmdRevokeRole(cdc),
	)...)

	return authzTxCmd
}

func GetCmdAddAssignRole(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "assign-role [addr] [role]",
		Short: "assign new role to the account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAssignRole(addr, types.AccountRole(args[1]), cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}

func GetCmdRevokeRole(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke-role [addr] [role]",
		Short: "revoke role from the account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeRole(addr, types.AccountRole(args[1]), cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}
