package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdProposeAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-account [address] [pub-key] [roles] [vendor-id]",
		Short: "Broadcast message ProposeAddAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argPubKey := args[1]
			argRoles := strings.Split(args[2], listSeparator)
			argVendorID, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeAddAccount(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argPubKey,
				argRoles,
				argVendorID,
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
