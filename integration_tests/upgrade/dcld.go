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

// ExecuteTxWithBin runs a `dcld tx ...` command via the binary at binPath,
// broadcasting in block mode so the returned TxResult is the confirmed result
// (mirrors the bash global `config broadcast-mode block`).
func ExecuteTxWithBin(binPath string, args ...string) (*utils.TxResult, error) {
	args = append(args,
		"--yes",
		"-o", "json",
		"--keyring-backend", "test",
		"--broadcast-mode", "block",
	)

	out, err := ExecuteCLIWithBin(binPath, args...)
	if err != nil {
		return nil, err
	}

	var result utils.TxResult
	if jerr := json.Unmarshal(out, &result); jerr != nil {
		return nil, fmt.Errorf("parse tx result: %w, output: %s", jerr, string(out))
	}

	return &result, nil
}

// ConfigureClient applies the standard dcld client config (chain-id, node,
// keyring-backend, broadcast-mode) for the binary at binPath. Equivalent to
// the `dcld config ...` block that opens every upgrade script.
func ConfigureClient(binPath string) error {
	cfgs := [][]string{
		{"config", "chain-id", ChainID},
		{"config", "output", "json"},
		{"config", "node", Node0Conn},
		{"config", "keyring-backend", "test"},
		{"config", "broadcast-mode", "block"},
	}
	for _, args := range cfgs {
		if _, err := ExecuteCLIWithBin(binPath, args...); err != nil {
			return err
		}
	}

	return nil
}
