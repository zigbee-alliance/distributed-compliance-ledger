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
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	upgradeInfoV120 = `{"binaries":{"linux/amd64":"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.0/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"}}`
	upgradeInfoV121 = `{"binaries":{"linux/amd64":"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"}}`
	upgradeInfoV141 = `{"binaries":{"linux/amd64":"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"}}`

	upgradeNameV120 = "v1.2.0"
	upgradeNameV121 = "v1.2.1"
	upgradeNameV141 = "v1.4.1"

	// A very large height so the upgrade is never actually executed during tests
	farFutureHeight = "10000000"
)

func TestUpgradeDemo(t *testing.T) {
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount
	jack := testconstants.JackAccount

	// Create a trustee account used as additional proposer/approver
	trusteeAccount := cliputils.CreateAccount(t, "Trustee")

	t.Run("ProposeApproveUpgrade_v1_2_0", func(t *testing.T) {
		// trusteeAccount proposes
		txResult, err := ProposeUpgrade(upgradeNameV120, farFutureHeight, trusteeAccount, ProposeUpgradeOpts{UpgradeInfo: upgradeInfoV120})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Verify proposed upgrade
		proposed, err := GetProposedUpgrade(upgradeNameV120)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, upgradeNameV120, proposed.Plan.Name)
		require.Equal(t, farFutureHeight, fmt.Sprintf("%d", proposed.Plan.Height))

		// alice approves
		txResult, err = ApproveUpgrade(upgradeNameV120, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// alice rejects (revotes)
		txResult, err = RejectUpgrade(upgradeNameV120, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// alice approves again
		txResult, err = ApproveUpgrade(upgradeNameV120, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Approved upgrade should NOT yet be in approved store (threshold not reached)
		approved, err := GetApprovedUpgrade(upgradeNameV120)
		require.NoError(t, err)
		require.Nil(t, approved)

		// Still in proposed
		proposed, err = GetProposedUpgrade(upgradeNameV120)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, upgradeNameV120, proposed.Plan.Name)

		// bob approves — threshold now reached
		txResult, err = ApproveUpgrade(upgradeNameV120, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Upgrade plan should now be scheduled
		plan, err := GetUpgradePlan()
		require.NoError(t, err)
		require.NotNil(t, plan)
		require.Equal(t, upgradeNameV120, plan.Name)
		require.Equal(t, farFutureHeight, fmt.Sprintf("%d", plan.Height))

		// Should be in approved store
		approved, err = GetApprovedUpgrade(upgradeNameV120)
		require.NoError(t, err)
		require.NotNil(t, approved)
		require.Equal(t, upgradeNameV120, approved.Plan.Name)

		// Should no longer be in proposed store
		proposed, err = GetProposedUpgrade(upgradeNameV120)
		require.NoError(t, err)
		require.Nil(t, proposed)
	})

	t.Run("ProposerCannotApproveOwnUpgrade", func(t *testing.T) {
		upgradeName := fmt.Sprintf("upgrade_%s", utils.RandString())

		txResult, err := ProposeUpgrade(upgradeName, farFutureHeight, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveUpgrade(upgradeName, alice)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code) // unauthorized
	})

	t.Run("CannotApproveTwice", func(t *testing.T) {
		upgradeName := fmt.Sprintf("upgrade_%s", utils.RandString())

		txResult, err := ProposeUpgrade(upgradeName, farFutureHeight, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveUpgrade(upgradeName, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveUpgrade(upgradeName, bob)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code) // unauthorized / already approved
	})

	t.Run("CannotProposeTwice", func(t *testing.T) {
		upgradeName := fmt.Sprintf("upgrade_%s", utils.RandString())

		txResult, err := ProposeUpgrade(upgradeName, farFutureHeight, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ProposeUpgrade(upgradeName, farFutureHeight, alice)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code) // proposed upgrade already exists
	})

	t.Run("UpgradeHeightInPast_Fails", func(t *testing.T) {
		upgradeName := fmt.Sprintf("upgrade_%s", utils.RandString())

		txResult, err := ProposeUpgrade(upgradeName, "1", alice)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code) // upgrade cannot be scheduled in the past
	})

	t.Run("ProposeAndRejectUpgrade_v1_2_1", func(t *testing.T) {
		// Use a fresh far-future height
		txResult, err := ProposeUpgrade(upgradeNameV121, farFutureHeight, trusteeAccount, ProposeUpgradeOpts{UpgradeInfo: upgradeInfoV121})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// alice approves
		txResult, err = ApproveUpgrade(upgradeNameV121, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed
		proposed, err := GetProposedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, upgradeNameV121, proposed.Plan.Name)

		// trusteeAccount rejects (revotes)
		txResult, err = RejectUpgrade(upgradeNameV121, trusteeAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// trusteeAccount approves again
		txResult, err = ApproveUpgrade(upgradeNameV121, trusteeAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// alice rejects
		txResult, err = RejectUpgrade(upgradeNameV121, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed (not enough rejections)
		proposed, err = GetProposedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, upgradeNameV121, proposed.Plan.Name)

		// Not yet rejected or approved
		rejected, err := GetRejectedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.Nil(t, rejected)

		approved, err := GetApprovedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.Nil(t, approved)

		// bob rejects — threshold reached
		txResult, err = RejectUpgrade(upgradeNameV121, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now in rejected store
		rejected, err = GetRejectedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.NotNil(t, rejected)
		require.Equal(t, upgradeNameV121, rejected.Plan.Name)

		// No longer in proposed
		proposed, err = GetProposedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.Nil(t, proposed)

		// Not approved
		approved, err = GetApprovedUpgrade(upgradeNameV121)
		require.NoError(t, err)
		require.Nil(t, approved)
	})

	t.Run("ProposeAndRejectByProposer_v1_4_1", func(t *testing.T) {
		h, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := fmt.Sprintf("%d", h+10000000)

		// jack proposes
		txResult, err := ProposeUpgrade(upgradeNameV141, planHeight, jack, ProposeUpgradeOpts{UpgradeInfo: upgradeInfoV141})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// jack rejects own proposal
		txResult, err = RejectUpgrade(upgradeNameV141, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Should not be in proposed
		proposed, err := GetProposedUpgrade(upgradeNameV141)
		require.NoError(t, err)
		require.Nil(t, proposed)

		// Should not be in rejected
		rejected, err := GetRejectedUpgrade(upgradeNameV141)
		require.NoError(t, err)
		require.Nil(t, rejected)

		// Should not be in approved
		approved, err := GetApprovedUpgrade(upgradeNameV141)
		require.NoError(t, err)
		require.Nil(t, approved)
	})
}
