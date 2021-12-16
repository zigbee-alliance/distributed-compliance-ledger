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
	//"net/http"
	"net/http"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `dclcli rest-server --chain-id dclchain`

	TODO: provide tests for error cases
*/

//nolint:funlen
func TestAuthDemo(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID)
	kr := suite.GetKeyring()

	jackName := testconstants.JackAccount
	aliceName := testconstants.AliceAccount
	bobName := testconstants.BobAccount

	// Query all active accounts
	inputAccounts, err := suite.GetAccounts()
	require.NoError(t, err)
	require.Equal(t, 4, len(inputAccounts))

	// build map with an acc address as a key
	accDataInitial := make(map[string]dclauthtypes.Account)
	for _, acc := range inputAccounts {
		accDataInitial[acc.GetAddress().String()] = acc
	}

	// Jack, Alice and Bob are predefined Trustees
	jackKeyInfo, err := kr.Key(jackName)
	require.NoError(t, err)
	require.Contains(t, accDataInitial, jackKeyInfo.GetAddress().String())
	jackSequence := accDataInitial[jackKeyInfo.GetAddress().String()].GetSequence()
	jackAccNum := accDataInitial[jackKeyInfo.GetAddress().String()].GetAccountNumber()

	aliceKeyInfo, err := kr.Key(aliceName)
	require.NoError(t, err)
	require.Contains(t, accDataInitial, aliceKeyInfo.GetAddress().String())
	aliceSequence := accDataInitial[aliceKeyInfo.GetAddress().String()].GetSequence()
	aliceAccNum := accDataInitial[aliceKeyInfo.GetAddress().String()].GetAccountNumber()

	bobKeyInfo, err := kr.Key(bobName)
	require.NoError(t, err)
	require.Contains(t, accDataInitial, bobKeyInfo.GetAddress().String())
	bobSequence := accDataInitial[bobKeyInfo.GetAddress().String()].GetSequence()
	bobAccNum := accDataInitial[bobKeyInfo.GetAddress().String()].GetAccountNumber()

	// Query all proposed accounts
	inputProposedAccounts, err := suite.GetProposedAccounts()
	require.NoError(t, err)
	require.Equal(t, 0, len(inputProposedAccounts))

	// Query all proposed accounts to revoke
	inputProposedAccountsToRevoke, err := suite.GetProposedAccountsToRevoke()
	require.NoError(t, err)
	require.Equal(t, 0, len(inputProposedAccountsToRevoke))

	_, testAccPubKey, testAccAddr := testdata.KeyTestPubAddr()

	/*
		// Create keys for new account
		name := utils.RandString()
		testAccountKeyInfo, _ := utils.CreateKey(name)

	*/

	// Jack proposes new account
	_, err = suite.ProposeAddAccount(
		jackName, testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, testconstants.VID,
		jackAccNum, jackSequence,
	)
	require.NoError(t, err)
	jackSequence += 1

	// Query all active accounts
	receivedAccounts, _ := suite.GetAccounts()
	require.Equal(t, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ := suite.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ := suite.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Alice approves new account
	_, err = suite.ApproveAddAccount(aliceName, testAccAddr, aliceAccNum, aliceSequence)
	require.NoError(t, err)
	aliceSequence += 1

	// Query all active accounts
	receivedAccounts, _ = suite.GetAccounts()
	require.Equal(t, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = suite.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = suite.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Get info for new account
	testAccount, err := suite.GetAccount(testAccAddr)
	require.NoError(t, err)
	require.Equal(t, testAccAddr, testAccount.GetAddress())
	require.Equal(t, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testAccount.GetRoles())

	// FIXME issue 99: enable once implemented
	/*
		// Publish model info by test account
		model := suite.NewMsgAddModel(testAccountKeyInfo.Address, testconstants.VID)
		_, _ = suite.AddModel(model, testAccountKeyInfo)

		// Check model is created
		receivedModel, _ := suite.GetModel(model.VID, model.PID)
		require.Equal(t, receivedModel.VID, model.VID)
		require.Equal(t, receivedModel.PID, model.PID)
		require.Equal(t, receivedModel.ProductName, model.ProductName)
	*/

	// Alice proposes to revoke new account
	_, err = suite.ProposeRevokeAccount(aliceName, testAccAddr, aliceAccNum, aliceSequence)
	require.NoError(t, err)
	aliceSequence += 1

	// Query all active accounts
	receivedAccounts, err = suite.GetAccounts()
	require.Equal(t, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = suite.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = suite.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke)+1, len(receivedProposedAccountsToRevoke))

	// Bob approves to revoke new account
	_, err = suite.ApproveRevokeAccount(bobName, testAccAddr, bobAccNum, bobSequence)
	require.NoError(t, err)
	bobSequence += 1

	// Query all active accounts
	receivedAccounts, err = suite.GetAccounts()
	require.Equal(t, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = suite.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = suite.GetProposedAccountsToRevoke()
	require.Equal(t, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Ensure that new account is not available anymore
	res, err := suite.GetAccount(testAccAddr)
	_ = res
	require.Equal(t, http.StatusNotFound, err)

	// FIXME issue 99: enable once implemented
	/*
		// Try to publish another model info by test account.
		// Ensure that the request is responded with not OK status code.
		model = suite.NewMsgAddModel(testAccountKeyInfo.Address, testconstants.VID)
		_, code = suite.AddModel(model, testAccountKeyInfo)
		require.NotEqual(t, http.StatusOK, code)
	*/
}
