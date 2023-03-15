package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

var _ = strconv.Itoa(0)

func CmdDeleteModelVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model-version [vid] [pid] [software-version]",
		Short: "Broadcast message DeleteModelVersion",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelVersion(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
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
