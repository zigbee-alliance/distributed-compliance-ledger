package cli

import (
	"fmt"

	"github.com/spf13/viper"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

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
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the compliance module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdModelInfo(storeKey, cdc),
		GetCmdModelInfoWithProof(storeKey, cdc),
		GetCmdModelInfoHeaders(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdModelInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "model-info [vid] [pid]",
		Short: "Query ModelInfo by combination of VID and PID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info/%s/%s", queryRoute, vid, pid), nil)
			if err != nil {
				fmt.Printf("could not query ModelInfo - %s/%s: %s\n", vid, pid, err)
				return nil
			}

			var out types.ModelInfo
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdModelInfoWithProof(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "model-info-with-proof [vid] [pid]",
		Short: "Query ModelInfo with proof by combination of VID and PID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/model_info/%s/%s", queryRoute, vid, pid))
			if err != nil {
				fmt.Printf("could not query ModelInfo - %s/%s: %s\n", vid, pid, err)
				return nil
			}

			var out types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdModelInfoHeaders(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model-info-headers",
		Short: "Query the list of all ModelInfo",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := types.NewQueryModelInfoHeadersParams(
				viper.GetInt(FlagSkip),
				viper.GetInt(FlagTake),
			)

			data := cdc.MustMarshalJSON(params)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info_headers", queryRoute), data)
			if err != nil {
				fmt.Printf("could not get query names: %s\n", err)
				return nil
			}

			var out types.QueryModelInfoHeadersResult
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of accounts to take")

	return cmd
}
