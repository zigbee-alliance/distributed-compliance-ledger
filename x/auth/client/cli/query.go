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
	"github.com/spf13/viper"
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
		GetCmdAccount(storeKey, cdc),
		GetCmdAccounts(storeKey, cdc),
		GetCmdProposedAccounts(storeKey, cdc),
	)...)

	return modelinfoQueryCmd
}

func GetCmdAccount(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Get account associated with the address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			res, height, err := cliCtx.QueryStore(types.GetAccountKey(address), queryRoute)
			if err != nil {
				return types.ErrAccountDoesNotExist(address)
			}

			var account types.Account
			cdc.MustUnmarshalBinaryBare(res, &account)

			out := cdc.MustMarshalJSON(types.ZBAccount(account))

			return cliCtx.PrintWithHeight(out, height)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

type AccountWithHeight struct {
	Result types.Account `json:"result"`
	Height int64         `json:"height"`
}

func GetCmdProposedAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-accounts",
		Short: "Get all proposed but not approved accounts",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()
			return cliCtx.QueryList(fmt.Sprintf("custom/%s/proposed_accounts", queryRoute), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of accounts to take")

	return cmd
}

func GetCmdAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-accounts",
		Short: "Get all accounts",
		Args:  cobra.ExactArgs(0),
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
