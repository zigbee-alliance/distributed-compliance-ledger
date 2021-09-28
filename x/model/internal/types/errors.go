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

	CodeModelAlreadyExists       sdk.CodeType = 501
	CodeModelDoesNotExist        sdk.CodeType = 502
	CodeOtaURLCannotBeSet        sdk.CodeType = 503
	CodeVendorProductsDoNotExist sdk.CodeType = 504
)

func ErrModelAlreadyExists(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelAlreadyExists,
		fmt.Sprintf("Model info associated with vid=%v and pid=%v already exists on the ledger", vid, pid))
}

func ErrModelDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelDoesNotExist,
		fmt.Sprintf("No model info associated with vid=%v and pid=%v exist on the ledger", vid, pid))
}

func ErrOtaURLCannotBeSet(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeOtaURLCannotBeSet,
		fmt.Sprintf("OTA URL cannot be set for model info associated with vid=%v and pid=%v "+
			"because OTA was not set for this model info initially", vid, pid))
}

func ErrVendorProductsDoNotExist(vid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorProductsDoNotExist,
		fmt.Sprintf("No vendor products associated with vid=%v exist on the ledger", vid))
}
