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
	require.Equal(t, msgCreateVendorInfo.CompanyPrefferedName, receivedVendorInfo.CompanyPrefferedName)
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
	require.Equal(t, msgUpdateVendorInfo.CompanyPrefferedName, receivedVendorInfo.CompanyPrefferedName)
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

func TestHandler_AddVendorInfoWithEmptyOptionalFields(t *testing.T) {
	setup := Setup(t)

	// add new msgCreateVendorInfo
	msgCreateVendorInfo := NewMsgCreateVendorInfo(setup.Vendor)
	msgCreateVendorInfo.CompanyPrefferedName = "" // Set empty CID

	_, err := setup.Handler(setup.Ctx, msgCreateVendorInfo)
	require.NoError(t, err)

	// query vendorinfo
	receivedVendorInfo, err := queryVendorInfo(setup, msgCreateVendorInfo.VendorID)
	require.NoError(t, err)

	// check
	require.Equal(t, "", receivedVendorInfo.CompanyPrefferedName)
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
		CompanyPrefferedName: testconstants.CompanyPreferredName,
		VendorName:           testconstants.VendorName,
		VendorLandingPageURL: testconstants.VendorLandingPageUrl,
	}
}

func NewMsgUpdateVendorInfo(signer sdk.AccAddress) *types.MsgUpdateVendorInfo {
	return &types.MsgUpdateVendorInfo{
		Creator:              signer.String(),
		VendorID:             testconstants.VendorID1,
		CompanyLegalName:     testconstants.CompanyLegalName + "/updated",
		CompanyPrefferedName: testconstants.CompanyPreferredName + "/updated",
		VendorName:           testconstants.VendorName + "/updated",
		VendorLandingPageURL: testconstants.VendorLandingPageUrl + "/updated",
	}
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}
