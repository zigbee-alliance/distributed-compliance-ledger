package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	// "strings".
	"github.com/spf13/cobra"

	// sdk "github.com/cosmos/cosmos-sdk/types".
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group compliance queries under a subcommand
	cmd := &cobra.Command{
		Use:                        dclcompltypes.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", dclcompltypes.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListComplianceInfo())
	cmd.AddCommand(CmdShowComplianceInfo())
	cmd.AddCommand(CmdListCertifiedModel())
	cmd.AddCommand(CmdShowCertifiedModel())
	cmd.AddCommand(CmdListRevokedModel())
	cmd.AddCommand(CmdShowRevokedModel())
	cmd.AddCommand(CmdListProvisionalModel())
	cmd.AddCommand(CmdShowProvisionalModel())
	cmd.AddCommand(CmdListDeviceSoftwareCompliance())
	cmd.AddCommand(CmdShowDeviceSoftwareCompliance())
	// this line is used by starport scaffolding # 1

	return cmd
}
