package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	To Run test you need:
		* prepare config with `genlocalconfig.sh`
		* update `/.zbld/config/genesis.json` to set `Administrator` role to the first account as described in Readme (#Genesis template)
		* run node with `zbld start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
*/

func TestCompliancetestDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo := utils.GetKeyInfo(test_constants.AccountName)

	//Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Publish model info
	modelInfo := utils.NewModelInfo(jackKeyInfo.Address)
	utils.PublishModelInfo(jackKeyInfo.Address, modelInfo)

	//Assign TestHouse role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.TestHouse)

	// Publish first testing result using Sign and Broadcast AddTestingResult message
	firstTestingResult := compliancetest.NewMsgAddTestingResult(
		modelInfo.VID,
		modelInfo.PID,
		utils.RandString(),
		jackKeyInfo.Address,
	)

	utils.SignAndBroadcastMessage(jackKeyInfo, firstTestingResult)

	// Check testing result is created
	receivedTestingResult := utils.GetTestingResult(firstTestingResult.VID, firstTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, firstTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, firstTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, firstTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].Owner, firstTestingResult.Signer)

	// Publish second model info
	secondModelInfo := utils.NewModelInfo(jackKeyInfo.Address)
	utils.PublishModelInfo(jackKeyInfo.Address, secondModelInfo)

	// Publish second testing result using POST
	secondTestingResult := compliancetest.NewTestingResult(
		secondModelInfo.VID,
		secondModelInfo.PID,
		utils.RandString(),
		jackKeyInfo.Address,
	)

	utils.PublishTestingResult(jackKeyInfo.Address, secondTestingResult)

	// Check testing result is created
	receivedTestingResult = utils.GetTestingResult(secondTestingResult.VID, secondTestingResult.PID)
	require.Equal(t, receivedTestingResult.VID, secondTestingResult.VID)
	require.Equal(t, receivedTestingResult.PID, secondTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].TestResult, secondTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[0].Owner, secondTestingResult.Owner)
}
