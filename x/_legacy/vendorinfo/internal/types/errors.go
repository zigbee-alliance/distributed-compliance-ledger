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
	Codespace sdk.CodespaceType = ModuleName

	CodeVendorDoesNotExist              sdk.CodeType = 701
	CodeMissingVendorIDForVendorAccount sdk.CodeType = 702
	CodeVendorInfoAlreadyExists         sdk.CodeType = 703
)

func ErrVendorInfoDoesNotExist(vendorID uint16) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorDoesNotExist,
		fmt.Sprintf("Vendor Account with VendorID %v does not exist on the ledger", vendorID))
}

func ErrMissingVendorIDForVendorAccount() sdk.Error {
	return sdk.NewError(Codespace, CodeMissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrVendorInfoAlreadyExists(vendorID interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorInfoAlreadyExists,
		fmt.Sprintf("Vendor info associated with VendorID=%v already exists on the ledger", vendorID))
}
