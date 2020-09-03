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

	CodeModelInfoAlreadyExists   sdk.CodeType = 501
	CodeModelInfoDoesNotExist    sdk.CodeType = 502
	CodeVendorProductsDoNotExist sdk.CodeType = 503
)

func ErrModelInfoAlreadyExists(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoAlreadyExists,
		fmt.Sprintf("Model info associated with the vid=%v and pid=%v already exists on the ledger", vid, pid))
}

func ErrModelInfoDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoDoesNotExist,
		fmt.Sprintf("No model info associated with the vid=%v and pid=%v on the ledger", vid, pid))
}

func ErrVendorProductsDoNotExist(vid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorProductsDoNotExist,
		fmt.Sprintf("No models associated with the vid=%v on the ledger", vid))
}
