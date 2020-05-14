package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authnextQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authentication extensions module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	authnextQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccounts(storeKey, cdc),
		GetCmdAccount(storeKey, cdc),
	)...)

	return authnextQueryCmd
}

func GetCmdAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Query the list of accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()
			return cliCtx.QueryList(fmt.Sprintf("custom/%s/accounts", queryRoute), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of accounts to take")

	return cmd
}

func GetCmdAccount(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Query account information associated with the address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account/%v", queryRoute, addr.String()), nil)
			if err != nil {
				return types.ErrAccountDoesNotExist(addr.String())
			}

			return cliCtx.PrintWithHeight(res, height)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
