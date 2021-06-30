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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `dclcli rest-server --chain-id dclchain`

	TODO: provide tests for error cases
*/

func TestModelinfoDemo(t *testing.T) {
	// Register new Vendor account
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Get all model infos
	inputModelInfos, _ := utils.GetModelInfos()

	// Get all vendors
	inputVendors, _ := utils.GetVendors()

	// Prepare model info
	firstModelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	VID := firstModelInfo.VID

	// Sign and Broadcast AddModelInfo message
	utils.SignAndBroadcastMessage(vendor, firstModelInfo)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(firstModelInfo.VID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.VID, firstModelInfo.VID)
	require.Equal(t, receivedModelInfo.Model.PID, firstModelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.ProductName, firstModelInfo.ProductName)
	require.Equal(t, receivedModelInfo.Model.Description, firstModelInfo.Description)

	// Publish second model info using POST command with passing name and passphrase. Same Vendor
	secondModelInfo := utils.NewMsgAddModelInfo(vendor.Address)
	secondModelInfo.VID = VID // Set same Vendor as for the first model
	_, _ = utils.AddModelInfo(secondModelInfo, vendor)

	// Check model is created
	receivedModelInfo, _ = utils.GetModelInfo(secondModelInfo.VID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.VID, secondModelInfo.VID)
	require.Equal(t, receivedModelInfo.Model.PID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.ProductName, secondModelInfo.ProductName)
	require.Equal(t, receivedModelInfo.Model.Description, secondModelInfo.Description)

	// Get all model infos
	modelInfos, _ := utils.GetModelInfos()
	require.Equal(t, utils.ParseUint(inputModelInfos.Total)+2, utils.ParseUint(modelInfos.Total))

	// Get all vendors
	vendors, _ := utils.GetVendors()
	require.Equal(t, utils.ParseUint(inputVendors.Total)+1, utils.ParseUint(vendors.Total))

	// Get vendor models
	vendorModels, _ := utils.GetVendorModels(VID)
	require.Equal(t, uint64(2), uint64(len(vendorModels.Products)))
	require.Equal(t, firstModelInfo.PID, vendorModels.Products[0].PID)
	require.Equal(t, secondModelInfo.PID, vendorModels.Products[1].PID)

	// Update second model info
	secondModelInfoUpdate := utils.NewMsgUpdateModelInfo(secondModelInfo.VID, secondModelInfo.PID, vendor.Address)
	_, _ = utils.UpdateModelInfo(secondModelInfoUpdate, vendor)

	// Check model is updated
	receivedModelInfo, _ = utils.GetModelInfo(secondModelInfo.VID, secondModelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.Description, secondModelInfoUpdate.Description)
}

func TestModelinfoDemo_Prepare_Sign_Broadcast(t *testing.T) {
	// Register new Vendor account
	vendor := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// Prepare model info
	modelInfo := utils.NewMsgAddModelInfo(vendor.Address)

	// Prepare Sing Broadcast
	addModelTransaction, _ := utils.PrepareAddModelInfoTransaction(modelInfo)
	_, _ = utils.SignAndBroadcastTransaction(vendor, addModelTransaction)

	// Check model is created
	receivedModelInfo, _ := utils.GetModelInfo(modelInfo.VID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.Model.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.Model.ProductName, modelInfo.ProductName)
}

/* Error cases */

func Test_AddModelinfo_ByNonVendor(t *testing.T) {
	// register new account
	testAccount := utils.CreateNewAccount(auth.AccountRoles{})

	// try to publish model info
	modelInfo := utils.NewMsgAddModelInfo(testAccount.Address)
	res, _ := utils.SignAndBroadcastMessage(testAccount, modelInfo)
	require.Equal(t, sdk.CodeUnauthorized, sdk.CodeType(res.Code))
}

func Test_AddModelinfo_Twice(t *testing.T) {
	// register new account
	testAccount := utils.CreateNewAccount(auth.AccountRoles{auth.Vendor})

	// publish model info
	modelInfo := utils.NewMsgAddModelInfo(testAccount.Address)
	res, _ := utils.AddModelInfo(modelInfo, testAccount)
	require.Equal(t, sdk.CodeOK, sdk.CodeType(res.Code))

	// publish second time
	res, _ = utils.AddModelInfo(modelInfo, testAccount)
	require.Equal(t, modelinfo.CodeModelInfoAlreadyExists, sdk.CodeType(res.Code))
}

func Test_GetModelinfo_ForUnknown(t *testing.T) {
	_, code := utils.GetModelInfo(common.RandUint16(), common.RandUint16())
	require.Equal(t, http.StatusNotFound, code)
}

func Test_GetModelinfo_ForInvalidVidPid(t *testing.T) {
	// zero vid
	_, code := utils.GetModelInfo(0, common.RandUint16())
	require.Equal(t, http.StatusBadRequest, code)

	// zero pid
	_, code = utils.GetModelInfo(common.RandUint16(), 0)
	require.Equal(t, http.StatusBadRequest, code)
}
