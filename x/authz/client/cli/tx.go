package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

	authzTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddAssignRole(cdc),
		GetCmdRevokeRole(cdc),
	)...)...)

	return authzTxCmd
}

func GetCmdAddAssignRole(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assign-role",
		Short: "Assign new role to the account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgAssignRole(addr, types.AccountRole(viper.GetString(FlagRole)), cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagRole, "", "Role to assign")

	cmd.MarkFlagRequired(FlagAddress)
	cmd.MarkFlagRequired(FlagRole)

	return cmd
}

func GetCmdRevokeRole(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-role",
		Short: "Revoke role from the account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeRole(addr, types.AccountRole(viper.GetString(FlagRole)), cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagRole, "", "Role to assign")

	cmd.MarkFlagRequired(FlagAddress)
	cmd.MarkFlagRequired(FlagRole)

	return cmd
}
