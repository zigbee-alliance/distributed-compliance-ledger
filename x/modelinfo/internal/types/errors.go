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
	CodeOtaURLCannotBeSet        sdk.CodeType = 503
	CodeVendorProductsDoNotExist sdk.CodeType = 504
)

func ErrModelInfoAlreadyExists(vid interface{}, pid interface{}, softwareVersion interface{}, hardwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoAlreadyExists,
		fmt.Sprintf("Model info associated with vid=%v, pid=%v,softwareVersion=%v and hardwareVersion=%v  already exists on the ledger",
			vid, pid, softwareVersion, hardwareVersion))
}

func ErrModelInfoDoesNotExist(vid interface{}, pid interface{}, softwareVersion interface{}, hardwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoDoesNotExist,
		fmt.Sprintf("No Model info associated with vid=%v, pid=%v,softwareVersion=%v and hardwareVersion=%v  already exists on the ledger",
			vid, pid, softwareVersion, hardwareVersion))
}

func ErrOtaURLCannotBeSet(vid interface{}, pid interface{}, softwareVersion interface{}, hardwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeOtaURLCannotBeSet,
		fmt.Sprintf("OTA URL cannot be set for model info associated with vid=%v, pid=%v, softwareVersion=%v and hardwareVersion=%v "+
			"because OTA was not set for this model info initially", vid, pid, softwareVersion, hardwareVersion))
}

func ErrVendorProductsDoNotExist(vid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorProductsDoNotExist,
		fmt.Sprintf("No vendor products associated with vid=%v exist on the ledger", vid))
}
