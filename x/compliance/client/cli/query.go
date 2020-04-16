package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		GetCmdCertifiedModel(storeKey, cdc),
		GetCmdAllCertifiedModels(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdCertifiedModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "certified-model [vid] [pid]",
		Short: "Query certified model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vid := args[0]
			pid := args[1]

			res, height, err := cliCtx.QueryStore([]byte(keeper.CertifiedModelId(vid, pid)), queryRoute)
			if err != nil || res == nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not get certified model VID:%v PID:%v", vid, pid))
			}

			var certifiedModel types.CertifiedModel
			cdc.MustUnmarshalBinaryBare(res, &certifiedModel)

			out, err := json.Marshal(certifiedModel)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not get certified model VID:%v PID:%v", vid, pid))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}

func GetCmdAllCertifiedModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-certified-models",
		Short: "Query the list of all certified models",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := pagination.ParsePaginationParamsFromFlags(cdc)
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_certified_models", queryRoute), params)
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
