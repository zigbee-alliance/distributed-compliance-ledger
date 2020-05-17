package rest_test

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"github.com/stretchr/testify/require"
	"testing"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: provide tests for error cases
*/

func TestCompliancetestDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.AccountName)

	// Register new Vendor account
	vendor, _ := utils.RegisterNewAccount()
	utils.AssignRole(vendor.Address, jackKeyInfo, authz.Vendor)

	// Register new TestHouse account
	testHouse, _ := utils.RegisterNewAccount()
	utils.AssignRole(testHouse.Address, jackKeyInfo, authz.TestHouse)

	// Register new TestHouse account
	secondTestHouse, _ := utils.RegisterNewAccount()
	utils.AssignRole(secondTestHouse.Address, jackKeyInfo, authz.TestHouse)

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
