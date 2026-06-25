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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// runLightFlowEnv is the env var that gates this suite. CI sets
// RUN_LIGHT_GO=1 when running the "light" test bucket (see
// integration_tests/run-all.sh). Local dev keeps the suite fast-skipping
// unless a developer opts in — the suite needs a live localnet plus the
// light_client_proxy container reachable on localhost:26620. Naming mirrors
// upgrade's runUpgradeFlowEnv / shouldRunUpgradeFlow pair.
const runLightFlowEnv = "RUN_LIGHT_GO"

func shouldRunLightFlow() bool {
	return os.Getenv(runLightFlowEnv) == "1"
}

// skipIfDisabled bails out of a test when RUN_LIGHT_GO=1 isn't set. Mirrors
// the upgrade package's shouldRunUpgradeFlow gate so `go test ./...` stays
// fast in regular development.
func skipIfDisabled(t *testing.T) {
	t.Helper()
	if !shouldRunLightFlow() {
		t.Skipf("set %s=1 to run the light client proxy tests", runLightFlowEnv)
	}
}

// TestMain chdirs to the repo root so cli/common.sh-relative paths and any
// future fixture lookups resolve the same as they do under run-all.sh.
// Matches the upgrade package's TestMain.
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
