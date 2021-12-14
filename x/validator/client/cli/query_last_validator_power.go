package cli

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListLastValidatorPower() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-last-powers",
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
		Use:   "last-powers",
		Short: "shows a LastValidatorPower",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			addr, err := sdk.ValAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			params := &types.QueryGetLastValidatorPowerRequest{
				Owner: addr.String(),
			}

			res, err := queryClient.LastValidatorPower(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded validator address")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
