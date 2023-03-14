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
	t.Helper()
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

	countTrustees := 1

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

		// propose account

		_, address, pubKey, err := proposeAddAccount(setup, trustee, types.AccountRoles{types.NodeAdmin})
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
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, testconstants.Info, pendingAccount.Approvals[0].Info)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[1].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account
	pendingAccount, found = setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee2))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(trustee3, address, testconstants.Info3)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info2, dclAccount.Approvals[1].Info)
	require.Equal(t, testconstants.Info3, dclAccount.Approvals[2].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)
	require.Equal(t, trustee3.String(), dclAccount.Approvals[2].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_NeededApprovalCountForManyTrustees(t *testing.T) {
	expectedApprovalCounts := [400]int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9, 10, 10, 11, 12, 12, 13, 14, 14, 15, 16, 16, 17, 18, 18, 19, 20, 20, 21, 22, 22, 23, 24, 24, 25, 26, 26, 27, 28, 28, 29, 30, 30, 31, 32, 32, 33, 34, 34, 35, 36, 36, 37, 38, 38, 39, 40, 40, 41, 42, 42, 43, 44, 44, 45, 46, 46, 47, 48, 48, 49, 50, 50, 51, 52, 52, 53, 54, 54, 55, 56, 56, 57, 58, 58, 59, 60, 60, 61, 62, 62, 63, 64, 64, 65, 66, 66, 67, 68, 68, 69, 70, 70, 71, 72, 72, 73, 74, 74, 75, 76, 76, 77, 78, 78, 79, 80, 80, 81, 82, 82, 83, 84, 84, 85, 86, 86, 87, 88, 88, 89, 90, 90, 91, 92, 92, 93, 94, 94, 95, 96, 96, 97, 98, 98, 99, 100, 100, 101, 102, 102, 103, 104, 104, 105, 106, 106, 107, 108, 108, 109, 110, 110, 111, 112, 112, 113, 114, 114, 115, 116, 116, 117, 118, 118, 119, 120, 120, 121, 122, 122, 123, 124, 124, 125, 126, 126, 127, 128, 128, 129, 130, 130, 131, 132, 132, 133, 134, 134, 135, 136, 136, 137, 138, 138, 139, 140, 140, 141, 142, 142, 143, 144, 144, 145, 146, 146, 147, 148, 148, 149, 150, 150, 151, 152, 152, 153, 154, 154, 155, 156, 156, 157, 158, 158, 159, 160, 160, 161, 162, 162, 163, 164, 164, 165, 166, 166, 167, 168, 168, 169, 170, 170, 171, 172, 172, 173, 174, 174, 175, 176, 176, 177, 178, 178, 179, 180, 180, 181, 182, 182, 183, 184, 184, 185, 186, 186, 187, 188, 188, 189, 190, 190, 191, 192, 192, 193, 194, 194, 195, 196, 196, 197, 198, 198, 199, 200, 200, 201, 202, 202, 203, 204, 204, 205, 206, 206, 207, 208, 208, 209, 210, 210, 211, 212, 212, 213, 214, 214, 215, 216, 216, 217, 218, 218, 219, 220, 220, 221, 222, 222, 223, 224, 224, 225, 226, 226, 227, 228, 228, 229, 230, 230, 231, 232, 232, 233, 234, 234, 235, 236, 236, 237, 238, 238, 239, 240, 240, 241, 242, 242, 243, 244, 244, 245, 246, 246, 247, 248, 248, 249, 250, 250, 251, 252, 252, 253, 254, 254, 255, 256, 256, 257, 258, 258, 259, 260, 260, 261, 262, 262, 263, 264, 264, 265, 266, 266, 267}

	setup := Setup(t)

	for i := 0; i < 400; i++ {
		storeTrustee(setup)
		require.Equal(t, expectedApprovalCounts[i], setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))
	}
}

func TestHandler_ProposeAndRejectAccount(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee1 rejects account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee1, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// ensure no pending account present
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeAddAndRejectAccount_ByAnotherTrustee(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// ensure pending account present
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeAddAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// propose new account
		_, _, _, err := proposeAddAccount(setup, signer, types.AccountRoles{types.NodeAdmin})
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeAddAccount_ForExistingActiveAccount(t *testing.T) {
	setup := Setup(t)

	// store 2 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)

	// propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// approve account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
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
		testconstants.Info,
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
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
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
		testconstants.Info,
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
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// try to approve account
		approveAddAccount := types.NewMsgApproveAddAccount(signer, address, testconstants.Info)
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
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// approve account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// ensure active account created
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// try to approve active account
	approveAddAccount = types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.ErrorIs(t, err, types.PendingAccountDoesNotExist)
}

func TestHandler_ApproveAddAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve unknown account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee, testconstants.Address1, testconstants.Info)
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
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// the same trustee tries to approve the account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee1, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RevokeAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup(t)

	countTrustees := 1

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// store account
		address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

		// ensure 1 trustee revocation approval is needed
		require.Equal(t, 1, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

		// propose to revoke account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee, address, testconstants.Info)
		_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.NoError(t, err)

		// ensure active account removed
		require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

		// ensure no pending account revocation created
		require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

		// ensure that account be in entity revoked account
		revokedAccount, isFound := setup.Keeper.GetRevokedAccount(setup.Ctx, address)
		require.True(t, isFound)
		require.Equal(t, address.String(), revokedAccount.Address)
		require.Equal(t, trustee.String(), revokedAccount.RevokeApprovals[0].Address)
		require.Equal(t, types.RevokedAccount_TrusteeVoting, revokedAccount.Reason)
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
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	revocation, found := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, true, revocation.HasRevocationFrom(trustee1))

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// ensure adding account to entity RevokedAccount
	require.True(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))

	// get revoked account and check fields
	revokedAccount, isFound := setup.Keeper.GetRevokedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address.String(), revokedAccount.Address)
	require.Equal(t, trustee1.String(), revokedAccount.RevokeApprovals[0].Address)
	require.Equal(t, trustee2.String(), revokedAccount.RevokeApprovals[1].Address)
	require.Equal(t, types.RevokedAccount_TrusteeVoting, revokedAccount.Reason)
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
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	revocation, found := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, true, revocation.HasRevocationFrom(trustee1))

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account revocation
	revocation, found = setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, true, revocation.HasRevocationFrom(trustee2))

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account revocation
	approveRevokeAccount = types.NewMsgApproveRevokeAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// ensure adding account to entity RevokedAccount
	require.True(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))

	// get revoked account and check fields
	revokedAccount, isFound := setup.Keeper.GetRevokedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address.String(), revokedAccount.Address)
	require.Equal(t, trustee1.String(), revokedAccount.RevokeApprovals[0].Address)
	require.Equal(t, trustee2.String(), revokedAccount.RevokeApprovals[1].Address)
	require.Equal(t, trustee3.String(), revokedAccount.RevokeApprovals[2].Address)
	require.Equal(t, types.RevokedAccount_TrusteeVoting, revokedAccount.Reason)
}

func TestHandler_ReAdding_RevokedAccount(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// ensure 2 trustee revocation approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	revocation, found := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), revocation.Address)
	require.Equal(t, true, revocation.HasRevocationFrom(trustee1))

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.NoError(t, err)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// ensure adding account to entity RevokedAccount
	require.True(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, testconstants.Info, pendingAccount.Approvals[0].Info)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure active account revoked in entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[1].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeRevokeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// propose new account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(signer, address, testconstants.Info)
		_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_ProposeRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// propose to revoke unknown account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee, testconstants.Address1, testconstants.Info)
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
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// trustee2 proposes to revoke the same account
	proposeRevokeAccount = types.NewMsgProposeRevokeAccount(trustee2, address, testconstants.Info)
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
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// try to approve account
		approveRevokeAccount := types.NewMsgApproveRevokeAccount(signer, address, testconstants.Info)
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
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, approveRevokeAccount)
	require.ErrorIs(t, err, types.PendingAccountRevocationDoesNotExist)
}

func TestHandler_ApproveRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve absent revocation of unknown account
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee, testconstants.Address1, testconstants.Info)
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
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// the same trustee tries to approve the account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(trustee1, address, testconstants.Info)
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
		testconstants.Info,
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
		testconstants.Info,
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
		testconstants.Info,
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, proposeVendorAccount)
	require.ErrorIs(t, err, types.MissingVendorIDForVendorAccount)
}

func TestHandler_RejectAccount_TwoRejectApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 3 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes account
	_, address, pubkey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// ensure account doesn't exist in entity RejectedAccount
	_, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// trustee3 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account must be created
	rejectedAccount, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address, rejectedAccount.GetAddress())
	require.Equal(t, pubkey, rejectedAccount.GetPubKey())

	// check for info, approvals, rejectedApprovals fields
	require.Equal(t, testconstants.Info, rejectedAccount.Approvals[0].Info)
	require.Equal(t, trustee1.String(), rejectedAccount.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedAccount.Rejects[0].Info)
	require.Equal(t, trustee2.String(), rejectedAccount.Rejects[0].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// ensure account doesn't exist in entity Account
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure account doesn't exist in entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_RejectAccount_ThreeRejectApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 7 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	trustee4 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 5 trustee approvals are needed
	require.Equal(t, 5, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee 1 proposes account
	_, address, pubkey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// ensure account doesn't exist in entity RejectedAccount
	_, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// trustee3 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// ensure account doesn't exist in entity RejectedAccount
	_, isFound = setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// trustee4 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee4, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account must be created
	rejectedAccount, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address, rejectedAccount.GetAddress())
	require.Equal(t, pubkey, rejectedAccount.GetPubKey())

	// check for info, approvals, rejectedApprovals fields
	require.Equal(t, testconstants.Info, rejectedAccount.Approvals[0].Info)
	require.Equal(t, trustee1.String(), rejectedAccount.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedAccount.Rejects[0].Info)
	require.Equal(t, trustee2.String(), rejectedAccount.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedAccount.Rejects[1].Info)
	require.Equal(t, trustee3.String(), rejectedAccount.Rejects[1].Address)
	require.Equal(t, testconstants.Info, rejectedAccount.Rejects[2].Info)
	require.Equal(t, trustee4.String(), rejectedAccount.Rejects[2].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// ensure account doesn't exist in entity Account
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure account doesn't exist in entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_RejectAccount_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccountWithVendorID(setup, types.Vendor, testconstants.VendorID1)

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(trustee1, address, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.NoError(t, err)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	for _, role := range []types.AccountRole{types.Vendor, types.CertificationCenter, types.NodeAdmin} {
		// store signer account
		signer := storeAccountWithVendorID(setup, role, testconstants.VendorID1)

		// reject new account
		rejectAddAccount := types.NewMsgRejectAddAccount(signer, address, testconstants.Info)
		_, err := setup.Handler(setup.Ctx, rejectAddAccount)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_RejectAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// store 1 trustee
	trustee := storeTrustee(setup)

	// reject unknown account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee, testconstants.Address1, testconstants.Info)
	_, err := setup.Handler(setup.Ctx, rejectAddAccount)
	require.ErrorIs(t, err, types.PendingAccountDoesNotExist)
}

func TestHandler_Duplicate_RejectAccountFromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// store 3 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes account
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account doesn't exist in entity RejecedAccount
	_, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// second time trustee2 try rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ApproveAccountAndRejectAccount_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// store 4 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes account
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 approves to add account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	pendingAcc, _ := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	prevRejectsLen := len(pendingAcc.Rejects)
	prevApprovalsLen := len(pendingAcc.Approvals)
	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	pendingAcc, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, len(pendingAcc.Rejects), prevRejectsLen+1)
	require.Equal(t, len(pendingAcc.Approvals), prevApprovalsLen-1)
}

func TestHandler_RejectAccountAndApproveAccount_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// store 4 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes account
	_, address, _, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	pendingAcc, _ := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	prevRejectsLen := len(pendingAcc.Rejects)
	prevApprovalsLen := len(pendingAcc.Approvals)
	// trustee2 approves to add account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	pendingAcc, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, len(pendingAcc.Rejects), prevRejectsLen-1)
	require.Equal(t, len(pendingAcc.Approvals), prevApprovalsLen+1)
}

func TestHandler_DoubleTimeRejectAccount(t *testing.T) {
	setup := Setup(t)

	// storee 3 trustee
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 proposes account
	_, address, pubkey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// trustee2 rejects to add account
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// trustee3 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account must be created
	rejectedAccountFirstTime, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address, rejectedAccountFirstTime.GetAddress())
	require.Equal(t, pubkey, rejectedAccountFirstTime.GetPubKey())

	// check for info, approvals, rejectedApprovals fields
	require.Equal(t, testconstants.Info, rejectedAccountFirstTime.Approvals[0].Info)
	require.Equal(t, trustee1.String(), rejectedAccountFirstTime.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectedAccountFirstTime.Rejects[0].Info)
	require.Equal(t, trustee2.String(), rejectedAccountFirstTime.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectedAccountFirstTime.Rejects[1].Info)
	require.Equal(t, trustee3.String(), rejectedAccountFirstTime.Rejects[1].Address)

	// trustee1 second time proposes account
	_, address, pubkey, err = proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin})
	require.NoError(t, err)

	// ensure that account not exist in <Rejected Account>
	_, isFound = setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// trustee2 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account doesn't exist in entity RejecedAccount
	_, isFound = setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.False(t, isFound)

	// trustee3 rejects to add account
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// rejected account must be created
	rejectAccountSecondTime, isFound := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, isFound)
	require.Equal(t, address, rejectAccountSecondTime.GetAddress())
	require.Equal(t, pubkey, rejectAccountSecondTime.GetPubKey())

	// check for info, approvals, rejectedApprovals fields
	require.Equal(t, testconstants.Info, rejectAccountSecondTime.Approvals[0].Info)
	require.Equal(t, trustee1.String(), rejectAccountSecondTime.Approvals[0].Address)
	require.Equal(t, testconstants.Info, rejectAccountSecondTime.Rejects[0].Info)
	require.Equal(t, trustee2.String(), rejectAccountSecondTime.Rejects[0].Address)
	require.Equal(t, testconstants.Info, rejectAccountSecondTime.Rejects[1].Info)
	require.Equal(t, trustee3.String(), rejectAccountSecondTime.Rejects[1].Address)
}

func TestHandler_CreateVendorAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup(t)

	countTrustees := 1

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.VendorAccountApprovalsPercent))

		// propose account

		_, address, pubKey, err := proposeAddAccount(setup, trustee, types.AccountRoles{types.Vendor})
		require.NoError(t, err)

		// ensure active account created
		account := setup.Keeper.GetAccount(setup.Ctx, address)
		require.Equal(t, address, account.GetAddress())
		require.Equal(t, pubKey, account.GetPubKey())

		// ensure no pending account created
		require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
	}
}

func TestHandler_CreateVendorAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 5 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.VendorAccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.Vendor})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, testconstants.Info, pendingAccount.Approvals[0].Info)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[1].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateVendorAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 8 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.VendorAccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.Vendor})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account
	pendingAccount, found = setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee2))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(trustee3, address, testconstants.Info3)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info2, dclAccount.Approvals[1].Info)
	require.Equal(t, testconstants.Info3, dclAccount.Approvals[2].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)
	require.Equal(t, trustee3.String(), dclAccount.Approvals[2].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateVendorAccountWithDifferentRole_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup(t)

	countTrustees := 1

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

		// propose account

		_, address, pubKey, err := proposeAddAccount(setup, trustee, types.AccountRoles{types.NodeAdmin, types.Vendor})
		require.NoError(t, err)

		// ensure active account created
		account := setup.Keeper.GetAccount(setup.Ctx, address)
		require.Equal(t, address, account.GetAddress())
		require.Equal(t, pubKey, account.GetPubKey())

		// ensure no pending account created
		require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
	}
}

func TestHandler_CreateVendorAccountWithDifferentRole_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin, types.Vendor})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, testconstants.Info, pendingAccount.Approvals[0].Info)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[1].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateVendorAccountWithDifferentRole_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin, types.Vendor})
	require.NoError(t, err)

	// ensure pending account created
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// ensure second approval added to pending account
	pendingAccount, found = setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee1))
	require.Equal(t, true, pendingAccount.HasApprovalFrom(trustee2))

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(trustee3, address, testconstants.Info3)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.GetAddress())
	require.Equal(t, pubKey, account.GetPubKey())

	// check for info field and approvals
	dclAccount, _ := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.Equal(t, testconstants.Info, dclAccount.Approvals[0].Info)
	require.Equal(t, testconstants.Info2, dclAccount.Approvals[1].Info)
	require.Equal(t, testconstants.Info3, dclAccount.Approvals[2].Info)
	require.Equal(t, trustee1.String(), dclAccount.Approvals[0].Address)
	require.Equal(t, trustee2.String(), dclAccount.Approvals[1].Address)
	require.Equal(t, trustee3.String(), dclAccount.Approvals[2].Address)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// check that account revoked from entity RevokedAccount
	require.False(t, setup.Keeper.IsRevokedAccountPresent(setup.Ctx, address))
}

func TestHandler_RejectAccount_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make account rejected

	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 4 trustee approvals are needed
	require.Equal(t, 4, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin, types.Vendor})
	require.NoError(t, err)

	// reject account by account Trustee2
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Proposed Account>, because we haven't enough reject approvals
	pendingAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)

	// check proposed account
	require.Equal(t, address.String(), pendingAccount.Address)
	require.Equal(t, pubKey, pendingAccount.GetPubKey())

	// reject account by account Trustee3
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Rejected Account>, because we have enough rejected approvals
	rejectedAccount, found := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, found)

	// check rejected account
	require.Equal(t, address.String(), rejectedAccount.Address)
	require.Equal(t, pubKey, rejectedAccount.GetPubKey())
}

func TestHandler_ApproveAccount_FourApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 3 more approvals

	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	trustee4 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 4 trustee approvals are needed
	require.Equal(t, 4, setup.Keeper.AccountApprovalsCount(setup.Ctx, types.AccountApprovalsPercent))

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.NodeAdmin, types.Vendor})
	require.NoError(t, err)

	// approve account by account Trustee2
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// approve account by account Trustee3
	approveAddAccount = types.NewMsgApproveAddAccount(trustee3, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Proposed Account>, because we haven't enough approvals
	proposedAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)

	// check proposed account
	require.Equal(t, address.String(), proposedAccount.Address)
	require.Equal(t, pubKey, proposedAccount.GetPubKey())

	// approve account by account Trustee4
	approveAddAccount = types.NewMsgApproveAddAccount(trustee4, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Account>, because we have enough approvals
	approvedAccount, found := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.True(t, found)

	// check account
	require.Equal(t, address.String(), approvedAccount.Address)
	require.Equal(t, pubKey, approvedAccount.GetPubKey())
}

func TestHandler_ApproveVendorAccount_TwoApprovalsAreNeeded_FourTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 4 trustees: 1 approval comes from propose => we need 1 more approval

	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.Vendor})
	require.NoError(t, err)

	// account should be in the entity <Proposed Account>, because we haven't enough approvals
	proposedAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)

	// check proposed account
	require.Equal(t, address.String(), proposedAccount.Address)
	require.Equal(t, pubKey, proposedAccount.GetPubKey())

	// approve account by account Trustee2
	approveAddAccount := types.NewMsgApproveAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, approveAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Account>, because we have enough approvals
	approvedAccount, found := setup.Keeper.GetAccountO(setup.Ctx, address)
	require.True(t, found)

	// check account
	require.Equal(t, address.String(), approvedAccount.Address)
	require.Equal(t, pubKey, approvedAccount.GetPubKey())
}

func TestHandler_RejectVendorAccount_ThreeRejectsAreNeeded_FourTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 4 trustees => we need 3 rejects

	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	trustee4 := storeTrustee(setup)

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.Vendor})
	require.NoError(t, err)

	// reject account by account Trustee2
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// reject account by account Trustee3
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Proposed Account>, because we haven't enough rejects
	proposedAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)

	// check proposed account
	require.Equal(t, address.String(), proposedAccount.Address)
	require.Equal(t, pubKey, proposedAccount.GetPubKey())

	// reject account by account Trustee4
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee4, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Rejected Account>, because we have enough rejects
	rejectedAccount, found := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, found)

	// check rejected account
	require.Equal(t, address.String(), rejectedAccount.Address)
	require.Equal(t, pubKey, rejectedAccount.GetPubKey())
}

func TestHandler_RejectVendorAccount_ThreeRejectsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees => we need 4 rejects

	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	trustee4 := storeTrustee(setup)
	trustee5 := storeTrustee(setup)

	// trustee1 propose account
	_, address, pubKey, err := proposeAddAccount(setup, trustee1, types.AccountRoles{types.Vendor})
	require.NoError(t, err)

	// reject account by account Trustee2
	rejectAddAccount := types.NewMsgRejectAddAccount(trustee2, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// reject account by account Trustee3
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee3, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// reject account by account Trustee4
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee4, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Proposed Account>, because we haven't enough rejects
	proposedAccount, found := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.True(t, found)

	// check proposed account
	require.Equal(t, address.String(), proposedAccount.Address)
	require.Equal(t, pubKey, proposedAccount.GetPubKey())

	// reject account by account Trustee5
	rejectAddAccount = types.NewMsgRejectAddAccount(trustee5, address, testconstants.Info2)
	_, err = setup.Handler(setup.Ctx, rejectAddAccount)
	require.NoError(t, err)

	// account should be in the entity <Rejected Account>, because we have enough rejects
	rejectedAccount, found := setup.Keeper.GetRejectedAccount(setup.Ctx, address)
	require.True(t, found)

	// check rejected account
	require.Equal(t, address.String(), rejectedAccount.Address)
	require.Equal(t, pubKey, rejectedAccount.GetPubKey())
}

func storeTrustee(setup TestSetup) sdk.AccAddress {
	return storeAccountWithVendorID(setup, types.Trustee, 0)
}

func storeAccountWithVendorID(setup TestSetup, role types.AccountRole, vendorID int32) sdk.AccAddress {
	_, pubKey, address := testdata.KeyTestPubAddr()
	ba := authtypes.NewBaseAccount(address, pubKey, 0, 0)
	account := types.NewAccount(ba, types.AccountRoles{role}, nil, nil, vendorID)
	account.AccountNumber = setup.Keeper.GetNextAccountNumber(setup.Ctx)
	setup.Keeper.SetAccount(setup.Ctx, account)

	return address
}

func proposeAddAccount(setup TestSetup, signer sdk.AccAddress, roles types.AccountRoles) (*sdk.Result, sdk.AccAddress, cryptotypes.PubKey, error) {
	_, pubKey, address := testdata.KeyTestPubAddr()
	proposeAddAccount, _ := types.NewMsgProposeAddAccount(
		signer,
		address,
		pubKey,
		roles,
		testconstants.VendorID1,
		testconstants.Info,
	)
	// TODO check the err here
	result, err := setup.Handler(setup.Ctx, proposeAddAccount)

	return result, address, pubKey, err
}
