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

// Package validator_test contains integration tests translated from validator-demo.sh.
//
// Note: The original shell script spins up a new Docker container, adds a brand new validator
// node, and then tests disable/enable/propose/approve/reject flows against that node.  These
// Go tests cover the same logical flows using the existing localnet nodes (node0 … node3) so
// that no Docker setup is required at the Go-test level.  The per-validator address that the
// shell script resolved via "dcld tendermint show-address" inside the container is replaced by
// querying all-nodes and picking the first node whose validator address is already known.
package validator

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestValidatorProposeRejectDisable translates the propose-and-reject disable-validator section
// and the sequential approve/reject/re-vote flows from validator-demo.sh.
//
// Prerequisites: a running localnet with at least one validator node visible via
//
//	dcld query validator all-nodes
func TestValidatorProposeRejectDisable(t *testing.T) {
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount
	jack := testconstants.JackAccount

	aliceAddr, err := getAddress(alice)
	require.NoError(t, err)
	bobAddr, err := getAddress(bob)
	require.NoError(t, err)

	// Pick the first known validator to run the test against.
	// The shell script used a freshly-started container validator; we reuse an existing one.
	validatorOwner, validatorAddress := resolveFirstValidator(t)
	t.Logf("Using validator owner=%s address=%s", validatorOwner, validatorAddress)

	t.Run("ProposeAndRejectDisableValidator_NoEffect", func(t *testing.T) {
		// Alice proposes to disable
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice rejects the proposal she just made
		txResult, err = RejectDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Should not be in proposed list
		out, err := QueryProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Should not be in rejected list (single proposer rejection clears it immediately)
		out, err = QueryRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Should not be in disabled list
		out, err = QueryDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("ProposeApproveDisableAndReEnable", func(t *testing.T) {
		// Alice proposes to disable
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed list contains the validator
		out, err := QueryAllProposedDisableNodes()
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)

		out, err = QueryProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)
		require.Contains(t, string(out), aliceAddr)

		// Bob approves — reaches threshold (ceil(2/3 * 3 trustees) = 2, proposer counts as 1),
		// validator becomes disabled and the proposal is removed.
		txResult, err = ApproveDisableNode(validatorAddress, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob cannot reject — proposal is gone (threshold was reached)
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Should no longer be in proposed list (threshold reached)
		out, err = QueryAllProposedDisableNodes()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"address":"%s"`, validatorAddress))

		// Should be in disabled list
		out, err = QueryDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)
		require.Contains(t, string(out), "false") // disabledByNodeAdmin = false

		// Node admin (validatorOwner) re-enables the validator
		txResult, err = EnableNode(validatorOwner)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Should no longer be disabled
		out, err = QueryDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("ProposeApproveRejectRejectFailsSecondTime", func(t *testing.T) {
		// Alice proposes
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob approves — reaches threshold, validator becomes disabled and proposal is removed.
		txResult, err = ApproveDisableNode(validatorAddress, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob cannot reject — proposal is gone (threshold was reached)
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Re-enable
		txResult, err = EnableNode(validatorOwner)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ProposeRejectByBobAndJack_GoesToRejectedList", func(t *testing.T) {
		// Alice proposes
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob rejects
		txResult, err = RejectDisableNode(validatorAddress, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob cannot reject twice
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Still in proposed list (not enough rejections yet)
		out, err := QueryProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), bobAddr)

		// Jack rejects — now enough rejections to move to rejected list
		txResult, err = RejectDisableNode(validatorAddress, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack cannot reject twice
		txBad, errBad = RejectDisableNode(validatorAddress, jack)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// No longer in proposed list
		out, err = QueryAllProposedDisableNodes()
		require.NoError(t, err)
		require.NotContains(t, string(out), validatorAddress)

		// Should be in rejected list
		out, err = QueryAllRejectedDisableNodes()
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), bobAddr)

		out, err = QueryRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)

		// Should NOT be disabled
		out, err = QueryDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Should NOT be jailed
		out, err = QueryNode(validatorAddress)
		require.NoError(t, err)
		require.NotContains(t, string(out), `"jailed":true`)
	})

	t.Run("RePropose_AfterRejected_StartsFresh", func(t *testing.T) {
		// Alice proposes again — should succeed even though it was previously rejected
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now in proposed list
		out, err := QueryProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), validatorAddress)
		require.Contains(t, string(out), aliceAddr)

		// Rejected list should no longer have it
		out, err = QueryRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Clean up: reject it to leave the network in a clean state
		txResult, err = RejectDisableNode(validatorAddress, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})
}

// resolveFirstValidator queries all validator nodes and returns (ownerAccountName, validatorAddress)
// for the first validator it finds whose owner key exists in the local keyring.
func resolveFirstValidator(t *testing.T) (ownerAccountName, validatorAddress string) {
	t.Helper()

	out, err := QueryAllNodes()
	require.NoError(t, err)
	outStr := string(out)

	// The validator "owner" field in all-nodes is a cosmosvaloper... address.
	// Use --bech val to get the cosmosvaloper address for each well-known key and
	// match against the all-nodes JSON. Prefer anna then jack so the validator under
	// test is not owned by alice or bob (the propose/approve signers).
	knownAccounts := []string{"anna", "jack", "alice", "bob"}
	for _, acc := range knownAccounts {
		valAddrOut, err := utils.ExecuteCLI("keys", "show", acc, "--bech", "val", "-a", "--keyring-backend", "test")
		if err != nil {
			continue
		}
		valAddr := strings.TrimSpace(string(valAddrOut))
		if valAddr != "" && strings.Contains(outStr, valAddr) {
			return acc, valAddr
		}
	}

	require.Fail(t, "could not find a known validator node admin account in the localnet")

	return "", ""
}

// getAddress returns the bech32 address for the given key name.
func getAddress(name string) (string, error) {
	out, err := utils.ExecuteCLI("keys", "show", name, "-a", "--keyring-backend", "test")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
