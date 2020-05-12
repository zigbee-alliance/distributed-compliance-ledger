package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Use:   "test-result",
		Short: "Query testing results for Model (identified by the `vid` and `pid`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
			if err_ != nil {
				return err_
			}

			pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(keeper.TestingResultId(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrTestingResultDoesNotExist(vid, pid)
			}

			var testingResult types.TestingResults
			cdc.MustUnmarshalBinaryBare(res, &testingResult)

			return cliCtx.EncodeAndPrintWithHeight(testingResult, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)

	return cmd
}
