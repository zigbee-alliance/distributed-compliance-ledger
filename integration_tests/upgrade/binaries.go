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
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// BinaryURLTemplate is the GitHub release download URL for a given dcld version.
const BinaryURLTemplate = "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v%s/dcld"

// BinaryPath returns the canonical local path for a given dcld version,
// regardless of whether it has been downloaded yet.
func BinaryPath(version string) string {
	return filepath.Join(BinariesDir, "dcld_v"+version)
}

// EnsureBinary downloads the dcld binary for `version` if it is not already
// present and executable on disk and applies the standard client config
// (chain-id, node, keyring-backend, broadcast-mode). The broadcast-mode
// matches the binary version — block for v0.12.x/v1.2.x, sync for v1.4+ —
// mirroring the `$DCLD_BIN config broadcast-mode ...` lines bash runs at the
// top of each upgrade script. Idempotent — safe to call repeatedly.
// Returns the absolute path to the binary.
func EnsureBinary(version string) (string, error) {
	if err := os.MkdirAll(BinariesDir, 0o755); err != nil {
		return "", fmt.Errorf("mkdir %s: %w", BinariesDir, err)
	}

	path := BinaryPath(version)

	if info, err := os.Stat(path); err == nil && info.Mode()&0o111 != 0 && info.Size() > 0 {
		if cerr := ConfigureClient(path); cerr != nil {
			return path, fmt.Errorf("configure client for %s: %w", path, cerr)
		}

		return path, nil
	}

	url := fmt.Sprintf(BinaryURLTemplate, version)
	resp, err := http.Get(url) //nolint:gosec,noctx
	if err != nil {
		return "", fmt.Errorf("download %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download %s: HTTP %d", url, resp.StatusCode)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", path, err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return "", fmt.Errorf("write %s: %w", path, err)
	}

	if cerr := ConfigureClient(path); cerr != nil {
		return path, fmt.Errorf("configure client for %s: %w", path, cerr)
	}

	return path, nil
}

// EnsureAllBinaries downloads every version in HistoricalVersions. Used as the
// per-test-run setup, equivalent to prepare-dcld-versions.sh.
func EnsureAllBinaries() error {
	for _, v := range HistoricalVersions {
		if _, err := EnsureBinary(v); err != nil {
			return err
		}
	}

	return nil
}
