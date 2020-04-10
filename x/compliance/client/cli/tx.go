package cli

import (
	"github.com/spf13/viper"
	"strconv"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const (
	FlagCID              = "cid"
	FlagCertificateId    = "certificate-id"
	FlagCertifiedDate    = "certified-date"
	FlagNewCID           = "new-cid"
	FlagNewCertificateId = "new-certificate-id"
	FlagNewCertifiedDate = "new-certified-date"
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
		GetCmdAddModelInfo(cdc),
		GetCmdUpdateModelInfo(cdc),
		//GetCmdDeleteModelInfo(cdc), Disable deletion
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdAddModelInfo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-model-info [vid] [pid] [name] [description] [sku] [firmware-version] [hardware-version] " +
			"[custom] [tis-or-trp-testing-completed]",
		Short: "add new ModelInfo",
		Args:  cobra.ExactArgs(9),
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
			custom := args[7]

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[8])
			if err != nil {
				return err
			}

			certificateID := viper.GetString(FlagCertificateId)

			var certifiedDate time.Time
			if certifiedDateStr := viper.GetString(FlagCertifiedDate); len(certifiedDateStr) != 0 {
				certifiedDate, err = time.Parse(time.RFC3339, certifiedDateStr)
				if err != nil {
					return err
				}
			}

			var cid int16
			if cidStr := viper.GetString(FlagCertifiedDate); len(cidStr) != 0 {
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

	cmd.Flags().String(FlagCID, "", "Device CID")
	cmd.Flags().String(FlagCertificateId, "", "ModelInfoID of certificate")
	cmd.Flags().String(FlagCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

//nolint dupl
func GetCmdUpdateModelInfo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-model-info [vid] [pid] [new-name] [new-owner] [new-description] [new-sku] [new-firmware-version] " +
			"[new-hardware-version] [new-custom] [new-tis-or-trp-testing-completed]",
		Short: "update existing ModelInfo",
		Args:  cobra.ExactArgs(10),
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

			newName := args[2]

			newOwner, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return err
			}

			newDescription := args[4]
			newSku := args[5]
			newFirmwareVersion := args[6]
			newHardwareVersion := args[7]
			newCustom := args[8]

			newTisOrTrpTestingCompleted, err := strconv.ParseBool(args[9])
			if err != nil {
				return err
			}

			newCertificateID := viper.GetString(FlagNewCertificateId)

			var newCertifiedDate time.Time
			if newCertifiedDateStr := viper.GetString(FlagNewCertifiedDate); len(newCertifiedDateStr) != 0 {
				newCertifiedDate, err = time.Parse(time.RFC3339, newCertifiedDateStr)
				if err != nil {
					return err
				}
			}

			var newCid int16
			if cidStr := viper.GetString(FlagNewCID); len(cidStr) != 0 {
				newCid, err = types.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpdateModelInfo(vid, pid, newCid, newName, newOwner, newDescription, newSku, newFirmwareVersion,
				newHardwareVersion, newCustom, newCertificateID, newCertifiedDate, newTisOrTrpTestingCompleted, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagNewCID, "", "Device CID")
	cmd.Flags().String(FlagNewCertificateId, "", "ModelInfoID of certificate")
	cmd.Flags().String(FlagNewCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

func GetCmdDeleteModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-model-info [vid] [pid]",
		Short: "delete existing ModelInfo",
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
