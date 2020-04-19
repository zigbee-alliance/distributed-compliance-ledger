package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
		GetCmdGetComplianceInfo(storeKey, cdc),
		GetCmdGetAllComplianceInfos(storeKey, cdc),
		GetCmdGetCertifiedModel(storeKey, cdc),
		GetCmdGetAllCertifiedModels(storeKey, cdc),
		GetCmdGetRevokedModel(storeKey, cdc),
		GetCmdGetAllRevokedModels(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdGetComplianceInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compliance-info [vid] [pid]",
		Short: "Query compliance info by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfo(queryRoute, cdc, args, "")
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")

	return cmd
}

func GetCmdGetAllComplianceInfos(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-compliance-info-records",
		Short: "Query the list of all compliance info records",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllComplianceInfoRecords(cdc, fmt.Sprintf("custom/%s/all_compliance_info_records", queryRoute))
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdGetCertifiedModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certified-model [vid] [pid]",
		Short: "Query certified model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfo(queryRoute, cdc, args, types.Certified)
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")

	return cmd
}

func GetCmdGetAllCertifiedModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-certified-models",
		Short: "Query the list of all certified models",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllComplianceInfoRecords(cdc, fmt.Sprintf("custom/%s/all_certified_models", queryRoute))
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdGetRevokedModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoked-model [vid] [pid]",
		Short: "Query revoked model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfo(queryRoute, cdc, args, types.Revoked)
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")

	return cmd
}

func GetCmdGetAllRevokedModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-models",
		Short: "Query the list of all revoked models",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllComplianceInfoRecords(cdc, fmt.Sprintf("custom/%s/all_revoked_models", queryRoute))
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func getComplianceInfo(queryRoute string, cdc *codec.Codec, args []string, state types.ComplianceState) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	vid := args[0]
	pid := args[1]
	certificationType := types.CertificationType(viper.GetString(FlagCertificationType))
	res, height, err := cliCtx.QueryStore([]byte(keeper.ComplianceInfoId(certificationType, vid, pid)), queryRoute)
	if err != nil || res == nil {
		return types.ErrComplianceInfoDoesNotExist(vid, pid)
	}

	var complianceInfo types.ComplianceInfo
	cdc.MustUnmarshalBinaryBare(res, &complianceInfo)

	if len(state) != 0 && complianceInfo.State != state {
		return types.ErrComplianceInfoDoesNotExist(vid, pid)
	}

	out, err := json.Marshal(complianceInfo)
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could encode result: %v", err))
	}

	return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
}

func getAllComplianceInfoRecords(cdc *codec.Codec, path string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	paginationParams := pagination.ParsePaginationParamsFromFlags()
	certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

	params := types.NewListQueryParams(certificationType, paginationParams.Skip, paginationParams.Take)

	res, height, err := cliCtx.QueryWithData(path, cliCtx.Codec.MustMarshalJSON(params))
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not query compliance info records: %s\n", err))
	}

	return cliCtx.PrintOutput(cli.NewReadResult(cdc, res, height))
}
