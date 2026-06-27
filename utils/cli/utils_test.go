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

package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFromFileExistingFile(t *testing.T) {
	const content = "certificate-content\nmultiple-lines\n"

	path := filepath.Join(t.TempDir(), "cert.pem")
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))

	got, err := ReadFromFile(path)
	require.NoError(t, err)
	require.Equal(t, content, got)
}

func TestReadFromFileEmptyFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.pem")
	require.NoError(t, os.WriteFile(path, []byte{}, 0o600))

	got, err := ReadFromFile(path)
	require.NoError(t, err)
	require.Equal(t, "", got)
}

// TestReadFromFileNotAPath covers the branch where the target does not exist on
// disk and is therefore returned verbatim (e.g. a PEM string passed inline).
func TestReadFromFileNotAPath(t *testing.T) {
	tests := []struct {
		name   string
		target string
	}{
		{"inline pem", "-----BEGIN CERTIFICATE-----\nMIIB...\n-----END CERTIFICATE-----"},
		{"plain string", "not-a-file"},
		{"empty string", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFromFile(tt.target)
			require.NoError(t, err)
			require.Equal(t, tt.target, got)
		})
	}
}

// TestReadFromFileReadError covers the branch where os.Stat succeeds but the
// subsequent read fails. Pointing at a directory triggers a read error while
// still passing the existence check.
func TestReadFromFileReadError(t *testing.T) {
	dir := t.TempDir()

	got, err := ReadFromFile(dir)
	require.Error(t, err)
	require.Equal(t, "", got)
}
