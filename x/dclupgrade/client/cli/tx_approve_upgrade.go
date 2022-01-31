package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func CmdApproveUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-upgrade [name]",
		Short: "Approve proposed upgrade with given name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			name := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveUpgrade(
				clientCtx.GetFromAddress().String(),
				name,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
