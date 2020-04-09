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
	FlagCertificateId    = "certificate-id"
	FlagCertifiedDate    = "certified-date"
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
		Use: "add-model-info [id] [name] [description] [sku] [firmware-version] [hardware-version] " +
			"[tis-or-trp-testing-completed]",
		Short: "add new ModelInfo",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			name := args[1]

			description := args[2]
			sku := args[3]
			firmwareVersion := args[4]
			hardwareVersion := args[5]

			tisOrTrpTestingCompleted, err := strconv.ParseBool(args[6])
			if err != nil {
				return err
			}

			certificateID := viper.GetString(FlagCertificateId)
			certifiedDateStr := viper.GetString(FlagCertifiedDate)

			var certifiedDate time.Time
			if len(certifiedDateStr) != 0 {
				certifiedDate, err = time.Parse(time.RFC3339, certifiedDateStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgAddModelInfo(id, name, description, sku, firmwareVersion, hardwareVersion,
				certificateID, certifiedDate, tisOrTrpTestingCompleted, cliCtx.GetFromAddress())

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagCertificateId, "", "ID of certificate")
	cmd.Flags().String(FlagCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

//nolint dupl
func GetCmdUpdateModelInfo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-model-info [id] [new-name] [new-owner] [new-description] [new-sku] [new-firmware-version] " +
			"[new-hardware-version] [new-tis-or-trp-testing-completed]",
		Short: "update existing ModelInfo",
		Args:  cobra.RangeArgs(9, 10),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]
			newName := args[1]

			newOwner, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			newDescription := args[3]
			newSku := args[4]
			newFirmwareVersion := args[5]
			newHardwareVersion := args[6]

			newTisOrTrpTestingCompleted, err := strconv.ParseBool(args[7])
			if err != nil {
				return err
			}

			newCertificateID := viper.GetString(FlagCertificateId)
			newCertifiedDateStr := viper.GetString(FlagCertifiedDate)

			var newCertifiedDate time.Time
			if len(newCertifiedDateStr) != 0 {
				newCertifiedDate, err = time.Parse(time.RFC3339, newCertifiedDateStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpdateModelInfo(id, newName, newOwner, newDescription, newSku, newFirmwareVersion,
				newHardwareVersion, newCertificateID, newCertifiedDate, newTisOrTrpTestingCompleted, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagNewCertificateId, "", "ID of certificate")
	cmd.Flags().String(FlagNewCertifiedDate, "", "Date of device certification in the RFC3339 format")

	return cmd
}

func GetCmdDeleteModelInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-model-info [id]",
		Short: "delete existing ModelInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			id := args[0]

			msg := types.NewMsgDeleteModelInfo(id, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
