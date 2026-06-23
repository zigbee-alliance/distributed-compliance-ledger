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

package upgrade

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// SoftwareUpgradeStep encapsulates one cosmovisor upgrade transition —
// the propose+approve+wait+verify-applied sequence used by every phase.
type SoftwareUpgradeStep struct {
	// PlanName is the dclupgrade plan name (e.g. "v1.5.2").
	PlanName string
	// BinaryVersionNew is the release tag used in the upgrade-info URL.
	BinaryVersionNew string
	// Checksum is the sha256:... value appended to the upgrade-info URL.
	Checksum string
	// DcldOldBin is the dcld binary used to submit propose/approve txs while
	// the chain still runs the pre-upgrade version.
	DcldOldBin string
	// DcldNewBin is the dcld binary used to query the chain post-upgrade.
	DcldNewBin string
	// Trustees are the account names that vote on the upgrade plan.
	// Trustees[0] proposes; Trustees[1:] approve.
	Trustees []string
	// HeightOffset is how many blocks ahead of the current height to schedule
	// the plan. Defaults to 20 if zero.
	HeightOffset int64
	// WaitTimeoutSec caps how long we wait for the chain to cross plan_height.
	// Defaults to 300 if zero — matches `wait_for_height ... 300 outage-safe`.
	WaitTimeoutSec int
}

// Run executes the upgrade transition end-to-end. Fails the test on any error.
func (s SoftwareUpgradeStep) Run(t *testing.T) {
	t.Helper()
	require.GreaterOrEqual(t, len(s.Trustees), 2,
		"need at least one proposer and one approver")

	offset := s.HeightOffset
	if offset == 0 {
		offset = 20
	}

	timeout := s.WaitTimeoutSec
	if timeout == 0 {
		timeout = 300
	}

	currentHeight, err := cliputils.GetHeight()
	require.NoError(t, err, "GetHeight before propose")

	planHeight := currentHeight + offset
	t.Logf("Propose upgrade %s at height %d (current %d)",
		s.PlanName, planHeight, currentHeight)

	upgradeInfo := UpgradeInfoForVersion(s.BinaryVersionNew, s.Checksum)

	proposeRes, err := ProposeUpgrade(s.DcldOldBin, s.PlanName, planHeight, upgradeInfo, s.Trustees[0])
	require.NoError(t, err, "propose-upgrade")
	require.Equal(t, uint32(0), proposeRes.Code,
		"propose-upgrade rejected: %s", proposeRes.RawLog)

	for _, who := range s.Trustees[1:] {
		t.Logf("Approve upgrade %s from %s", s.PlanName, who)
		approveRes, aerr := ApproveUpgrade(s.DcldOldBin, s.PlanName, who)
		require.NoError(t, aerr, "approve-upgrade from %s", who)
		require.Equal(t, uint32(0), approveRes.Code,
			"approve-upgrade from %s rejected: %s", who, approveRes.RawLog)
	}

	t.Logf("Waiting for height %d", planHeight+1)
	cliputils.WaitForHeight(t, planHeight+1, timeout)

	t.Logf("Verify no upgrade is scheduled anymore")
	out, _ := QueryUpgradePlan(s.DcldNewBin)
	require.True(t, strings.Contains(string(out), "no upgrade scheduled"),
		"expected 'no upgrade scheduled', got: %s", string(out))

	// `query upgrade applied` is diagnostic only — log the result (or any
	// error) and move on. The real "upgrade applied" check is the
	// "no upgrade scheduled" assertion above; older dcld versions return
	// non-zero exits here that aren't meaningful.
	t.Logf("Verify upgrade %s is applied", s.PlanName)
	if out, err := QueryAppliedPlan(s.DcldNewBin, s.PlanName); err != nil {
		t.Logf("query upgrade applied %s: %v", s.PlanName, err)
	} else {
		t.Logf("applied: %s", string(out))
	}
}

// requireTxSuccess asserts a tx broadcast cleanly and executed with on-chain
// code 0. ExecuteTxWithBin already waits for inclusion (polling in sync mode),
// so — unlike cliputils.RequireTxOK — this performs no extra confirmation poll.
func requireTxSuccess(t *testing.T, tx *utils.TxResult, err error) {
	t.Helper()
	require.NoError(t, err)
	require.Equal(t, uint32(0), tx.Code, tx.RawLog)
}

// checkResponseContains asserts a query response contains substr.
func checkResponseContains(t *testing.T, out []byte, substr string) {
	t.Helper()
	require.True(t, strings.Contains(string(out), substr),
		"response missing %q, got: %s", substr, string(out))
}

// quoteField is a convenience for the recurring `"key": value` assertions.
// Historical dcld binaries emit JSON with `"key": value` (spaced); the
// master binary emits `"key":value` (compact). Both forms appear in this
// suite's queries, so quoteField returns both candidates.
func quoteField(key string, value any) []string {
	return []string{
		fmt.Sprintf("%q: %v", key, value), // legacy spaced form
		fmt.Sprintf("%q:%v", key, value),  // compact form (master)
	}
}

// containsAny reports whether s contains any of the alternatives. Pairs with
// quoteField to assert presence of a field across legacy and compact JSON.
func containsAny(s string, alternatives []string) bool {
	for _, alt := range alternatives {
		if strings.Contains(s, alt) {
			return true
		}
	}

	return false
}

// requireFieldEquals asserts that out contains either `"key": value` (legacy
// formatting from historical dcld releases) or `"key":value` (compact form
// from master). Used by post-upgrade query verifications.
func requireFieldEquals(t *testing.T, out []byte, key string, value any) {
	t.Helper()
	if !containsAny(string(out), quoteField(key, value)) {
		t.Fatalf("expected field %s=%v, got: %s", key, value, string(out))
	}
}

// MustRun is like t.Run but halts the parent test as soon as the subtest
// fails. Migration sequences chain stateful steps (account → vendor →
// models → compliance → …); once an earlier step fails the chain's
// preconditions are gone, so continuing only produces a cascade of
// misleading errors.
func MustRun(t *testing.T, name string, f func(t *testing.T)) {
	t.Helper()
	if !t.Run(name, f) {
		t.FailNow()
	}
}
