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
package vendorinfo

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func TestHandler_AddVendorInfo(t *testing.T) {
	setup := Setup()

	// add new vendorInfo
	vendorInfo := TestMsgAddVendorInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, vendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query vendorInfo
	receivedVendorInfo := queryVendorInfo(setup, vendorInfo.VendorId)

	// check
	require.Equal(t, vendorInfo.VendorId, receivedVendorInfo.VendorId)
	require.Equal(t, vendorInfo.VendorName, receivedVendorInfo.VendorName)
	require.Equal(t, vendorInfo.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(t, vendorInfo.CompanyPreferredName, receivedVendorInfo.CompanyPreferredName)
	require.Equal(t, vendorInfo.VendorLandingPageUrl, receivedVendorInfo.VendorLandingPageUrl)

}

func TestHandler_UpdateVendorInfo(t *testing.T) {
	setup := Setup()

	// try update vendor for non existant vendor
	msgUpdateVendorInfo := TestMsgUpdateVendorInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.Equal(t, types.CodeVendorDoesNotExist, result.Code)

	// add new vendorInfo
	msgAddVendorInfo := TestMsgAddVendorInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgAddVendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// update existing vendorInfo
	result = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query updated vendorInfo
	recievedVendorInfo := queryVendorInfo(setup, msgUpdateVendorInfo.VendorId)

	// check
	require.Equal(t, msgUpdateVendorInfo.VendorId, recievedVendorInfo.VendorId)
	require.Equal(t, msgUpdateVendorInfo.VendorName, recievedVendorInfo.VendorName)
	require.Equal(t, msgUpdateVendorInfo.CompanyLegalName, recievedVendorInfo.CompanyLegalName)
	require.Equal(t, msgUpdateVendorInfo.CompanyPreferredName, recievedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgUpdateVendorInfo.VendorLandingPageUrl, recievedVendorInfo.VendorLandingPageUrl)
}

func TestHandler_OnlyOwnerCanUpdateVendorInfo(t *testing.T) {
	setup := Setup()

	// add new vendorInfo
	msgAddVendorInfo := TestMsgAddVendorInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddVendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.Vendor} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorId3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// update existing VendorInfo by not owner
		msgUpdateVendorInfo := TestMsgUpdateVendorInfo(testconstants.Address3)
		result = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}

	// owner update existing VendorInfo
	msgUpdateVendorInfo := TestMsgUpdateVendorInfo(setup.Vendor)
	result = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)
}

func TestHandler_AddVendorInfoWithEmptyOptionalFields(t *testing.T) {
	setup := Setup()

	// add new vendorInfo
	vendorInfo := TestMsgAddVendorInfo(setup.Vendor)
	vendorInfo.CompanyLegalName = ""

	result := setup.Handler(setup.Ctx, vendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query vendorInfo
	receivedVendorInfo := queryVendorInfo(setup, testconstants.VID)

	// check
	require.Equal(t, receivedVendorInfo.CompanyLegalName, "")

}

func TestHandler_AddVendorByNonVendor(t *testing.T) {
	setup := Setup()

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse} {
		// store account
		account := auth.NewAccount(testconstants.Address3, testconstants.PubKey3, auth.AccountRoles{role}, testconstants.VendorId3)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// add new VendorInfo
		VendorInfo := TestMsgAddVendorInfo(testconstants.Address3)
		result := setup.Handler(setup.Ctx, VendorInfo)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_PartiallyUpdateVendor(t *testing.T) {
	setup := Setup()

	// add new vendorInfo
	msgAddVendorInfo := TestMsgAddVendorInfo(setup.Vendor)
	result := setup.Handler(setup.Ctx, msgAddVendorInfo)

	// owner update Description of existing vendorInfo
	msgUpdateVendorInfo := TestMsgUpdateVendorInfo(setup.Vendor)

	msgUpdateVendorInfo.CompanyPreferredName = "New Preferred Name"
	msgUpdateVendorInfo.VendorLandingPageUrl = "https://new.example.com"
	result = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query VendorInfo
	receivedVendorInfo := queryVendorInfo(setup, msgUpdateVendorInfo.VendorId)

	// check
	require.Equal(t, receivedVendorInfo.CompanyPreferredName, msgUpdateVendorInfo.CompanyPreferredName)
	require.Equal(t, receivedVendorInfo.VendorLandingPageUrl, msgUpdateVendorInfo.VendorLandingPageUrl)
}

func queryVendorInfo(setup TestSetup, vid uint16) types.VendorInfo {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryVendor, fmt.Sprintf("%v", vid)},
		abci.RequestQuery{},
	)

	var receivedVendorInfo types.VendorInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedVendorInfo)

	return receivedVendorInfo
}
