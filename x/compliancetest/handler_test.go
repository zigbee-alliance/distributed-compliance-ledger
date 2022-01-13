// Copyright 2022 DSR Corporation
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

package compliancetest

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)
	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type ModelKeeperMock struct {
	mock.Mock
}

func (m *ModelKeeperMock) GetModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
) (val modeltypes.ModelVersion, found bool) {
	args := m.Called(ctx, vid, pid, softwareVersion)
	return args.Get(0).(modeltypes.ModelVersion), args.Bool(1)
}

var _ types.ModelKeeper = &ModelKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	ModelKeeper   *ModelKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	TestHouse sdk.AccAddress
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func (setup *TestSetup) AddModelVersion(
	vid int32, pid int32, softwareVersion uint32, softwareVersionString string,
) (int32, int32, uint32, string) {
	modelVersion := NewModelVersion(vid, pid, softwareVersion, softwareVersionString)

	setup.ModelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(*modelVersion, true)

	// return just for convenient re-assignment
	return vid, pid, softwareVersion, softwareVersionString
}

func (setup *TestSetup) SetNoModelVersionForKey(
	vid int32,
	pid int32,
	softwareVersion uint32,
) {
	modelKeeper := setup.ModelKeeper

	modelKeeper.On(
		"GetModelVersion",
		mock.Anything, vid, pid, softwareVersion,
	).Return(modeltypes.ModelVersion{}, false)
}

func Setup(t *testing.T) *TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	modelKeeper := &ModelKeeperMock{}
	keeper, ctx := testkeeper.CompliancetestKeeper(t, dclauthKeeper, modelKeeper)

	testHouse := GenerateAccAddress()

	setup := &TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		ModelKeeper:   modelKeeper,
		Handler:       NewHandler(*keeper),
		TestHouse:     testHouse,
	}

	setup.AddAccount(testHouse, []dclauthtypes.AccountRole{dclauthtypes.TestHouse})

	return setup
}

func TestHandler_AddTestingResult(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// add new testing result
	testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString, setup.TestHouse)
	_, err := setup.Handler(setup.Ctx, testingResult)
	require.NoError(t, err)

	// query testing result
	receivedTestingResult, err := queryTestingResults(setup, vid, pid, softwareVersion)
	require.NoError(t, err)

	// check
	require.Equal(t, vid, receivedTestingResult.Vid)
	require.Equal(t, pid, receivedTestingResult.Pid)
	require.Equal(t, softwareVersion, receivedTestingResult.SoftwareVersion)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	CheckTestingResult(t, testingResult, receivedTestingResult.Results[0])
}

func TestHandler_AddTestingResultByNonTestHouse(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.Vendor,
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role})

		// add new testing result by non TestHouse
		testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString, accAddress)
		_, err := setup.Handler(setup.Ctx, testingResult)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_AddTestingResultForUnknownModel(t *testing.T) {
	setup := Setup(t)

	// Set absence of model version
	setup.SetNoModelVersionForKey(testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion)

	// add new testing result
	testingResult := NewMsgAddTestingResult(
		testconstants.Vid,
		testconstants.Pid,
		testconstants.SoftwareVersion,
		testconstants.SoftwareVersionString,
		setup.TestHouse,
	)
	_, err := setup.Handler(setup.Ctx, testingResult)
	require.Error(t, err)
	require.True(t, modeltypes.ErrModelVersionDoesNotExist.Is(err))
}

func TestHandler_AddSeveralTestingResultsForOneModel(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	for i, accAddress := range []sdk.AccAddress{
		testconstants.Address1,
		testconstants.Address2,
		testconstants.Address3,
	} {
		// store account
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.TestHouse})

		// add new testing result
		testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString, accAddress)
		_, err := setup.Handler(setup.Ctx, testingResult)
		require.NoError(t, err)

		// query testing result
		receivedTestingResult, err := queryTestingResults(setup, vid, pid, softwareVersion)
		require.NoError(t, err)

		// check
		require.Equal(t, vid, receivedTestingResult.Vid)
		require.Equal(t, pid, receivedTestingResult.Pid)
		require.Equal(t, i+1, len(receivedTestingResult.Results))
		CheckTestingResult(t, testingResult, receivedTestingResult.Results[i])
	}
}

func TestHandler_AddSeveralTestingResultsForDifferentModels(t *testing.T) {
	setup := Setup(t)

	for i := 1; i < 5; i++ {
		// add model version
		vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
			int32(i), int32(i), uint32(i), fmt.Sprint(i))

		// add new testing result
		testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString, setup.TestHouse)
		_, err := setup.Handler(setup.Ctx, testingResult)
		require.NoError(t, err)

		// query testing result
		receivedTestingResult, err := queryTestingResults(setup, vid, pid, softwareVersion)
		require.NoError(t, err)

		// check
		require.Equal(t, vid, receivedTestingResult.Vid)
		require.Equal(t, pid, receivedTestingResult.Pid)
		require.Equal(t, 1, len(receivedTestingResult.Results))
		CheckTestingResult(t, testingResult, receivedTestingResult.Results[0])
	}
}

func TestHandler_AddTestingResultTwiceForSameModelAndSameTestHouse(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// add new testing result
	testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString, setup.TestHouse)
	_, err := setup.Handler(setup.Ctx, testingResult)
	require.NoError(t, err)

	// add testing result second time
	testingResult.TestResult = "Second Testing Result"
	_, err = setup.Handler(setup.Ctx, testingResult)
	require.NoError(t, err)

	// query testing result
	receivedTestingResult, err := queryTestingResults(setup, vid, pid, softwareVersion)
	require.NoError(t, err)

	// check
	require.Equal(t, 2, len(receivedTestingResult.Results))
}

func TestHandler_AddTestingResultWithInvalidSoftwareVersionString(t *testing.T) {
	setup := Setup(t)

	// add model version
	vid, pid, softwareVersion, softwareVersionString := setup.AddModelVersion(
		testconstants.Vid, testconstants.Pid, testconstants.SoftwareVersion, testconstants.SoftwareVersionString)

	// add new testing result
	testingResult := NewMsgAddTestingResult(vid, pid, softwareVersion, softwareVersionString+"-modified", setup.TestHouse)
	_, err := setup.Handler(setup.Ctx, testingResult)
	require.Error(t, err)
	require.True(t, types.ErrModelVersionStringDoesNotMatch.Is(err))
}

func queryTestingResults(
	setup *TestSetup,
	vid int32,
	pid int32,
	softwareVersion uint32,
) (*types.TestingResults, error) {

	req := &types.QueryGetTestingResultsRequest{
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
	}

	resp, err := setup.Keeper.TestingResults(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return &resp.TestingResults, nil
}

func NewMsgAddTestingResult(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer sdk.AccAddress,
) *types.MsgAddTestingResult {

	return &types.MsgAddTestingResult{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		TestResult:            testconstants.TestResult,
		TestDate:              testconstants.TestDate,
	}
}

func CheckTestingResult(
	t *testing.T,
	expectedTestingResult *types.MsgAddTestingResult,
	receivedTestingResult *types.TestingResult,
) {
	require.Equal(t, expectedTestingResult.Signer, receivedTestingResult.Owner)
	require.Equal(t, expectedTestingResult.TestResult, receivedTestingResult.TestResult)
	require.Equal(t, expectedTestingResult.TestDate, receivedTestingResult.TestDate)
}

func NewModelVersion(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
) *modeltypes.ModelVersion {

	return &modeltypes.ModelVersion{
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl,
		Creator:                      GenerateAccAddress().String(),
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}
