// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:testpackage
package keeper

//nolint:goimports
import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestKeeper_TestingResultGetSet(t *testing.T) {
	setup := Setup()

	// check if testing result present
	require.False(t, setup.CompliancetestKeeper.IsTestingResultsPresents(
		setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion))

	// empty testing result before its created
	receivedTestingResult := setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)
	require.Equal(t, 0, len(receivedTestingResult.Results))

	// create testing results
	testingResult := DefaultTestingResult()
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	// check if testing result present
	require.True(t, setup.CompliancetestKeeper.IsTestingResultsPresents(
		setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion))

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)
	require.Equal(t, receivedTestingResult.VID, testconstants.VID)
	require.Equal(t, receivedTestingResult.PID, testconstants.PID)
	require.Equal(t, receivedTestingResult.SoftwareVersion, testconstants.SoftwareVersion)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)

	// add second testing result
	secondTestingResult := DefaultTestingResult()
	secondTestingResult.Owner = testconstants.Address2
	secondTestingResult.TestResult = "Second Test Result"
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, secondTestingResult)

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(
		setup.Ctx, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)
	require.Equal(t, receivedTestingResult.VID, testconstants.VID)
	require.Equal(t, receivedTestingResult.PID, testconstants.PID)
	require.Equal(t, receivedTestingResult.SoftwareVersion, testconstants.SoftwareVersion)
	require.Equal(t, 2, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
	CheckTestingResult(t, receivedTestingResult.Results[1], secondTestingResult)
}
