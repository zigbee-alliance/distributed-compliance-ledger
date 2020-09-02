//nolint:testpackage
package keeper

//nolint:goimports
import (
	"testing"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"github.com/stretchr/testify/require"
)

func TestKeeper_TestingResultGetSet(t *testing.T) {
	setup := Setup()

	// check if testing result present
	require.False(t, setup.CompliancetestKeeper.IsTestingResultsPresents(
		setup.Ctx, testconstants.VID, testconstants.PID))

	// empty testing result before its created
	receivedTestingResult := setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID)
	require.Equal(t, 0, len(receivedTestingResult.Results))

	// create testing results
	testingResult := DefaultTestingResult()
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	// check if testing result present
	require.True(t, setup.CompliancetestKeeper.IsTestingResultsPresents(
		setup.Ctx, testconstants.VID, testconstants.PID))

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID)
	require.Equal(t, receivedTestingResult.VID, testconstants.VID)
	require.Equal(t, receivedTestingResult.PID, testconstants.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)

	// add second testing result
	secondTestingResult := DefaultTestingResult()
	secondTestingResult.Owner = testconstants.Address2
	secondTestingResult.TestResult = "Second Test Result"
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, secondTestingResult)

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID)
	require.Equal(t, receivedTestingResult.VID, testconstants.VID)
	require.Equal(t, receivedTestingResult.PID, testconstants.PID)
	require.Equal(t, 2, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
	CheckTestingResult(t, receivedTestingResult.Results[1], secondTestingResult)
}
