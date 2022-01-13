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

package ante

import (
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// TODO issue 99: add GetParams and GetModuleAddress stubs in dclauth.AccountKeeper.
type AccountKeeper authante.AccountKeeper

// which is
/*
// AccountKeeper defines the contract needed for AccountKeeper related APIs.
// Interface provides support to use non-sdk AccountKeeper for AnteHandler's decorators.
type AccountKeeper interface {
	GetParams(ctx sdk.Context) (params types.Params)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) dclauthtypes.Account
	SetAccount(ctx sdk.Context, acc dclauthtypes.Account)
	GetModuleAddress(moduleName string) sdk.AccAddress
}
*/
