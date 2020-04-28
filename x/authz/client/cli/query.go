package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the authorization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	modelinfoQueryCmd.AddCommand(client.GetCommands(
		GetCmdAccountRoles(storeKey, cdc),
	)...)

	return modelinfoQueryCmd
}

func GetCmdAccountRoles(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "account-roles [addr]",
		Short: "Query roles by account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			addr := args[0]

			address, err := sdk.AccAddressFromBech32(addr)
			if err != nil {
				return sdk.ErrInvalidAddress(addr)
			}

			res, height, err := cliCtx.Context().QueryStore(address.Bytes(), queryRoute)
			if err != nil {
				fmt.Printf("could not query AccountRoles - %s \n", addr)
				return nil
			}

			var accountRoles types.AccountRoles
			cdc.MustUnmarshalBinaryBare(res, &accountRoles)

			return cliCtx.EncodeAndPrintWithHeight(accountRoles, height)
		},
	}
}
