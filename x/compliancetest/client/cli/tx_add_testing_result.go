package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

var _ = strconv.Itoa(0)

func CmdAddTestingResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-test-result",
		Short: "Add new testing result for Model (identified by the `vid`, `pid` and `SoftwareVersion`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argVid, err := cast.ToInt32E(viper.GetString(FlagVid))
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(viper.GetString(FlagPid))
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint32E(viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}
			argSoftwareVersionString := viper.GetString(FlagSoftwareVersionString)
			argTestResult, err_ := cli.ReadFromFile(viper.GetString(FlagTestResult))
			if err_ != nil {
				return err_
			}
			argTestDate := viper.GetString(FlagTestDate)

			msg := types.NewMsgAddTestingResult(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
				argSoftwareVersionString,
				argTestResult,
				argTestDate,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagVid, "", "Model vendor ID")
	cmd.Flags().String(FlagPid, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().String(FlagSoftwareVersionString, "", "Model software version string")
	cmd.Flags().StringP(FlagTestResult, FlagTestResultShortcut, "",
		"Test result (string or path to file containing data)")
	cmd.Flags().StringP(FlagTestDate, FlagTestDateShortcut, "", "Date of test result (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z")

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagTestResult)
	_ = cmd.MarkFlagRequired(FlagTestDate)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	cli.AddTxFlagsToCmd(cmd)

	return cmd
}
