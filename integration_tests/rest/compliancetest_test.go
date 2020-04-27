package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
	TODO: provide tests for error cases
*/

func /*Test*/CompliancetestDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(test_constants.AccountName)

	//Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(jackKeyInfo.Address)
	_, _ = utils.PublishModelInfo(modelInfo)

	//Assign TestHouse role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.TestHouse)

	// Publish first testing result using Sign and Broadcast AddTestingResult message
	firstTestingResult := utils.NewMsgAddTestingResult(modelInfo.VID, modelInfo.PID, jackKeyInfo.Address)
	utils.SignAndBroadcastMessage(jackKeyInfo, firstTestingResult)

	// Check testing result is created
	receivedTestingResult, _ := utils.GetTestingResult(firstTestingResult.VID, firstTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, firstTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, firstTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].TestDate, firstTestingResult.TestDate)
	require.Equal(t, receivedTestingResult.Results[0].Owner, firstTestingResult.Signer)

	// Publish second model info
	secondModelInfo := utils.NewMsgAddModelInfo(jackKeyInfo.Address)
	_, _ = utils.PublishModelInfo(secondModelInfo)

	// Publish second testing result using POST
	secondTestingResult := utils.NewMsgAddTestingResult(secondModelInfo.VID, secondModelInfo.PID, jackKeyInfo.Address)
	_, _ = utils.PublishTestingResult(secondTestingResult)

	// Check testing result is created
	receivedTestingResult, _ = utils.GetTestingResult(secondTestingResult.VID, secondTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, secondTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, secondTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, secondTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].TestDate, secondTestingResult.TestDate)
	require.Equal(t, receivedTestingResult.Results[0].Owner, secondTestingResult.Signer)
}