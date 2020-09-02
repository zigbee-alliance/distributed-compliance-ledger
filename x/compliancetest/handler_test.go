//nolint:testpackage
package compliancetest

import (
	"fmt"
	"testing"

	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandler_AddTestingResult(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query testing result
	receivedTestingResult := queryTestingResult(setup, vid, pid)

	// check
	require.Equal(t, receivedTestingResult.VID, vid)
	require.Equal(t, receivedTestingResult.PID, pid)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
}

func TestHandler_AddTestingResultByNonTestHouse(t *testing.T) {
	setup := Setup()
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	for _, role := range []auth.AccountRole{auth.Vendor, auth.ZBCertificationCenter, auth.NodeAdmin} {
		// store account
		account := auth.NewAccount(test_constants.Address3, test_constants.PubKey3, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new testing result by non TestHouse
		testingResult := TestMsgAddTestingResult(test_constants.Address3, vid, pid)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_AddTestingResultForUnknownModel(t *testing.T) {
	setup := Setup()

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, test_constants.VID, test_constants.PID)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, modelinfo.CodeModelInfoDoesNotExist, result.Code)
}

func TestHandler_AddSeveralTestingResultsForOneModel(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	for i, th := range []sdk.AccAddress{test_constants.Address1, test_constants.Address2, test_constants.Address3} {
		// store account
		account := auth.NewAccount(th, test_constants.PubKey1, auth.AccountRoles{auth.TestHouse})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new testing result
		testingResult := TestMsgAddTestingResult(th, vid, pid)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query testing result
		receivedTestingResult := queryTestingResult(setup, vid, pid)

		// check
		require.Equal(t, receivedTestingResult.VID, vid)
		require.Equal(t, receivedTestingResult.PID, pid)
		require.Equal(t, i+1, len(receivedTestingResult.Results))
		CheckTestingResult(t, receivedTestingResult.Results[i], testingResult)
	}
}

func TestHandler_AddSeveralTestingResultsForDifferentModels(t *testing.T) {
	setup := Setup()

	for i := uint16(1); i < uint16(5); i++ {
		// add model
		vid, pid := addModel(setup, i, i)

		// add new testing result
		testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query testing result
		receivedTestingResult := queryTestingResult(setup, vid, pid)

		// check
		require.Equal(t, receivedTestingResult.VID, vid)
		require.Equal(t, receivedTestingResult.PID, pid)
		require.Equal(t, 1, len(receivedTestingResult.Results))
		CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
	}
}

func TestHandler_AddTestingResultTwiceForSameModelAndSameTestHouse(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add testing result second time
	testingResult.TestResult = "Second Testing Result"
	result = setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query testing result
	receivedTestingResult := queryTestingResult(setup, vid, pid)

	// check
	require.Equal(t, 2, len(receivedTestingResult.Results))
}

func queryTestingResult(setup TestSetup, vid uint16, pid uint16) types.TestingResults {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryTestingResult, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedTestingResult types.TestingResults
	_ = setup.Cdc.UnmarshalJSON(result, &receivedTestingResult)

	return receivedTestingResult
}

func addModel(setup TestSetup, vid uint16, pid uint16) (uint16, uint16) {
	modelInfo := modelinfo.ModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      test_constants.CID,
		Owner:                    test_constants.Owner,
		Name:                     test_constants.Name,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		Custom:                   test_constants.Custom,
		TisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
	}

	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)

	return vid, pid
}
