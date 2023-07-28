package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	// "github.com/cosmos/cosmos-sdk/client/flags".
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        pkitypes.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", pkitypes.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdProposeAddX509RootCert())
	cmd.AddCommand(CmdApproveAddX509RootCert())
	cmd.AddCommand(CmdAddX509Cert())
	cmd.AddCommand(CmdProposeRevokeX509RootCert())
	cmd.AddCommand(CmdApproveRevokeX509RootCert())
	cmd.AddCommand(CmdRevokeX509Cert())
	cmd.AddCommand(CmdRejectAddX509RootCert())
	cmd.AddCommand(CmdAddPkiRevocationDistributionPoint())
	cmd.AddCommand(CmdUpdatePkiRevocationDistributionPoint())
	cmd.AddCommand(CmdDeletePkiRevocationDistributionPoint())
	cmd.AddCommand(CmdAssignVid())
	// this line is used by starport scaffolding # 1

	return cmd
}
