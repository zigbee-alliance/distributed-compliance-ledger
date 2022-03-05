package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListProposedDisableValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-disable-validators",
		Short: "Query the list of all proposed disable validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposedDisableValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedDisableValidatorAll(context.Background(), params)
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

func CmdShowProposedDisableValidator() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "proposed-disable-validator --address [address]",
		Short: "Query proposed disable validator by address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.ProposedDisableValidator

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.ProposedDisableValidatorKeyPrefix,
				types.ProposedDisableValidatorKey(address),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Validator address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}