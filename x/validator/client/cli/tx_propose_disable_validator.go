package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdProposeDisableValidator() *cobra.Command {
	var (
		address string
		info    string
	)

	cmd := &cobra.Command{
		Use:   "propose-disable-validator --address [address] --info [info] [flags]",
		Short: "Proposes disabling of the Validator node by a Trustee.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddr := clientCtx.GetFromAddress()
			addr, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeDisableValidator(
				fromAddr,
				sdk.ValAddress(addr),
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

	cmd.Flags().StringVar(&address, FlagAddress, "", "Bech32 encoded validator address")
	cmd.Flags().StringVar(&info, FlagInfo, "", "Optional information/notes for approval or proposal validator")

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
