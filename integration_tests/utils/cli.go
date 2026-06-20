package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
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
	cmd := exec.CommandContext(context.Background(), "dcld", args...)
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

// IsNotFound reports whether the dcld CLI output is the literal `"Not Found"`
// response emitted by utils/cli.QueryWithProof when a single-item query misses.
// All single-item Get* helpers use this to translate the sentinel response into
// a nil result without an error.
func IsNotFound(out []byte) bool {
	return strings.Contains(string(out), "Not Found")
}

// StripPagination removes the top-level "pagination" field from a CLI JSON
// response. The Cosmos SDK prints proto messages via protojson, which encodes
// PageResponse.Total as a quoted string ("0"). encoding/json then refuses to
// unmarshal that back into the uint64 field on the gogoproto-generated
// QueryAll*Response, so we drop the pagination wrapper entirely before
// decoding — list helpers never read it anyway.
//
// Returns the input unchanged if the JSON is malformed or has no pagination
// key, so callers can use it unconditionally.
func StripPagination(out []byte) []byte {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(out, &fields); err != nil {
		return out
	}
	if _, ok := fields["pagination"]; !ok {
		return out
	}
	delete(fields, "pagination")
	stripped, err := json.Marshal(fields)
	if err != nil {
		return out
	}

	return stripped
}

// quotedIntFields are the JSON keys whose values are uint64/int64 in
// gogoproto-generated structs but get emitted as quoted strings by the Cosmos
// SDK's protojson printer. NormalizeProtoJSON unquotes them so encoding/json
// can decode the response into the typed proto wrapper without choking on the
// type mismatch.
var quotedIntFields = map[string]struct{}{
	"account_number": {}, // cosmos.auth.BaseAccount.account_number (uint64)
	"sequence":       {}, // cosmos.auth.BaseAccount.sequence       (uint64)
	"time":           {}, // dclauth.Grant.time                     (int64)
	"height":         {}, // cosmos.upgrade.Plan.height             (int64)
	"otaFileSize":    {}, // model.ModelVersion.otaFileSize         (uint64)
	"total":          {}, // cosmos.base.query.PageResponse.total   (uint64)
}

// NormalizeProtoJSON walks out and converts every "<key>":"<digits>" entry
// whose key is in quotedIntFields into "<key>":<digits>. Returns the original
// bytes unchanged if the JSON is malformed or no rewrites are needed.
func NormalizeProtoJSON(out []byte) []byte {
	var root interface{}
	if err := json.Unmarshal(out, &root); err != nil {
		return out
	}
	changed := false
	walkAndUnquoteInts(&root, &changed)
	if !changed {
		return out
	}
	rewritten, err := json.Marshal(root)
	if err != nil {
		return out
	}

	return rewritten
}

// walkAndUnquoteInts recursively rewrites string-encoded ints in place for any
// map key that appears in quotedIntFields.
func walkAndUnquoteInts(node *interface{}, changed *bool) {
	switch v := (*node).(type) {
	case map[string]interface{}:
		for k, child := range v {
			if _, ok := quotedIntFields[k]; ok {
				if s, isStr := child.(string); isStr && looksLikeInt(s) {
					if n, err := strconv.ParseInt(s, 10, 64); err == nil {
						v[k] = json.Number(strconv.FormatInt(n, 10))
						*changed = true

						continue
					}
				}
			}
			walkAndUnquoteInts(&child, changed)
			v[k] = child
		}
	case []interface{}:
		for i := range v {
			walkAndUnquoteInts(&v[i], changed)
		}
	}
}

// looksLikeInt reports whether s is a non-empty optionally-signed integer.
func looksLikeInt(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	if s[0] == '-' || s[0] == '+' {
		if len(s) == 1 {
			return false
		}
		i = 1
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
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
