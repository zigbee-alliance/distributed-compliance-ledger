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

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	vendorinfoQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the vendorinfo module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	vendorinfoQueryCmd.AddCommand(client.GetCommands(
		GetCmdVendor(storeKey, cdc),
		GetCmdVendors(storeKey, cdc),
	)...)

	return vendorinfoQueryCmd
}

func GetCmdVendor(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendor",
		Short: "Get vendor details for the given vendorId",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vendorId, vidErr := conversions.ParseVID(viper.GetString(FlagVID))
			if vidErr != nil {
				return vidErr
			}

			res, height, err := cliCtx.QueryStore(types.GetVendorInfoKey(vendorId), queryRoute)
			if err != nil || res == nil {
				return types.ErrVendorInfoDoesNotExist(vendorId)
			}

			var vendorInfo types.VendorInfo
			cdc.MustUnmarshalBinaryBare(res, &vendorInfo)

			// the trick to prevent appending of `type` field by cdc
			out := cdc.MustMarshalJSON(vendorInfo)

			return cliCtx.PrintWithHeight(out, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Unique ID assigned to the vendor")

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}

func GetCmdVendors(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-vendors",
		Short: "Get information about all vendors",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/%s", queryRoute, keeper.QueryAllVendors), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, pagination.FlagSkipUsage)
	cmd.Flags().Int(pagination.FlagTake, 0, pagination.FlagTakeUsage)

	return cmd
}
