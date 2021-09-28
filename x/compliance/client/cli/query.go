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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
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
		Use:   "compliance-info",
		Short: "Query compliance info for Model (identified by the `vid`, `pid` and `certification_type`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfo(queryRoute, cdc)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Certification type (zb` is the only supported value now)")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)

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

	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdGetCertifiedModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "certified-model",
		Short: "Gets a boolean if the given Model (identified by the `vid`, `pid`, `softwareVersion` and " +
			"`certification_type`) is compliant to ZB standards",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfoInState(queryRoute, cdc, types.CodeCertified)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")

	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Certification type (zb` is the only supported value now)")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)

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

	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdGetRevokedModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoked-model",
		Short: "Gets a boolean if the given Model (identified by the `vid`, `pid` and `certification_type`) is revoked",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getComplianceInfoInState(queryRoute, cdc, types.CodeRevoked)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Certification type (zb` is the only supported value now)")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)

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

	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Requested certification type. `zb` is the default and the only supported value now")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func getComplianceInfo(queryRoute string, cdc *codec.Codec) error {
	cliCtx := cli.NewCLIContext().WithCodec(cdc)

	vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
	if err_ != nil {
		return err_
	}

	pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
	if err_ != nil {
		return err_
	}

	softwareVersion, err_ := conversions.ParseUInt32FromString("SoftwareVersion", viper.GetString(FlagSoftwareVersion))
	if err_ != nil {
		return err_
	}

	certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

	res, height, err := cliCtx.QueryStore(types.GetComplianceInfoKey(certificationType, vid, pid, softwareVersion), queryRoute)
	if err != nil || res == nil {
		return types.ErrComplianceInfoDoesNotExist(vid, pid, softwareVersion, certificationType)
	}

	var complianceInfo types.ComplianceInfo

	cdc.MustUnmarshalBinaryBare(res, &complianceInfo)

	return cliCtx.EncodeAndPrintWithHeight(complianceInfo, height)
}

func getComplianceInfoInState(queryRoute string, cdc *codec.Codec, status types.SoftwareVersionCertificationStatus) error {
	cliCtx := cli.NewCLIContext().WithCodec(cdc)

	vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
	if err_ != nil {
		return err_
	}

	pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
	if err_ != nil {
		return err_
	}

	softwareVersion, err_ := conversions.ParseUInt32FromString("SoftwareVersion", viper.GetString(FlagSoftwareVersion))
	if err_ != nil {
		return err_
	}
	certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

	isInState := types.ComplianceInfoInState{Value: false}

	res, height, err := cliCtx.QueryStore(types.GetComplianceInfoKey(certificationType, vid, pid, softwareVersion), queryRoute)
	if res != nil {
		var complianceInfo types.ComplianceInfo

		cdc.MustUnmarshalBinaryBare(res, &complianceInfo)

		isInState.Value = complianceInfo.SoftwareVersionCertificationStatus == status
	}

	if err != nil {
		return types.ErrComplianceInfoDoesNotExist(vid, pid, softwareVersion, certificationType)
	}

	return cliCtx.EncodeAndPrintWithHeight(isInState, height)
}

func getAllComplianceInfoRecords(cdc *codec.Codec, path string) error {
	cliCtx := cli.NewCLIContext().WithCodec(cdc)

	paginationParams := pagination.ParsePaginationParamsFromFlags()
	certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

	params := types.NewListQueryParams(certificationType, paginationParams.Skip, paginationParams.Take)

	return cliCtx.QueryList(path, params)
}
