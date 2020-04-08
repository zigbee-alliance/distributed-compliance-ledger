package cli

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authorization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccountRoles(storeKey, cdc),
		GetCmdAccountRolesWithProof(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdAccountRoles(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "account-roles <addr>",
		Short: "Query AccountRoles by account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/account_roles/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not query AccountRoles - %s \n", addr)
				return nil
			}

			var out types.AccountRoles
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdAccountRolesWithProof(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "account-roles-with-proof <addr>",
		Short: "Query AccountRoles with proof by account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, _, err := cliCtx.QueryStore([]byte(addr), queryRoute)
			if err != nil {
				fmt.Printf("could not query AccountRoles - %s \n", addr)
				return nil
			}

			var out types.AccountRoles
			cdc.MustUnmarshalBinaryBare(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
