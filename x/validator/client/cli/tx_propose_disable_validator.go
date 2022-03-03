package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdProposeDisableValidator() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "propose-disable-validator --address [address]",
		Short: "Proposes disabling of the Validator node by a Trustee.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeDisableValidator(
				clientCtx.GetFromAddress().String(),
				address,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Bench32 encoded validator address or owner account")

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
