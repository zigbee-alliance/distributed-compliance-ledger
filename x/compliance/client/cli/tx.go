package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	types "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	// "github.com/cosmos/cosmos-sdk/client/flags".
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCertifyModel())
	cmd.AddCommand(CmdRevokeModel())
	cmd.AddCommand(CmdProvisionModel())
	cmd.AddCommand(CmdUpdateComplianceInfo())
	cmd.AddCommand(CmdDeleteComplianceInfo())
	// this line is used by starport scaffolding # 1

	return cmd
}
