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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

/*
	Account Role
*/

type AccountRole string

const (
	Vendor                AccountRole = "Vendor"
	TestHouse             AccountRole = "TestHouse"
	ZBCertificationCenter AccountRole = "ZBCertificationCenter"
	Trustee               AccountRole = "Trustee"
	NodeAdmin             AccountRole = "NodeAdmin"
)

var Roles = AccountRoles{Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin}

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
	Pending Account
*/
type PendingAccount struct {
	Address   sdk.AccAddress   `json:"address"`
	PubKey    crypto.PubKey    `json:"public_key"`
	Roles     AccountRoles     `json:"roles"`
	VendorId  uint16           `json:"vendorId"`
	Approvals []sdk.AccAddress `json:"approvals"`
}

// NewPendingAccount creates a new PendingAccount object.
func NewPendingAccount(address sdk.AccAddress, pubKey crypto.PubKey,
	roles AccountRoles, vendorId uint16, approval sdk.AccAddress) PendingAccount {
	return PendingAccount{
		Address:   address,
		PubKey:    pubKey,
		Roles:     roles,
		VendorId:  vendorId,
		Approvals: []sdk.AccAddress{approval},
	}
}

// String implements fmt.Stringer.
func (pendAcc PendingAccount) String() string {
	bytes, err := json.Marshal(pendAcc)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate checks for errors on the vesting and module account parameters.
func (pendAcc PendingAccount) Validate() sdk.Error {
	if pendAcc.Address == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Pending Account: Value: %s. Error: Missing Address", pendAcc.Address))
	}

	if pendAcc.PubKey == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Pending Account: Value: %s. Error: Missing PubKey", pendAcc.PubKey))
	}

	if err := pendAcc.Roles.Validate(); err != nil {
		return err
	}

	return nil
}

//nolint:interfacer
func (pendAcc PendingAccount) HasApprovalFrom(address sdk.AccAddress) bool {
	for _, approval := range pendAcc.Approvals {
		if approval.Equals(address) {
			return true
		}
	}

	return false
}

/*
	Account
*/
type Account struct {
	Address       sdk.AccAddress `json:"address"`
	PubKey        crypto.PubKey  `json:"public_key"`
	AccountNumber uint64         `json:"account_number"`
	Sequence      uint64         `json:"sequence"`
	Roles         AccountRoles   `json:"roles"`
	VendorId      uint16         `json:"vendorId"`
}

// NewAccount creates a new Account object.
func NewAccount(address sdk.AccAddress, pubKey crypto.PubKey, roles AccountRoles, vendorId uint16) Account {
	return Account{
		Address:  address,
		PubKey:   pubKey,
		Roles:    roles,
		VendorId: vendorId,
	}
}

// String implements fmt.Stringer.
func (acc Account) String() string {
	bytes, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate checks for errors on the vesting and module account parameters.
func (acc Account) Validate() error {
	if acc.Address == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Account: Value: %s. Error: Missing Address", acc.Address))
	}

	if acc.PubKey == nil {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid Account: Value: %s. Error: Missing PubKey", acc.PubKey))
	}

	if err := acc.Roles.Validate(); err != nil {
		return err
	}

	// If creating an account with Vendor Role, we need to have a associated VendorId
	if acc.HasRole(Vendor) && acc.VendorId <= 0 {
		return ErrMissingVendorIdForVendorAccount()
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

func (acc Account) GetAddress() sdk.AccAddress {
	return acc.Address
}

// SetAddress - Implements sdk.Account.
func (acc *Account) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return sdk.ErrInvalidAddress("Cannot override Account address")
	}

	acc.Address = addr

	return nil
}

// GetPubKey - Implements sdk.Account.
func (acc Account) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// SetPubKey - Implements sdk.Account.
func (acc *Account) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey

	return nil
}

// GetCoins - Implements sdk.Account.
func (acc *Account) GetCoins() sdk.Coins {
	return nil
}

// SetCoins - Implements sdk.Account.
func (acc *Account) SetCoins(coins sdk.Coins) error {
	return nil
}

// GetAccountNumber - Implements Account.
func (acc *Account) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

// SetAccountNumber - Implements Account.
func (acc *Account) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber

	return nil
}

// GetSequence - Implements sdk.Account.
func (acc *Account) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence - Implements sdk.Account.
func (acc *Account) SetSequence(seq uint64) error {
	acc.Sequence = seq

	return nil
}

// SpendableCoins returns the total set of spendable coins. For a base account,
// this is simply the base coins.
func (acc *Account) SpendableCoins(_ time.Time) sdk.Coins {
	return nil
}

/*
	Pending Account Revocation
*/
type PendingAccountRevocation struct {
	Address   sdk.AccAddress   `json:"address"`
	Approvals []sdk.AccAddress `json:"approvals"`
}

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
