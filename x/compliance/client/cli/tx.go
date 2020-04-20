package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const (
	FlagCertificationType = "certification-type"
	FlagReason  = "reason"
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
		Use:   "certify-model [vid] [pid] [certification-date]",
		Short: "Certify an existing model. Note that the corresponding Model Info and test results must be present on ledger.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			vid, err := conversions.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(args[1])
			if err != nil {
				return err
			}

			certificationDate, err_ := time.Parse(time.RFC3339, args[2])
			if err_ != nil {
				return sdk.ErrInternal(fmt.Sprintf("Invalid certification-date: Parsing Error: %v must be RFC3339 date", err_))
			}

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))
			reason := viper.GetString(FlagReason)

			msg := types.NewMsgCertifyModel(vid, pid, certificationDate, certificationType, reason, cliCtx.GetFromAddress())

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Certification type. `zb` is the default and the only supported value now")
	cmd.Flags().String(FlagReason, "", "Optional comment describing the reason of certification")

	return cmd
}

//nolint dupl
func GetCmdRevokeModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-model [vid] [pid] [revocation-date]",
		Short: "Revoke compliance of an existing model.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			vid, err := conversions.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(args[1])
			if err != nil {
				return err
			}

			revocationDate, err_ := time.Parse(time.RFC3339, args[2])
			if err_ != nil {
				return sdk.ErrInternal(fmt.Sprintf("Invalid revocation-date: Parsing Error: %v must be RFC3339 date", err_))
			}

			certificationType := types.CertificationType(viper.GetString(FlagCertificationType))
			reason := viper.GetString(FlagReason)

			msg := types.NewMsgRevokeModel(vid, pid, revocationDate, certificationType, reason, cliCtx.GetFromAddress())

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagCertificationType, "", "Certification type. `zb` is the default and the only supported value now")
	cmd.Flags().String(FlagReason, "", "Optional comment describing the reason of revocation")

	return cmd
}
