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

	pkicli "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
)

// offlineQueryCtx returns a command context carrying a client context with no
// RPC node configured. Executing a query against it drives the RunE closure
// through its happy-path branches up to the (failing) query call, without
// requiring a running network.
func offlineQueryCtx() context.Context {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	clientCtx := client.Context{}.
		WithCodec(codec.NewProtoCodec(interfaceRegistry)).
		WithInterfaceRegistry(interfaceRegistry)

	return context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)
}

// execCmdOffline runs a command with an empty --node so the client context stays
// offline (no real network connection is attempted) plus any extra args.
func execCmdOffline(cmd *cobra.Command, extraArgs ...string) error {
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	cmd.SetArgs(append([]string{"--" + flags.FlagNode + "="}, extraArgs...))

	return cmd.ExecuteContext(offlineQueryCtx())
}

// TestCmdListCertificatesOffline covers the happy-path branches of the
// CmdListCertificates RunE closure up to the gRPC query call, which fails
// because no RPC node is configured.
func TestCmdListCertificatesOffline(t *testing.T) {
	err := execCmdOffline(pkicli.CmdListCertificates())
	require.Error(t, err)
	require.Contains(t, err.Error(), "no RPC client")
}

// TestCmdListCertificatesPaginationError covers the ReadPageRequest error branch:
// --page and --offset cannot be used together.
func TestCmdListCertificatesPaginationError(t *testing.T) {
	err := execCmdOffline(pkicli.CmdListCertificates(),
		"--"+flags.FlagPage+"=2",
		"--"+flags.FlagOffset+"=1",
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "page and offset cannot be used together")
}

// TestCmdListAllCertificatesBySubjectOffline covers the CmdListAllCertificatesBySubject
// RunE closure up to the QueryWithProof store query, which fails offline.
func TestCmdListAllCertificatesBySubjectOffline(t *testing.T) {
	err := execCmdOffline(pkicli.CmdListAllCertificatesBySubject(),
		"--"+pkicli.FlagSubject+"=test-subject",
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no RPC client")
}
