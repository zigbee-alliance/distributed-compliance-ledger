package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func CmdProposeUpgrade() *cobra.Command {
	var (
		name          string
		upgradeHeight int64
		upgradeInfo   string
	)

	cmd := &cobra.Command{
		Use:   "propose-upgrade --name [name] --upgrade-height [height] --upgrade-info [info] [flags]",
		Short: "Propose Upgrade with given name at given height",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			plan := types.Plan{Name: name, Height: upgradeHeight, Info: upgradeInfo}

			msg := types.NewMsgProposeUpgrade(
				clientCtx.GetFromAddress().String(),
				plan,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().StringVar(&name, FlagName, "", "Upgrade name")
	cmd.Flags().Int64Var(&upgradeHeight, FlagUpgradeHeight, 0, "Height at which upgrade must be performed")
	cmd.Flags().StringVar(&upgradeInfo, FlagUpgradeInfo, "", "Any application specific upgrade info")

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagUpgradeHeight)

	return cmd
}
