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

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/dclauth module sentinel errors.
var (
	AccountAlreadyExists                  = errors.Register(ModuleName, 101, "account already exists")
	AccountDoesNotExist                   = errors.Register(ModuleName, 102, "account not found")
	PendingAccountAlreadyExists           = errors.Register(ModuleName, 103, "pending account already exists")
	PendingAccountDoesNotExist            = errors.Register(ModuleName, 104, "pending account not found")
	PendingAccountRevocationAlreadyExists = errors.Register(ModuleName, 105, "pending account revocation already exists")
	PendingAccountRevocationDoesNotExist  = errors.Register(ModuleName, 106, "pending account revocation not found")
	MissingVendorIDForVendorAccount       = errors.Register(ModuleName, 107, "no Vendor ID provided")
	MissingRoles                          = errors.Register(ModuleName, 108, "no roles provided")
)

func ErrAccountAlreadyExists(address interface{}) error {
	return errors.Wrapf(AccountAlreadyExists,
		"Account associated with the address=%v already exists on the ledger", address)
}

func ErrAccountDoesNotExist(address interface{}) error {
	return errors.Wrapf(AccountDoesNotExist,
		"No account associated with the address=%v on the ledger", address)
}

func ErrPendingAccountAlreadyExists(address interface{}) error {
	return errors.Wrapf(PendingAccountAlreadyExists,
		"Pending account associated with the address=%v already exists on the ledger", address)
}

func ErrPendingAccountDoesNotExist(address interface{}) error {
	return errors.Wrapf(PendingAccountDoesNotExist,
		"No pending account associated with the address=%v on the ledger", address)
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) error {
	return errors.Wrapf(PendingAccountRevocationAlreadyExists,
		"Pending account revocation associated with the address=%v already exists on the ledger", address)
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) error {
	return errors.Wrapf(PendingAccountRevocationDoesNotExist,
		"No pending account revocation associated with the address=%v on the ledger", address)
}

func ErrMissingVendorIDForVendorAccount() error {
	return errors.Wrapf(MissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrMissingRoles() error {
	return errors.Wrapf(MissingRoles,
		"No roles provided")
}
