package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/*
	To Run test you need:
		* prepare config with `genlocalconfig.sh`
		* update `/.zbld/config/genesis.json` to set `Administrator` role to the first account as described in Readme (#Genesis template)
		* run node with `zbld start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
	TODO: provide tests for error cases
*/

func TestComplianceDemo(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo := utils.GetKeyInfo(test_constants.AccountName)

	// Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Get all certified models
	inputCertifiedModels := utils.GetAllCertifiedModels()

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(jackKeyInfo.Address)
	utils.PublishModelInfo(modelInfo)

	// Assign TestHouse role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.TestHouse)

	// Publish testing result
	testingResult := utils.NewMsgAddTestingResult(modelInfo.VID, modelInfo.PID, jackKeyInfo.Address)
	utils.PublishTestingResult(testingResult)

	// Assign ZBCertificationCenter role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.ZBCertificationCenter)

	// Certify model
	certifyModelMsg := compliance.NewMsgCertifyModel(modelInfo.VID, modelInfo.PID, time.Now().UTC(),
		test_constants.CertificationType, jackKeyInfo.Address)
	utils.PublishCertifiedModel(certifyModelMsg)

	// Check model is certified
	certifiedModel := utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID)
	require.Equal(t, certifiedModel.VID, modelInfo.VID)
	require.Equal(t, certifiedModel.PID, modelInfo.PID)
	require.Equal(t, certifiedModel.CertificationType, certifyModelMsg.CertificationType)
	require.Equal(t, certifiedModel.CertificationDate, certifyModelMsg.CertificationDate)

	// Get all certified models
	certifiedModels := utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total)+1, utils.ParseUint(certifiedModels.Total))
}
