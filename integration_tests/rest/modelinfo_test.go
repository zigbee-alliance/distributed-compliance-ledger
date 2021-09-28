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

package rest_test

import (
	"net/http"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `dclcli rest-server --chain-id dclchain`

	TODO: provide tests for error cases
*/

func TestModelDemo(t *testing.T) {
	// Register new Vendor account
	vid := common.RandUint16()
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid)

	// Get all model infos
	inputModels, _ := utils.GetModels()

	// Get all vendors
	inputVendors, _ := utils.GetVendors()

	// Prepare model info
	firstModel := utils.NewMsgAddModel(vendor.Address, vid)
	VID := firstModel.VID

	// Sign and Broadcast AddModel message
	utils.SignAndBroadcastMessage(vendor, firstModel)

	// Check model is created
	receivedModel, _ := utils.GetModel(firstModel.VID, firstModel.PID)
	require.Equal(t, receivedModel.VID, firstModel.VID)
	require.Equal(t, receivedModel.PID, firstModel.PID)
	require.Equal(t, receivedModel.ProductName, firstModel.ProductName)
	require.Equal(t, receivedModel.ProductLabel, firstModel.ProductLabel)

	// Publish second model info using POST command with passing name and passphrase. Same Vendor
	secondModel := utils.NewMsgAddModel(vendor.Address, vid)
	_, _ = utils.AddModel(secondModel, vendor)

	// Check model is created
	receivedModel, _ = utils.GetModel(secondModel.VID, secondModel.PID)
	require.Equal(t, receivedModel.VID, secondModel.VID)
	require.Equal(t, receivedModel.PID, secondModel.PID)
	require.Equal(t, receivedModel.ProductName, secondModel.ProductName)
	require.Equal(t, receivedModel.ProductLabel, secondModel.ProductLabel)

	// Get all model infos
	models, _ := utils.GetModels()
	require.Equal(t, utils.ParseUint(inputModels.Total)+2, utils.ParseUint(models.Total))

	// Get all vendors
	vendors, _ := utils.GetVendors()
	require.Equal(t, utils.ParseUint(inputVendors.Total)+1, utils.ParseUint(vendors.Total))

	// Get vendor models
	vendorModels, _ := utils.GetVendorModels(VID)
	require.Equal(t, uint64(2), uint64(len(vendorModels.Products)))
	require.Equal(t, firstModel.PID, vendorModels.Products[0].PID)
	require.Equal(t, secondModel.PID, vendorModels.Products[1].PID)

	// Update second model info
	secondModelUpdate := utils.NewMsgUpdateModel(secondModel.VID, secondModel.PID, vendor.Address)
	_, _ = utils.UpdateModel(secondModelUpdate, vendor)

	// Check model is updated
	receivedModel, _ = utils.GetModel(secondModel.VID, secondModel.PID)
	require.Equal(t, receivedModel.ProductLabel, secondModelUpdate.ProductLabel)
}

func TestModelDemo_Prepare_Sign_Broadcast(t *testing.T) {
	// Register new Vendor account
	vid := common.RandUint16()
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid)

	// Prepare model info
	model := utils.NewMsgAddModel(vendor.Address, vid)

	// Prepare Sign Broadcast
	addModelTransaction, _ := utils.PrepareAddModelTransaction(model)
	_, _ = utils.SignAndBroadcastTransaction(vendor, addModelTransaction)

	// Check model is created
	receivedModel, _ := utils.GetModel(model.VID, model.PID)
	require.Equal(t, receivedModel.VID, model.VID)
	require.Equal(t, receivedModel.PID, model.PID)
	require.Equal(t, receivedModel.ProductName, model.ProductName)
}

/* Error cases */

func Test_AddModel_ByNonVendor(t *testing.T) {
	// register new account
	vid := common.RandUint16()
	testAccount := utils.CreateNewAccount(auth.AccountRoles{}, vid)

	// try to publish model info
	model := utils.NewMsgAddModel(testAccount.Address, vid)
	res, _ := utils.SignAndBroadcastMessage(testAccount, model)
	require.Equal(t, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
}

func Test_AddModel_ByDifferentVendor(t *testing.T) {
	// register new account
	vid := common.RandUint16()
	testAccount := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid+1)

	// try to publish model info
	model := utils.NewMsgAddModel(testAccount.Address, vid)
	res, _ := utils.SignAndBroadcastMessage(testAccount, model)
	require.Equal(t, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
}

func Test_AddModel_Twice(t *testing.T) {
	// register new account
	vid := common.RandUint16()
	testAccount := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor}, vid)

	// publish modelMsg info
	modelMsg := utils.NewMsgAddModel(testAccount.Address, vid)
	res, _ := utils.AddModel(modelMsg, testAccount)
	require.Equal(t, sdk.CodeOK, sdk.CodeType(res.Code))

	// publish second time
	res, _ = utils.AddModel(modelMsg, testAccount)
	require.Equal(t, model.CodeModelAlreadyExists, sdk.CodeType(res.Code))
}

func Test_GetModel_ForUnknown(t *testing.T) {
	_, code := utils.GetModel(common.RandUint16(), common.RandUint16())
	require.Equal(t, http.StatusNotFound, code)
}

func Test_GetModel_ForInvalidVidPid(t *testing.T) {
	// zero vid
	_, code := utils.GetModel(0, common.RandUint16())
	require.Equal(t, http.StatusBadRequest, code)

	// zero pid
	_, code = utils.GetModel(common.RandUint16(), 0)
	require.Equal(t, http.StatusBadRequest, code)
}
