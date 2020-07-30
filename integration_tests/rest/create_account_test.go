package rest_test

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
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
	// Query all proposed accounts
	inputProposedAccounts, _ := utils.GetProposedAccounts()

	// Query all accounts
	inputAccounts, _ := utils.GetAccounts()

	// Create keys for new account
	name := utils.RandString()
	testAccountKeyInfo, _ := utils.CreateKey(name)

	// Jack and Alice are predefined Trustees
	jackKeyInfo, _ := utils.GetKeyInfo(testconstants.JackAccount)
	aliceKeyInfo, _ := utils.GetKeyInfo(testconstants.AliceAccount)

	// Jack propose new account
	utils.ProposeAccount(testAccountKeyInfo, jackKeyInfo, auth.AccountRoles{auth.Vendor})

	// Query all proposed accounts
	receivedProposedAccounts, _ := utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items) + 1, len(receivedProposedAccounts.Items))

	// Query all accounts
	receivedAccounts, _ := utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items), len(receivedAccounts.Items))

	// Alice approve new account
	utils.ApproveAccount(testAccountKeyInfo, aliceKeyInfo)

	// Query all proposed accounts
	receivedProposedAccounts, _ = utils.GetProposedAccounts()
	require.Equal(t, len(inputProposedAccounts.Items), len(receivedProposedAccounts.Items))

	// Query all accounts
	receivedAccounts, _ = utils.GetAccounts()
	require.Equal(t, len(inputAccounts.Items) + 1, len(receivedAccounts.Items))

	// Get info for test account
	testAccountInfo, _ := utils.GetAccount(testAccountKeyInfo.Address)
	require.Equal(t, testAccountKeyInfo.Address, testAccountInfo.Address)
	require.Equal(t, auth.AccountRoles{auth.Vendor}, testAccountInfo.Roles)

	// Publish model info by test account
	modelInfo := utils.NewMsgAddModelInfo(testAccountKeyInfo.Address)
	_, _ = utils.PublishModelInfo(modelInfo, testAccountKeyInfo)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(modelInfo.VID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
}
