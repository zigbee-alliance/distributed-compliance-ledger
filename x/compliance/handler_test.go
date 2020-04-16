package compliance

import (
	"fmt"
	constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_CertifyModel(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certified model
	receivedTestingResult := queryCertifiedModel(setup, vid, pid)

	// check
	require.Equal(t, receivedTestingResult.VID, certifyModelMsg.VID)
	require.Equal(t, receivedTestingResult.PID, certifyModelMsg.PID)
	require.Equal(t, receivedTestingResult.CertificationDate, certifyModelMsg.CertificationDate)
	require.Equal(t, receivedTestingResult.CertificationType, certifyModelMsg.CertificationType)
}

func TestHandler_CertifyModelByDifferentRoles(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	cases := []authz.AccountRole{
		authz.Administrator,
		authz.Vendor,
		authz.TestHouse,
	}

	for _, tc := range cases {
		address := constants.Address2
		setup.AuthzKeeper.AssignRole(setup.Ctx, address, tc)

		// try to certify model
		certifyModelMsg := msgCertifyModel(address, vid, pid)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CertifyModelForUnknownModel(t *testing.T) {
	setup := Setup()

	// try to certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, constants.VID, constants.PID)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, modelinfo.CodeModelInfoDoesNotExist, result.Code)
}

func TestHandler_CertifyModelForModelWithoutTestingResults(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, constants.VID, constants.PID)

	// try to certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	result := setup.Handler(setup.Ctx, certifyModelMsg)
	require.Equal(t, compliancetest.CodeTestingResultDoesNotExist, result.Code)
}

func TestHandler_CertifyModelTwice(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	cases := []struct {
		expectedCode sdk.CodeType
		address      sdk.AccAddress
	}{
		{sdk.CodeOK, setup.ZBCertificationCenter(constants.Address1)},
		{types.CodeDeviceComplianceAlreadyExists, setup.ZBCertificationCenter(constants.Address1)}, // same address
		{types.CodeDeviceComplianceAlreadyExists, setup.ZBCertificationCenter(constants.Address2)}, // different address
	}

	for _, tc := range cases {
		certifyModelMsg := msgCertifyModel(tc.address, vid, pid)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, tc.expectedCode, result.Code)
	}
}

func TestHandler_CertifyDifferentModels(t *testing.T) {
	setup := Setup()

	for i := int16(1); i < int16(5); i++ {
		// add model amd testing result
		vid, pid := addModel(setup, i, i)
		addTestingResult(setup, vid, pid)

		// add new testing result
		certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
		result := setup.Handler(setup.Ctx, certifyModelMsg)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certified model
		receivedModel := queryCertifiedModel(setup, vid, pid)

		// check
		require.Equal(t, receivedModel.VID, certifyModelMsg.VID)
		require.Equal(t, receivedModel.PID, certifyModelMsg.PID)
		require.Equal(t, receivedModel.CertificationDate, certifyModelMsg.CertificationDate)
		require.Equal(t, receivedModel.CertificationType, certifyModelMsg.CertificationType)
	}
}

func TestHandler_CertifyModelForEmptyCertificationType(t *testing.T) {
	setup := Setup()

	// add model amd testing result
	vid, pid := addModel(setup, constants.VID, constants.PID)
	addTestingResult(setup, vid, pid)

	// certify model
	certifyModelMsg := msgCertifyModel(setup.CertificationCenter, vid, pid)
	certifyModelMsg.CertificationType = ""
	setup.Handler(setup.Ctx, certifyModelMsg)

	// query certified model
	receivedTestingResult := queryCertifiedModel(setup, vid, pid)

	// check default type is set
	require.Equal(t, receivedTestingResult.CertificationType, types.ZbCertificationType)
}

func queryCertifiedModel(setup TestSetup, vid int16, pid int16) types.CertifiedModel {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryCertifiedModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var model types.CertifiedModel
	_ = setup.Cdc.UnmarshalJSON(result, &model)
	return model
}

func addModel(setup TestSetup, vid int16, pid int16) (int16, int16) {
	modelInfo := modelinfo.ModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      constants.CID,
		Owner:                    constants.Owner,
		Name:                     constants.Name,
		Description:              constants.Description,
		SKU:                      constants.Sku,
		FirmwareVersion:          constants.FirmwareVersion,
		HardwareVersion:          constants.HardwareVersion,
		Custom:                   constants.Custom,
		TisOrTrpTestingCompleted: constants.TisOrTrpTestingCompleted,
	}

	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)
	return vid, pid
}

func addTestingResult(setup TestSetup, vid int16, pid int16) (int16, int16) {
	testingResult := compliancetest.TestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: constants.TestResult,
		Owner:      constants.Owner,
	}

	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)
	return vid, pid
}

func msgCertifyModel(signer sdk.AccAddress, vid int16, pid int16) MsgCertifyModel {
	return MsgCertifyModel{
		VID:               vid,
		PID:               pid,
		CertificationDate: constants.CertificationDate,
		CertificationType: constants.CertificationType,
		Signer:            signer,
	}
}
