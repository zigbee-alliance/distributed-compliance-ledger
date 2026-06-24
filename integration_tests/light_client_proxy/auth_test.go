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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestLightClientProxyAuth exercises the dcld auth module against the
// light client proxy.
//
// Flow:
//
//  1. Generate two random users (user1 + user2). Both are bech32 addresses
//     only — no on-chain account exists yet.
//  2. Query non-existent records via the light client proxy → "Not Found".
//  3. Issue list (`all-*`) queries via the proxy → proxy rejects with the
//     "List queries don't work with a Light Client Proxy" payload.
//  4. Propose + approve user1's account against the full node.
//  5. Re-query the light client proxy → user1's record is now visible.
//  6. user2 is still "Not Found".
//  7. Attempt a write against the proxy → proxy rejects with the
//     "Write requests don't work with a Light Client Proxy" payload.
//
// Steps run sequentially; each step's preconditions come from the previous
// one (e.g. step 5 needs the propose/approve from step 4).
//
//nolint:funlen
func TestLightClientProxyAuth(t *testing.T) {
	skipIfDisabled(t)

	user1Name := utils.RandString()
	user1Addr, user1Pub, err := createKey(user1Name)
	require.NoError(t, err, "create key %s", user1Name)
	require.NotEmpty(t, user1Addr)
	require.NotEmpty(t, user1Pub)

	user2Name := utils.RandString()
	user2Addr, _, err := createKey(user2Name)
	require.NoError(t, err, "create key %s", user2Name)
	require.NotEmpty(t, user2Addr)

	// 1. Non-existent records via the proxy: every single-record query
	//    returns "Not Found".
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		for _, q := range []string{
			"account", "proposed-account", "proposed-account-to-revoke",
		} {
			out, qerr := queryWithRetry(LightClientProxyAddr,
				"query", "auth", q, "--address", user1Addr,
			)
			require.NoError(t, qerr, "query %s", q)
			require.True(t,
				strings.Contains(string(out), "Not Found"),
				"expected Not Found for %s, got: %s", q, string(out),
			)
		}
	})

	// 2. List queries are rejected by the proxy with the expected payload.
	mustRun(t, "ListQueries_Rejected", func(t *testing.T) {
		t.Helper()
		assertListQueriesRejected(t, "auth",
			"all-accounts", "all-proposed-accounts", "all-proposed-accounts-to-revoke")
	})

	// 3. Write user1's account against the full node: Jack proposes
	//    (NodeAdmin role), Alice approves.
	//
	//    --node FullNodeAddr is explicit so the suite doesn't depend on what
	//    ~/.dcl/config/client.toml currently points at — the upgrade suite
	//    flips it as part of its own flow.
	mustRun(t, "ProposeAndApprove_User1", func(t *testing.T) {
		t.Helper()
		tx, terr := utils.ExecuteTx(
			"tx", "auth", "propose-add-account",
			"--address", user1Addr,
			"--pubkey", user1Pub,
			"--roles", "NodeAdmin",
			"--from", "jack",
			"--node", FullNodeAddr,
		)
		require.NoError(t, terr)
		require.Equal(t, uint32(0), tx.Code, "propose-add-account: %s", tx.RawLog)

		tx, terr = utils.ExecuteTx(
			"tx", "auth", "approve-add-account",
			"--address", user1Addr,
			"--from", "alice",
			"--node", FullNodeAddr,
		)
		require.NoError(t, terr)
		require.Equal(t, uint32(0), tx.Code, "approve-add-account: %s", tx.RawLog)
	})

	// 4. user1 is now visible through the proxy.
	//    queryUntilContains polls through the proxy's post-write sync window —
	//    after Jack proposes and Alice approves, the proxy needs a few seconds
	//    to catch up to the new block (we poll up to 30s).
	mustRun(t, "Found_User1_AfterAdd", func(t *testing.T) {
		t.Helper()
		out, qerr := queryUntilContains(LightClientProxyAddr, user1Addr,
			"query", "auth", "account", "--address", user1Addr,
		)
		require.NoError(t, qerr)
		require.True(t,
			strings.Contains(string(out), user1Addr),
			"expected proxy to surface %s, got: %s", user1Addr, string(out),
		)
	})

	// 5. user2 was never proposed — proxy still says Not Found for every
	//    single-record query.
	mustRun(t, "NotFound_User2_AfterAdd", func(t *testing.T) {
		t.Helper()
		for _, q := range []string{
			"account", "proposed-account", "proposed-account-to-revoke",
		} {
			out, qerr := queryWithRetry(LightClientProxyAddr,
				"query", "auth", q, "--address", user2Addr,
			)
			require.NoError(t, qerr, "query %s", q)
			require.True(t,
				strings.Contains(string(out), "Not Found"),
				"expected Not Found for %s, got: %s", q, string(out),
			)
		}
	})

	// 6. Writes through the proxy are rejected with the expected payload.
	//    The tx is intentionally well-formed — the rejection comes from the
	//    proxy itself, not from cosmos-sdk message validation.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		out, _ := executeCLIWithNode(LightClientProxyAddr,
			"tx", "auth", "propose-add-account",
			"--address", user1Addr,
			"--pubkey", user1Pub,
			"--roles", "NodeAdmin",
			"--from", "jack",
			"--yes",
			"-o", "json",
			"--keyring-backend", "test",
		)
		// Exit code can be non-zero; ignore err. The rejection payload lands
		// in stdout/stderr either way.
		require.True(t,
			strings.Contains(string(out), writeRejection),
			"expected %q, got: %s", writeRejection, string(out),
		)
	})
}
