package cli

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
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
				fmt.Printf("Model Not Found")
				return nil
			}

			var modelInfo types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &modelInfo)

			out, err := json.Marshal(modelInfo)
			if err != nil {
				fmt.Printf("Could not query model VID:%v PID:%v : %s\n", vid, pid, err)
				return nil
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
				return nil
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, res, height))
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of accounts to take")

	return cmd
}

func GetCmdVendors(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vendors",
		Short: "Query the list of Vendors",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := parsePaginationParams(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vendors", queryRoute), params)
			if err != nil {
				fmt.Printf("Could not get query vendors: %s\n", err)
				return nil
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, res, height))
		},
	}
}

func GetCmdVendorModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vendor-models [vid]",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			params := parsePaginationParams(cdc)

			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vendor_models/%s", queryRoute, vid), params)
			if err != nil {
				fmt.Printf("could not get query names: %s\n", err)
				return nil
			}

			return cliCtx.PrintOutput(types.NewReadResult(cdc, res, height))
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
