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
	"os/exec"
	"strings"
)

// InitPoolOpts mirrors the positional args of bash `init_pool`. All fields
// are optional; zero values produce the same defaults as bash.
type InitPoolOpts struct {
	// PatchConfig controls whether the consensus timeouts are patched down to
	// 1s. Defaults to "yes" if empty.
	PatchConfig string
	// InitTarget is the Makefile target invoked to generate node configs.
	// Defaults to "localnet_init".
	InitTarget string
	// BinaryVersion seeds the localnet with a specific dcld release. Empty
	// means use the host-built binary.
	BinaryVersion string
}

// runPoolHelper invokes one of the helper functions defined in pool.sh / common.sh
// from a transient bash subshell, with the same environment a `run-all.sh`
// invocation would set up.
func runPoolHelper(call string) error {
	script := fmt.Sprintf(
		`source integration_tests/cli/common.sh; source integration_tests/pool.sh; %s`,
		call,
	)

	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		"LOCALNET_DIR="+LocalnetDir,
		"DETAILED_OUTPUT_TARGET=/dev/stdout",
	)

	return cmd.Run()
}

// InitPool starts the localnet via the bash helpers. Shells out to pool.sh
// rather than reimplementing make / docker-compose orchestration in Go.
func InitPool(opts InitPoolOpts) error {
	patch := opts.PatchConfig
	if patch == "" {
		patch = "yes"
	}

	args := []string{patch}
	if opts.InitTarget != "" || opts.BinaryVersion != "" {
		args = append(args, defaultIfEmpty(opts.InitTarget, "localnet_init"))
	}
	if opts.BinaryVersion != "" {
		args = append(args, opts.BinaryVersion)
	}

	return runPoolHelper("init_pool " + strings.Join(args, " "))
}

// CleanupPool tears down the localnet.
func CleanupPool() error {
	return runPoolHelper("cleanup_pool")
}

// readPersistentPeersFromHost extracts the `persistent_peers = "..."` value
// from the host-side config.toml file produced by `make localnet_init`.
// Equivalent to bash:
//
//	cat <file> | grep -o -E 'persistent_peers = ".*"'
//
// Used when seeding fresh observer / validator-demo containers with the same
// peer list as node0.
func readPersistentPeersFromHost(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if val, ok := strings.CutPrefix(line, "persistent_peers = "); ok {
			return strings.Trim(val, `"`), nil
		}
	}

	return "", fmt.Errorf("persistent_peers not found in %s", path)
}

func defaultIfEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}

	return s
}
