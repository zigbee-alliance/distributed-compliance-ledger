// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lightclientproxy

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	// FullNodeAddr is the localnet full node endpoint (node0). Writes must
	// be sent here — the light client proxy rejects them.
	FullNodeAddr = "tcp://localhost:26657"

	// LightClientProxyAddr is the light-client proxy endpoint exposed by
	// docker-compose. Single-record queries are served here; list queries
	// (`all-*`) and writes return a "doesn't work with a Light Client Proxy"
	// rejection payload (see listQueryRejection / writeRejection).
	LightClientProxyAddr = "tcp://localhost:26620"

	// queryWithRetry budget. The proxy emits transport-level errors during
	// its cold-start window — "EOF", "connection reset by peer",
	// "connection refused", "i/o timeout" — before its `cosmovisor run
	// light dclchain` command finishes verifying trust headers and binds
	// port 26620. Without retry, the first query of a fresh localnet races
	// the proxy startup and fails before it ever gets a chance.
	queryRetryAttempts = 15
	queryRetryDelay    = 2 * time.Second

	// listQueryRejection / writeRejection are the user-facing payloads the
	// proxy returns when it refuses to serve a request.
	listQueryRejection = "List queries don't work with a Light Client Proxy"
	writeRejection     = "Write requests don't work with a Light Client Proxy"
)

// runDcld shells out to `dcld <args...>`. Returns combined stdout+stderr.
// Errors include the full output to make CI logs actionable.
func runDcld(args ...string) ([]byte, error) {
	cmd := exec.Command("dcld", args...)
	// Some `keys add` paths prompt for a BIP39 passphrase; feeding two
	// newlines unblocks them. Same idiom the upgrade package uses.
	cmd.Stdin = bytes.NewBufferString("\n\n")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("dcld %s: %w, output: %s",
			strings.Join(args, " "), err, string(out))
	}

	return out, nil
}

// executeCLIWithNode runs `dcld <args...> --node <node>`. The light client
// proxy and the full node share an installed binary and global config; the
// only difference is which endpoint each invocation targets. We pass --node
// per call instead of mutating ~/.dcl/config/client.toml so concurrent suites
// don't trample each other.
func executeCLIWithNode(node string, args ...string) ([]byte, error) {
	args = append(args, "--node", node)

	return runDcld(args...)
}

// queryWithRetry executes a `dcld query …` against `node` and retries on
// transient transport errors from the light client proxy (EOF,
// connection reset/refused, i/o timeout — see needsRetry).
func queryWithRetry(node string, args ...string) ([]byte, error) {
	out, err := executeCLIWithNode(node, args...)
	for i := 0; i < queryRetryAttempts; i++ {
		// Either the subprocess errored out or the body itself carries the
		// retry token. Both happen when the proxy is warming up.
		if !needsRetry(out, err) {
			break
		}
		time.Sleep(queryRetryDelay)
		out, err = executeCLIWithNode(node, args...)
	}

	return out, err
}

// queryUntilContains polls the proxy until the response body contains
// `expected`, or the retry budget is exhausted. After a write to the full
// node the proxy needs to sync the new block's state before it can serve
// the new value, and during that window it returns "Not Found" — which
// queryWithRetry treats as a valid response, since other tests assert
// "Not Found" as a positive outcome. queryRetryAttempts × queryRetryDelay
// = 30s of total wait.
//
// Returns the final response (which may or may not contain `expected` if
// the budget runs out — let the caller decide whether to fail loudly).
func queryUntilContains(node, expected string, args ...string) ([]byte, error) {
	out, err := queryWithRetry(node, args...)
	for i := 0; i < queryRetryAttempts; i++ {
		if err == nil && strings.Contains(string(out), expected) {
			break
		}
		time.Sleep(queryRetryDelay)
		out, err = queryWithRetry(node, args...)
	}

	return out, err
}

// needsRetry reports whether queryWithRetry should try again — covers the
// transport-level errors the proxy emits during its cold-start window.
func needsRetry(out []byte, err error) bool {
	body := string(out)
	if err != nil {
		body += " " + err.Error()
	}

	for _, token := range []string{
		"EOF",
		"connection reset by peer",
		"connection refused",
		"i/o timeout",
	} {
		if strings.Contains(body, token) {
			return true
		}
	}

	return false
}

// createKey runs `dcld keys add <name> --keyring-backend test`, then `keys
// show -a` / `-p` to capture the bech32 address and pubkey JSON.
//
// The localnet keyring is the file-backed `test` backend (see
// genlocalnetconfig.sh: `dcld config keyring-backend test`), so we always
// pass --keyring-backend explicitly — relying on the global config would
// break if another concurrent test happened to flip it.
func createKey(name string) (address, pubkey string, err error) {
	if _, err = runDcld("keys", "add", name, "--keyring-backend", "test"); err != nil {
		return "", "", fmt.Errorf("keys add %s: %w", name, err)
	}

	addrOut, err := runDcld("keys", "show", name, "-a", "--keyring-backend", "test")
	if err != nil {
		return "", "", fmt.Errorf("keys show %s -a: %w", name, err)
	}

	pubOut, err := runDcld("keys", "show", name, "-p", "--keyring-backend", "test")
	if err != nil {
		return "", "", fmt.Errorf("keys show %s -p: %w", name, err)
	}

	return trimWS(addrOut), trimWS(pubOut), nil
}

func trimWS(b []byte) string {
	return strings.TrimRight(string(b), "\n\r \t")
}

// randomUint16 returns a random uint16 (0..32767). Used to generate
// VIDs/PIDs/software versions that don't collide with prior runs against
// the same localnet.
func randomUint16() int {
	var b [2]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(fmt.Errorf("rand.Read: %w", err))
	}

	return int(binary.BigEndian.Uint16(b[:]) >> 1)
}

// proposeVendorAccount creates a fresh key under `name` and proposes it
// as a Vendor with `vid` against the full node. No explicit approval is
// needed — Vendor uses the 1/3 quorum so jack's proposer vote already
// meets threshold on a 3-trustee genesis chain.
func proposeVendorAccount(t *testing.T, name string, vid int) (address string) {
	t.Helper()

	addr, pub, err := createKey(name)
	require.NoError(t, err, "create key %s", name)

	tx, err := utils.ExecuteTx(
		"tx", "auth", "propose-add-account",
		"--address", addr,
		"--pubkey", pub,
		"--roles", "Vendor",
		"--vid", fmt.Sprintf("%d", vid),
		"--from", "jack",
		"--node", FullNodeAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, "propose vendor %s: %s", name, tx.RawLog)

	return addr
}

// proposeAndApproveAccount creates a fresh key under `name`, proposes it
// with `roles` from jack, then has alice approve. Used for the
// CertificationCenter account in the compliance test. Two approvals
// satisfy the 2/3 quorum on the 3-trustee genesis chain.
func proposeAndApproveAccount(t *testing.T, name, roles string) (address string) {
	t.Helper()

	addr, pub, err := createKey(name)
	require.NoError(t, err, "create key %s", name)

	tx, err := utils.ExecuteTx(
		"tx", "auth", "propose-add-account",
		"--address", addr,
		"--pubkey", pub,
		"--roles", roles,
		"--from", "jack",
		"--node", FullNodeAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, "propose %s: %s", name, tx.RawLog)

	tx, err = utils.ExecuteTx(
		"tx", "auth", "approve-add-account",
		"--address", addr,
		"--from", "alice",
		"--node", FullNodeAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, "approve %s: %s", name, tx.RawLog)

	return addr
}

// addModelAndVersion adds a model + its first model-version against the
// full node.
func addModelAndVersion(t *testing.T, vid, pid, softwareVersion int, softwareVersionString, vendorAccount string) {
	t.Helper()

	tx, err := utils.ExecuteTx(
		"tx", "model", "add-model",
		"--vid", fmt.Sprintf("%d", vid),
		"--pid", fmt.Sprintf("%d", pid),
		"--deviceTypeID", "1",
		"--productName", "TestProduct",
		"--productLabel", "TestingProductLabel",
		"--partNumber", "1",
		"--commissioningCustomFlow", "0",
		"--enhancedSetupFlowOptions", "0",
		"--from", vendorAccount,
		"--node", FullNodeAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, "add-model %d/%d: %s", vid, pid, tx.RawLog)

	tx, err = utils.ExecuteTx(
		"tx", "model", "add-model-version",
		"--cdVersionNumber", "1",
		"--maxApplicableSoftwareVersion", "10",
		"--minApplicableSoftwareVersion", "1",
		"--vid", fmt.Sprintf("%d", vid),
		"--pid", fmt.Sprintf("%d", pid),
		"--softwareVersion", fmt.Sprintf("%d", softwareVersion),
		"--softwareVersionString", softwareVersionString,
		"--from", vendorAccount,
		"--node", FullNodeAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, "add-model-version %d/%d/%d: %s",
		vid, pid, softwareVersion, tx.RawLog)
}

// assertContains is a t.Helper wrapper around strings.Contains with a useful
// failure message. Used by every "expected Not Found / List queries don't
// work / Write requests don't work" assertion.
func assertContains(t *testing.T, out []byte, expected, label string) {
	t.Helper()
	require.True(t, strings.Contains(string(out), expected),
		"expected %q in %s, got: %s", expected, label, string(out))
}

// assertRejectionContains tolerates both the on-success and on-error paths
// for proxy rejections — the rejection payload may land in stdout/stderr
// regardless of subprocess exit code, so we coalesce both and check.
func assertRejectionContains(t *testing.T, out []byte, err error, expected, label string) {
	t.Helper()

	text := string(out)
	if err != nil {
		text += " " + err.Error()
	}
	require.True(t, strings.Contains(text, expected),
		"expected %q in %s, got: %s", expected, label, text)
}

// mustRun is `t.Run` + `t.FailNow()` on failure, so the cascade halts at
// the first failure instead of producing misleading follow-on errors.
// Same pattern as upgrade.MustRun; ported here because each Test* in this
// package is a linear chain (NotFound → Seed → Found → Write_Rejected)
// and the later steps depend on earlier ones writing to chain state.
func mustRun(t *testing.T, name string, f func(t *testing.T)) {
	t.Helper()
	if !t.Run(name, f) {
		t.FailNow()
	}
}

// containsAnyLocal reports whether `out` contains any of the alternatives.
// Used to accept both spaced (`"key": value`) and compact (`"key":value`) JSON
// formats — historical dcld releases emit the spaced form, master emits the
// compact form, and tests should tolerate either.
func containsAnyLocal(out []byte, alternatives ...string) bool {
	s := string(out)
	for _, alt := range alternatives {
		if strings.Contains(s, alt) {
			return true
		}
	}

	return false
}
