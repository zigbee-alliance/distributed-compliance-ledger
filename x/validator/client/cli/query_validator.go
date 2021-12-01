package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-validator",
		Short: "list all Validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorAll(context.Background(), params)
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
		Use:   "show-validator [owner]",
		Short: "shows a Validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]

			params := &types.QueryGetValidatorRequest{
				Owner: argOwner,
			}

			res, err := queryClient.Validator(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
