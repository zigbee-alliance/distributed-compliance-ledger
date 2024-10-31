package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"

	// "strings".
	"github.com/spf13/cobra"
	// sdk "github.com/cosmos/cosmos-sdk/types".
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(_ string) *cobra.Command {
	// Group pki queries under a subcommand
	cmd := &cobra.Command{
		Use:                        pkitypes.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", pkitypes.ModuleName),
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
	cmd.AddCommand(CmdListRejectedCertificate())
	cmd.AddCommand(CmdShowRejectedCertificate())
	cmd.AddCommand(CmdListPkiRevocationDistributionPoint())
	cmd.AddCommand(CmdShowPkiRevocationDistributionPoint())
	cmd.AddCommand(CmdShowPkiRevocationDistributionPointsByIssuerSubjectKeyID())
	cmd.AddCommand(CmdListNocRootCertificates())
	cmd.AddCommand(CmdShowNocRootCertificates())
	cmd.AddCommand(CmdShowNocCertificatesByVidAndSkid())
	cmd.AddCommand(CmdListNocCertificates())
	cmd.AddCommand(CmdShowNocCertificates())
	cmd.AddCommand(CmdShowNocCertificatesBySubject())
	cmd.AddCommand(CmdListNocIcaCertificates())
	cmd.AddCommand(CmdShowNocIcaCertificates())
	cmd.AddCommand(CmdListRevokedNocRootCertificates())
	cmd.AddCommand(CmdShowRevokedNocRootCertificates())
	cmd.AddCommand(CmdListCertificates())
	cmd.AddCommand(CmdShowCertificates())
	cmd.AddCommand(CmdListRevokedNocIcaCertificates())
	cmd.AddCommand(CmdShowRevokedNocIcaCertificates())
	// this line is used by starport scaffolding # 1

	return cmd
}
