package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-nodes",
		Short: "list all Validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorAll(context.Background(), params)
			if cli.IsKeyNotFoundRPCError(err) {
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

func CmdShowValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "shows a Validator",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.ValAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				owner, err2 := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
				if err2 != nil {
					return err2
				}
				addr = sdk.ValAddress(owner)
			}

			var res types.Validator

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ValidatorKeyPrefix,
				types.ValidatorKey(addr),
				&res,
			)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bech32 encoded validator address or owner account")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
