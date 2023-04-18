package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags".
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
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

	cmd.AddCommand(CmdCreateModel())
	cmd.AddCommand(CmdUpdateModel())
	cmd.AddCommand(CmdDeleteModel())
	cmd.AddCommand(CmdCreateModelVersion())
	cmd.AddCommand(CmdUpdateModelVersion())
	cmd.AddCommand(CmdDeleteModelVersion())
	// this line is used by starport scaffolding # 1

	return cmd
}
