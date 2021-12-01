package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListLastValidatorPower() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-last-validator-power",
		Short: "list all LastValidatorPower",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLastValidatorPowerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LastValidatorPowerAll(context.Background(), params)
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

func CmdShowLastValidatorPower() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-last-validator-power [owner]",
		Short: "shows a LastValidatorPower",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]

			params := &types.QueryGetLastValidatorPowerRequest{
				Owner: argOwner,
			}

			res, err := queryClient.LastValidatorPower(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
