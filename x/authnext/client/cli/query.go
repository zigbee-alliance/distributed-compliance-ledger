package cli

import (
	"fmt"

	"github.com/spf13/viper"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

const (
	FlagSkip = "skip"
	FlagTake = "take"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authentication extensions module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccountHeaders(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdAccountHeaders(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-headers [start] [count]",
		Short: "List all ModelInfo IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := types.NewQueryAccountHeadersParams(
				viper.GetInt(FlagSkip),
				viper.GetInt(FlagTake),
			)

			data := cdc.MustMarshalJSON(params)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account_headers", queryRoute), data)
			if err != nil {
				fmt.Printf("could not get account headers\n")
				return nil
			}

			var out types.QueryAccountHeadersResult
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of accounts to take")

	return cmd
}
