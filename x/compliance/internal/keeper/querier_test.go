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
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

func TestQuerier_QueryComplianceInfo(t *testing.T) {
	setup := Setup()

	// add certified model
	certifiedModel := DefaultCertifiedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, certifiedModel)

	// query compliance info and check
	receivedComplianceInfo, _ := getComplianceInfo(setup, certifiedModel.VID, certifiedModel.PID, certifiedModel.SoftwareVersion)
	CheckComplianceInfo(t, certifiedModel, receivedComplianceInfo)

	// add revoked model
	revokedModel := DefaultRevokedModel()
	setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, revokedModel)

	// query compliance info and check
	receivedComplianceInfo, _ = getComplianceInfo(setup, revokedModel.VID, revokedModel.PID, revokedModel.SoftwareVersion)
	CheckComplianceInfo(t, revokedModel, receivedComplianceInfo)
}

func TestQuerier_QueryComplianceInfoForUnknownModel(t *testing.T) {
	setup := Setup()

	// query compliance info and check
	_, err := getComplianceInfo(setup, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)

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
	receivedComplianceInfo, _ := getCertifiedModel(setup, certifiedModel.VID, certifiedModel.PID, certifiedModel.SoftwareVersion)

	// check
	require.True(t, receivedComplianceInfo.Value)
}

func TestQuerier_QueryCertifiedModelForUnknown(t *testing.T) {
	setup := Setup()

	// query certified model
	_, err := getCertifiedModel(setup, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)

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
	_, err := getCertifiedModel(setup, revokedModel.VID, revokedModel.PID, revokedModel.SoftwareVersion)

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
	receivedComplianceInfo, _ := getRevokedModel(setup, revokedModel.VID, revokedModel.PID, revokedModel.SoftwareVersion)

	// check
	require.True(t, receivedComplianceInfo.Value)
}

func TestQuerier_QueryRevokedModelForUnknown(t *testing.T) {
	setup := Setup()

	// query revoked model
	_, err := getRevokedModel(setup, testconstants.VID, testconstants.PID, testconstants.SoftwareVersion)

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
	_, err := getRevokedModel(setup, certifiedModel.VID, certifiedModel.PID, certifiedModel.SoftwareVersion)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeComplianceInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryAllModels(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstID := PopulateStoreWithMixedModels(setup, count)

	params := types.NewListQueryParams("", 0, 0)

	receivedInfos := getComplianceInfos(setup, params)

	// check
	require.Equal(t, count, receivedInfos.Total)
	require.Equal(t, count, len(receivedInfos.Items))

	for i, item := range receivedInfos.Items {
		require.Equal(t, uint16(i)+firstID, item.VID)
		require.Equal(t, uint16(i)+firstID, item.PID)
	}
}

func TestQuerier_QueryAllModelsInState(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstID := PopulateStoreWithMixedModels(setup, count)

	params := types.NewListQueryParams("", 0, 0)

	cases := []struct {
		firstID       uint16
		count         int
		receivedInfos types.ListComplianceInfoKeyItems
	}{
		{firstID, count / 2, getCertifiedModels(setup, params)},                 // query certified model
		{firstID + uint16(count/2), count / 2, getRevokedModels(setup, params)}, // query revoked models
	}

	for _, tc := range cases {
		// check
		require.Equal(t, tc.count, tc.receivedInfos.Total)
		require.Equal(t, tc.count, len(tc.receivedInfos.Items))

		for i, item := range tc.receivedInfos.Items {
			require.Equal(t, uint16(i)+tc.firstID, item.VID)
			require.Equal(t, uint16(i)+tc.firstID, item.PID)
		}
	}
}

func TestQuerier_QueryAllModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstID := PopulateStoreWithMixedModels(setup, count)

	// query all certified models skip=1 take=2
	skip := 1
	take := 2
	params := types.NewListQueryParams("", skip, take)

	// query all certified models skip=1 take=2
	receivedInfos := getComplianceInfos(setup, params)

	// check
	require.Equal(t, count, receivedInfos.Total)
	require.Equal(t, take, len(receivedInfos.Items))

	for i, item := range receivedInfos.Items {
		require.Equal(t, uint16(skip)+uint16(i)+firstID, item.VID)
		require.Equal(t, uint16(skip)+uint16(i)+firstID, item.PID)
	}
}

func TestQuerier_QueryAllModelsInStateWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 8

	// add 4 certified and 4 revoked models
	firstID := PopulateStoreWithMixedModels(setup, count)

	// query all certified models skip=1 take=2
	skip := 1
	take := 2
	params := types.NewListQueryParams("", skip, take)

	cases := []struct {
		firstID       uint16
		count         int
		receivedInfos types.ListComplianceInfoKeyItems
	}{
		{firstID, count / 2, getCertifiedModels(setup, params)},                 // query certified model
		{firstID + uint16(count/2), count / 2, getRevokedModels(setup, params)}, // query revoked models
	}

	for _, tc := range cases {
		// check
		require.Equal(t, tc.count, tc.receivedInfos.Total)
		require.Equal(t, take, len(tc.receivedInfos.Items))

		for i, item := range tc.receivedInfos.Items {
			require.Equal(t, uint16(skip)+uint16(i)+tc.firstID, item.VID)
			require.Equal(t, uint16(skip)+uint16(i)+tc.firstID, item.PID)
		}
	}
}

func getComplianceInfo(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32) (types.ComplianceInfo, sdk.Error) {
	return getSingle(setup, vid, pid, softwareVersion, QueryComplianceInfo)
}

func getCertifiedModel(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32) (types.ComplianceInfoInState, sdk.Error) {
	return getSingleInState(setup, vid, pid, softwareVersion, QueryCertifiedModel)
}

func getRevokedModel(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32) (types.ComplianceInfoInState, sdk.Error) {
	return getSingleInState(setup, vid, pid, softwareVersion, QueryRevokedModel)
}

func getComplianceInfos(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoItems {
	return getAll(setup, params, QueryAllComplianceInfoRecords)
}

func getCertifiedModels(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoKeyItems {
	return getAllInState(setup, params, QueryAllCertifiedModels)
}

func getRevokedModels(setup TestSetup, params types.ListQueryParams) types.ListComplianceInfoKeyItems {
	return getAllInState(setup, params, QueryAllRevokedModels)
}

func getSingle(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, state string) (types.ComplianceInfo, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{state, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion), fmt.Sprintf("%v", types.ZbCertificationType)},
		abci.RequestQuery{},
	)
	if err != nil {
		return types.ComplianceInfo{}, err
	}

	var receivedComplianceInfo types.ComplianceInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedComplianceInfo)

	return receivedComplianceInfo, nil
}

func getSingleInState(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, state string) (types.ComplianceInfoInState, sdk.Error) {
	result, err := setup.Querier(
		setup.Ctx,
		[]string{state, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion), fmt.Sprintf("%v", types.ZbCertificationType)},
		abci.RequestQuery{},
	)
	if err != nil {
		return types.ComplianceInfoInState{}, err
	}

	var receivedComplianceInfo types.ComplianceInfoInState
	_ = setup.Cdc.UnmarshalJSON(result, &receivedComplianceInfo)

	return receivedComplianceInfo, nil
}

func getAll(setup TestSetup, params types.ListQueryParams, state string) types.ListComplianceInfoItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{state},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModels types.ListComplianceInfoItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModels)

	return receiveModels
}

func getAllInState(setup TestSetup, params types.ListQueryParams, state string) types.ListComplianceInfoKeyItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{state},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModels types.ListComplianceInfoKeyItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModels)

	return receiveModels
}
