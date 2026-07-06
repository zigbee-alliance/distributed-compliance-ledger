package cli_test

import (
	"context"
	"io"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	dclupgradecli "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/client/cli"
)

// listUpgradeCmds maps each list command constructor to a readable name so the
// query-list RunE closures can be exercised with a single table of cases.
var listUpgradeCmds = map[string]func() *cobra.Command{
	"proposed": dclupgradecli.CmdListProposedUpgrade,
	"approved": dclupgradecli.CmdListApprovedUpgrade,
	"rejected": dclupgradecli.CmdListRejectedUpgrade,
}

// offlineQueryCtx returns a command context carrying a client context that has
// no RPC node configured. Executing a list query against it drives the RunE
// closure through GetClientQueryContext, ReadPageRequest, the query-client
// construction and the (failing) gRPC call without needing a running network.
func offlineQueryCtx() context.Context {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	clientCtx := client.Context{}.
		WithCodec(codec.NewProtoCodec(interfaceRegistry)).
		WithInterfaceRegistry(interfaceRegistry)

	return context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)
}

// execListCmd runs a list command with an empty --node so the client context
// stays offline (no real network connection is attempted) plus any extra args.
func execListCmd(cmd *cobra.Command, extraArgs ...string) error {
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	cmd.SetArgs(append([]string{"--" + flags.FlagNode + "="}, extraArgs...))

	return cmd.ExecuteContext(offlineQueryCtx())
}

// TestCmdListUpgradeOffline covers the happy-path branches of the list RunE
// closures up to the query call, which fails because no RPC node is configured.
func TestCmdListUpgradeOffline(t *testing.T) {
	for name, newCmd := range listUpgradeCmds {
		name, newCmd := name, newCmd
		t.Run(name, func(t *testing.T) {
			err := execListCmd(newCmd())
			require.Error(t, err)
			require.Contains(t, err.Error(), "no RPC client")
		})
	}
}

// TestCmdListUpgradePaginationError covers the ReadPageRequest error branch:
// --page and --offset cannot be used together.
func TestCmdListUpgradePaginationError(t *testing.T) {
	for name, newCmd := range listUpgradeCmds {
		name, newCmd := name, newCmd
		t.Run(name, func(t *testing.T) {
			err := execListCmd(newCmd(),
				"--"+flags.FlagPage+"=2",
				"--"+flags.FlagOffset+"=1",
			)
			require.Error(t, err)
			require.Contains(t, err.Error(), "page and offset cannot be used together")
		})
	}
}
