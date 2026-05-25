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
// Subtests run in order against a shared chain. Each phase adds more.
//
// Phase ordering (see plan in the migration PR):
//
//	Phase 2 — 01_InitializeV0_12      (this commit)
//	Phase 2 — 03_UpgradeTo1_2         (next)
//	Phase 2 — 05_UpgradeTo1_4_3       (next)
//	Phase 2 — 06_UpgradeTo1_4_4       (next)
//	Phase 2 — 07_UpgradeTo1_5_1       (next)
//	Phase 1 — 08_UpgradeTo1_5_2       (done — runs once phase 2 completes)
//	Phase 1 — 09_UpgradeTo1_6_0       (done — runs once phase 2 completes)
//	Phase 3 — 02_RollbackV0_12        (alt path)
//	Phase 3 — 04_RollbackV1_2         (alt path)
//	Phase 4 — 10_UpgradeTo_Master
//	Phase 4 — 11_AddNewNodeAfterUpgrade
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

	state := DefaultBashState()

	// Phase 2: script 01 — initialize the chain at v0.12.0 and seed all
	// downstream prerequisite state.
	t.Run("01_InitializeV0_12", func(t *testing.T) {
		runInitV0_12(t, state)
	})

	// Phase 3: script 02 — wrong-plan-name upgrade attempt that no-ops.
	t.Run("02_RollbackV0_12", func(t *testing.T) {
		runRollback012(t, state)
	})

	// Phase 2: script 03 — upgrade 0.12 → 1.2, plus 1.2-era seed data.
	t.Run("03_UpgradeTo1_2", func(t *testing.T) {
		runUpgrade012To12(t, state)
	})

	// Phase 3: script 04 — second wrong-plan-name attempt against v1.2.
	t.Run("04_RollbackV1_2", func(t *testing.T) {
		runRollback12(t, state)
	})

	// Phase 2: script 05 — upgrade 1.2 → 1.4.3, plus NOC certs + revocation points.
	t.Run("05_UpgradeTo1_4_3", func(t *testing.T) {
		runUpgrade12To143(t, state)
	})

	// Phase 2: script 06 — upgrade 1.4.3 → 1.4.4, plus DA certs + NOC revoke.
	t.Run("06_UpgradeTo1_4_4", func(t *testing.T) {
		runUpgrade143To144(t, state)
	})

	// Phase 2: script 07 — upgrade 1.4.4 → 1.5.1. Final Phase 2 script.
	t.Run("07_UpgradeTo1_5_1", func(t *testing.T) {
		runUpgrade144To151(t, state)
	})

	// Phase 1: scripts 08 and 09 — chain state from 07 enables these to run.
	t.Run("08_UpgradeTo1_5_2", func(t *testing.T) {
		runUpgrade151To152(t, state)
	})

	t.Run("09_UpgradeTo1_6_0", func(t *testing.T) {
		runUpgrade152To160(t, state)
	})

	// Phase 4: script 10 — build master image, upgrade 1.6 → master.
	t.Run("10_UpgradeTo_Master", func(t *testing.T) {
		runUpgrade160ToMaster(t, state)
	})

	// Phase 4: script 11 — fresh observer joins post-upgrade chain and
	// catches up through cosmovisor.
	t.Run("11_AddNewNodeAfterUpgrade", func(t *testing.T) {
		runAddNewNodeAfterUpgrade(t, state)
	})
}
