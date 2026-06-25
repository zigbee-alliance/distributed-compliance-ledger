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
func TestLightClientProxyAuth(t *testing.T) {
	skipIfDisabled(t)

	user1 := requireUserKey(t, "")
	user2 := requireUserKey(t, "")

	authSingleRecordCmds := []string{
		"account", "proposed-account", "proposed-account-to-revoke",
	}

	// 1. Non-existent records via the proxy: every single-record query
	//    returns "Not Found".
	mustRun(t, "NotFound_BeforeAdd", func(t *testing.T) {
		t.Helper()
		for _, cmd := range authSingleRecordCmds {
			assertNotFoundOnProxy(t, cmd, AuthByAddress(cmd, user1.address)...)
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
		tx, err := ProposeAddAccountArgs{
			Address: user1.address, Pubkey: user1.pubkey, Roles: "NodeAdmin",
		}.Send("jack")
		requireTxOK(t, tx, err, "propose-add-account")

		tx, err = ApproveAddAccountArgs{Address: user1.address}.Send("alice")
		requireTxOK(t, tx, err, "approve-add-account")
	})

	// 4. user1 is now visible through the proxy.
	//    queryUntilContains polls through the proxy's post-write sync window —
	//    after Jack proposes and Alice approves, the proxy needs a few seconds
	//    to catch up to the new block (we poll up to 30s).
	mustRun(t, "Found_User1_AfterAdd", func(t *testing.T) {
		t.Helper()
		out, qerr := queryUntilContains(LightClientProxyAddr, user1.address,
			AuthByAddress("account", user1.address)...,
		)
		require.NoError(t, qerr)
		require.True(t,
			strings.Contains(string(out), user1.address),
			"expected proxy to surface %s, got: %s", user1.address, string(out),
		)
	})

	// 5. user2 was never proposed — proxy still says Not Found for every
	//    single-record query.
	mustRun(t, "NotFound_User2_AfterAdd", func(t *testing.T) {
		t.Helper()
		for _, cmd := range authSingleRecordCmds {
			assertNotFoundOnProxy(t, cmd, AuthByAddress(cmd, user2.address)...)
		}
	})

	// 6. Writes through the proxy are rejected with the expected payload.
	//    The tx is intentionally well-formed — the rejection comes from the
	//    proxy itself, not from cosmos-sdk message validation.
	mustRun(t, "Write_Rejected", func(t *testing.T) {
		t.Helper()
		args := ProposeAddAccountArgs{
			Address: user1.address, Pubkey: user1.pubkey, Roles: "NodeAdmin",
		}.Build()
		args = append(args, "--from", "jack")
		assertWriteRejected(t, "propose-add-account", args...)
	})
}
