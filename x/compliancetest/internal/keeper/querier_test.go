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

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	test_constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
)

func TestQuerier_QueryTestingResult(t *testing.T) {
	setup := Setup()

	// add testing result
	testingResult := DefaultTestingResult()
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryTestingResult, fmt.Sprintf("%v", testingResult.VID),
			fmt.Sprintf("%v", testingResult.PID), fmt.Sprintf("%v", testingResult.SoftwareVersion)},
		abci.RequestQuery{},
	)

	var receivedTestingResult types.TestingResults
	_ = setup.Cdc.UnmarshalJSON(result, &receivedTestingResult)

	// check
	require.Equal(t, receivedTestingResult.VID, testingResult.VID)
	require.Equal(t, receivedTestingResult.PID, testingResult.PID)
	require.Equal(t, receivedTestingResult.SoftwareVersion, testingResult.SoftwareVersion)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
}

func TestQuerier_QueryTestingResultForUnknown(t *testing.T) {
	setup := Setup()

	// query unknown testing result
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryTestingResult, fmt.Sprintf("%v", test_constants.VID), fmt.Sprintf("%v", test_constants.PID),
			fmt.Sprintf("%v", test_constants.SoftwareVersion)},
		abci.RequestQuery{},
	)
	require.Equal(t, types.CodeTestingResultsDoNotExist, err.Code())
}
