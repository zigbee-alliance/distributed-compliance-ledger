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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeAccountAlreadyExists                  sdk.CodeType = 101
	CodeAccountDoesNotExist                   sdk.CodeType = 102
	CodePendingAccountAlreadyExists           sdk.CodeType = 103
	CodePendingAccountDoesNotExist            sdk.CodeType = 104
	CodePendingAccountRevocationAlreadyExists sdk.CodeType = 105
	CodePendingAccountRevocationDoesNotExist  sdk.CodeType = 106
	CodeMissingVendorIdForVendorAccount       sdk.CodeType = 107
)

func ErrAccountAlreadyExists(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExists(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExists(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}

func ErrMissingVendorIdForVendorAccount() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeMissingVendorIdForVendorAccount,
		fmt.Sprintf("No Vendor ID is provided in the Vendor Role for the new account"))
}
