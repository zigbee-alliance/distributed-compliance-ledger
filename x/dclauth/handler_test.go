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
package dclauth

import (
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type TestSetup struct {
	// Cdc     *amino.Codec
	Ctx     sdk.Context
	Keeper  keeper.Keeper
	Handler sdk.Handler
	// Querier sdk.Querier
}

func Setup(t *testing.T) TestSetup {
	k, ctx := testkeeper.DclauthKeeper(t)

	setup := TestSetup{
		Ctx:     ctx,
		Keeper:  *k,
		Handler: NewHandler(*k),
	}

	return setup
}

func TestHandler_CreateAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup(t)

	countTrustees := 2

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

		// propose account
		_, address, pubKey, err := proposeAddAccount(setup, trustee)
		require.NoError(t, err)

		// ensure active account created
		account := setup.Keeper.GetAccount(setup.Ctx, address)
		require.Equal(t, address, account.GetAddress())
		require.Equal(t, pubKey, account.GetPubKey())

		// ensure no pending account created
		require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
	}
}

func TestHandler_CreateAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, []string{trustee1.String()}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, []string{trustee1.String()}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account
	pendingAccount, found = setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, []string{trustee1.String(), trustee2.String()}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(trustee3, address)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeAddAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// propose new account
		_, _, _, err := proposeAddAccount(setup, signer)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeAddAccount_ForExistingActiveAccount(t *testing.T) {
	setup := Setup(t)

	// store 2 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)

	// propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure active account created
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// propose existing active account
	proposeAddAccount, err := types.NewMsgProposeAddAccount(
		trustee2,
		address,
		pubKey,
		types.AccountRoles{types.Vendor},
		testconstants.VendorID1,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeAddAccount)
	require.ErrorIs(t, err, types.AccountAlreadyExists)
}

func TestHandler_ProposeAddAccount_ForExistingPendingAccount(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// trustee1 proposes account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// trustee2 proposes the same account
	proposeAddAccount, err := types.NewMsgProposeAddAccount(
		trustee2,
		address,
		pubKey,
		types.AccountRoles{types.Vendor},
		testconstants.VendorID1,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeAddAccount)
	require.ErrorIs(t, err, types.PendingAccountAlreadyExists)
}

func TestHandler_ApproveAddAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// propose account
	_, address, _, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// try to approve account
		approveAddAccount := types.NewMsgApproveAddAccount(signer, address)
		_, err = setup.Handler(setup.Ctx, approveAddAccount)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ApproveAddAccount_ForExistingActiveAccount(t *testing.T) {
	setup := Setup(t)

	// store 2 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)

	// propose account
	_, address, _, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure active account created
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// try to approve active account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.ErrorIs(t, err, types.PendingAccountDoesNotExist)
}

func TestHandler_ApproveAddAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve unknown account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee, testconstants.Address1)
	_, err := setup.Handler(setup.Ctx, approveAddAccount)
	require.ErrorIs(t, err, types.PendingAccountDoesNotExist)
}

func TestHandler_ApproveAddAccount_ForDuplicateApproval(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// propose account
	_, address, _, err := proposeAddAccount(setup, trustee1)
	require.NoError(t, err)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// the same trustee tries to approve the account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee1, address)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RevokeAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup(t)

	countTrustees := 2

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// store account
		address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

		// ensure 1 trustee revocation approval is needed
		require.Equal(t, 1, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

		// propose to revoke account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee, address)
		_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.NoError(t, err)

		// ensure active account removed
		require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

		// ensure no pending account revocation created
		require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
	}
}

func TestHandler_RevokeAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// ensure 2 trustee revocation approvals are needed
	require.Equal(t, 2, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	revocation, found := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, []string{trustee1.String()}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
}

func TestHandler_RevokeAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// ensure 3 trustee revocation approvals are needed
	require.Equal(t, 3, keeper.AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	revocation, found := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, []string{trustee1.String()}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account revocation
	revocation, found = setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, []string{trustee1.String(), trustee2.String()}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account revocation
	approveRevokeAccount = types.NewMsgApproveRevokeAccount(trustee3, address)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
}

func TestHandler_ProposeRevokeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// propose new account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(signer, address)
		_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// propose to revoke unknown account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee, testconstants.Address1)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.ErrorIs(t, err, types.AccountDoesNotExist)
}

func TestHandler_ProposeRevokeAccount_ForExistingPendingAccountRevocation(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// trustee2 proposes to revoke the same account
	proposeRevokeAccount = types.NewMsgProposeRevokeAccount(trustee2, address)
	_, err = setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.ErrorIs(t, err, types.PendingAccountRevocationAlreadyExists)
}

func TestHandler_ApproveRevokeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// try to approve account
		approveRevokeAccount := types.NewMsgApproveRevokeAccount(signer, address)
		_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ApproveRevokeAccount_ForAbsentPendingAccountRevocation(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// approve absent revocation of active account
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee, address)
	_, err := setup.Handler(setup.Ctx, approveRevokeAccount)
	require.ErrorIs(t, err, types.PendingAccountRevocationDoesNotExist)
}

func TestHandler_ApproveRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve absent revocation of unknown account
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee, testconstants.Address1)
	_, err := setup.Handler(setup.Ctx, approveRevokeAccount)
	require.ErrorIs(t, err, types.PendingAccountRevocationDoesNotExist)
}

func TestHandler_ApproveRevokeAccount_ForDuplicateApproval(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// propose account revocation
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// the same trustee tries to approve the account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee1, address)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ProposeAddAccount_VendorIDNotRequiredForNonVendorAccounts(t *testing.T) {
	setup := Setup(t)
	// store trustee
	trustee := storeTrustee(setup)
	_, pubKey, address := testdata.KeyTestPubAddr()
	proposeTrusteeAccount, err := types.NewMsgProposeAddAccount(
		trustee,
		address,
		pubKey,
		types.AccountRoles{types.Trustee},
		0,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeTrusteeAccount)
	require.NoError(t, err)

	_, pubKey, address = testdata.KeyTestPubAddr()
	proposeCertificationCenterAccount, err := types.NewMsgProposeAddAccount(
		trustee,
		address,
		pubKey,
		types.AccountRoles{types.CertificationCenter},
		0,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeCertificationCenterAccount)
	require.NoError(t, err)
}

func TestHandler_ProposeAddAccount_VendorIDRequiredForVendorAccounts(t *testing.T) {
	setup := Setup(t)
	// store trustee
	trustee := storeTrustee(setup)
	_, pubKey, address := testdata.KeyTestPubAddr()
	proposeVendorAccount, err := types.NewMsgProposeAddAccount(
		trustee,
		address,
		pubKey,
		types.AccountRoles{types.Vendor},
		0,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeVendorAccount)
	require.ErrorIs(t, err, types.MissingVendorIDForVendorAccount)
}

func storeTrustee(setup TestSetup) sdk.AccAddress {
	return storeAccountWithVendorID(setup, types.Trustee, 0)
}

func storeAccountWithVendorID(setup TestSetup, role types.AccountRole, vendorID int32) sdk.AccAddress {
	_, pubKey, address := testdata.KeyTestPubAddr()
	ba := authtypes.NewBaseAccount(address, pubKey, 0, 0)
	account := types.NewAccount(ba, types.AccountRoles{role}, vendorID)
	account.AccountNumber = setup.Keeper.GetNextAccountNumber(setup.Ctx)
	setup.Keeper.SetAccount(setup.Ctx, account)

	return address
}

func proposeAddAccount(setup TestSetup, signer sdk.AccAddress) (*sdk.Result, sdk.AccAddress, cryptotypes.PubKey, error) {
	_, pubKey, address := testdata.KeyTestPubAddr()
	proposeAddAccount, _ := types.NewMsgProposeAddAccount(
		signer,
		address,
		pubKey,
		types.AccountRoles{types.Vendor},
		testconstants.VendorID1,
	)
	// TODO check the err here
	result, err := setup.Handler(setup.Ctx, proposeAddAccount)

	return result, address, pubKey, err
}
