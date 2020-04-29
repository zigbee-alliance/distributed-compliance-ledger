package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/spf13/cobra"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliancetest transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(client.PostCommands(
		GetCmdAddTestingResult(cdc),
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdAddTestingResult(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-test-result [vid] [pid] [test-result-string-or-path] [test-date]",
		Short: "Add new testing result for Model (identified by the `vid` and `pid`)",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(args[1])
			if err != nil {
				return err
			}

			testResult, err_ := cliCtx.ReadFromFile(args[2])
			if err_ != nil {
				return err_
			}

			testDate, err_ := time.Parse(time.RFC3339, args[3])
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestDate \"%v\": it must be RFC3339 encoded date", args[3]))
			}

			msg := types.NewMsgAddTestingResult(vid, pid, testResult, testDate, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}
