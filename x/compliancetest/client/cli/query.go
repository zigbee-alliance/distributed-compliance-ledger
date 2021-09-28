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
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
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
		Short: "Query testing results for Model (identified by the `vid`, `pid` and `softwareVersion`)",
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

			softwareVersion, err_ := conversions.ParseUInt32FromString("softwareVersion", viper.GetString(FlagSoftwareVersion))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetTestingResultsKey(vid, pid, softwareVersion), queryRoute)
			if err != nil || res == nil {
				return types.ErrTestingResultDoesNotExist(vid, pid, softwareVersion)
			}

			var testingResult types.TestingResults
			cdc.MustUnmarshalBinaryBare(res, &testingResult)

			return cliCtx.EncodeAndPrintWithHeight(testingResult, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model Software Version")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)

	return cmd
}
