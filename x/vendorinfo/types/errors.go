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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vendorinfo module sentinel errors.
var (
	DefaultCodespace string = ModuleName

	CodeVendorDoesNotExist              = sdkerrors.Register(ModuleName, 701, "Code vendor does not exist")
	CodeMissingVendorIDForVendorAccount = sdkerrors.Register(ModuleName, 702, "Code missing vendor id for vendor account")
	CodeVendorInfoAlreadyExists         = sdkerrors.Register(ModuleName, 703, "Code vendorinfo already exists")
)
