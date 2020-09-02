//nolint:testpackage
package keeper

import (
	"fmt"
	"testing"

	test_constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestQuerier_QueryTestingResult(t *testing.T) {
	setup := Setup()

	// add testing result
	testingResult := DefaultTestingResult()
	setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)

	// query testing result
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryTestingResult, fmt.Sprintf("%v", testingResult.VID), fmt.Sprintf("%v", testingResult.PID)},
		abci.RequestQuery{},
	)

	var receivedTestingResult types.TestingResults
	_ = setup.Cdc.UnmarshalJSON(result, &receivedTestingResult)

	// check
	require.Equal(t, receivedTestingResult.VID, testingResult.VID)
	require.Equal(t, receivedTestingResult.PID, testingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
}

func TestQuerier_QueryTestingResultForUnknown(t *testing.T) {
	setup := Setup()

	// query unknown testing result
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryTestingResult, fmt.Sprintf("%v", test_constants.VID), fmt.Sprintf("%v", test_constants.PID)},
		abci.RequestQuery{},
	)
	require.Equal(t, types.CodeTestingResultsDoNotExist, err.Code())
}
