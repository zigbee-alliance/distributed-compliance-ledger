package cli

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authorization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	modelinfoQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccounts(storeKey, cdc),
		GetCmdAccount(storeKey, cdc),
	)...)

	return modelinfoQueryCmd
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
	return &cobra.Command{
		Use:   "account [address]",
		Short: "Query account information associated with the address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return sdk.ErrInvalidAddress(args[0])
			}

			res, height, err := cliCtx.QueryStore(types.GetAccountKey(address), queryRoute)
			if err != nil {
				return types.ErrAccountDoesNotExist(address)
			}

			var account types.Account
			cdc.MustUnmarshalBinaryBare(res, &account)

			out, err := cdc.MarshalJSON(account)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not encode result: %v", err))
			}

			return cliCtx.PrintWithHeight(out, height)
		},
	}
}
