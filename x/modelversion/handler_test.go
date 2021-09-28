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
package modelversion

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

func TestHandler_AddModelVersion(t *testing.T) {
	setup := Setup()

	// add new model
	addModel(setup)
	// add new model version
	modelVersion := TestMsgAddModelVersion(setup.Vendor)
	result := setup.Handler(setup.Ctx, modelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query model
	receivedModelVersion := queryModelVersion(setup, modelVersion.VID, modelVersion.PID, modelVersion.SoftwareVersion)

	// check
	require.Equal(t, receivedModelVersion.VID, modelVersion.VID)
	require.Equal(t, receivedModelVersion.PID, modelVersion.PID)
	require.Equal(t, receivedModelVersion.SoftwareVersion, modelVersion.SoftwareVersion)
}

func TestHandler_UpdateModelVersion(t *testing.T) {
	setup := Setup()

	// try updating non existing model version
	msgUpdateModelVersion := TestMsgUpdateModelVersion(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Equal(t, types.CodeModelVersionDoesNotExist, result.Code)

	// add new model
	addModel(setup)

	// add new model version
	msgAddModelVersion := TestMsgAddModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing model
	result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query updated model
	receivedModel := queryModelVersion(setup, msgUpdateModelVersion.VID, msgUpdateModelVersion.PID, msgUpdateModelVersion.SoftwareVersion)

	// check
	require.Equal(t, receivedModel.VID, msgUpdateModelVersion.VID)
	require.Equal(t, receivedModel.PID, msgUpdateModelVersion.PID)
	require.Equal(t, receivedModel.ReleaseNotesURL, msgUpdateModelVersion.ReleaseNotesURL)
}

func TestHandler_OnlyAccountWithVendorIDCanUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	addModel(setup)
	// add new model version
	msgAddModel := TestMsgAddModelVersion(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorId3)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// update existing model by not owner
		msgUpdateModelVersion := TestMsgUpdateModelVersion(testconstants.Address3)
		result = setup.Handler(setup.Ctx, msgUpdateModelVersion)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}

	// owner update existing model
	msgUpdateModel := TestMsgUpdateModelVersion(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdateModel)
	require.Equal(t, sdk.CodeOK, result.Code)
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

func addModel(setup TestSetup) {
	model := model.Model{
		VID:          testconstants.VendorId1,
		PID:          testconstants.PID,
		DeviceTypeID: testconstants.DeviceTypeID,
		ProductName:  testconstants.ProductName,
		ProductLabel: testconstants.ProductLabel,
		PartNumber:   testconstants.PartNumber,
	}

	setup.ModelKeeper.SetModel(setup.Ctx, model)

}
