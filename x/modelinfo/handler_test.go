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
	receivedModelInfo := queryModelInfo(setup, modelInfo.VID, modelInfo.PID)

	// check
	// TODO do a smart check

	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	/*require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.CID, modelInfo.CID)
	require.Equal(t, receivedModelInfo.Version, modelInfo.Version)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
	require.Equal(t, receivedModelInfo.Description, modelInfo.Description)
	require.Equal(t, receivedModelInfo.SKU, modelInfo.SKU)
	require.Equal(t, receivedModelInfo.HardwareVersion, modelInfo.HardwareVersion)
	require.Equal(t, receivedModelInfo.FirmwareVersion, modelInfo.FirmwareVersion)
	require.Equal(t, receivedModelInfo.OtaURL, modelInfo.OtaURL)
	require.Equal(t, receivedModelInfo.OtaChecksum, modelInfo.OtaChecksum)
	require.Equal(t, receivedModelInfo.OtaChecksumType, modelInfo.OtaChecksumType)
	require.Equal(t, receivedModelInfo.Custom, modelInfo.Custom)
	require.Equal(t, receivedModelInfo.TisOrTrpTestingCompleted, modelInfo.TisOrTrpTestingCompleted)
	require.Equal(t, receivedModelInfo.Owner, modelInfo.Signer)*/
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
	receivedModelInfo := queryModelInfo(setup, msgUpdateModelInfo.VID, msgUpdateModelInfo.PID)

	// check
	require.Equal(t, receivedModelInfo.VID, msgAddModelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, msgAddModelInfo.PID)
	require.Equal(t, receivedModelInfo.CID, msgUpdateModelInfo.CID)

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
	receivedModelInfo := queryModelInfo(setup, testconstants.VID, testconstants.PID)

	// check

	require.Equal(t, receivedModelInfo.CID, uint16(0))
	require.Equal(t, receivedModelInfo.OtaURL, "")
	require.Equal(t, receivedModelInfo.OtaChecksum, "")
	require.Equal(t, receivedModelInfo.OtaChecksumType, "")

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
	receivedModelInfo := queryModelInfo(setup, msgUpdateModelInfo.VID, msgUpdateModelInfo.PID)

	// check
	require.Equal(t, receivedModelInfo.CID, msgAddModelInfo.CID)
	require.Equal(t, receivedModelInfo.Description, msgUpdateModelInfo.Description)
	require.Equal(t, receivedModelInfo.OtaURL, msgAddModelInfo.OtaURL)
}

func queryModelInfo(setup TestSetup, vid uint16, pid uint16) types.ModelInfo {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryModel, fmt.Sprintf("%v", vid), fmt.Sprintf("%v", pid)},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelInfo)

	return receivedModelInfo
}
