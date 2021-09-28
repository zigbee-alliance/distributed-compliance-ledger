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
	require.Equal(t, receivedModel.VID, msgAddModel.VID)
	require.Equal(t, receivedModel.PID, msgAddModel.PID)
	require.Equal(t, receivedModel.DeviceTypeID, msgUpdateModel.DeviceTypeID)
}

func TestHandler_OnlyOwnerCanUpdateModel(t *testing.T) {
	setup := Setup()

	// add new model
	msgAddModel := TestMsgAddModel(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddModel)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorId3)
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
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorId3)
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
