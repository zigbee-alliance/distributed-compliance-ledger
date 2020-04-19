package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryComplianceInfo(t *testing.T) {
	setup := Setup()

	// add certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, certifiedModel)

	// query compliance info and check
	receivedComplianceInfo, _ := getComplianceInfo(setup, certifiedModel.VID, certifiedModel.PID)
	CheckComplianceInfo(t, certifiedModel, receivedComplianceInfo)

	// add revoked model
	revokedModel := DefaultRevokedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, revokedModel)

	// query compliance info and check
	receivedComplianceInfo, _ = getComplianceInfo(setup, revokedModel.VID, revokedModel.PID)
	CheckComplianceInfo(t, revokedModel, receivedComplianceInfo)
}

func TestQuerier_QueryComplianceInfoForUnknownModel(t *testing.T) {
	setup := Setup()

	// query compliance info and check
	_, err := getComplianceInfo(setup, test_constants.VID, test_constants.PID)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())

}

func TestQuerier_QueryCertifiedModel(t *testing.T) {
	setup := Setup()

	// add certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, certifiedModel)

	// query certified model
	receivedComplianceInfo, _ := getCertifiedModel(setup, certifiedModel.VID, certifiedModel.PID)

	// check
	CheckComplianceInfo(t, certifiedModel, receivedComplianceInfo)
}

func TestQuerier_QueryCertifiedModelForUnknown(t *testing.T) {
	setup := Setup()

	// query certified model
	_, err := getCertifiedModel(setup, test_constants.VID, test_constants.PID)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryCertifiedModelForModelInRevokedState(t *testing.T) {
	setup := Setup()

	// add revoked model
	revokedModel := DefaultRevokedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, revokedModel)

	// query certified model
	_, err := getCertifiedModel(setup, revokedModel.VID, revokedModel.PID)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryRevokedModel(t *testing.T) {
	setup := Setup()

	// add revoked model
	revokedModel := DefaultRevokedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, revokedModel)

	// query revoked model
	receivedComplianceInfo, _ := getRevokedModel(setup, revokedModel.VID, revokedModel.PID)

	// check
	CheckComplianceInfo(t, revokedModel, receivedComplianceInfo)
}

func TestQuerier_QueryRevokedModelForUnknown(t *testing.T) {
	setup := Setup()

	// query revoked model
	_, err := getRevokedModel(setup, test_constants.VID, test_constants.PID)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryRevokedModelForModelInRevokedState(t *testing.T) {
	setup := Setup()

	// add certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, certifiedModel)

	// query revoked model
	_, err := getRevokedModel(setup, certifiedModel.VID, certifiedModel.PID)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryAllModels(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstId := PopulateStoreWithMixedModels(setup, count)

	params := types.NewListQueryParams(types.EmptyCertificationType, 0, 0)

	cases := []struct {
		firstId       int16
		count         int
		receivedInfos types.ListComplianceInfoItems
	}{
		{firstId, count, getComplianceInfos(setup, params)},                    // query compliance infos
		{firstId, count / 2, getCertifiedModels(setup, params)},                // query certified model
		{firstId + int16(count/2), count / 2, getRevokedModels(setup, params)}, // query revoked models
	}

	for _, tc := range cases {
		// check
		require.Equal(t, tc.count, tc.receivedInfos.Total)
		require.Equal(t, tc.count, len(tc.receivedInfos.Items))

		for i, item := range tc.receivedInfos.Items {
			require.Equal(t, int16(i)+tc.firstId, item.VID)
			require.Equal(t, int16(i)+tc.firstId, item.PID)
		}
	}
}

func TestQuerier_QueryAllModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstId := PopulateStoreWithMixedModels(setup, count)

	// query all certified models skip=1 take=2
	skip := 1
	take := 2
	params := types.NewListQueryParams(types.EmptyCertificationType, skip, take)

	cases := []struct {
		firstId       int16
		count         int
		receivedInfos types.ListComplianceInfoItems
	}{
		{firstId, count, getComplianceInfos(setup, params)},                    // query compliance infos
		{firstId, count / 2, getCertifiedModels(setup, params)},                // query certified model
		{firstId + int16(count/2), count / 2, getRevokedModels(setup, params)}, // query revoked models
	}

	for _, tc := range cases {
		// check
		require.Equal(t, tc.count, tc.receivedInfos.Total)
		require.Equal(t, take, len(tc.receivedInfos.Items))

		for i, item := range tc.receivedInfos.Items {
			require.Equal(t, int16(skip)+int16(i)+tc.firstId, item.VID)
			require.Equal(t, int16(skip)+int16(i)+tc.firstId, item.PID)
		}
	}
}

func getComplianceInfo(setup TestSetup, vid int16, pid int16) (types.ComplianceInfo, sdk.Error) {
	return getSingle(setup, vid, pid, QueryComplianceInfo)
}

func getCertifiedModel(setup TestSetup, vid int16, pid int16) (types.ComplianceInfo, sdk.Error) {
	return getSingle(setup, vid, pid, QueryCertifiedModel)
}

func getRevokedModel(setup TestSetup, vid int16, pid int16) (types.ComplianceInfo, sdk.Error) {
	return getSingle(setup, vid, pid, QueryRevokedModel)
}

func getComplianceInfos(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoItems {
	return getAll(setup, params, QueryAllComplianceInfoRecords)
}

func getCertifiedModels(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoItems {
	return getAll(setup, params, QueryAllCertifiedModels)
}

func getRevokedModels(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoItems {
	return getAll(setup, params, QueryAllRevokedModels)
}

func getSingle(setup TestSetup, vid int16, pid int16, state string) (types.ComplianceInfo, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{state, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(types.SingleQueryParams{CertificationType: types.ZbCertificationType})},
	)

	if err != nil {
		return types.ComplianceInfo{}, err
	}

	var receivedComplianceInfo types.ComplianceInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedComplianceInfo)

	return receivedComplianceInfo, nil
}

func getAll(setup TestSetup, params types.ListQueryParams, state string) types.ListComplianceInfoItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{state},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.ListComplianceInfoItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}
