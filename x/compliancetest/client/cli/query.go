package cli

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
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
	return &cobra.Command{
		Use:   "test-result [vid] [pid]",
		Short: "Query testing result by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			vid := args[0]
			pid := args[1]

			res, height, err := cliCtx.QueryStore([]byte(keeper.TestingResultId(vid, pid)), queryRoute)
			if err != nil || res == nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not query testing result VID:%v PID:%v", vid, pid))
			}

			var testingResult types.TestingResults
			cdc.MustUnmarshalBinaryBare(res, &testingResult)

			out, err := json.Marshal(testingResult)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not query testing result VID:%v PID:%v", vid, pid))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}
