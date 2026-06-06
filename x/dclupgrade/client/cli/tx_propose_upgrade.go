// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		info          string
	)

	cmd := &cobra.Command{
		Use:   "propose-upgrade --name [name] --upgrade-height [upgrade-height] --upgrade-info [upgrade-info] --info [info] [flags]",
		Short: "Propose upgrade with given name at given height",
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
				info,
			)

			err = msg.ValidateBasicCLI()

			if err == nil {
				// validate basic will be called in GenerateOrBroadcastTxCLI
				err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			}

			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().StringVar(&name, FlagName, "", "Upgrade name")
	cmd.Flags().Int64Var(&upgradeHeight, FlagUpgradeHeight, 0, "Height at which upgrade must be performed")
	cmd.Flags().StringVar(&upgradeInfo, FlagUpgradeInfo, "", "Any application specific upgrade info")
	cmd.Flags().StringVar(&info, FlagInfo, "", FlagInfoUsage)

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagUpgradeHeight)

	return cmd
}
