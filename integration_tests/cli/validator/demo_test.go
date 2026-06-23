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

// Package validator contains integration tests for the validator module's
// disable/enable/propose/approve/reject flows.
//
// Tests run against the existing localnet nodes (node0 … node3) — no
// per-test Docker setup. The per-validator address is resolved by querying
// all-nodes and picking the first node whose validator address is already
// known.
package validator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	validatortypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// proposedHasVoter reports whether the proposed-disable record has addr as the
// creator, among its approvals, or among its rejects.
func proposedHasVoter(p *validatortypes.ProposedDisableValidator, addr string) bool {
	if p == nil {
		return false
	}
	if p.Creator == addr {
		return true
	}
	for _, a := range p.Approvals {
		if a != nil && a.Address == addr {
			return true
		}
	}
	for _, a := range p.Rejects {
		if a != nil && a.Address == addr {
			return true
		}
	}

	return false
}

// rejectedHasVoter reports whether the rejected-disable record has addr in its
// approvals (proposer) or rejects.
func rejectedHasVoter(r *validatortypes.RejectedDisableValidator, addr string) bool {
	if r == nil {
		return false
	}
	for _, a := range r.Approvals {
		if a != nil && a.Address == addr {
			return true
		}
	}
	for _, a := range r.Rejects {
		if a != nil && a.Address == addr {
			return true
		}
	}

	return false
}

// TestValidatorProposeRejectDisable covers the propose-and-reject
// disable-validator path and the sequential approve/reject/re-vote flows.
//
// Prerequisites: a running localnet with at least one validator node visible via
//
//	dcld query validator all-nodes
func TestValidatorProposeRejectDisable(t *testing.T) {
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount
	jack := testconstants.JackAccount

	aliceAddr, err := cliputils.GetAddress(alice)
	require.NoError(t, err)
	bobAddr, err := cliputils.GetAddress(bob)
	require.NoError(t, err)

	// Pick the first known validator to run the test against.
	// The shell script used a freshly-started container validator; we reuse an existing one.
	validatorOwner, validatorAddress := resolveFirstValidator(t)
	t.Logf("Using validator owner=%s address=%s", validatorOwner, validatorAddress)

	t.Run("QueryUnknownNode_NotFound", func(t *testing.T) {
		// A never-added address has neither a node nor a last-power record
		// (validator-demo.sh:113-119).
		name := utils.RandString()
		require.NoError(t, cliputils.AddKey(name))
		unknownAddr, err := cliputils.GetAddress(name)
		require.NoError(t, err)

		node, err := GetNode(unknownAddr)
		require.NoError(t, err)
		require.Nil(t, node)

		power, err := GetLastPower(unknownAddr)
		require.NoError(t, err)
		require.Nil(t, power)
	})

	t.Run("ProposeAndRejectDisableValidator_NoEffect", func(t *testing.T) {
		// Alice proposes to disable
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Alice rejects the proposal she just made
		txResult, err = RejectDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Should not be in proposed list
		proposed, err := GetProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, proposed)

		// Should not be in rejected list (single proposer rejection clears it immediately)
		rejected, err := GetRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, rejected)

		// Should not be in disabled list
		disabled, err := GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, disabled)
	})

	t.Run("ProposeApproveDisableAndReEnable", func(t *testing.T) {
		// Alice proposes to disable
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Proposed list contains the validator
		allProposed, err := GetAllProposedDisableNodes()
		require.NoError(t, err)
		require.True(t, containsProposedByAddress(allProposed, validatorAddress))

		proposed, err := GetProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, validatorAddress, proposed.Address)
		require.True(t, proposedHasVoter(proposed, aliceAddr))

		// Bob approves — reaches threshold (ceil(2/3 * 3 trustees) = 2, proposer counts as 1),
		// validator becomes disabled and the proposal is removed.
		txResult, err = ApproveDisableNode(validatorAddress, bob)
		cliputils.RequireTxOK(t, txResult, err)

		// Bob cannot reject — proposal is gone (threshold was reached)
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Should no longer be in proposed list (threshold reached)
		allProposed, err = GetAllProposedDisableNodes()
		require.NoError(t, err)
		require.False(t, containsProposedByAddress(allProposed, validatorAddress))

		// Should be in disabled list
		disabled, err := GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, disabled)
		require.Equal(t, validatorAddress, disabled.Address)
		require.False(t, disabled.DisabledByNodeAdmin)

		// Node admin (validatorOwner) re-enables the validator
		txResult, err = EnableNode(validatorOwner)
		cliputils.RequireTxOK(t, txResult, err)

		// Should no longer be disabled
		disabled, err = GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, disabled)
	})

	t.Run("NodeAdminSelfDisableAndReEnable", func(t *testing.T) {
		// The node admin (validator owner) disables its own validator
		// (validator-demo.sh:248-294). Unlike the trustee-voting path, this sets
		// disabledByNodeAdmin=true with no approvals, and jails the validator
		// synchronously.
		txResult, err := DisableNode(validatorOwner)
		cliputils.RequireTxOK(t, txResult, err)

		disabled, err := GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, disabled)
		require.Equal(t, validatorAddress, disabled.Address)
		require.True(t, disabled.DisabledByNodeAdmin)
		require.Empty(t, disabled.Approvals)

		// The validator is jailed while disabled.
		v, err := GetNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.True(t, v.Jailed)

		// The node admin re-enables: the disabled record clears and the
		// validator is unjailed (handler calls Unjail + RemoveDisabledValidator).
		txResult, err = EnableNode(validatorOwner)
		cliputils.RequireTxOK(t, txResult, err)

		disabled, err = GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, disabled)

		v, err = GetNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.False(t, v.Jailed)
	})

	t.Run("ProposeApproveRejectRejectFailsSecondTime", func(t *testing.T) {
		// Alice proposes
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Bob approves — reaches threshold, validator becomes disabled and proposal is removed.
		txResult, err = ApproveDisableNode(validatorAddress, bob)
		cliputils.RequireTxOK(t, txResult, err)

		// Bob cannot reject — proposal is gone (threshold was reached)
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Re-enable
		txResult, err = EnableNode(validatorOwner)
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("ProposeRejectByBobAndJack_GoesToRejectedList", func(t *testing.T) {
		// Alice proposes
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Bob rejects
		txResult, err = RejectDisableNode(validatorAddress, bob)
		cliputils.RequireTxOK(t, txResult, err)

		// Bob cannot reject twice
		txBad, errBad := RejectDisableNode(validatorAddress, bob)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Still in proposed list (not enough rejections yet)
		proposed, err := GetProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, validatorAddress, proposed.Address)
		require.True(t, proposedHasVoter(proposed, aliceAddr))
		require.True(t, proposedHasVoter(proposed, bobAddr))

		// Jack rejects — now enough rejections to move to rejected list
		txResult, err = RejectDisableNode(validatorAddress, jack)
		cliputils.RequireTxOK(t, txResult, err)

		// Jack cannot reject twice
		txBad, errBad = RejectDisableNode(validatorAddress, jack)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// No longer in proposed list
		allProposed, err := GetAllProposedDisableNodes()
		require.NoError(t, err)
		require.False(t, containsProposedByAddress(allProposed, validatorAddress))

		// Should be in rejected list
		allRejected, err := GetAllRejectedDisableNodes()
		require.NoError(t, err)
		require.True(t, containsRejectedByAddress(allRejected, validatorAddress))

		rejected, err := GetRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, rejected)
		require.Equal(t, validatorAddress, rejected.Address)
		require.True(t, rejectedHasVoter(rejected, aliceAddr))
		require.True(t, rejectedHasVoter(rejected, bobAddr))

		// Should NOT be disabled
		disabled, err := GetDisabledNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, disabled)

		// Should NOT be jailed (validator must still exist on chain)
		v, err := GetNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.False(t, v.Jailed)
	})

	t.Run("RePropose_AfterRejected_StartsFresh", func(t *testing.T) {
		// Alice proposes again — should succeed even though it was previously rejected
		txResult, err := ProposeDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Now in proposed list
		proposed, err := GetProposedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, validatorAddress, proposed.Address)
		require.True(t, proposedHasVoter(proposed, aliceAddr))

		// Rejected list should no longer have it
		rejected, err := GetRejectedDisableNode(validatorAddress)
		require.NoError(t, err)
		require.Nil(t, rejected)

		// Clean up: reject it to leave the network in a clean state
		txResult, err = RejectDisableNode(validatorAddress, alice)
		cliputils.RequireTxOK(t, txResult, err)
	})
}

// resolveFirstValidator queries all validator nodes and returns (ownerAccountName, validatorAddress)
// for the first validator it finds whose owner key exists in the local keyring.
func resolveFirstValidator(t *testing.T) (ownerAccountName, validatorAddress string) {
	t.Helper()

	nodes, err := GetAllNodes()
	require.NoError(t, err)

	// The validator "owner" field in all-nodes is a cosmosvaloper... address.
	// Use --bech val to get the cosmosvaloper address for each well-known key and
	// match against the on-chain owners. Prefer anna then jack so the validator
	// under test is not owned by alice or bob (the propose/approve signers).
	knownAccounts := []string{"anna", "jack", "alice", "bob"}
	for _, acc := range knownAccounts {
		valAddrOut, err := utils.ExecuteCLI("keys", "show", acc, "--bech", "val", "-a", "--keyring-backend", "test")
		if err != nil {
			continue
		}
		valAddr := strings.TrimSpace(string(valAddrOut))
		if valAddr != "" && containsValidatorByOwner(nodes, valAddr) {
			return acc, valAddr
		}
	}

	require.Fail(t, "could not find a known validator node admin account in the localnet")

	return "", ""
}
