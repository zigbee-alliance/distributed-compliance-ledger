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
package auth

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func TestHandler_CreateAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup()

	countTrustees := 2

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, AccountApprovalsCount(setup.Ctx, setup.Keeper))

		// propose account
		result, address, pubkey := proposeAddAccount(setup, trustee)
		require.Equal(t, sdk.CodeOK, result.Code)

		// ensure active account created
		account := setup.Keeper.GetAccount(setup.Ctx, address)
		require.Equal(t, address, account.Address)
		require.Equal(t, pubkey, account.PubKey)

		// ensure no pending account created
		require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
	}
}

func TestHandler_CreateAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	result, address, pubkey := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account created
	pendingAccount := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.Equal(t, address, pendingAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.Address)
	require.Equal(t, pubkey, account.PubKey)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	result, address, pubkey := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account created
	pendingAccount := setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.Equal(t, address, pendingAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure second approval added to pending account
	pendingAccount = setup.Keeper.GetPendingAccount(setup.Ctx, address)
	require.Equal(t, address, pendingAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1, trustee2}, pendingAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(address, trustee3)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.Address)
	require.Equal(t, pubkey, account.PubKey)

	// ensure pending account removed
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeAddAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer := storeAccount(setup, role, testconstants.VendorId1)

		// propose new account
		result, _, _ := proposeAddAccount(setup, signer)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ProposeAddAccount_ForExistingActiveAccount(t *testing.T) {
	setup := Setup()

	// store 2 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)

	// propose account
	result, address, pubkey := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure active account created
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// propose existing active account
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		sdk.MustBech32ifyAccPub(pubkey),
		types.AccountRoles{types.Vendor},
		testconstants.VendorId1,
		trustee2,
	)
	result = setup.Handler(setup.Ctx, proposeAddAccount)
	require.Equal(t, types.CodeAccountAlreadyExists, result.Code)
}

func TestHandler_ProposeAddAccount_ForExistingPendingAccount(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// trustee1 proposes account
	result, address, pubkey := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// trustee2 proposes the same account
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		sdk.MustBech32ifyAccPub(pubkey),
		types.AccountRoles{types.Vendor},
		testconstants.VendorId1,
		trustee2,
	)
	result = setup.Handler(setup.Ctx, proposeAddAccount)
	require.Equal(t, types.CodePendingAccountAlreadyExists, result.Code)
}

func TestHandler_ApproveAddAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// propose account
	result, address, _ := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer := storeAccount(setup, role, testconstants.VendorId1)

		// try to approve account
		approveAddAccount := types.NewMsgApproveAddAccount(address, signer)
		result := setup.Handler(setup.Ctx, approveAddAccount)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveAddAccount_ForExistingActiveAccount(t *testing.T) {
	setup := Setup()

	// store 2 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)

	// propose account
	result, address, _ := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure active account created
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// try to approve active account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, types.CodePendingAccountDoesNotExist, result.Code)
}

func TestHandler_ApproveAddAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup()

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve unknown account
	approveAddAccount := types.NewMsgApproveAddAccount(testconstants.Address1, trustee)
	result := setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, types.CodePendingAccountDoesNotExist, result.Code)
}

func TestHandler_ApproveAddAccount_ForDuplicateApproval(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// propose account
	result, address, _ := proposeAddAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account created
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, address))

	// the same trustee tries to approve the account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee1)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_RevokeAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup()

	countTrustees := 2

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// store account
		address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

		// ensure 1 trustee revocation approval is needed
		require.Equal(t, 1, AccountApprovalsCount(setup.Ctx, setup.Keeper))

		// propose to revoke account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee)
		result := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.Equal(t, sdk.CodeOK, result.Code)

		// ensure active account removed
		require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

		// ensure no pending account revocation created
		require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
	}
}

func TestHandler_RevokeAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// ensure 2 trustee revocation approvals are needed
	require.Equal(t, 2, AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee1)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account revocation created
	revocation := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.Equal(t, address, revocation.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
}

func TestHandler_RevokeAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 4 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// ensure 3 trustee revocation approvals are needed
	require.Equal(t, 3, AccountApprovalsCount(setup.Ctx, setup.Keeper))

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee1)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account revocation created
	revocation := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.Equal(t, address, revocation.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure second approval added to pending account revocation
	revocation = setup.Keeper.GetPendingAccountRevocation(setup.Ctx, address)
	require.Equal(t, address, revocation.Address)
	require.Equal(t, []sdk.AccAddress{trustee1, trustee2}, revocation.Approvals)

	// ensure active account still exists
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account revocation
	approveRevokeAccount = types.NewMsgApproveRevokeAccount(address, trustee3)
	result = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be removed
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// ensure pending account revocation removed
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))
}

func TestHandler_ProposeRevokeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer := storeAccount(setup, role, testconstants.VendorId1)

		// propose new account
		proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, signer)
		result := setup.Handler(setup.Ctx, proposeRevokeAccount)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ProposeRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup()

	// store 1 trustee
	trustee := storeTrustee(setup)

	// propose to revoke unknown account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(testconstants.Address1, trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, types.CodeAccountDoesNotExist, result.Code)
}

func TestHandler_ProposeRevokeAccount_ForExistingPendingAccountRevocation(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee1)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// trustee2 proposes to revoke the same account
	proposeRevokeAccount = types.NewMsgProposeRevokeAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, types.CodePendingAccountRevocationAlreadyExists, result.Code)
}

func TestHandler_ApproveRevokeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// trustee1 proposes to revoke account
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee1)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer := storeAccount(setup, role, testconstants.VendorId1)

		// try to approve account
		approveRevokeAccount := types.NewMsgApproveRevokeAccount(address, signer)
		result := setup.Handler(setup.Ctx, approveRevokeAccount)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveRevokeAccount_ForAbsentPendingAccountRevocation(t *testing.T) {
	setup := Setup()

	// store 1 trustee
	trustee := storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// approve absent revocation of active account
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(address, trustee)
	result := setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, types.CodePendingAccountRevocationDoesNotExist, result.Code)
}

func TestHandler_ApproveRevokeAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup()

	// store 1 trustee
	trustee := storeTrustee(setup)

	// approve absent revocation of unknown account
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(testconstants.Address1, trustee)
	result := setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, types.CodePendingAccountRevocationDoesNotExist, result.Code)
}

func TestHandler_ApproveRevokeAccount_ForDuplicateApproval(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// store account
	address := storeAccount(setup, types.Vendor, testconstants.VendorId1)

	// propose account revocation
	proposeRevokeAccount := types.NewMsgProposeRevokeAccount(address, trustee1)
	result := setup.Handler(setup.Ctx, proposeRevokeAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure pending account revocation created
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, address))

	// the same trustee tries to approve the account revocation
	approveRevokeAccount := types.NewMsgApproveRevokeAccount(address, trustee1)
	result = setup.Handler(setup.Ctx, approveRevokeAccount)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func storeTrustee(setup TestSetup) sdk.AccAddress {
	return storeAccount(setup, types.Trustee, 0)
}

func storeAccount(setup TestSetup, role types.AccountRole, vendorId uint16) sdk.AccAddress {
	address, pubkey, _ := testconstants.TestAddress()
	account := types.NewAccount(address, pubkey, types.AccountRoles{role}, vendorId)
	account.AccountNumber = setup.Keeper.GetNextAccountNumber(setup.Ctx)
	setup.Keeper.SetAccount(setup.Ctx, account)

	return address
}

func proposeAddAccount(setup TestSetup, signer sdk.AccAddress) (sdk.Result, sdk.AccAddress, crypto.PubKey) {
	address, pubkey, pubkeyStr := testconstants.TestAddress()
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		pubkeyStr,
		types.AccountRoles{types.Vendor},
		testconstants.VendorId1,
		signer,
	)
	result := setup.Handler(setup.Ctx, proposeAddAccount)

	return result, address, pubkey
}
