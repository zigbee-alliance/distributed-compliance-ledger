package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	// "strings".
	"github.com/spf13/cobra"

	// sdk "github.com/cosmos/cosmos-sdk/types".
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group validator queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListValidator())
	cmd.AddCommand(CmdShowValidator())

	// XXX issue 99: do we need that actually (auto-scaffolded)
	// TODO verify commands implementation if we need them
	cmd.AddCommand(CmdListLastValidatorPower())
	cmd.AddCommand(CmdShowLastValidatorPower())
	// this line is used by starport scaffolding # 1

	return cmd
}
