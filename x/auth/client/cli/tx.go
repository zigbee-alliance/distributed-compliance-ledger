package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
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
		GetCmdCreateAccount(cdc),
		GetCmdAddAssignRole(cdc),
		GetCmdRevokeRole(cdc),
	)...)

	return authzTxCmd
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

			authtypes.NewBaseAccountWithAddress(addr)

			var roles types.AccountRoles
			if rolesStr:= viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, types.AccountRole(role))
				}
			}

			msg := types.NewMsgAddAccount(addr, viper.GetString(FlagPubKey), roles, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "Bench32 encoded account public key")
	cmd.Flags().String(FlagRoles, "", fmt.Sprintf("The list of roles (split by comma) to assign to account (supported roles: %v)", types.Roles))

	cmd.MarkFlagRequired(FlagAddress)
	cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}

func GetCmdAddAssignRole(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "assign-role [addr] [role]",
		Short: "Assign new role to the account",
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
		Short: "Revoke role from the account",
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
