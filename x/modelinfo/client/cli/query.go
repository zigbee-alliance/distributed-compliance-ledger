package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
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
	cmd := &cobra.Command{
		Use:   "model [vid] [pid]",
		Short: "Query Model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, height, err := cliCtx.QueryStore(keeper.ModelInfoId(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrModelInfoDoesNotExist(vid, pid)
			}

			var modelInfo types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &modelInfo)

			return cliCtx.EncodeAndPrintWithHeight(modelInfo, height)
		},
	}

	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)
	return cmd
}

func GetCmdAllModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-models",
		Short: "Query the list of all Models",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()
			return cliCtx.QueryList(fmt.Sprintf("custom/%s/all_models", queryRoute), params)
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
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()
			return cliCtx.QueryList(fmt.Sprintf("custom/%s/vendors", queryRoute), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of vendors to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of vendors to take")

	return cmd
}

func GetCmdVendorModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendor-models [vid]",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid := args[0]

			res, height, err := cliCtx.QueryStore(keeper.VendorProductsId(vid), queryRoute)
			if err != nil || res == nil {
				return types.ErrVendorProductsDoNotExist(vid)
			}

			var vendorProducts types.VendorProducts
			cdc.MustUnmarshalBinaryBare(res, &vendorProducts)

			return cliCtx.EncodeAndPrintWithHeight(vendorProducts, height)
		},
	}

	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)
	return cmd
}
