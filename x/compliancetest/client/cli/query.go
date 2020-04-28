package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the compliancetest module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdTestingResult(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdTestingResult(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-result [vid] [pid]",
		Short: "Query testing results for Model (identified by the `vid` and `pid`)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, height, err := cliCtx.QueryStore(keeper.TestingResultId(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrTestingResultDoesNotExist(vid, pid)
			}

			var testingResult types.TestingResults
			cdc.MustUnmarshalBinaryBare(res, &testingResult)

			return cliCtx.EncodeAndPrintWithHeight(testingResult, height)
		},
	}

	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)
	return cmd
}
