package cli

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the modelinfo module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	modelinfoQueryCmd.AddCommand(client.GetCommands(
		GetCmdModel(storeKey, cdc),
		GetCmdAllModels(storeKey, cdc),
		GetCmdVendors(storeKey, cdc),
		GetCmdVendorModels(storeKey, cdc),
	)...)

	return modelinfoQueryCmd
}

func GetCmdModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "model [vid] [pid]",
		Short: "Query Model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, height, err := cliCtx.QueryStore([]byte(keeper.ModelInfoId(vid, pid)), queryRoute)
			if err != nil || res == nil {
				return types.ErrModelInfoDoesNotExist(vid, pid)
			}

			var modelInfo types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &modelInfo)

			out, err := json.Marshal(modelInfo)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could encode result: %v", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}

func GetCmdAllModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-models",
		Short: "Query the list of all Models",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := pagination.ParseAndMarshalPaginationParamsFromFlags(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_models", queryRoute), params)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not query models: %s\n", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, res, height))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdVendors(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendors",
		Short: "Query the list of Vendors",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := pagination.ParseAndMarshalPaginationParamsFromFlags(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vendors", queryRoute), params)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not get query vendors: %s\n", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, res, height))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of vendors to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of vendors to take")

	return cmd
}

func GetCmdVendorModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vendor-models [vid]",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]

			res, height, err := cliCtx.QueryStore([]byte(keeper.VendorProductsId(vid)), queryRoute)
			if err != nil || res == nil {
				return types.ErrVendorProductsDoNotExist(vid)
			}

			var vendorProducts types.VendorProducts
			cdc.MustUnmarshalBinaryBare(res, &vendorProducts)

			out, err := json.Marshal(vendorProducts)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could encode result: %v", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}
