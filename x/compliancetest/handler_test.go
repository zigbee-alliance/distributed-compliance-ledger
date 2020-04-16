package compliancetest

import (
	"fmt"
	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_AddTestingResult(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	// add new testing result
	testHouse := setup.TestHouse(test_constants.Address1)
	testingResult := TestMsgAddTestingResult(testHouse, vid, pid)
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

	cases := []sdk.AccAddress{
		setup.Vendor(test_constants.Address1),
		setup.Administrator(test_constants.Address1),
	}

	for _, tc := range cases {
		// add new testing result by non TestHouse
		testingResult := TestMsgAddTestingResult(tc, vid, pid)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_AddTestingResultForUnknownModel(t *testing.T) {
	setup := Setup()
	testHouse := setup.TestHouse(test_constants.Address1)

	// add new testing result
	testingResult := TestMsgAddTestingResult(testHouse, test_constants.VID, test_constants.PID)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, modelinfo.CodeModelInfoDoesNotExist, result.Code)
}

func TestHandler_AddSeveralTestingResultsForOneModel(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	testHouses := []sdk.AccAddress{
		setup.TestHouse(test_constants.Address1),
		setup.TestHouse(test_constants.Address2),
		setup.TestHouse(test_constants.Address3),
	}

	for i, th := range testHouses {
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
		CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
	}
}

func TestHandler_AddSeveralTestingResultsForDifferentModels(t *testing.T) {
	setup := Setup()
	testHouse := setup.TestHouse(test_constants.Address1)

	for i := int16(1); i < int16(5); i++ {
		// add model
		vid, pid := addModel(setup, i, i)

		// add new testing result
		testingResult := TestMsgAddTestingResult(testHouse, vid, pid)
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
	testHouse := setup.TestHouse(test_constants.Address1)

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)

	// add new testing result
	testingResult := TestMsgAddTestingResult(testHouse, vid, pid)
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

func queryTestingResult(setup TestSetup, vid int16, pid int16) types.TestingResults {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryTestingResult, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedTestingResult types.TestingResults
	_ = setup.Cdc.UnmarshalJSON(result, &receivedTestingResult)
	return receivedTestingResult
}

func addModel(setup TestSetup, vid int16, pid int16) (int16, int16) {
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
