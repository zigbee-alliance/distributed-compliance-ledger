// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliancetest transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddTestingResult(cdc),
	)...)...)

	return complianceTxCmd
}

func GetCmdAddTestingResult(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-test-result",
		Short: "Add new testing result for Model (identified by the `vid` and `pid`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			testResult, err_ := cliCtx.ReadFromFile(viper.GetString(FlagTestResult))
			if err_ != nil {
				return err_
			}

			testDate, err_ := time.Parse(time.RFC3339, viper.GetString(FlagTestDate))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestDate \"%v\": "+
					"it must be RFC3339 encoded date", viper.GetString(FlagTestDate)))
			}

			msg := types.NewMsgAddTestingResult(vid, pid, testResult, testDate, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().StringP(FlagTestResult, FlagTestResultShortcut, "",
		"Test result (string or path to file containing data)")
	cmd.Flags().StringP(FlagTestDate, FlagTestDateShortcut, "", "Date of test result (rfc3339 encoded)")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagTestResult)
	_ = cmd.MarkFlagRequired(FlagTestDate)

	return cmd
}
