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

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	Codespace string = ModuleName

	CodeValidatorAlreadyExist uint32 = 601
	CodeValidatorDoesNotExist uint32 = 602
	CodePoolIsFull            uint32 = 603
	CodeAccountAlreadyHasNode uint32 = 604
)

// TODO issue 99: either remove completely or replace cosmos's staking errors

/*
func ErrValidatorExists(owner sdk.ConsAddress) *sdkerrors.Error {
	return sdkerrors.Register(Codespace, CodeValidatorAlreadyExist,
		fmt.Sprintf("Validator with the consensus address %v already exists on the ledger", owner))
}

func ErrValidatorDoesNotExist(owner sdk.ValAddress) *sdkerrors.Error {
	return sdkerrors.Register(Codespace, CodeValidatorDoesNotExist,
		fmt.Sprintf("No validator associated with the account address %v exists on the ledger", owner))
}
*/

func ErrPoolIsFull() *sdkerrors.Error {
	return sdkerrors.Register(Codespace, CodePoolIsFull,
		fmt.Sprintf("Pool ledger already contains maximum number of active nodes: \"%v\"", MaxNodes))
}

/*
func ErrAccountAlreadyHasNode(owner sdk.AccAddress) *sdkerrors.Error {
	return sdkerrors.Register(Codespace, CodeAccountAlreadyHasNode,
		fmt.Sprintf("Validator associated with the account address %v already exists on the ledger", owner))
}
*/
