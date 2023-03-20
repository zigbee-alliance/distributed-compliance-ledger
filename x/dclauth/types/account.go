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
	VendorAdmin         AccountRole = "VendorAdmin"
)

var Roles = AccountRoles{Vendor, CertificationCenter, Trustee, NodeAdmin, VendorAdmin}

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
	GetApprovals() []*Grant
	GetRejects() []*Grant
}

// NewAccount creates a new Account object.
func NewAccount(ba *authtypes.BaseAccount, roles AccountRoles, approvals []*Grant, rejects []*Grant, vendorID int32) *Account {
	return &Account{
		BaseAccount: ba,
		Roles:       roles,
		Approvals:   approvals,
		Rejects:     rejects,
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

func (acc Account) GetApprovals() []*Grant {
	return acc.Approvals
}

func (acc Account) GetRejects() []*Grant {
	return acc.Rejects
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

func (acc Account) HasOnlyVendorRole(targetRole AccountRole) bool {
	if len(acc.Roles) == 1 && acc.Roles[0] == targetRole {
		return true
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
func NewPendingAccount(acc *Account, approval sdk.AccAddress, info string, time int64) *PendingAccount {
	pendingAccount := &PendingAccount{
		Account: acc,
	}

	pendingAccount.Approvals = []*Grant{
		{
			Address: approval.String(),
			Time:    time,
			Info:    info,
		},
	}

	return pendingAccount
}

// NewRevokedAccount creates a new RevokedAccount object.
func NewRevokedAccount(acc *Account, approvals []*Grant) *RevokedAccount {
	revokedAccount := &RevokedAccount{
		Account: acc,
	}

	revokedAccount.RevokeApprovals = approvals

	return revokedAccount
}

func (acc PendingAccount) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range acc.Approvals {
		if approval.Address == addrStr {
			return true
		}
	}

	return false
}

func (acc PendingAccount) HasRejectApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, rejectApproval := range acc.Rejects {
		if rejectApproval.Address == addrStr {
			return true
		}
	}

	return false
}

/*
	Pending Account Revocation
*/

// NewPendingAccountRevocation creates a new PendingAccountRevocation object.
func NewPendingAccountRevocation(address sdk.AccAddress,
	info string, time int64, approval sdk.AccAddress,
) PendingAccountRevocation {
	pendingAccountRevocation := PendingAccountRevocation{
		Address: address.String(),
	}
	pendingAccountRevocation.Approvals = []*Grant{
		{
			Address: approval.String(),
			Time:    time,
			Info:    info,
		},
	}

	return pendingAccountRevocation
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

func (revoc PendingAccountRevocation) HasRevocationFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approvals := range revoc.Approvals {
		if approvals.Address == addrStr {
			return true
		}
	}

	return false
}
