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
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

func TestCompliancetestDemo(t *testing.T) {
	// Register new Vendor account
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Register new TestHouse account
	testHouse := utils.CreateNewAccount(auth.AccountRoles{auth.TestHouse})

	// Register new TestHouse account
	secondTestHouse := utils.CreateNewAccount(auth.AccountRoles{auth.TestHouse})

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	_, _ = utils.PublishModelInfo(modelInfo, vendor)

	// Publish first testing result using Sign and Broadcast AddTestingResult message
	firstTestingResult := utils.NewMsgAddTestingResult(modelInfo.VID, modelInfo.PID, testHouse.Address)
	utils.SignAndBroadcastMessage(testHouse, firstTestingResult)

	// Check testing result is created
	receivedTestingResult, _ := utils.GetTestingResult(firstTestingResult.VID, firstTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, firstTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, firstTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].TestDate, firstTestingResult.TestDate)
	require.Equal(t, receivedTestingResult.Results[0].Owner, firstTestingResult.Signer)

	// Publish second model info
	secondModelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	_, _ = utils.PublishModelInfo(secondModelInfo, vendor)

	// Publish second testing result using POST
	secondTestingResult := utils.NewMsgAddTestingResult(secondModelInfo.VID, secondModelInfo.PID, testHouse.Address)
	_, _ = utils.PublishTestingResult(secondTestingResult, testHouse)

	// Check testing result is created
	receivedTestingResult, _ = utils.GetTestingResult(secondTestingResult.VID, secondTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, secondTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, secondTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, secondTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].TestDate, secondTestingResult.TestDate)
	require.Equal(t, receivedTestingResult.Results[0].Owner, secondTestingResult.Signer)

	// Publish new testing result for second model
	thirdTestingResult := utils.NewMsgAddTestingResult(secondModelInfo.VID, secondModelInfo.PID, secondTestHouse.Address)
	_, _ = utils.PublishTestingResult(thirdTestingResult, secondTestHouse)

	// Check testing result is created
	receivedTestingResult, _ = utils.GetTestingResult(secondTestingResult.VID, secondTestingResult.PID)
	require.Equal(t, 2, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].Owner, secondTestingResult.Signer)
	require.Equal(t, receivedTestingResult.Results[0].TestResult, secondTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[1].Owner, thirdTestingResult.Signer)
	require.Equal(t, receivedTestingResult.Results[1].TestResult, thirdTestingResult.TestResult)
}
