package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

var _ = strconv.Itoa(0)

func CmdAddTestingResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-testing-result [vid] [pid] [software-version] [software-version-string] [test-result] [test-date]",
		Short: "Broadcast message AddTestingResult",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint32E(args[2])
			if err != nil {
				return err
			}
			argSoftwareVersionString := args[3]
			argTestResult := args[4]
			argTestDate := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
