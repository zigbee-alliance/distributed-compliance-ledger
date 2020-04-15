package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_TestingResultGetSet(t *testing.T) {
	setup := Setup()

	// check if testing result present
	require.False(t, setup.CompliancetestKeeper.IsTestingResultsPresents(setup.Ctx, test_constants.VID, test_constants.PID))

	// empty testing result before its created
	receivedTestingResult := setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, 0, len(receivedTestingResult.Results))

	// create testing results
	testingResult := DefaultTestingResult()
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	// check if testing result present
	require.True(t, setup.CompliancetestKeeper.IsTestingResultsPresents(setup.Ctx, test_constants.VID, test_constants.PID))

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, test_constants.VID, receivedTestingResult.VID)
	require.Equal(t, test_constants.PID, receivedTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].Owner, testingResult.Owner)
	require.Equal(t, receivedTestingResult.Results[0].TestResult, testingResult.TestResult)

	// add second testing result
	secondTestingResult := DefaultTestingResult()
	secondTestingResult.Owner = test_constants.Address2
	secondTestingResult.TestResult = "Second Test Result"
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, secondTestingResult)

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, test_constants.VID, receivedTestingResult.VID)
	require.Equal(t, test_constants.PID, receivedTestingResult.PID)
	require.Equal(t, 2, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0].Owner, testingResult.Owner)
	require.Equal(t, receivedTestingResult.Results[0].TestResult, testingResult.TestResult)
	require.Equal(t, receivedTestingResult.Results[1].Owner, secondTestingResult.Owner)
	require.Equal(t, receivedTestingResult.Results[1].TestResult, secondTestingResult.TestResult)
}
