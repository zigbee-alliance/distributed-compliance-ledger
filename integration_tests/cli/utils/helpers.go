package cliputils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// isAccountActive checks whether an account address is active on-chain.
func isAccountActive(addr string) bool {
	out, err := utils.ExecuteCLI("query", "auth", "account", "--address", addr, "-o", "json")
	if err != nil {
		return false
	}

	return !strings.Contains(string(out), "Not Found")
}

// CreateAccount generates a new key, proposes and approves the account via jack/alice/bob,
// and returns the account name. roles is a comma-separated string e.g. "NodeAdmin" or "Vendor".
// It is robust against varying trustee counts: it tries up to three approvals (jack, alice, bob)
// and verifies the account is active before returning.
func CreateAccount(t *testing.T, roles string) string {
	t.Helper()

	name := utils.RandString()

	// Delete existing key if present (keys accumulate across test runs against a shared keyring).
	utils.ExecuteCLI("keys", "delete", name, "--keyring-backend", "test", "-y") //nolint:errcheck

	_, err := utils.ExecuteCLI("keys", "add", name, "--keyring-backend", "test", "--no-backup")
	require.NoError(t, err)

	addrOut, err := utils.ExecuteCLI("keys", "show", name, "-a", "--keyring-backend", "test")
	require.NoError(t, err)
	addr := stripNewline(addrOut)

	pubkeyOut, err := utils.ExecuteCLI("keys", "show", name, "-p", "--keyring-backend", "test")
	require.NoError(t, err)
	pubkey := stripNewline(pubkeyOut)

	txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
		"--address", addr,
		"--pubkey", pubkey,
		"--roles", roles,
		"--from", testconstants.JackAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "propose-add-account failed: %s", txResult.RawLog)

	if isAccountActive(addr) {
		return name
	}

	// Alice approves (required for >= 3 trustees with AccountApprovalsPercent=2/3).
	txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
		"--address", addr,
		"--from", testconstants.AliceAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "approve-add-account (alice) failed: %s", txResult.RawLog)

	if isAccountActive(addr) {
		return name
	}

	// Bob approves (required for >= 4 trustees where ceil(2/3*4)=3 approvals are needed).
	txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
		"--address", addr,
		"--from", testconstants.BobAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "approve-add-account (bob) failed: %s", txResult.RawLog)

	require.True(t, isAccountActive(addr),
		"account %s (%s) is still not active after jack+alice+bob approvals; "+
			"too many trustees on chain for automatic activation", name, addr)

	return name
}

// CreateVendorAccount creates a vendor account with the given name and vid.
// An optional single pid_ranges argument may be passed.
// It is robust against varying trustee counts: it tries up to three approvals (jack, alice, bob)
// and verifies the account is active before returning.
func CreateVendorAccount(t *testing.T, name string, vid int, pidRanges ...string) string {
	t.Helper()

	// Delete existing key if present (keys accumulate across test runs against a shared keyring).
	utils.ExecuteCLI("keys", "delete", name, "--keyring-backend", "test", "-y") //nolint:errcheck

	_, err := utils.ExecuteCLI("keys", "add", name, "--keyring-backend", "test", "--no-backup")
	require.NoError(t, err)

	addrOut, err := utils.ExecuteCLI("keys", "show", name, "-a", "--keyring-backend", "test")
	require.NoError(t, err)
	addr := stripNewline(addrOut)

	pubkeyOut, err := utils.ExecuteCLI("keys", "show", name, "-p", "--keyring-backend", "test")
	require.NoError(t, err)
	pubkey := stripNewline(pubkeyOut)

	args := []string{
		"tx", "auth", "propose-add-account",
		"--address", addr,
		"--pubkey", pubkey,
		"--roles", "Vendor",
		"--vid", strconv.Itoa(vid),
		"--from", testconstants.JackAccount,
	}
	if len(pidRanges) > 0 {
		args = append(args, "--pid_ranges", pidRanges[0])
	}

	txResult, err := utils.ExecuteTx(args...)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "propose-add-account (vendor) failed: %s", txResult.RawLog)

	// With 3 trustees, VendorAccountApprovalsPercent=1/3 → ceil(1)=1 → Jack's proposal activates
	// directly. With 4+ trustees, ceil(1/3*N) > 1 → need additional approvals.
	if isAccountActive(addr) {
		return name
	}

	// Alice approves (needed when 4-6 trustees, ceil(1/3*N) = 2).
	txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
		"--address", addr,
		"--from", testconstants.AliceAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "approve-add-account (vendor, alice) failed: %s", txResult.RawLog)

	if isAccountActive(addr) {
		return name
	}

	// Bob approves (needed when 7+ trustees, ceil(1/3*N) = 3).
	txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
		"--address", addr,
		"--from", testconstants.BobAccount,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "approve-add-account (vendor, bob) failed: %s", txResult.RawLog)

	require.True(t, isAccountActive(addr),
		"vendor account %s (vid=%d, addr=%s) is still not active after jack+alice+bob approvals; "+
			"too many trustees on chain for automatic activation", name, vid, addr)

	return name
}

// CreateModelAndVersion adds a model and a model version for the given vid/pid/sv/svs
// using the provided userAddr (account name) as the signer.
func CreateModelAndVersion(t *testing.T, vid, pid, sv int, svs, userAddr string) {
	t.Helper()

	txResult, err := utils.ExecuteTx("tx", "model", "add-model",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--deviceTypeID", "1",
		"--productName", "TestProduct",
		"--productLabel", "TestingProductLabel",
		"--partNumber", "1",
		"--commissioningCustomFlow", "0",
		"--enhancedSetupFlowOptions", "0",
		"--from", userAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "add-model failed: %s", txResult.RawLog)

	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	txResult, err = utils.ExecuteTx("tx", "model", "add-model-version",
		"--vid", strconv.Itoa(vid),
		"--pid", strconv.Itoa(pid),
		"--softwareVersion", strconv.Itoa(sv),
		"--softwareVersionString", svs,
		"--cdVersionNumber", "1",
		"--maxApplicableSoftwareVersion", "10",
		"--minApplicableSoftwareVersion", "1",
		"--from", userAddr,
	)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "add-model-version failed: %s", txResult.RawLog)

	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)
}

// GetHeight queries the node status and returns the latest block height.
func GetHeight() (int64, error) {
	out, err := utils.ExecuteCLI("status")
	if err != nil {
		return 0, err
	}

	var status struct {
		SyncInfo struct {
			LatestBlockHeight string `json:"latest_block_height"`
		} `json:"SyncInfo"`
	}
	if err := json.Unmarshal(out, &status); err != nil {
		return 0, fmt.Errorf("failed to parse status: %w, output: %s", err, string(out))
	}

	h, err := strconv.ParseInt(status.SyncInfo.LatestBlockHeight, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse latest_block_height: %w", err)
	}

	return h, nil
}

// WaitForHeight polls until the chain reaches target height or timeoutSec is exceeded.
func WaitForHeight(t *testing.T, target int64, timeoutSec int) {
	t.Helper()

	deadline := time.Now().Add(time.Duration(timeoutSec) * time.Second)
	for {
		h, err := GetHeight()
		if err == nil && h >= target {
			return
		}
		if time.Now().After(deadline) {
			require.Failf(t, "WaitForHeight timed out",
				"height %d not reached within %d seconds (current: %d, err: %v)",
				target, timeoutSec, h, err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// stripNewline removes trailing newline/whitespace from CLI output bytes.
func stripNewline(b []byte) string {
	s := string(b)
	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r' || s[len(s)-1] == ' ') {
		s = s[:len(s)-1]
	}
	return s
}
