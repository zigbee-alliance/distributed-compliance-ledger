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

	cmd.AddCommand(CmdListApprovedCertificates()) // TODO: use store-based index of cert Ids
	cmd.AddCommand(CmdShowApprovedCertificates())
	cmd.AddCommand(CmdListProposedCertificate()) // TODO: use store-based index of cert Ids
	cmd.AddCommand(CmdShowProposedCertificate())
	cmd.AddCommand(CmdShowChildCertificates())
	cmd.AddCommand(CmdListProposedCertificateRevocation()) // TODO: use store-based index of cert Ids
	cmd.AddCommand(CmdShowProposedCertificateRevocation())
	cmd.AddCommand(CmdListRevokedCertificates()) // TODO: use store-based index of cert Ids
	cmd.AddCommand(CmdShowRevokedCertificates())
	cmd.AddCommand(CmdShowApprovedRootCertificates())
	cmd.AddCommand(CmdShowRevokedRootCertificates())
	cmd.AddCommand(CmdShowApprovedCertificatesBySubject())
	// this line is used by starport scaffolding # 1

	return cmd
}
