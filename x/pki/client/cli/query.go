package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group pki queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListApprovedCertificates())
	cmd.AddCommand(CmdShowApprovedCertificates())
	cmd.AddCommand(CmdListProposedCertificate())
	cmd.AddCommand(CmdShowProposedCertificate())
	cmd.AddCommand(CmdListChildCertificates())
	cmd.AddCommand(CmdShowChildCertificates())
	cmd.AddCommand(CmdListProposedCertificateRevocation())
	cmd.AddCommand(CmdShowProposedCertificateRevocation())
	cmd.AddCommand(CmdListRevokedCertificates())
	cmd.AddCommand(CmdShowRevokedCertificates())
	// this line is used by starport scaffolding # 1

	return cmd
}
