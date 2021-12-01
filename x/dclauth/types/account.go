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

package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
)

/*
	Account Role
*/

type AccountRole string

const (
	Vendor              AccountRole = "Vendor"
	TestHouse           AccountRole = "TestHouse"
	CertificationCenter AccountRole = "CertificationCenter"
	Trustee             AccountRole = "Trustee"
	NodeAdmin           AccountRole = "NodeAdmin"
)

var Roles = AccountRoles{Vendor, TestHouse, CertificationCenter, Trustee, NodeAdmin}

func (role AccountRole) Validate() sdk.Error {
	for _, r := range Roles {
		if role == r {
			return nil
		}
	}

	return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Account Role: %v. Supported roles: [%v]", role, Roles))
}

/*
	List of Account Roles
*/

type AccountRoles []AccountRole

// Validate checks for errors on the account roles.
func (roles AccountRoles) Validate() sdk.Error {
	for _, role := range roles {
		if err := role.Validate(); err != nil {
			return err
		}
	}

	return nil
}

/*
	Account
*/

// NewAccount creates a new Account object.
func NewAccount(ba *BaseAccount, roles AccountRoles, vendorID uint64) *Account {
	return &Account{
		BaseAccount: ba,
		Roles:       roles,
		VendorID:    vendorID,
	}
}

// Validate checks for errors on the vesting and module account parameters.
func (acc Account) Validate() error {
	err = acc.BaseAccount.Validate()

	if err != nil {
		if acc.Address == nil {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Account: Value: %s. Error: Missing Address", acc.Address))
		}

		if acc.PubKey == nil {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid Account: Value: %s. Error: Missing PubKey", acc.PubKey))
		}

		return err
	}

	if err := acc.Roles.Validate(); err != nil {
		return err
	}

	// If creating an account with Vendor Role, we need to have a associated VendorID
	if acc.HasRole(Vendor) && acc.VendorID <= 0 {
		return ErrMissingVendorIDForVendorAccount()
	}

	return nil
}

func (acc Account) HasRole(targetRole AccountRole) bool {
	for _, role := range acc.Roles {
		if role == targetRole {
			return true
		}
	}

	return false
}

/*
	Pending Account
*/

// NewPendingAccount creates a new PendingAccount object.
func NewPendingAccount(acc *Account, approval sdk.AccAddress) *PendingAccount {
	acc = &PendingAccount{
		Account:   acc,
		Approvals: []sdk.AccAddress{approval.String()},
	}

	return acc

}

//nolint:interfacer
func (acc PendingAccount) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range acc.Approvals {
		if approval.Equals(addrStr) {
			return true
		}
	}

	return false
}

/*
	Pending Account Revocation
*/

// NewPendingAccountRevocation creates a new PendingAccountRevocation object.
func NewPendingAccountRevocation(address sdk.AccAddress, approval sdk.AccAddress) PendingAccountRevocation {
	return PendingAccountRevocation{
		Address:   address,
		Approvals: []sdk.AccAddress{approval},
	}
}

// String implements fmt.Stringer.
func (revoc PendingAccountRevocation) String() string {
	bytes, err := json.Marshal(revoc)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate checks for errors on the vesting and module account parameters.
func (revoc PendingAccountRevocation) Validate() sdk.Error {
	if revoc.Address == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Pending Account Revocation: Value: %s. Error: Missing Address", revoc.Address))
	}

	return nil
}

//nolint:interfacer
func (revoc PendingAccountRevocation) HasApprovalFrom(address sdk.AccAddress) bool {
	for _, approval := range revoc.Approvals {
		if approval.Equals(address) {
			return true
		}
	}

	return false
}

// GenesisAccountsIterator implements genesis account iteration.
type GenesisAccountsIterator struct{}

// IterateGenesisAccounts iterates over all the genesis accounts found in
// appGenesis and invokes a callback on each genesis account. If any call
// returns true, iteration stops.
func (GenesisAccountsIterator) IterateGenesisAccounts(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, cb func(exported.GenesisAccount) (stop bool),
) {
	for _, account := range GetGenesisStateFromAppState(cdc, appState).AccountList {
		if cb(account) {
			break
		}
	}
}
