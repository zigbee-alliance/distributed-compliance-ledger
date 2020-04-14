package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authnextTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authentication extensions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authnextTxCmd.AddCommand(client.PostCommands()...)

	return authnextTxCmd
}
