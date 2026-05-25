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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// RandomString returns a short hex-encoded random suffix suitable for use as
// a unique account name. Equivalent to the bash `random_string` helper.
func RandomString() string {
	var b [6]byte
	if _, err := rand.Read(b[:]); err != nil {
		// crypto/rand failure on a CI runner is fatal; panic so tests don't
		// silently fall back to a static suffix.
		panic(fmt.Errorf("rand.Read: %w", err))
	}

	return "u" + hex.EncodeToString(b[:])
}

// CreateKey runs `keys add NAME` with the test keyring and returns the key's
// bech32 address and pubkey JSON. Equivalent to the bash pattern:
//
//	dcld keys add <name>
//	dcld keys show <name> -a
//	dcld keys show <name> -p
func CreateKey(binPath, name string) (address, pubkey string, err error) {
	if _, err = ExecuteCLIWithBin(binPath,
		"keys", "add", name,
		"--keyring-backend", "test",
	); err != nil {
		return "", "", fmt.Errorf("keys add %s: %w", name, err)
	}

	addrOut, err := ExecuteCLIWithBin(binPath,
		"keys", "show", name, "-a",
		"--keyring-backend", "test",
	)
	if err != nil {
		return "", "", fmt.Errorf("keys show %s -a: %w", name, err)
	}

	pubOut, err := ExecuteCLIWithBin(binPath,
		"keys", "show", name, "-p",
		"--keyring-backend", "test",
	)
	if err != nil {
		return "", "", fmt.Errorf("keys show %s -p: %w", name, err)
	}

	return TrimTrailingWS(string(addrOut)), TrimTrailingWS(string(pubOut)), nil
}

// TrimTrailingWS strips trailing newlines/whitespace from a CLI output line.
// Promoted from the test files since several non-test helpers need it.
func TrimTrailingWS(s string) string {
	return strings.TrimRight(s, "\n\r \t")
}

// ProposeAddAccountArgs gathers the optional fields for ProposeAddAccount.
type ProposeAddAccountArgs struct {
	VID   int    // -1 to omit
	Roles string // comma-separated role names (e.g. "Vendor", "Trustee")
}

// ProposeAddAccount runs `tx auth propose-add-account`.
func ProposeAddAccount(binPath, address, pubkey, from string, extra ProposeAddAccountArgs) (*utils.TxResult, error) {
	args := []string{
		"tx", "auth", "propose-add-account",
		"--address", address,
		"--pubkey", pubkey,
	}
	if extra.VID >= 0 {
		args = append(args, "--vid", fmt.Sprintf("%d", extra.VID))
	}
	if extra.Roles != "" {
		args = append(args, "--roles", extra.Roles)
	}
	args = append(args, "--from", from)

	return ExecuteTxWithBin(binPath, args...)
}

// ApproveAddAccount runs `tx auth approve-add-account`.
func ApproveAddAccount(binPath, address, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "auth", "approve-add-account",
		"--address", address,
		"--from", from,
	)
}

// CreateAndApproveAccount creates a new key at `name`, proposes it with
// `proposer`, and obtains approvals from each of `approvers` until the account
// becomes active. Returns the on-chain address. Pass vid=-1 to omit the --vid
// flag.
//
// The approval loop short-circuits as soon as the account moves from pending
// to active: the underlying threshold depends on trustee count and role (e.g.
// vendor activates after one approval with 3 trustees), and extra approvals
// against an already-active account return "pending account not found".
//
// Fatals via the supplied *testing.T on any tx code != 0 (RawLog is logged).
func CreateAndApproveAccount(t TestingHelper, binPath, name, role string, vid int, proposer string, approvers []string) string {
	t.Helper()

	addr, pub, err := CreateKey(binPath, name)
	if err != nil {
		t.Fatalf("create key %s: %v", name, err)
	}

	tx, err := ProposeAddAccount(binPath, addr, pub, proposer, ProposeAddAccountArgs{
		VID: vid, Roles: role,
	})
	if err != nil {
		t.Fatalf("propose-add-account %s: %v", name, err)
	}
	if tx.Code != 0 {
		t.Fatalf("propose-add-account %s: code=%d log=%s", name, tx.Code, tx.RawLog)
	}

	if isAccountActive(binPath, addr) {
		return addr
	}

	for _, who := range approvers {
		tx, err = ApproveAddAccount(binPath, addr, who)
		if err != nil {
			t.Fatalf("approve-add-account %s from %s: %v", name, who, err)
		}
		if tx.Code != 0 {
			t.Fatalf("approve-add-account %s from %s: code=%d log=%s",
				name, who, tx.Code, tx.RawLog)
		}

		if isAccountActive(binPath, addr) {
			return addr
		}
	}

	return addr
}

// isAccountActive reports whether an account address has cleared its pending
// state and is queryable via `query auth account`. Errors are treated as
// "not active" — callers either keep approving or fail with the original
// approve error.
func isAccountActive(binPath, addr string) bool {
	out, err := ExecuteCLIWithBin(binPath, "query", "auth", "account", "--address", addr, "-o", "json")
	if err != nil {
		return false
	}

	return !strings.Contains(string(out), "Not Found")
}

// TestingHelper is the subset of *testing.T used by package-level helpers, so
// they can be called from both _test.go files and non-test code if needed.
type TestingHelper interface {
	Helper()
	Fatalf(format string, args ...any)
}

// ProposeRevokeAccount runs `tx auth propose-revoke-account`.
func ProposeRevokeAccount(binPath, address, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "auth", "propose-revoke-account",
		"--address", address,
		"--from", from,
	)
}

// ApproveRevokeAccount runs `tx auth approve-revoke-account`.
func ApproveRevokeAccount(binPath, address, from string) (*utils.TxResult, error) {
	return ExecuteTxWithBin(binPath,
		"tx", "auth", "approve-revoke-account",
		"--address", address,
		"--from", from,
	)
}
