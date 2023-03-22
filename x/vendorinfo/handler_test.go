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
package vendorinfo

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	Vendor   sdk.AccAddress
	VendorID int32
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vendorID int32,
) {
	for _, role := range roles {
		setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	setup.DclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)

	setup.DclauthKeeper.On("HasVendorID", mock.Anything, accAddress, vendorID).Return(true)
	setup.DclauthKeeper.On("HasVendorID", mock.Anything, accAddress, mock.Anything).Return(false)
}

func Setup(t *testing.T) TestSetup {
	t.Helper()
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := testkeeper.VendorinfoKeeper(t, dclauthKeeper)

	vendor := GenerateAccAddress()
	vendorID := testconstants.VendorID1

	setup := TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		Handler:       NewHandler(*keeper),
		Vendor:        vendor,
		VendorID:      vendorID,
	}

	setup.AddAccount(vendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vendorID)

	return setup
}

func TestHandler_AddVendorInfo(t *testing.T) {
	setup := Setup(t)

	// add new vendorinfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// query vendorinfo
	receivedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check
	require.Equal(t, msgCreateVendorInfo.VendorID, receivedVendorInfo.VendorID)
	require.Equal(t, msgCreateVendorInfo.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(t, msgCreateVendorInfo.CompanyPreferredName, receivedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgCreateVendorInfo.Creator, receivedVendorInfo.Creator)
	require.Equal(t, msgCreateVendorInfo.VendorLandingPageURL, receivedVendorInfo.VendorLandingPageURL)
	require.Equal(t, msgCreateVendorInfo.VendorName, receivedVendorInfo.VendorName)
}

func queryVendorInfo(
	setup TestSetup,
	vid int32,
) (*types.VendorInfo, error) {
	req := &types.QueryGetVendorInfoRequest{
		VendorID: vid,
	}

	resp, err := setup.Keeper.VendorInfo(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.VendorInfo, nil
}

func TestHandler_UpdateVendorInfo(t *testing.T) {
	setup := Setup(t)

	// try update not present vendorinfo
	msgUpdateVendorInfo := NewMsgUpdateVendorInfo(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	println(err.Error())
	require.Error(t, err)

	// add new vendorinfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// update existing vendorinfo
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	// query updated vendorinfo
	receivedVendorInfo, err := queryVendorInfo(setup, msgUpdateVendorInfo.VendorID)
	require.NoError(t, err)

	// check
	require.Equal(t, msgUpdateVendorInfo.VendorID, receivedVendorInfo.VendorID)
	require.Equal(t, msgUpdateVendorInfo.CompanyLegalName, receivedVendorInfo.CompanyLegalName)
	require.Equal(t, msgUpdateVendorInfo.CompanyPreferredName, receivedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgUpdateVendorInfo.Creator, receivedVendorInfo.Creator)
	require.Equal(t, msgUpdateVendorInfo.VendorLandingPageURL, receivedVendorInfo.VendorLandingPageURL)
	require.Equal(t, msgUpdateVendorInfo.VendorName, receivedVendorInfo.VendorName)
}

func TestHandler_OnlyOwnerCanUpdateVendorInfo(t *testing.T) {
	setup := Setup(t)

	// add new vendorinfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// update existing vendorinfo by user without Vendor role
		msgUpdateVendorInfo := NewMsgUpdateVendorInfo(accAddress)
		_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
		require.Error(t, err)
	}

	anotherVendor := GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	// update existing vendorinfo by vendor with another VendorID
	msgUpdateVendorInfo := NewMsgUpdateVendorInfo(anotherVendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))

	// update existing vendorinfo by owner
	msgUpdateVendorInfo = NewMsgUpdateVendorInfo(setup.Vendor)
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)
}

func TestHandler_UpdateVendorInfoWithEmptyOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add new msgCreateVendorInfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)

	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// query vendorinfo
	vendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	msgUpdateVendorInfo := &types.MsgUpdateVendorInfo{
		Creator:              vendorInfo.Creator,
		VendorID:             vendorInfo.VendorID,
		CompanyLegalName:     "",
		CompanyPreferredName: "",
		VendorName:           "",
		VendorLandingPageURL: "",
	}
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	// query the updated vendorinfo
	updatedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	require.Equal(t, updatedVendorInfo.Creator, vendorInfo.Creator)
	require.Equal(t, updatedVendorInfo.VendorID, vendorInfo.VendorID)
	require.Equal(t, updatedVendorInfo.CompanyLegalName, vendorInfo.CompanyLegalName)
	require.Equal(t, updatedVendorInfo.CompanyPreferredName, vendorInfo.CompanyPreferredName)
	require.Equal(t, updatedVendorInfo.VendorName, vendorInfo.VendorName)
	require.Equal(t, updatedVendorInfo.VendorLandingPageURL, vendorInfo.VendorLandingPageURL)
}

func TestHandler_UpdateVendorInfoWithAllOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add new msgCreateVendorInfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)

	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// query vendorinfo
	vendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	companyLegalName := "1"
	companyPreferredName := "2"
	vendorName := "3"
	vendorLandingPageURL := "4"

	msgUpdateVendorInfo := &types.MsgUpdateVendorInfo{
		Creator:              vendorInfo.Creator,
		VendorID:             vendorInfo.VendorID,
		CompanyLegalName:     companyLegalName,
		CompanyPreferredName: companyPreferredName,
		VendorName:           vendorName,
		VendorLandingPageURL: vendorLandingPageURL,
	}
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	// query the updated vendorinfo
	updatedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	require.Equal(t, updatedVendorInfo.Creator, vendorInfo.Creator)
	require.Equal(t, updatedVendorInfo.VendorID, vendorInfo.VendorID)
	require.Equal(t, updatedVendorInfo.CompanyLegalName, companyLegalName)
	require.Equal(t, updatedVendorInfo.CompanyPreferredName, companyPreferredName)
	require.Equal(t, updatedVendorInfo.VendorName, vendorName)
	require.Equal(t, updatedVendorInfo.VendorLandingPageURL, vendorLandingPageURL)
}

func TestHandler_AddVendorInfoWithEmptyOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add new msgCreateVendorInfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)
	msgCreateVendorInfo.CompanyPreferredName = "" // Set empty CID

	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// query vendorinfo
	receivedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check
	require.Equal(t, "", receivedVendorInfo.CompanyPreferredName)
}

func TestHandler_AddVendorInfoByNonVendor(t *testing.T) {
	setup := Setup(t)

	for _, role := range []dclauthtypes.AccountRole{
		dclauthtypes.CertificationCenter,
		dclauthtypes.Trustee,
		dclauthtypes.NodeAdmin,
	} {
		accAddress := GenerateAccAddress()
		setup.AddAccount(accAddress, []dclauthtypes.AccountRole{role}, setup.VendorID)

		// add new vendorinfo
		vendorinfo := NewMsgCreateVendorInfo(accAddress)
		_, err := setup.Handler(setup.Ctx, vendorinfo)
		require.Error(t, err)
		require.True(t, sdkerrors.ErrUnauthorized.Is(err))
	}
}

func TestHandler_AddVendorInfoByVendorAdmin(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()

	// add a new account with VendoeAdmin role
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, setup.VendorID)

	msgCreateVendorInfo := NewMsgCreateVendorInfo(accAddress)
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	newVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check if stored correctly
	require.Equal(t, msgCreateVendorInfo.Creator, newVendorInfo.Creator)
	require.Equal(t, msgCreateVendorInfo.VendorID, newVendorInfo.VendorID)
	require.Equal(t, msgCreateVendorInfo.CompanyLegalName, newVendorInfo.CompanyLegalName)
	require.Equal(t, msgCreateVendorInfo.CompanyPreferredName, newVendorInfo.CompanyPreferredName)
	require.Equal(t, msgCreateVendorInfo.VendorName, newVendorInfo.VendorName)
	require.Equal(t, msgCreateVendorInfo.VendorLandingPageURL, newVendorInfo.VendorLandingPageURL)
}

func TestHandler_AddAndUpdateVendorInfoByVendorAdmin(t *testing.T) {
	setup := Setup(t)

	accAddress := GenerateAccAddress()

	// add a new account with VendoeAdmin role
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, setup.VendorID)

	msgCreateVendorInfo := NewMsgCreateVendorInfo(accAddress)
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	msgUpdateVendorInfo := NewMsgUpdateVendorInfo(accAddress)
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	updatedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check if updated correctly
	require.Equal(t, msgUpdateVendorInfo.Creator, updatedVendorInfo.Creator)
	require.Equal(t, msgUpdateVendorInfo.VendorID, updatedVendorInfo.VendorID)
	require.Equal(t, msgUpdateVendorInfo.CompanyLegalName, updatedVendorInfo.CompanyLegalName)
	require.Equal(t, msgUpdateVendorInfo.CompanyPreferredName, updatedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgUpdateVendorInfo.VendorName, updatedVendorInfo.VendorName)
	require.Equal(t, msgUpdateVendorInfo.VendorLandingPageURL, updatedVendorInfo.VendorLandingPageURL)
}

func TestHandled_AddVendorInfoByVendorUpdateVendorInfoByVendorAdmin(t *testing.T) {
	setup := Setup(t)

	vendorAccAddress := GenerateAccAddress()
	vendorAdminAccAddress := GenerateAccAddress()

	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID)
	setup.AddAccount(vendorAdminAccAddress, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, setup.VendorID+2)

	// create vendorInfo by Vendor
	msgCreateVendorInfo := NewMsgCreateVendorInfo(vendorAccAddress)
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// update vendorInfo by VendorAdmin
	msgUpdateVendorInfo := NewMsgUpdateVendorInfo(vendorAdminAccAddress)
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	updatedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check if updated correctly
	require.Equal(t, msgUpdateVendorInfo.Creator, updatedVendorInfo.Creator)
	require.Equal(t, msgUpdateVendorInfo.VendorID, updatedVendorInfo.VendorID)
	require.Equal(t, msgUpdateVendorInfo.CompanyLegalName, updatedVendorInfo.CompanyLegalName)
	require.Equal(t, msgUpdateVendorInfo.CompanyPreferredName, updatedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgUpdateVendorInfo.VendorName, updatedVendorInfo.VendorName)
	require.Equal(t, msgUpdateVendorInfo.VendorLandingPageURL, updatedVendorInfo.VendorLandingPageURL)
}

func TestHandled_AddVendorInfoByVendorAdminUpdateVendorInfoByVendor(t *testing.T) {
	setup := Setup(t)

	vendorAccAddress := GenerateAccAddress()
	vendorAdminAccAddress := GenerateAccAddress()

	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, setup.VendorID)
	setup.AddAccount(vendorAdminAccAddress, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, setup.VendorID+1)

	// create vendorInfo by VendorAdmin
	msgCreateVendorInfo := NewMsgCreateVendorInfo(vendorAdminAccAddress)
	msgCreateVendorInfo.VendorID = setup.VendorID
	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// update vendorInfo by Vendor
	msgUpdateVendorInfo := NewMsgUpdateVendorInfo(vendorAccAddress)
	_, err = setup.Handler(setup.Ctx, msgUpdateVendorInfo)
	require.NoError(t, err)

	updatedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check if updated correctly
	require.Equal(t, msgUpdateVendorInfo.Creator, updatedVendorInfo.Creator)
	require.Equal(t, msgUpdateVendorInfo.VendorID, updatedVendorInfo.VendorID)
	require.Equal(t, msgUpdateVendorInfo.CompanyLegalName, updatedVendorInfo.CompanyLegalName)
	require.Equal(t, msgUpdateVendorInfo.CompanyPreferredName, updatedVendorInfo.CompanyPreferredName)
	require.Equal(t, msgUpdateVendorInfo.VendorName, updatedVendorInfo.VendorName)
	require.Equal(t, msgUpdateVendorInfo.VendorLandingPageURL, updatedVendorInfo.VendorLandingPageURL)
}

func TestHandler_AddVendorInfoByVendorWithAnotherVendorId(t *testing.T) {
	setup := Setup(t)

	anotherVendor := GenerateAccAddress()
	setup.AddAccount(anotherVendor, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID2)

	// add new vendorinfo
	vendorinfo := NewMsgCreateVendorInfo(anotherVendor)
	_, err := setup.Handler(setup.Ctx, vendorinfo)
	require.Error(t, err)
	require.True(t, sdkerrors.ErrUnauthorized.Is(err))
}

func NewMsgCreateVendorInfo(signer sdk.AccAddress) *types.MsgCreateVendorInfo {
	return &types.MsgCreateVendorInfo{
		Creator:              signer.String(),
		VendorID:             testconstants.VendorID1,
		CompanyLegalName:     testconstants.CompanyLegalName,
		CompanyPreferredName: testconstants.CompanyPreferredName,
		VendorName:           testconstants.VendorName,
		VendorLandingPageURL: testconstants.VendorLandingPageURL,
	}
}

func NewMsgUpdateVendorInfo(signer sdk.AccAddress) *types.MsgUpdateVendorInfo {
	return &types.MsgUpdateVendorInfo{
		Creator:              signer.String(),
		VendorID:             testconstants.VendorID1,
		CompanyLegalName:     testconstants.CompanyLegalName + "/updated",
		CompanyPreferredName: testconstants.CompanyPreferredName + "/updated",
		VendorName:           testconstants.VendorName + "/updated",
		VendorLandingPageURL: testconstants.VendorLandingPageURL + "/updated",
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()

	return accAddress
}
