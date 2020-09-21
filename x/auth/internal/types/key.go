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

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of the module.
	ModuleName = "auth"

	// StoreKey to be used when creating the KVStore.
	StoreKey = "acc" // it differs from ModuleName to be compatible with cosmos transaction builder and handler.
)

var (
	PendingAccountPrefix           = []byte{0x01} // prefix for each key to a pending account
	AccountPrefix                  = []byte{0x02} // prefix for each key to an account
	PendingAccountRevocationPrefix = []byte{0x03} // prefix for each key to a pending account revocation

	PendingAccountsTotalKey           = []byte{0x11} // key for a total number of pending accounts
	AccountsTotalKey                  = []byte{0x12} // key for a total number of accounts
	PendingAccountRevocationsTotalKey = []byte{0x13} // key for a total number of pending account revocations

	// Only increases in contrast to total.
	AccountNumberCounterKey = []byte("globalAccountNumber") // key for account number counter
)

// nolint: interfacer
// Key builder for Pending Account.
func GetPendingAccountKey(addr sdk.AccAddress) []byte {
	// We are using `String` instead of `Bytes` in order to provide
	// more comfortable querying of ranges by prefixes
	return append(PendingAccountPrefix, addr.String()...)
}

// nolint: interfacer
// Key builder for Account.
func GetAccountKey(addr sdk.AccAddress) []byte {
	return append(AccountPrefix, addr.String()...)
}

// nolint: interfacer
// Key builder for Pending Account Revocation.
func GetPendingAccountRevocationKey(addr sdk.AccAddress) []byte {
	return append(PendingAccountRevocationPrefix, addr.String()...)
}
