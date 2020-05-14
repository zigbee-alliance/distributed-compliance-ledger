package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

//nolint dupl
func GetCmdCertifyModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certify-model",
		Short: "Certify an existing model. Note that the corresponding model info and test results must be present on ledger",
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

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

			certificationDate, err_ := time.Parse(time.RFC3339, viper.GetString(FlagCertificationDate))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid CertificationDate \"%v\": it must be RFC3339 date. Error: %v", viper.GetString(FlagRevocationDate), err_.Error()))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgCertifyModel(vid, pid, certificationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut,"", "Certification type (zb` is the only supported value now)")
	cmd.Flags().StringP(FlagCertificationDate, FlagCertificationDateShortcut, "", "The date of model certification (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "", "Optional comment describing the reason of certification")

	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)
	cmd.MarkFlagRequired(FlagCertificationType)
	cmd.MarkFlagRequired(FlagCertificationDate)

	return cmd
}

//nolint dupl
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

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))

			revocationDate, err_ := time.Parse(time.RFC3339, viper.GetString(FlagRevocationDate))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid CertificationDate \"%v\": it must be RFC3339 date. Error: %v", viper.GetString(FlagRevocationDate), err_.Error()))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgRevokeModel(vid, pid, revocationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "", "Certification type (zb` is the only supported value now)")
	cmd.Flags().StringP(FlagRevocationDate, FlagCertificationDateShortcut,"", "The date of model revocation (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "", "Optional comment describing the reason of revocation")

	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)
	cmd.MarkFlagRequired(FlagCertificationType)
	cmd.MarkFlagRequired(FlagRevocationDate)

	return cmd
}
