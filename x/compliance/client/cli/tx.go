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
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliance transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdCertifyModel(cdc),
		GetCmdRevokeModel(cdc),
	)...)...)

	return complianceTxCmd
}

func GetCmdCertifyModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "certify-model",
		Short: "Certify an existing model. Note that the corresponding model version and " +
			"test results must be present on ledger",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			softwareVersion, err := conversions.ParseUInt32FromString("SoftwareVersion", viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}

			softwareVersionString := viper.GetString(FlagSoftwareVersionString)

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

			certificationDate, err_ := time.Parse(time.RFC3339, viper.GetString(FlagCertificationDate))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid CertificationDate \"%v\": "+
					"it must be RFC3339 date. Error: %v", viper.GetString(FlagRevocationDate), err_.Error()))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgCertifyModel(vid, pid, softwareVersion, softwareVersionString, certificationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().String(FlagSoftwareVersionString, "", "Model software version string")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Certification type (zb` is the only supported value now)")
	cmd.Flags().StringP(FlagCertificationDate, FlagCertificationDateShortcut, "",
		"The date of model certification (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "",
		"Optional comment describing the reason of certification")
	cmd.Flags().String(FlagIsProvisional, "",
		"boolean flag to specify if this is only a provisional certification. This is set to true for a SoftwareVersion when going into certification testing")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagCertificationDate)

	return cmd
}

func GetCmdRevokeModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-model",
		Short: "Revoke compliance of an existing model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			softwareVersion, err := conversions.ParseUInt32FromString("SoftwareVersion", viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

			revocationDate, err_ := time.Parse(time.RFC3339, viper.GetString(FlagRevocationDate))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid CertificationDate \"%v\": "+
					"it must be RFC3339 date. Error: %v", viper.GetString(FlagRevocationDate), err_.Error()))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgRevokeModel(vid, pid, softwareVersion, revocationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")

	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "",
		"Certification type (zb` is the only supported value now)")
	cmd.Flags().StringP(FlagRevocationDate, FlagCertificationDateShortcut, "",
		"The date of model revocation (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "",
		"Optional comment describing the reason of revocation")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagRevocationDate)

	return cmd
}
