package cli

import (
	"context"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListValidatorMissedBlockBitArray() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-validator-missed-block-bit-array",
		Short: "list all ValidatorMissedBlockBitArray",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorMissedBlockBitArrayRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorMissedBlockBitArrayAll(context.Background(), params)
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

func CmdShowValidatorMissedBlockBitArray() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-missed-block-bit-array [address] [index]",
		Short: "shows a ValidatorMissedBlockBitArray",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argIndex, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetValidatorMissedBlockBitArrayRequest{
				Address: argAddress,
				Index:   argIndex,
			}

			res, err := queryClient.ValidatorMissedBlockBitArray(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
