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

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

/*
	Account Role
*/

type AccountRole string

const (
	Vendor              AccountRole = "Vendor"
	CertificationCenter AccountRole = "CertificationCenter"
	Trustee             AccountRole = "Trustee"
	NodeAdmin           AccountRole = "NodeAdmin"
)

var Roles = AccountRoles{Vendor, CertificationCenter, Trustee, NodeAdmin}

func (role AccountRole) Validate() error {
	for _, r := range Roles {
		if role == r {
			return nil
		}
	}

	return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid Account Role: %v. Supported roles: [%v]", role, Roles)
}

/*
	List of Account Roles
*/

type AccountRoles []AccountRole

/*
	Account
*/

type DCLAccountI interface {
	authtypes.AccountI

	GetRoles() []AccountRole
	GetVendorID() int32
}

// NewAccount creates a new Account object.
func NewAccount(ba *authtypes.BaseAccount, roles AccountRoles, vendorID int32) *Account {
	return &Account{
		BaseAccount: ba,
		Roles:       roles,
		VendorID:    vendorID,
	}
}

// Validate checks for errors on the vesting and module account parameters.
func (acc Account) Validate() error {
	err := acc.BaseAccount.Validate()
	if err != nil {
		return err
	}

	for _, role := range acc.Roles {
		if err := role.Validate(); err != nil {
			return err
		}
	}

	// If creating an account with Vendor Role, we need to have a associated VendorID
	if acc.HasRole(Vendor) && acc.VendorID <= 0 {
		return ErrMissingVendorIDForVendorAccount()
	}

	return nil
}

func (acc Account) GetRoles() []AccountRole {
	return acc.Roles
}

func (acc Account) GetVendorID() int32 {
	return acc.VendorID
}

func (acc Account) HasRole(targetRole AccountRole) bool {
	for _, role := range acc.Roles {
		if role == targetRole {
			return true
		}
	}

	return false
}

func (acc Account) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

/*
	Pending Account
*/

// NewPendingAccount creates a new PendingAccount object.
func NewPendingAccount(acc *Account, approval sdk.AccAddress) *PendingAccount {
	return &PendingAccount{
		Account:   acc,
		Approvals: []string{approval.String()},
	}
}

//nolint:interfacer
func (acc PendingAccount) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range acc.Approvals {
		if approval == addrStr {
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
		Address:   address.String(),
		Approvals: []string{approval.String()},
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
func (revoc PendingAccountRevocation) Validate() error {
	if revoc.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"Invalid Pending Account Revocation: Value: %s. Error: Missing Address", revoc.Address,
		)
	}

	return nil
}

//nolint:interfacer
func (revoc PendingAccountRevocation) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range revoc.Approvals {
		if approval == addrStr {
			return true
		}
	}

	return false
}
