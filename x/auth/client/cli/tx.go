package cli

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authorization subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdProposeAddAccount(cdc),
		GetCmdApproveAddAccount(cdc),
	)...)...)

	return authTxCmd
}

func GetCmdProposeAddAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-account",
		Short: "Propose a new account with the given address, public key and roles",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			pubkey := viper.GetString(FlagPubKey)
			_, err = sdk.GetAccPubKeyBech32(pubkey)
			if err != nil {
				return err
			}

			var roles types.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, types.AccountRole(role))
				}
			}

			msg := types.NewMsgProposeAddAccount(address, pubkey, roles, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "Bench32 encoded account public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles, comma-separated, assigning to the account (supported roles: %v)",
			types.Roles))

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}

func GetCmdApproveAddAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-add-account",
		Short: "Approve the proposed account associated with the given address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveAddAccount(address, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
