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
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// ExecuteCLIWithBin runs an arbitrary dcld command using the binary at binPath.
// Mirrors utils.ExecuteCLI but allows targeting a historical dcld release.
// Empty newlines are piped on stdin so commands that prompt for a passphrase
// don't block on EOF.
func ExecuteCLIWithBin(binPath string, args ...string) ([]byte, error) {
	cmd := exec.Command(binPath, args...)
	cmd.Stdin = bytes.NewBufferString("\n\n")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("%s %s: %w, output: %s",
			binPath, strings.Join(args, " "), err, string(out))
	}

	strOut := string(out)
	if idx := strings.Index(strOut, "{"); idx > 0 {
		strOut = strOut[idx:]
	}

	return []byte(strOut), nil
}

// ExecuteTxWithBin runs a `dcld tx ...` command via the binary at binPath.
//
// Broadcast mode is taken from the binary's persistent client config (set by
// ConfigureClient at EnsureBinary time), so we don't pass --broadcast-mode
// here. Cosmos-SDK retired `block` mode in dcld v1.4.3+, which is why the
// config is `block` for v0.12.x/v1.2.x and `sync` for v1.4+. In sync mode the
// broadcast only returns ante-handler/mempool acceptance, so we poll
// `query tx` for the in-block result — keeping the contract that the
// returned TxResult.Code reflects on-chain execution in either mode.
func ExecuteTxWithBin(binPath string, args ...string) (*utils.TxResult, error) {
	args = append(args,
		"--yes",
		"-o", "json",
		"--keyring-backend", "test",
	)

	out, err := ExecuteCLIWithBin(binPath, args...)
	if err != nil {
		return nil, err
	}

	var result utils.TxResult
	if jerr := json.Unmarshal(out, &result); jerr != nil {
		return nil, fmt.Errorf("parse tx result: %w, output: %s", jerr, string(out))
	}

	// Poll for on-chain confirmation when the binary is in sync mode. Block
	// mode already returns the confirmed result, so polling there is skipped.
	if binPathSupportsOnlySyncMode(binPath) && result.Code == 0 && result.TxHash != "" {
		confirmedOut, awaitErr := utils.AwaitTxConfirmation(result.TxHash)
		if awaitErr != nil {
			return &result, awaitErr
		}

		var confirmed utils.TxResult
		if jerr := json.Unmarshal(confirmedOut, &confirmed); jerr == nil {
			return &confirmed, nil
		}
	}

	return &result, nil
}

// ConfigureClient applies the standard dcld client config (chain-id, node,
// keyring-backend, broadcast-mode) for the binary at binPath. Equivalent to
// the `dcld config ...` block that opens every upgrade script.
//
// Broadcast-mode follows the binary version: `block` for v0.12.x/v1.2.x
// (matches bash scripts 01-03), `sync` for v1.4+ (matches bash scripts 04+,
// where cosmos-sdk removed `block` mode). All dcld binaries share the same
// `~/.dcl/config/client.toml`, so the last call wins — in practice the
// upgrade test downloads binaries in version order, and the v1.4+ download
// flips the global config to sync. v0.12/v1.2 binaries used after that point
// still work because they also support sync mode.
func ConfigureClient(binPath string) error {
	mode := "block"
	if binPathSupportsOnlySyncMode(binPath) {
		mode = "sync"
	}

	cfgs := [][]string{
		{"config", "chain-id", ChainID},
		{"config", "output", "json"},
		{"config", "node", Node0Conn},
		{"config", "keyring-backend", "test"},
		{"config", "broadcast-mode", mode},
	}
	for _, args := range cfgs {
		if _, err := ExecuteCLIWithBin(binPath, args...); err != nil {
			return err
		}
	}

	return nil
}

// binPathSupportsOnlySyncMode reports whether the dcld binary at binPath
// rejects `--broadcast-mode block`. Cosmos-SDK retired block mode in the
// release that ships with dcld v1.4.3; v0.12.x and v1.2.x still accept it.
// The path format from EnsureBinary is `<dir>/dcld_v<major>.<minor>.<patch>`.
func binPathSupportsOnlySyncMode(binPath string) bool {
	idx := strings.LastIndex(binPath, "_v")
	if idx < 0 {
		return false
	}
	version := binPath[idx+2:]
	parts := strings.SplitN(version, ".", 3)
	if len(parts) < 2 {
		return false
	}

	var major, minor int
	if _, err := fmt.Sscanf(parts[0], "%d", &major); err != nil {
		return false
	}
	if _, err := fmt.Sscanf(parts[1], "%d", &minor); err != nil {
		return false
	}

	// >= v1.4 — covers 1.4.3, 1.4.4, 1.5.x, 1.6.x, plus the master build.
	if major > 1 {
		return true
	}

	return major == 1 && minor >= 4
}
