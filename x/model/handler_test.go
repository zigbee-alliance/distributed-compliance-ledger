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

//nolint:testpackage,lll
package model

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

func TestHandler_AddModel(t *testing.T) {
	setup := Setup()

	// add new model
	model := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, model)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModel := queryModel(setup, model.VID, model.PID)

	// check
	require.Equal(t, receivedModel.VID, model.VID)
	require.Equal(t, receivedModel.PID, model.PID)
	require.Equal(t, receivedModel.DeviceTypeID, model.DeviceTypeID)
}

func TestHandler_UpdateModel(t *testing.T) {
	setup := Setup()

	// try update not present model
	msgUpdateModel := TestMsgUpdateModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgUpdateModel)
	require.Equal(t, types.CodeModelDoesNotExist, result.Code)

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query updated model
	receivedModel := queryModel(setup, msgUpdateModel.VID, msgUpdateModel.PID)

	// check
	// Mutable Fields ProductName,ProductLable,PartNumber,CommissioningCustomFlowUrl,
	// CommissioningModeInitialStepsInstruction,CommissioningModeSecondaryStepsInstruction,UserManualUrl,SupportUrl,SupportUrl
	require.Equal(t, receivedModel.VID, msgAddModel.VID)
	require.Equal(t, receivedModel.PID, msgAddModel.PID)
	require.Equal(t, receivedModel.ProductLabel, msgUpdateModel.ProductLabel)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorID3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// update existing model by not owner
		msgUpdateModel := TestMsgUpdateModel(testconstants.Address3)
		result = setup.Handler(setup.Ctx, msgUpdateModel)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}

	// owner update existing model
	msgUpdateModel := TestMsgUpdateModel(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_AddModelWithEmptyOptionalFields(t *testing.T) {
	setup := Setup()

	// add new model
	model := TestMsgAddModel(setup.Vendor)
	model.DeviceTypeID = 0 // Set empty CID

	result := setup.Handler(setup.Ctx, model)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModel := queryModel(setup, testconstants.VID, testconstants.PID)

	// check
	require.Equal(t, receivedModel.DeviceTypeID, uint16(0))
}

func TestHandler_AddModelByNonVendor(t *testing.T) {
	setup := Setup()

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorID3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new model
		model := TestMsgAddModel(testconstants.Address3)
		result := setup.Handler(setup.Ctx, model)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_PartiallyUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)

	// owner update Description of existing model
	msgUpdateModel := TestMsgUpdateModel(setup.Vendor)
	msgUpdateModel.DeviceTypeID = 0
	msgUpdateModel.ProductLabel = "New Description"
	result = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModel := queryModel(setup, msgUpdateModel.VID, msgUpdateModel.PID)

	// check
	// Mutable Fields ProductName,ProductLable,PartNumber,CommissioningCustomFlowUrl,
	// CommissioningModeInitialStepsInstruction,CommissioningModeSecondaryStepsInstruction,UserManualUrl,SupportUrl,SupportUrl
	require.Equal(t, receivedModel.DeviceTypeID, msgAddModel.DeviceTypeID)
	require.Equal(t, receivedModel.ProductLabel, msgUpdateModel.ProductLabel)
}

func queryModel(setup TestSetup, vid uint16, pid uint16) types.Model {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedModel types.Model
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModel)

	return receivedModel
}

// ----------------------------------------------------------------------------
// Model Version Tests --------------------------------------------------------
// ----------------------------------------------------------------------------

func TestHandler_AddModelVersion(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add new model version
	msgAddModelVersion := TestMsgAddModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model version
	receivedModelVersion := queryModelVersion(setup, msgAddModelVersion.VID, msgAddModelVersion.PID, msgAddModelVersion.SoftwareVersion)

	// check
	require.Equal(t, receivedModelVersion.VID, msgAddModelVersion.VID)
	require.Equal(t, receivedModelVersion.PID, msgAddModelVersion.PID)
	require.Equal(t, receivedModelVersion.SoftwareVersion, msgAddModelVersion.SoftwareVersion)
	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgAddModelVersion.SoftwareVersionString)
}

func TestHandler_UpdateModelVersion(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add new model version
	msgAddModelVersion := TestMsgAddModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model version
	msgUpdateModelVersion := TestMsgUpdateModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model version
	receivedModelVersion := queryModelVersion(setup, msgUpdateModelVersion.VID, msgUpdateModelVersion.PID, msgUpdateModelVersion.SoftwareVersion)

	// check
	// Mutable Fields SoftwareVersionValid,OtaUrl,MinApplicableSoftwareVersion,MaxApplicableSoftwareVersion,ReleaseNotesUrl
	require.Equal(t, receivedModelVersion.VID, msgAddModelVersion.VID)
	require.Equal(t, receivedModelVersion.PID, msgAddModelVersion.PID)
	require.Equal(t, receivedModelVersion.SoftwareVersion, msgUpdateModelVersion.SoftwareVersion)

	require.NotEqual(t, receivedModelVersion.SoftwareVersionString, msgUpdateModelVersion.SoftwareVersionString)
	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgAddModelVersion.SoftwareVersionString)

	require.NotEqual(t, receivedModelVersion.CDVersionNumber, msgUpdateModelVersion.CDVersionNumber)
	require.Equal(t, receivedModelVersion.CDVersionNumber, msgAddModelVersion.CDVersionNumber)

	require.NotEqual(t, receivedModelVersion.FirmwareDigests, msgUpdateModelVersion.FirmwareDigests)
	require.Equal(t, receivedModelVersion.FirmwareDigests, msgAddModelVersion.FirmwareDigests)

	require.Equal(t, receivedModelVersion.SoftwareVersionValid, msgUpdateModelVersion.SoftwareVersionValid)
	require.Equal(t, receivedModelVersion.OtaURL, msgUpdateModelVersion.OtaURL)

	require.NotEqual(t, receivedModelVersion.OtaFileSize, msgUpdateModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgAddModelVersion.OtaFileSize)

	require.NotEqual(t, receivedModelVersion.OtaChecksum, msgUpdateModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaChecksum, msgAddModelVersion.OtaChecksum)

	require.NotEqual(t, receivedModelVersion.OtaChecksumType, msgUpdateModelVersion.OtaChecksumType)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgAddModelVersion.OtaChecksumType)

	require.Equal(t, receivedModelVersion.MinApplicableSoftwareVersion, msgUpdateModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.MaxApplicableSoftwareVersion, msgUpdateModelVersion.MaxApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.ReleaseNotesURL, msgUpdateModelVersion.ReleaseNotesURL)
}

//nolint:funclen
func TestHandler_PartiallyUpdateModelVersion(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add new model version
	msgAddModelVersion := TestMsgAddModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	msgUpdateModelVersion := MsgUpdateModelVersion{
		ModelVersion: msgAddModelVersion.ModelVersion,
		Signer:       setup.Vendor,
	}

	// Just update ReleaseNotesURL
	msgUpdateModelVersion.ReleaseNotesURL = "https://new.releasenotes.url"
	result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model version
	receivedModelVersion := queryModelVersion(setup, msgUpdateModelVersion.VID, msgUpdateModelVersion.PID, msgUpdateModelVersion.SoftwareVersion)
	// Mutable Fields SoftwareVersionValid,OtaUrl,MinApplicableSoftwareVersion,MaxApplicableSoftwareVersion,ReleaseNotesUrl
	require.Equal(t, receivedModelVersion.VID, msgAddModelVersion.VID)
	require.Equal(t, receivedModelVersion.PID, msgAddModelVersion.PID)
	require.Equal(t, receivedModelVersion.SoftwareVersion, msgAddModelVersion.SoftwareVersion)
	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgAddModelVersion.SoftwareVersionString)
	require.Equal(t, receivedModelVersion.CDVersionNumber, msgAddModelVersion.CDVersionNumber)
	require.Equal(t, receivedModelVersion.FirmwareDigests, msgAddModelVersion.FirmwareDigests)
	require.Equal(t, receivedModelVersion.OtaURL, msgAddModelVersion.OtaURL)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgAddModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.OtaChecksum, msgAddModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgAddModelVersion.OtaChecksumType)
	require.Equal(t, receivedModelVersion.MinApplicableSoftwareVersion, msgAddModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.MaxApplicableSoftwareVersion, msgAddModelVersion.MaxApplicableSoftwareVersion)
	// Only field that is changed should be releaseNotesURL
	require.NotEqual(t, receivedModelVersion.ReleaseNotesURL, msgAddModelVersion.ReleaseNotesURL)
	require.Equal(t, receivedModelVersion.ReleaseNotesURL, msgUpdateModelVersion.ReleaseNotesURL)

	// Update few more fields and check
	msgUpdateModelVersion.ReleaseNotesURL = "https://v2.new.releasenotes.url"
	msgUpdateModelVersion.OtaURL = "https://new.ota.url"
	msgUpdateModelVersion.MinApplicableSoftwareVersion = 2

	result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model version
	receivedModelVersion = queryModelVersion(setup, msgUpdateModelVersion.VID, msgUpdateModelVersion.PID, msgUpdateModelVersion.SoftwareVersion)
	// Mutable Fields SoftwareVersionValid,OtaUrl,MinApplicableSoftwareVersion,MaxApplicableSoftwareVersion,ReleaseNotesUrl
	require.Equal(t, receivedModelVersion.VID, msgAddModelVersion.VID)
	require.Equal(t, receivedModelVersion.PID, msgAddModelVersion.PID)
	require.Equal(t, receivedModelVersion.SoftwareVersion, msgAddModelVersion.SoftwareVersion)
	require.Equal(t, receivedModelVersion.SoftwareVersionString, msgAddModelVersion.SoftwareVersionString)
	require.Equal(t, receivedModelVersion.CDVersionNumber, msgAddModelVersion.CDVersionNumber)
	require.Equal(t, receivedModelVersion.FirmwareDigests, msgAddModelVersion.FirmwareDigests)
	require.Equal(t, receivedModelVersion.OtaFileSize, msgAddModelVersion.OtaFileSize)
	require.Equal(t, receivedModelVersion.OtaChecksum, msgAddModelVersion.OtaChecksum)
	require.Equal(t, receivedModelVersion.OtaChecksumType, msgAddModelVersion.OtaChecksumType)
	require.Equal(t, receivedModelVersion.MaxApplicableSoftwareVersion, msgAddModelVersion.MaxApplicableSoftwareVersion)
	// Only field that is changed should be releaseNotesURL, otaURL, minApplicableSoftwareVersion
	require.NotEqual(t, receivedModelVersion.ReleaseNotesURL, msgAddModelVersion.ReleaseNotesURL)
	require.Equal(t, receivedModelVersion.ReleaseNotesURL, msgUpdateModelVersion.ReleaseNotesURL)
	require.NotEqual(t, receivedModelVersion.OtaURL, msgAddModelVersion.OtaURL)
	require.Equal(t, receivedModelVersion.OtaURL, msgUpdateModelVersion.OtaURL)
	require.NotEqual(t, receivedModelVersion.MinApplicableSoftwareVersion, msgAddModelVersion.MinApplicableSoftwareVersion)
	require.Equal(t, receivedModelVersion.MinApplicableSoftwareVersion, msgUpdateModelVersion.MinApplicableSoftwareVersion)
}

func TestHandler_OnlyOwnerCanUpdateModelVersion(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add new model version
	msgAddModelVersion := TestMsgAddModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update model version with wrong signer
	// store account

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorID3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// update existing model version by not same vendorID
		msgUpdateModelVersion := MsgUpdateModelVersion{
			ModelVersion: msgAddModelVersion.ModelVersion,
			Signer:       testconstants.Address3,
		}

		result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func queryModelVersion(setup TestSetup, vid uint16, pid uint16, softwareVersion uint32) types.ModelVersion {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModelVersion, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid), fmt.Sprintf("%v", softwareVersion)},
		abci.RequestQuery{},
	)

	var receivedModelVersion types.ModelVersion
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelVersion)

	return receivedModelVersion
}
