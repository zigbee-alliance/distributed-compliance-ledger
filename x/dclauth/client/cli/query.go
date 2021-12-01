package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group dclauth queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdShowAccount())
	cmd.AddCommand(CmdListPendingAccount())
	cmd.AddCommand(CmdListPendingAccountRevocation())

	// TODO issue 99: do we actually need that
	cmd.AddCommand(CmdShowAccountStat())

	// TODO issue 99: do we need the following ones ???
	// cmd.AddCommand(CmdShowPendingAccount())
	// cmd.AddCommand(CmdShowPendingAccountRevocation())
	// this line is used by starport scaffolding # 1

	return cmd
}
