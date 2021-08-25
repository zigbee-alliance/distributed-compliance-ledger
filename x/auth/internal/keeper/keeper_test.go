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
package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func TestKeeper_AccountGetSet(t *testing.T) {
	setup := Setup()

	// check if account present
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))

	// no account before its created
	require.Panics(t, func() {
		setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	})

	// store account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee}, 0)
	setup.Keeper.SetAccount(setup.Ctx, account)

	// check if account present
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))

	// get account
	receivedAccount := setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	require.Equal(t, account.Address, receivedAccount.Address)
	require.Equal(t, account.PubKey, receivedAccount.PubKey)
	require.Equal(t, account.Roles, receivedAccount.Roles)
	require.Equal(t, account.Sequence, receivedAccount.Sequence)
	require.Equal(t, account.AccountNumber, receivedAccount.AccountNumber)

	// check if account has role
	require.True(t, setup.Keeper.HasRole(setup.Ctx, testconstants.Address1, types.Trustee))

	// get all accounts
	accounts := setup.Keeper.GetAllAccounts(setup.Ctx)
	require.Equal(t, 1, len(accounts))
	require.Equal(t, account.Address, accounts[0].Address)

	// count accounts with role
	require.Equal(t, 1, setup.Keeper.CountAccountsWithRole(setup.Ctx, types.Trustee))
	require.Equal(t, 0, setup.Keeper.CountAccountsWithRole(setup.Ctx, types.Vendor))

	// delete account
	setup.Keeper.DeleteAccount(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_PendingAccountGetSet(t *testing.T) {
	setup := Setup()

	// check if pending account present
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))

	// no pending account before its created
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	})

	// store pending account
	pendAcc := types.NewPendingAccount(
		testconstants.Address1,
		testconstants.PubKey1,
		types.AccountRoles{types.Trustee},
		testconstants.VendorId1,
		testconstants.Address2,
	)

	setup.Keeper.SetPendingAccount(setup.Ctx, pendAcc)

	// check if pending account present
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))

	// get pending account
	receivedPendAcc := setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	require.Equal(t, pendAcc.Address, receivedPendAcc.Address)
	require.Equal(t, pendAcc.PubKey, receivedPendAcc.PubKey)
	require.Equal(t, pendAcc.Roles, receivedPendAcc.Roles)
	require.Equal(t, pendAcc.Approvals, receivedPendAcc.Approvals)

	// delete pending account
	setup.Keeper.DeletePendingAccount(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_PendingAccountRevocationGetSet(t *testing.T) {
	setup := Setup()

	// check if pending account revocation present
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))

	// no pending account revocation before its created
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	})

	// store pending account revocation
	revocation := types.NewPendingAccountRevocation(
		testconstants.Address1,
		testconstants.Address2,
	)

	setup.Keeper.SetPendingAccountRevocation(setup.Ctx, revocation)

	// check if pending account revocation present
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))

	// get pending account revocation
	receivedRevocation := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	require.Equal(t, revocation.Address, receivedRevocation.Address)
	require.Equal(t, revocation.Approvals, receivedRevocation.Approvals)

	// delete pending account revocation
	setup.Keeper.DeletePendingAccountRevocation(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_AccountNumber(t *testing.T) {
	setup := Setup()

	for i := uint64(0); i < 5; i++ {
		require.Equal(t, i, setup.Keeper.GetNextAccountNumber(setup.Ctx))
	}
}
