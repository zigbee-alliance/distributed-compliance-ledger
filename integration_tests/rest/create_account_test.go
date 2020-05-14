package rest_test

//nolint:goimports
import (
	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
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

func Test_CreateNewAccount(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(test_constants.AccountName)

	// Create keys for test account
	testAccountKeyInfo, _ := utils.CreateKey(utils.RandString())

	// Register test account on the ledger
	res, _ := utils.CreateAccount(testAccountKeyInfo, jackKeyInfo)
	require.NotNil(t, res)

	// Get info for test account
	testAccountInfo, _ := utils.GetAccountInfo(testAccountKeyInfo.Address)
	require.Equal(t, testAccountKeyInfo.Address, testAccountInfo.Address)
	require.Equal(t, testAccountKeyInfo.PublicKey, testAccountInfo.PublicKey)

	// Assign Vendor role to test account
	utils.AssignRole(testAccountKeyInfo.Address, jackKeyInfo, auth.Vendor)

	// Publish model info by test account
	modelInfo := utils.NewMsgAddModelInfo(testAccountKeyInfo.Address)
	_, _ = utils.PublishModelInfo(modelInfo, testAccountKeyInfo)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(modelInfo.VID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
}
