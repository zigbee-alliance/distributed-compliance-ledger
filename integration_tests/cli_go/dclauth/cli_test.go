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

package dclauth_test_cli

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli_go/helpers"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestAuthDemoCLI(t *testing.T) {
	suite := utils.SetupTest(t, testconstants.ChainID, false)

	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	user1 := helpers.CreateAccountInfo(&suite)

	// Propose user1 account by jack
	txResult, err := ProposeAccount(user1.Address, user1.Key, dclauthtypes.NodeAdmin, jack)
	require.NoError(suite.T, err)
	require.Equal(suite.T, txResult.Code, uint32(0))

	// Approve user1 account by alice
	txResult, err = ApproveAccount(user1.Address, alice)
	require.NoError(suite.T, err)
	require.Equal(suite.T, txResult.Code, uint32(0))

	// await transaction is written
	_, err = helpers.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(suite.T, err)

	// Query list of all active accounts
	accounts, err := QueryAccounts()
	require.NoError(suite.T, err)
	require.True(suite.T, AccountIsInList(user1.Address, accounts.Account))

	// Query user1 account
	account, err := QueryAccount(user1.Address)
	require.NoError(suite.T, err)
	require.Equal(suite.T, account.Address, user1.Address)
}
