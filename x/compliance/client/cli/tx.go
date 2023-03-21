package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	// "github.com/cosmos/cosmos-sdk/client/flags".
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        dclcompltypes.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", dclcompltypes.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCertifyModel())
	cmd.AddCommand(CmdRevokeModel())
	cmd.AddCommand(CmdProvisionModel())
	// this line is used by starport scaffolding # 1

	return cmd
}
