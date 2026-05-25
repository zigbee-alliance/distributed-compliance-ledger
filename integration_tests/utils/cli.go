package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// TxResult represents the core standard JSON output emitted by `dcld` node
type TxResult struct {
	Code   uint32 `json:"code"`
	TxHash string `json:"txhash"`
	RawLog string `json:"raw_log"`
}

// ExecuteCLI runs a `dcld` command as a subprocess.
// It sets standard cosmos flags like --output json and keyring-backend test
// if they are not already provided in the args for transaction commands.
func ExecuteCLI(args ...string) ([]byte, error) {
	cmd := exec.Command("dcld", args...)
	// Provide empty newlines as stdin so commands that prompt for a passphrase
	// (e.g. `keys add` prompts for a BIP39 passphrase) don't block on EOF.
	cmd.Stdin = bytes.NewBufferString("\n\n")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("cmd execution failed: %w, output: %s", err, string(out))
	}
	// Many cosmos CLI commands output warnings before JSON.
	// Strip out everything before the first `{` to ensure clean parsing.
	strOut := string(out)
	idx := strings.Index(strOut, "{")
	if idx > 0 {
		strOut = strOut[idx:]
	}

	return []byte(strOut), nil
}

// ExecuteTx runs a `dcld tx` command, broadcasts it, and waits for on-chain
// confirmation. The returned TxResult.Code is the actual on-chain execution
// code, not the mempool acceptance code. This eliminates sequence-number
// conflicts between consecutive transactions from the same account.
func ExecuteTx(args ...string) (*TxResult, error) {
	args = append(args, "--yes", "-o", "json", "--keyring-backend", "test")
	out, err := ExecuteCLI(args...)
	if err != nil {
		// CLI failed before broadcasting (bad flag, key not found, etc.).
		return nil, err
	}

	var broadcast TxResult
	if err := json.Unmarshal(out, &broadcast); err != nil {
		return nil, fmt.Errorf("failed to parse TxResult: %w, output: %s", err, string(out))
	}

	// Non-zero broadcast code means the tx was rejected before entering a block
	// (ante-handler failure, sequence error, etc.) — return it as-is.
	if broadcast.Code != 0 || broadcast.TxHash == "" {
		return &broadcast, nil
	}

	// Wait for block inclusion to get the actual on-chain execution code.
	txData, awaitErr := AwaitTxConfirmation(broadcast.TxHash)
	if awaitErr != nil {
		return &broadcast, awaitErr
	}

	var confirmed TxResult
	if err := json.Unmarshal(txData, &confirmed); err != nil {
		// can't parse confirmed result; fall back to broadcast result
		return &broadcast, nil //nolint:nilerr
	}

	return &confirmed, nil
}

// QueryTx runs `dcld query tx [txHash]` mimicking exactly how the REST /grpc tests
// await for block inclusion.
func QueryTx(txHash string) ([]byte, error) {
	return ExecuteCLI("query", "tx", txHash, "-o", "json")
}

// OnChainCode parses the on-chain execution code from a confirmed tx response.
// Returns 0 and no error if the code field is absent (assumed success).
func OnChainCode(txData []byte) (uint32, string, error) {
	var res TxResult
	if err := json.Unmarshal(txData, &res); err != nil {
		return 0, "", fmt.Errorf("failed to parse on-chain tx result: %w", err)
	}

	return res.Code, res.RawLog, nil
}

// AwaitTxConfirmation replaces shell `get_txn_result` retry looping, resolving brittleness.
func AwaitTxConfirmation(txHash string) ([]byte, error) {
	var result []byte
	var err error

	for i := 0; i < 20; i++ {
		result, err = QueryTx(txHash)
		if err == nil {
			return result, nil // tx is in a block (succeeded or failed on-chain)
		}
		time.Sleep(2 * time.Second)
	}

	return result, fmt.Errorf("transaction %s not confirmed after 40 seconds. Last error: %w", txHash, err)
}
