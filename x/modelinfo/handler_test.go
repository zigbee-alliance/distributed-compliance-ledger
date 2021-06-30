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
package modelinfo

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

func TestHandler_AddModel(t *testing.T) {
	setup := Setup()

	// add new model
	modelInfo := TestMsgAddModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, modelInfo.VID, modelInfo.PID,
		modelInfo.SoftwareVersion, modelInfo.HardwareVersion)

	// check
	require.Equal(t, receivedModelInfo.Model.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.Model.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.CID, modelInfo.CID)
}

func TestHandler_UpdateModel(t *testing.T) {
	setup := Setup()

	// try update not present model
	msgUpdateModelInfo := TestMsgUpdateModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgUpdateModelInfo)
	require.Equal(t, types.CodeModelInfoDoesNotExist, result.Code)

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, msgUpdateModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query updated model
	receivedModelInfo := queryModelInfo(setup, msgUpdateModelInfo.Model.VID, msgUpdateModelInfo.Model.PID,
		msgUpdateModelInfo.Model.SoftwareVersion, msgUpdateModelInfo.Model.HardwareVersion)

	// check
	require.Equal(t, receivedModelInfo.Model.VID, msgUpdateModelInfo.Model.VID)
	require.Equal(t, receivedModelInfo.Model.PID, msgUpdateModelInfo.Model.PID)
	require.Equal(t, receivedModelInfo.Model.CID, msgUpdateModelInfo.Model.CID)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// update existing model by not owner
		msgUpdateModelInfo := TestMsgUpdateModelInfo(testconstants.Address3)
		result = setup.Handler(setup.Ctx, msgUpdateModelInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}

	// owner update existing model
	msgUpdateModelInfo := TestMsgUpdateModelInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdateModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_AddModelWithEmptyOptionalFields(t *testing.T) {
	setup := Setup()

	// add new model
	modelInfo := TestMsgAddModelInfo(setup.Vendor)
	modelInfo.CID = 0              // Set empty CID
	modelInfo.OtaURL = ""          // Set empty OtaURL
	modelInfo.OtaChecksum = ""     // Set empty OtaChecksum
	modelInfo.OtaChecksumType = "" // Set empty OtaChecksumType

	result := setup.Handler(setup.Ctx, modelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, testconstants.VID, testconstants.PID,
		testconstants.SoftwareVersion, testconstants.HardwareVersion)

	// check
	require.Equal(t, receivedModelInfo.Model.CID, uint16(0))
	require.Equal(t, receivedModelInfo.Model.OtaURL, "")
	require.Equal(t, receivedModelInfo.Model.OtaChecksum, "")
	require.Equal(t, receivedModelInfo.Model.OtaChecksumType, "")
}

func TestHandler_AddModelByNonVendor(t *testing.T) {
	setup := Setup()

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new model
		modelInfo := TestMsgAddModelInfo(testconstants.Address3)
		result := setup.Handler(setup.Ctx, modelInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_PartiallyUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModelInfo := TestMsgAddModelInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModelInfo)

	// owner update Description of existing model
	msgUpdateModelInfo := TestMsgUpdateModelInfo(setup.Vendor)
	msgUpdateModelInfo.CID = 0
	msgUpdateModelInfo.Description = "New Description"
	msgUpdateModelInfo.OtaURL = ""
	result = setup.Handler(setup.Ctx, msgUpdateModelInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelInfo := queryModelInfo(setup, msgUpdateModelInfo.VID, msgUpdateModelInfo.PID,
		msgUpdateModelInfo.SoftwareVersion, msgUpdateModelInfo.HardwareVersion)

	// check
	require.Equal(t, receivedModelInfo.Model.CID, msgAddModelInfo.CID)
	require.Equal(t, receivedModelInfo.Model.Description, msgUpdateModelInfo.Description)
	require.Equal(t, receivedModelInfo.Model.OtaURL, msgAddModelInfo.OtaURL)
}

func queryModelInfo(setup TestSetup, vid uint16, pid uint16,
	softwareVersion uint32, hardwareVersion uint32) types.ModelInfo {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{
			keeper.QueryModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid),
			fmt.Sprintf("%v", softwareVersion), fmt.Sprintf("%v", hardwareVersion),
		},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelInfo)

	return receivedModelInfo
}
