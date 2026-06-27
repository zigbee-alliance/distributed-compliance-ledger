package cli_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	dclgenutilcli "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/client/cli"
)

// validMnemonic is a well-formed bip39 mnemonic used to exercise the --recover path.
const validMnemonic = "decide praise business actor peasant farm drastic weather extend front hurt later song give verb rhythm worry fun pond reform school tumble august one"

var testMbm = module.NewBasicManager(
	staking.AppModuleBasic{},
	genutil.AppModuleBasic{},
)

// initTestSetup wires up the client/server contexts the InitCmd RunE closure
// pulls out of the command context. chainID, when non-empty, is set on the
// client context to exercise the chain-id resolution branch.
func initTestSetup(t *testing.T, chainID string) (context.Context, string) {
	t.Helper()

	home := t.TempDir()
	cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, tmlog.NewNopLogger())

	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	clientCtx := client.Context{}.
		WithCodec(marshaler).
		WithHomeDir(home)
	if chainID != "" {
		clientCtx = clientCtx.WithChainID(chainID)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

	return ctx, home
}

// TestInitCmdDefaultChainID covers the happy path where no chain-id is provided,
// so a random one is generated.
func TestInitCmdDefaultChainID(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})

	require.NoError(t, cmd.ExecuteContext(ctx))
	require.FileExists(t, filepath.Join(home, "config", "genesis.json"))
	require.FileExists(t, filepath.Join(home, "config", "config.toml"))
}

// TestInitCmdChainIDFromClientCtx covers the branch where the chain-id is taken
// from the client context.
func TestInitCmdChainIDFromClientCtx(t *testing.T) {
	ctx, home := initTestSetup(t, "dclchain-1")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})

	require.NoError(t, cmd.ExecuteContext(ctx))

	genFile := filepath.Join(home, "config", "genesis.json")
	content, err := os.ReadFile(genFile)
	require.NoError(t, err)
	require.Contains(t, string(content), "dclchain-1")
}

// TestInitCmdChainIDFromFlag covers the branch where the chain-id is taken from
// the --chain-id flag, which takes precedence over the client context value.
func TestInitCmdChainIDFromFlag(t *testing.T) {
	ctx, home := initTestSetup(t, "ignored-chain")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{
		"appnode-test",
		fmt.Sprintf("--%s=flag-chain-9", flags.FlagChainID),
	})

	require.NoError(t, cmd.ExecuteContext(ctx))

	content, err := os.ReadFile(filepath.Join(home, "config", "genesis.json"))
	require.NoError(t, err)
	require.Contains(t, string(content), "flag-chain-9")
	require.NotContains(t, string(content), "ignored-chain")
}

// TestInitCmdRecoverValidMnemonic covers the --recover path with a valid mnemonic.
func TestInitCmdRecoverValidMnemonic(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetIn(strings.NewReader(validMnemonic + "\n"))
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	cmd.SetArgs([]string{
		"appnode-test",
		fmt.Sprintf("--%s=true", dclgenutilcli.FlagRecover),
	})

	require.NoError(t, cmd.ExecuteContext(ctx))
	require.FileExists(t, filepath.Join(home, "config", "genesis.json"))
}

// TestInitCmdRecoverInvalidMnemonic covers the --recover path where the supplied
// mnemonic is invalid and the command errors out.
func TestInitCmdRecoverInvalidMnemonic(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetIn(strings.NewReader("not a valid mnemonic\n"))
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	cmd.SetArgs([]string{
		"appnode-test",
		fmt.Sprintf("--%s=true", dclgenutilcli.FlagRecover),
	})

	err := cmd.ExecuteContext(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid mnemonic")
}

// TestInitCmdGenesisExists covers the guard that refuses to overwrite an
// existing genesis.json when --overwrite is not set.
func TestInitCmdGenesisExists(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})
	require.NoError(t, cmd.ExecuteContext(ctx))

	// Second run without --overwrite must fail because genesis.json now exists.
	cmd = dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "genesis.json file already exists")
}

// TestInitCmdOverwrite covers the --overwrite path, which reads the existing
// genesis doc from file before re-exporting it.
func TestInitCmdOverwrite(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})
	require.NoError(t, cmd.ExecuteContext(ctx))

	cmd = dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{
		"appnode-test",
		fmt.Sprintf("--%s=true", dclgenutilcli.FlagOverwrite),
	})
	require.NoError(t, cmd.ExecuteContext(ctx))
	require.FileExists(t, filepath.Join(home, "config", "genesis.json"))
}

// TestInitCmdOverwriteCorruptGenesis covers the path where an existing genesis
// file is present but cannot be parsed, so reading the genesis doc fails.
func TestInitCmdOverwriteCorruptGenesis(t *testing.T) {
	ctx, home := initTestSetup(t, "")

	cmd := dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{"appnode-test"})
	require.NoError(t, cmd.ExecuteContext(ctx))

	// Corrupt the genesis file so GenesisDocFromFile fails on the next run.
	genFile := filepath.Join(home, "config", "genesis.json")
	require.NoError(t, os.WriteFile(genFile, []byte("{ not valid json"), 0o600))

	cmd = dclgenutilcli.InitCmd(testMbm, home)
	cmd.SetArgs([]string{
		"appnode-test",
		fmt.Sprintf("--%s=true", dclgenutilcli.FlagOverwrite),
	})
	err := cmd.ExecuteContext(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Failed to read genesis doc from file")
}
