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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelVersionQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the model version module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	modelVersionQueryCmd.AddCommand(client.GetCommands(
		GetCmdModelVersion(storeKey, cdc),
		GetCmdAllModelVersions(storeKey, cdc),
	)...)

	return modelVersionQueryCmd
}

func GetCmdModelVersion(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-model-version",
		Short: "Query Model Version by combination of Vendor ID, Product ID and Software Version",
		Args:  cobra.NoArgs,
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

			softwareVersion, err_ := conversions.ParseUInt32FromString(FlagSoftwareVersion, viper.GetString(FlagSoftwareVersion))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetModelVersionKey(vid, pid, softwareVersion), queryRoute)
			if err != nil || res == nil {
				return types.ErrModelVersionDoesNotExist(vid, pid, softwareVersion)
			}

			var modelVersion types.ModelVersion
			cdc.MustUnmarshalBinaryBare(res, &modelVersion)

			return cliCtx.EncodeAndPrintWithHeight(modelVersion, height)
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

func GetCmdAllModelVersions(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-model-versions",
		Short: "Query the list of all versions for a given Device Model",
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

			res, height, err := cliCtx.QueryStore(types.GetModelKey(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrNoModelVersionsExist(vid, pid)
			}

			var modelVersions types.ModelVersions
			cdc.MustUnmarshalBinaryBare(res, &modelVersions)

			return cliCtx.EncodeAndPrintWithHeight(modelVersions, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model Vendor ID")
	cmd.Flags().String(FlagPID, "", "Model Product ID")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)

	return cmd
}
