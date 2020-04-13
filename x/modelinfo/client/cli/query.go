package cli

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

const (
	FlagSkip = "skip"
	FlagTake = "take"
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
				return sdk.ErrInternal(fmt.Sprintf("Could not query model VID:%v PID:%v", vid, pid))
			}

			var modelInfo types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &modelInfo)

			out, err := json.Marshal(modelInfo)
			if err != nil {
				fmt.Printf("Could not query model VID:%v PID:%v\n", vid, pid)
				return err
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, out, height))
		},
	}
}

func GetCmdAllModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-models",
		Short: "Query the list of all Models",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := parsePaginationParams(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_models", queryRoute), params)
			if err != nil {
				fmt.Printf("Could not query models: %s\n", err)
				return err
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, res, height))
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdVendors(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendors",
		Short: "Query the list of Vendors",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := parsePaginationParams(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vendors", queryRoute), params)
			if err != nil {
				fmt.Printf("Could not get query vendors: %s\n", err)
				return err
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, res, height))
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of vendors to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of vendors to take")

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
				return sdk.ErrInternal(fmt.Sprintf("Could not query vendor models VID:%v\n", vid))
			}

			var vendorProducts types.VendorProducts
			cdc.MustUnmarshalBinaryBare(res, &vendorProducts)

			out, err := json.Marshal(vendorProducts)
			if err != nil {
				fmt.Printf("Could not query vendor products VID:%v", vid)
				return err
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, out, height))
		},
	}
}

func parsePaginationParams(cdc *codec.Codec) []byte {
	params := types.NewPaginationParams(
		viper.GetInt(FlagSkip),
		viper.GetInt(FlagTake),
	)
	return cdc.MustMarshalJSON(params)
}
