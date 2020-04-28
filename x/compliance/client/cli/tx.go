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

const (
	FlagCertificationType = "certification-type"
	FlagReason            = "reason"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliance transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(client.PostCommands(
		GetCmdCertifyModel(cdc),
		GetCmdRevokeModel(cdc),
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdCertifyModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certify-model [vid] [pid] [certification-type] [certification-date]",
		Short: "Certify an existing model. Note that the corresponding model info and test results must be present on ledger",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(args[1])
			if err != nil {
				return err
			}

			certificationType := types.CertificationType(args[2])

			certificationDate, err_ := time.Parse(time.RFC3339, args[3])
			if err_ != nil {
				return sdk.ErrInternal(fmt.Sprintf("Invalid certification-date: Parsing Error: %v must be RFC3339 date", err_))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgCertifyModel(vid, pid, certificationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagReason, "", "Optional comment describing the reason of certification")

	return cmd
}

//nolint dupl
func GetCmdRevokeModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-model [vid] [pid] [certification-type] [revocation-date]",
		Short: "Revoke compliance of an existing model",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(args[1])
			if err != nil {
				return err
			}

			certificationType := types.CertificationType(args[2])

			revocationDate, err_ := time.Parse(time.RFC3339, args[3])
			if err_ != nil {
				return sdk.ErrInternal(fmt.Sprintf("Invalid revocation-date: Parsing Error: %v must be RFC3339 date", err_))
			}

			reason := viper.GetString(FlagReason)

			msg := types.NewMsgRevokeModel(vid, pid, revocationDate, certificationType, reason, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagReason, "", "Optional comment describing the reason of revocation")

	return cmd
}
