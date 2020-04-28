package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagCID         = "cid"
	FlagDescription = "description"
	FlagCustom      = "custom"
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
		Use: "add-model [vid] [pid] [name] [description-string-or-path] [sku] [firmware-version] [hardware-version] " +
			"[tis-or-trp-testing-completed]",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(8),
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

			name := args[2]

			description, err_ := cliCtx.ReadFromFile(args[3])
			if err_ != nil {
				return err_
			}

			sku := args[4]
			firmwareVersion := args[5]
			hardwareVersion := args[6]

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(args[7])
			if err_ != nil {
				return sdk.ErrInternal(fmt.Sprintf("Invalid tis-or-trp-testing-completed: Parsing Error: %v must be boolean", tisOrTrpTestingCompleted))
			}

			custom := viper.GetString(FlagCustom)

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgAddModelInfo(vid, pid, cid, name, description, sku, firmwareVersion, hardwareVersion,
				custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().String(FlagCustom, "", "Custom information")

	return cmd
}

//nolint dupl
func GetCmdUpdateModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model [vid] [pid] [tis-or-trp-testing-completed]",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(3),
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

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(args[2])
			if err_ != nil {
				return err_
			}

			description := viper.GetString(FlagDescription)

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			custom := viper.GetString(FlagCustom)

			msg := types.NewMsgUpdateModelInfo(vid, pid, cid, description, custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().String(FlagDescription, "", "Model description")
	cmd.Flags().String(FlagCustom, "", "Custom information")

	return cmd
}

func GetCmdDeleteModel(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-model [vid] [pid]",
		Short: "Delete existing ModelInfo",
		Args:  cobra.ExactArgs(1),
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

			msg := types.NewMsgDeleteModelInfo(vid, pid, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}
