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

package rest_test

import (
	"net/http"
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"github.com/stretchr/testify/require"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

//nolint:funlen
func TestAuthDemo(t *testing.T) {
	// Query all active accounts
	inputAccounts, _ := utils.GetAccounts()

	// Query all proposed accounts
	inputProposedAccounts, _ := utils.GetProposedAccounts()

	// Query all proposed accounts to revoke
	inputProposedAccountsToRevoke, _ := utils.GetProposedAccountsToRevoke()

	// Create keys for new account
	name := utils.RandString()
	testAccountKeyInfo, _ := utils.CreateKey(name)

	// Jack, Alice and Bob are predefined Trustees
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.JackAccount)
	aliceKeyInfo, _ := utils.GetKeyInfo(testconstants.AliceAccount)
	bobKeyInfo, _ := utils.GetKeyInfo(testconstants.BobAccount)

	// Jack proposes new account
	utils.ProposeAddAccount(testAccountKeyInfo, jackKeyInfo, auth.AccountRoles{auth.Vendor})

	// Query all active accounts
	receivedAccounts, _ := utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items), len(receivedAccounts.Items))

	// Query all proposed accounts
	receivedProposedAccounts, _ := utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items)+1, len(receivedProposedAccounts.Items))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ := utils.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke.Items), len(receivedProposedAccountsToRevoke.Items))

	// Alice approves new account
	utils.ApproveAddAccount(testAccountKeyInfo, aliceKeyInfo)

	// Query all active accounts
	receivedAccounts, _ = utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items)+1, len(receivedAccounts.Items))

	// Query all proposed accounts
	receivedProposedAccounts, _ = utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items), len(receivedProposedAccounts.Items))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = utils.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke.Items), len(receivedProposedAccountsToRevoke.Items))

	// Get info for new account
	testAccount, _ := utils.GetAccount(testAccountKeyInfo.Address)
	require.Equal(t, testAccountKeyInfo.Address, testAccount.Address)
	require.Equal(t, auth.AccountRoles{auth.Vendor}, testAccount.Roles)

	// Publish model info by test account
	modelInfo := utils.NewMsgAddModelInfo(testAccountKeyInfo.Address)
	_, _ = utils.PublishModelInfo(modelInfo, testAccountKeyInfo)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(modelInfo.VID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)

	// Alice proposes to revoke new account
	utils.ProposeRevokeAccount(testAccountKeyInfo, aliceKeyInfo)

	// Query all active accounts
	receivedAccounts, _ = utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items)+1, len(receivedAccounts.Items))

	// Query all proposed accounts
	receivedProposedAccounts, _ = utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items), len(receivedProposedAccounts.Items))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = utils.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke.Items)+1, len(receivedProposedAccountsToRevoke.Items))

	// Bob approves to revoke new account
	utils.ApproveRevokeAccount(testAccountKeyInfo, bobKeyInfo)

	// Query all active accounts
	receivedAccounts, _ = utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items), len(receivedAccounts.Items))

	// Query all proposed accounts
	receivedProposedAccounts, _ = utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items), len(receivedProposedAccounts.Items))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = utils.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke.Items), len(receivedProposedAccountsToRevoke.Items))

	// Ensure that new account is not available anymore
	_, code := utils.GetAccount(testAccountKeyInfo.Address)
	require.Equal(t, http.StatusNotFound, code)

	// Try to publish another model info by test account.
	// Ensure that the request is responded with not OK status code.
	modelInfo = utils.NewMsgAddModelInfo(testAccountKeyInfo.Address)
	_, code = utils.PublishModelInfo(modelInfo, testAccountKeyInfo)
	require.NotEqual(t, http.StatusOK, code)
}
