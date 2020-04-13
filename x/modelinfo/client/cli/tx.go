package cli

import (
	"github.com/spf13/viper"
	"strconv"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const (
	FlagCID           = "cid"
	FlagCertificateId = "certificate-id"
	FlagCertifiedDate = "certified-date"
	FlagCustom        = "custom"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Modelinfo transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	modelinfoTxCmd.AddCommand(client.PostCommands(
		GetCmdAddModel(cdc),
		GetCmdUpdateModel(cdc),
		//GetCmdDeleteModel(cdc), Disable deletion
	)...)

	return modelinfoTxCmd
}

//nolint dupl
func GetCmdAddModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-model [vid] [pid] [name] [description] [sku] [firmware-version] [hardware-version] " +
			"[tis-or-trp-testing-completed]",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			vid, err := types.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := types.ParsePID(args[1])
			if err != nil {
				return err
			}

			name := args[2]

			description := args[3]
			sku := args[4]
			firmwareVersion := args[5]
			hardwareVersion := args[6]

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[7])
			if err != nil {
				return err
			}

			custom := viper.GetString(FlagCustom)
			certificateID := viper.GetString(FlagCertificateId)

			var certifiedDate time.Time
			if certifiedDateStr := viper.GetString(FlagCertifiedDate); len(certifiedDateStr) != 0 {
				certifiedDate, err = time.Parse(time.RFC3339, certifiedDateStr)
				if err != nil {
					return err
				}
			}

			var cid int16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = types.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgAddModelInfo(vid, pid, cid, name, description, sku, firmwareVersion, hardwareVersion,
				custom, certificateID, certifiedDate, tisOrTrpTestingCompleted, cliCtx.GetFromAddress())

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().String(FlagCustom, "", "Custom information")
	cmd.Flags().String(FlagCertificateId, "", "ID of certificate")
	cmd.Flags().String(FlagCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

//nolint dupl
func GetCmdUpdateModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-model [vid] [pid] [name] [description] [sku] [firmware-version] [hardware-version] " +
			"[tis-or-trp-testing-completed]",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			vid, err := types.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := types.ParsePID(args[1])
			if err != nil {
				return err
			}

			name := args[2]

			description := args[3]
			sku := args[4]
			firmwareVersion := args[5]
			hardwareVersion := args[6]

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[7])
			if err != nil {
				return err
			}

			custom := viper.GetString(FlagCustom)
			certificateID := viper.GetString(FlagCertificateId)

			var certifiedDate time.Time
			if certifiedDateStr := viper.GetString(FlagCertifiedDate); len(certifiedDateStr) != 0 {
				certifiedDate, err = time.Parse(time.RFC3339, certifiedDateStr)
				if err != nil {
					return err
				}
			}

			var cid int16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = types.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpdateModelInfo(vid, pid, cid, name, description, sku, firmwareVersion, hardwareVersion,
				custom, certificateID, certifiedDate, tisOrTrpTestingCompleted, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().String(FlagCustom, "", "Custom information")
	cmd.Flags().String(FlagCertificateId, "", "ID of certificate")
	cmd.Flags().String(FlagCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

func GetCmdDeleteModel(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-model [vid] [pid]",
		Short: "Delete existing ModelInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			vid, err := types.ParseVID(args[0])
			if err != nil {
				return err
			}

			pid, err := types.ParsePID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelInfo(vid, pid, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
