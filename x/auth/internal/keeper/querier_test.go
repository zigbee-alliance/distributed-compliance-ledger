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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func TestQuerier_QueryAccount(t *testing.T) {
	setup := Setup()

	// store account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee}, 0)
	setup.Keeper.SetAccount(setup.Ctx, account)

	// query account
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAccount},
		abci.RequestQuery{Data: queryAccountParams(setup, account.Address)},
	)

	var receivedAccount types.Account
	_ = setup.Cdc.UnmarshalJSON(result, &receivedAccount)

	// check
	require.Equal(t, account.Address, receivedAccount.Address)
	require.Equal(t, account.PubKey, receivedAccount.PubKey)
	require.Equal(t, account.Roles, receivedAccount.Roles)
}

func TestQuerier_QueryAccount_ForNotFound(t *testing.T) {
	setup := Setup()

	// query pending certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryAccount},
		abci.RequestQuery{Data: queryAccountParams(setup, testconstants.Address1)},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeAccountDoesNotExist, err.Code())
}

func TestQuerier_QueryAllAccounts(t *testing.T) {
	setup := Setup()

	// store active account
	account1 := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee}, 0)
	setup.Keeper.SetAccount(setup.Ctx, account1)

	// store second active account
	account2 := types.NewAccount(testconstants.Address2, testconstants.PubKey2, types.AccountRoles{types.Vendor}, testconstants.VendorId2)
	setup.Keeper.SetAccount(setup.Ctx, account2)

	// store pending account
	pendAcc := types.NewPendingAccount(
		testconstants.Address3,
		testconstants.PubKey3,
		types.AccountRoles{types.Vendor},
		testconstants.VendorId3,
		testconstants.Address1,
	)
	setup.Keeper.SetPendingAccount(setup.Ctx, pendAcc)

	// query active accounts
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllAccounts},
		abci.RequestQuery{Data: queryListEmptyQueryParams(setup)},
	)

	var listAccounts types.ListAccounts
	_ = setup.Cdc.UnmarshalJSON(result, &listAccounts)

	// check
	require.Equal(t, 2, len(listAccounts.Items))
	require.Equal(t, account1, listAccounts.Items[0])
	require.Equal(t, account2, listAccounts.Items[1])
}

func TestQuerier_QueryAllPendingAccounts(t *testing.T) {
	setup := Setup()

	// store active account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee}, 0)
	setup.Keeper.SetAccount(setup.Ctx, account)

	// store pending account
	pendAcc1 := types.NewPendingAccount(
		testconstants.Address2,
		testconstants.PubKey2,
		types.AccountRoles{types.Trustee},
		testconstants.VendorId2,
		testconstants.Address1,
	)
	setup.Keeper.SetPendingAccount(setup.Ctx, pendAcc1)

	// store second pending account
	pendAcc2 := types.NewPendingAccount(
		testconstants.Address3,
		testconstants.PubKey3,
		types.AccountRoles{types.Vendor},
		testconstants.VendorId3,
		testconstants.Address1,
	)
	setup.Keeper.SetPendingAccount(setup.Ctx, pendAcc2)

	// query pending accounts
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllPendingAccounts},
		abci.RequestQuery{Data: queryListEmptyQueryParams(setup)},
	)

	var listPendingAccounts types.ListPendingAccounts
	_ = setup.Cdc.UnmarshalJSON(result, &listPendingAccounts)

	// check
	require.Equal(t, 2, len(listPendingAccounts.Items))
	require.Equal(t, pendAcc1, listPendingAccounts.Items[0])
	require.Equal(t, pendAcc2, listPendingAccounts.Items[1])
}

func TestQuerier_QueryAllPendingAccountRevocations(t *testing.T) {
	setup := Setup()

	// store active account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee}, 0)
	setup.Keeper.SetAccount(setup.Ctx, account)

	// store pending account revocation
	revocation1 := types.NewPendingAccountRevocation(
		testconstants.Address2,
		testconstants.Address1,
	)
	setup.Keeper.SetPendingAccountRevocation(setup.Ctx, revocation1)

	// store second pending account revocation
	revocation2 := types.NewPendingAccountRevocation(
		testconstants.Address3,
		testconstants.Address1,
	)
	setup.Keeper.SetPendingAccountRevocation(setup.Ctx, revocation2)

	// query pending account revocations
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllPendingAccountRevocations},
		abci.RequestQuery{Data: queryListEmptyQueryParams(setup)},
	)

	var listPendingAccountRevocations types.ListPendingAccountRevocations
	_ = setup.Cdc.UnmarshalJSON(result, &listPendingAccountRevocations)

	// check
	require.Equal(t, 2, len(listPendingAccountRevocations.Items))
	require.Equal(t, revocation1, listPendingAccountRevocations.Items[0])
	require.Equal(t, revocation2, listPendingAccountRevocations.Items[1])
}

func queryAccountParams(setup TestSetup, address sdk.AccAddress) []byte {
	params := types.NewQueryAccountParams(address)

	return setup.Cdc.MustMarshalJSON(params)
}

func queryListEmptyQueryParams(setup TestSetup) []byte {
	paginationParams := pagination.NewPaginationParams(0, 0)

	return setup.Cdc.MustMarshalJSON(paginationParams)
}
