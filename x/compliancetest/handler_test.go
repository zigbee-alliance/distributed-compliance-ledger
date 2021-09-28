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
package compliancetest

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	test_constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
)

func TestHandler_AddTestingResult(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, test_constants.VID, test_constants.PID, test_constants.SoftwareVersion, test_constants.SoftwareVersionString)

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid, softwareVersion, softwareVersionString)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query testing result
	receivedTestingResult := queryTestingResult(setup, vid, pid, softwareVersion)

	// check
	require.Equal(t, receivedTestingResult.VID, vid)
	require.Equal(t, receivedTestingResult.PID, pid)
	require.Equal(t, receivedTestingResult.SoftwareVersion, softwareVersion)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, receivedTestingResult.Results[0], testingResult)
}

func TestHandler_AddTestingResultByNonTestHouse(t *testing.T) {
	setup := Setup()
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, test_constants.VID, test_constants.PID, test_constants.SoftwareVersion, test_constants.SoftwareVersionString)

	for _, role := range []auth.AccountRole{auth.Vendor, auth.ZBCertificationCenter, auth.NodeAdmin} {
		// store account
		account := auth.NewAccount(test_constants.Address3, test_constants.PubKey3, auth.AccountRoles{role}, test_constants.VendorId3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new testing result by non TestHouse
		testingResult := TestMsgAddTestingResult(test_constants.Address3, vid, pid, softwareVersion, softwareVersionString)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_AddTestingResultForUnknownModel(t *testing.T) {
	setup := Setup()

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, test_constants.VID, test_constants.PID,
		test_constants.SoftwareVersion, test_constants.SoftwareVersionString)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, modelversion.CodeModelVersionDoesNotExist, result.Code)
}

func TestHandler_AddSeveralTestingResultsForOneModel(t *testing.T) {
	setup := Setup()

	// add model
	vid, pid := addModel(setup, test_constants.VID, test_constants.PID)
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, test_constants.VID, test_constants.PID, test_constants.SoftwareVersion, test_constants.SoftwareVersionString)

	for i, th := range []sdk.AccAddress{test_constants.Address1, test_constants.Address2, test_constants.Address3} {
		// store account
		account := auth.NewAccount(th, test_constants.PubKey1, auth.AccountRoles{auth.TestHouse}, 0)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new testing result
		testingResult := TestMsgAddTestingResult(th, vid, pid, softwareVersion, softwareVersionString)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query testing result
		receivedTestingResult := queryTestingResult(setup, vid, pid, softwareVersion)

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
		// add model version
		_, _, softwareVersion, softwareVersionString :=
			addModelVersion(setup, i, i, uint32(i), fmt.Sprint(i))

		// add new testing result
		testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid, softwareVersion, softwareVersionString)
		result := setup.Handler(setup.Ctx, testingResult)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query testing result
		receivedTestingResult := queryTestingResult(setup, vid, pid, softwareVersion)

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
	// add model version
	_, _, softwareVersion, softwareVersionString :=
		addModelVersion(setup, test_constants.VID, test_constants.PID, test_constants.SoftwareVersion, test_constants.SoftwareVersionString)

	// add new testing result
	testingResult := TestMsgAddTestingResult(setup.TestHouse, vid, pid, softwareVersion, softwareVersionString)
	result := setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add testing result second time
	testingResult.TestResult = "Second Testing Result"
	result = setup.Handler(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query testing result
	receivedTestingResult := queryTestingResult(setup, vid, pid, softwareVersion)

	// check
	require.Equal(t, 2, len(receivedTestingResult.Results))
}

func queryTestingResult(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32) types.TestingResults {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryTestingResult, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion)},
		abci.RequestQuery{},
	)

	var receivedTestingResult types.TestingResults
	_ = setup.Cdc.UnmarshalJSON(result, &receivedTestingResult)

	return receivedTestingResult
}

func addModel(setup TestSetup, vid uint16, pid uint16) (uint16, uint16) {
	model := model.Model{
		VID:          vid,
		PID:          pid,
		DeviceTypeID: test_constants.DeviceTypeID,
		ProductName:  test_constants.ProductName,
		ProductLabel: test_constants.ProductLabel,
		PartNumber:   test_constants.PartNumber,
	}

	setup.ModelKeeper.SetModel(setup.Ctx, model)

	return vid, pid
}

func addModelVersion(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32, softwareVersionString string) (uint16, uint16, uint32, string) {
	modelVersion := modelversion.ModelVersion{
		VID:                          vid,
		PID:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CDVersionNumber:              test_constants.CDVersionNumber,
		MinApplicableSoftwareVersion: test_constants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: test_constants.MaxApplicableSoftwareVersion,
	}

	setup.ModelversionKeeper.SetModelVersion(setup.Ctx, modelVersion)

	return vid, pid, softwareVersion, softwareVersionString
}
