package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	// "strings".
	"github.com/spf13/cobra"

	// sdk "github.com/cosmos/cosmos-sdk/types".
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group dclauth queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.CmdName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdShowAccount())
	cmd.AddCommand(CmdListPendingAccount())
	cmd.AddCommand(CmdListPendingAccountRevocation())

	cmd.AddCommand(CmdShowPendingAccount())
	cmd.AddCommand(CmdShowPendingAccountRevocation())
	cmd.AddCommand(CmdListRevokedAccount())
	cmd.AddCommand(CmdShowRevokedAccount())
	// this line is used by starport scaffolding # 1

	return cmd
}
