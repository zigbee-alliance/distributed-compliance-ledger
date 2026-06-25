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
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// runUpgradeFlowEnv gates the destructive Docker-orchestrating subtests. CI
// sets it (via run-all.sh upgrade_go); local dev keeps the test set as a
// fast-skipping suite unless the developer opts in.
const runUpgradeFlowEnv = "RUN_UPGRADE_GO"

// shouldRunUpgradeFlow reports whether the suite should attempt the Docker /
// pool-touching parts of the migration. When false, every subtest skips with
// a clear message so `go test ./...` stays fast.
func shouldRunUpgradeFlow() bool {
	return os.Getenv(runUpgradeFlowEnv) == "1"
}

// TestMain runs once per package. The bash helpers (pool.sh, common.sh) and
// the fixture paths under integration_tests/constants/... are all
// repo-root-relative, matching how run-all.sh invokes everything. `go test`
// runs with the package directory as CWD, so we chdir to the repo root here
// so every downstream relative path resolves the same as it does under
// run-all.sh.
func TestMain(m *testing.M) {
	if err := chdirToRepoRoot(); err != nil {
		fmt.Fprintf(os.Stderr, "chdir to repo root: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func chdirToRepoRoot() error {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("runtime.Caller failed")
	}
	root, err := filepath.Abs(filepath.Join(filepath.Dir(thisFile), "..", ".."))
	if err != nil {
		return err
	}

	return os.Chdir(root)
}

// TestUpgradeSequence is the single entry point for the upgrade migration.
// Subtests run in order against a shared chain. Each subtest adds more.
func TestUpgradeSequence(t *testing.T) {
	if !shouldRunUpgradeFlow() {
		t.Skipf("set %s=1 to run the full upgrade migration sequence", runUpgradeFlowEnv)
	}

	// Download every historical binary referenced by the migration sequence.
	require.NoError(t, EnsureAllBinaries(), "ensure historical binaries")

	// Start localnet at v0.12.0 — the starting point of the upgrade chain.
	require.NoError(t, InitPool(InitPoolOpts{
		PatchConfig:   "yes",
		InitTarget:    "localnet_init_latest_stable_release",
		BinaryVersion: BinaryPath("0.12.0"),
	}), "init_pool v0.12.0")

	t.Cleanup(func() {
		_ = CleanupPool()
	})

	// The validator-demo container is created in the first subtest and reused
	// through every later subtest, so it must be torn down at the end of the
	// whole sequence (not when that subtest returns).
	t.Cleanup(func() {
		DockerCleanup(ValidatorDemoContainerName)
	})

	state := DefaultBashState()

	// Initialize the chain at v0.12.0 and seed all downstream prerequisite
	// state.
	MustRun(t, "01_InitializeV0_12", func(t *testing.T) {
		t.Helper()
		runInitV0_12(t, state)
	})

	// Wrong-plan-name upgrade attempt that no-ops.
	MustRun(t, "02_RollbackV0_12", func(t *testing.T) {
		t.Helper()
		runRollback012(t, state)
	})

	// Upgrade 0.12 → 1.2, plus 1.2-era seed data.
	MustRun(t, "03_UpgradeTo1_2", func(t *testing.T) {
		t.Helper()
		runUpgrade012To12(t, state)
	})

	// Second wrong-plan-name attempt against v1.2.
	MustRun(t, "04_RollbackV1_2", func(t *testing.T) {
		t.Helper()
		runRollback12(t, state)
	})

	// Upgrade 1.2 → 1.4.3, plus NOC certs + revocation points.
	MustRun(t, "05_UpgradeTo1_4_3", func(t *testing.T) {
		t.Helper()
		runUpgrade12To143(t, state)
	})

	// Upgrade 1.4.3 → 1.4.4, plus DA certs + NOC revoke.
	MustRun(t, "06_UpgradeTo1_4_4", func(t *testing.T) {
		t.Helper()
		runUpgrade143To144(t, state)
	})

	// Upgrade 1.4.4 → 1.5.1.
	MustRun(t, "07_UpgradeTo1_5_1", func(t *testing.T) {
		t.Helper()
		runUpgrade144To151(t, state)
	})

	// Chain state from the 1.5.1 step enables the 1.5.2 and 1.6.0 upgrades.
	MustRun(t, "08_UpgradeTo1_5_2", func(t *testing.T) {
		t.Helper()
		runUpgrade151To152(t, state)
	})

	MustRun(t, "09_UpgradeTo1_6_0", func(t *testing.T) {
		t.Helper()
		runUpgrade152To160(t, state)
	})

	// Build master image, upgrade 1.6 → master.
	MustRun(t, "10_UpgradeTo_Master", func(t *testing.T) {
		t.Helper()
		runUpgrade160ToMaster(t, state)
	})

	// Fresh observer joins post-upgrade chain and catches up through
	// cosmovisor.
	MustRun(t, "11_AddNewNodeAfterUpgrade", func(t *testing.T) {
		t.Helper()
		runAddNewNodeAfterUpgrade(t, state)
	})
}
