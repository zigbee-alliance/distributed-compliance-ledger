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
	var (
		name string
		info string
	)

	cmd := &cobra.Command{
		Use:   "approve-upgrade --name [name] --info [info] [flags]",
		Short: "Approve proposed upgrade with given name",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveUpgrade(
				clientCtx.GetFromAddress().String(),
				name,
				info,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().StringVar(&name, FlagName, "", "Upgrade name")
	cmd.Flags().StringVar(&info, FlagInfo, "", FlagInfoUsage)

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagName)

	return cmd
}
