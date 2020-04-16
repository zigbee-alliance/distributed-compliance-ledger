package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryCertifiedModel(t *testing.T) {
	setup := Setup()

	// add certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetCertifiedModel(setup.Ctx, certifiedModel)

	// query certified model
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryCertifiedModel, fmt.Sprintf("%v", certifiedModel.VID), fmt.Sprintf("%v", certifiedModel.PID)},
		abci.RequestQuery{},
	)

	var receivedModels types.CertifiedModel
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModels)

	// check
	require.Equal(t, receivedModels.VID, certifiedModel.VID)
	require.Equal(t, receivedModels.PID, certifiedModel.PID)
	require.Equal(t, receivedModels.CertificationDate, certifiedModel.CertificationDate)
	require.Equal(t, receivedModels.CertificationType, certifiedModel.CertificationType)
}

func TestQuerier_QueryCertifiedModelForUnknown(t *testing.T) {
	setup := Setup()

	// query certified model
	result, err := setup.Querier(
		setup.Ctx,
		[]string{QueryCertifiedModel, fmt.Sprintf("%v", test_constants.VID), fmt.Sprintf("%v", test_constants.PID)},
		abci.RequestQuery{},
	)

	// check
	require.Nil(t, result)
	require.NotNil(t, err)
	require.Equal(t, types.CodeDeviceComplianceDoesNotExist, err.Code())
}

func TestQuerier_QueryAllCertifiedModels(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 certified models
	firstId := PopulateStoreWithCertifiedModels(setup, count)

	// query all certified models
	params := pagination.NewPaginationParams(0, 0)
	receiveModels := getCertifiedModels(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, count, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
		require.Equal(t, int16(i)+firstId, item.VID)
		require.Equal(t, int16(i)+firstId, item.PID)
	}
}

func TestQuerier_QueryAllCertifiedModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 certified models
	firstId := PopulateStoreWithCertifiedModels(setup, count)

	// query all certified models skip=1 take=2
	skip := 1
	take := 2
	params := pagination.NewPaginationParams(skip, take)
	receiveModels := getCertifiedModels(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, take, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
		require.Equal(t, int16(skip)+int16(i)+firstId, item.VID)
	}
}

func getCertifiedModels(setup TestSetup, params pagination.PaginationParams) types.ListCertifiedModelItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllCertifiedModels},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.ListCertifiedModelItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}
