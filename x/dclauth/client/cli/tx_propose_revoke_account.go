package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

var _ = strconv.Itoa(0)

func CmdProposeRevokeAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-revoke-account",
		Short: "Broadcast message ProposeRevokeAccount",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argInfo := viper.GetString(FlagInfo)

			msg := types.NewMsgProposeRevokeAccount(
				clientCtx.GetFromAddress(),
				argAddress,
				argInfo,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bech32 encoded account address")
	cmd.Flags().String(FlagInfo, "", FlagInfoUsage)
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
