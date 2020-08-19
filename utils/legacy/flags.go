package legacy

import "github.com/spf13/cobra"
import "github.com/cosmos/cosmos-sdk/client/flags"

// PostCommands adds common flags for commands to post tx
//
// Removed from `client/flags/flags.go` in Cosmos
func PostCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		flags.AddTxFlagsToCmd(c)
	}
	return cmds
}