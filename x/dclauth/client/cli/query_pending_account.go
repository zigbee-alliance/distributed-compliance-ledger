package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func CmdListPendingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-accounts",
		Short: "list all PendingAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPendingAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PendingAccountAll(context.Background(), params)
			if cli.IsKeyNotFoundRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPendingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposed-account",
		Short: "shows a PendingAccount",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			argAddress, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			var res types.PendingAccount
			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.PendingAccountKeyPrefix,
				types.PendingAccountKey(argAddress),
				&res,
			)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bech32 encoded account address")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
