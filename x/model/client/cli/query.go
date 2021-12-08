package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group model queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListVendorProducts())
	cmd.AddCommand(CmdShowVendorProducts())
	cmd.AddCommand(CmdListModel())
	cmd.AddCommand(CmdShowModel())
	cmd.AddCommand(CmdListModelVersion())
	cmd.AddCommand(CmdShowModelVersion())
	cmd.AddCommand(CmdListModelVersions())
	cmd.AddCommand(CmdShowModelVersions())
	// this line is used by starport scaffolding # 1

	return cmd
}
