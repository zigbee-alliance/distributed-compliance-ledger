package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
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
	// this line is used by starport scaffolding # 1

	return cmd
}
