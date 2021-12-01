package cli

import (
	"encoding/json"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdCreateValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator [pub-key] [description]",
		Short: "Broadcast message CreateValidator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPubKey := args[0]
			argDescription := new(types.Description)
			err = json.Unmarshal([]byte(args[1]), argDescription)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateValidator(
				clientCtx.GetFromAddress().String(),
				argPubKey,
				argDescription,
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
