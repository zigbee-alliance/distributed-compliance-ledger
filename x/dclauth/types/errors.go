// Copyright 2022 DSR Corporation
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
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dclauth module sentinel errors.
var (
	AccountAlreadyExists                  = sdkerrors.Register(ModuleName, 101, "account already exists")
	AccountDoesNotExist                   = sdkerrors.Register(ModuleName, 102, "account not found")
	PendingAccountAlreadyExists           = sdkerrors.Register(ModuleName, 103, "pending account already exists")
	PendingAccountDoesNotExist            = sdkerrors.Register(ModuleName, 104, "pending account not found")
	PendingAccountRevocationAlreadyExists = sdkerrors.Register(ModuleName, 105, "pending account revocation already exists")
	PendingAccountRevocationDoesNotExist  = sdkerrors.Register(ModuleName, 106, "pending account revocation not found")
	MissingVendorIDForVendorAccount       = sdkerrors.Register(ModuleName, 107, "no Vendor ID provided")
	MissingRoles                          = sdkerrors.Register(ModuleName, 108, "no roles provided")
)

func ErrAccountAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(AccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(AccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(PendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(PendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) error {
	return sdkerrors.Wrapf(PendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) error {
	return sdkerrors.Wrapf(PendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}

func ErrMissingVendorIDForVendorAccount() error {
	return sdkerrors.Wrapf(MissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrMissingRoles() error {
	return sdkerrors.Wrapf(MissingRoles,
		"No roles provided")
}
