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

var (
	PoolIsFull                               = sdkerrors.Register(ModuleName, 601, "maximum number of active nodes reached")
	ErrProposedDisableValidatorAlreadyExists = sdkerrors.Register(ModuleName, 801, "disable validator propose already exists")
	ErrProposedDisableValidatorDoesNotExist  = sdkerrors.Register(ModuleName, 802, "disable validator propose does not exist")
	ErrDisabledValidatorAlreadytExists       = sdkerrors.Register(ModuleName, 803, "disabled validator already exist")
	ErrDisabledValidatorDoesNotExist         = sdkerrors.Register(ModuleName, 804, "disabled validator does not exist")
)

func ErrPoolIsFull() error {
	return sdkerrors.Wrapf(PoolIsFull,
		fmt.Sprintf("Pool ledger already contains maximum number of active nodes: \"%v\"", MaxNodes))
}

func NewErrProposedDisableValidatorAlreadyExists(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrProposedDisableValidatorAlreadyExists,
		"Disable proposal with validator address=%v already exists on the ledger",
		name,
	)
}

func NewErrProposedDisableValidatorDoesNotExist(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrProposedDisableValidatorDoesNotExist,
		"Disable proposal with validator address=%v does not exist on the ledger",
		name,
	)
}

func NewErrDisabledValidatorAlreadyExists(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrDisabledValidatorAlreadytExists,
		"Disabled validator with address=%v already exists on the ledger",
		name,
	)
}

func NewErrDisabledValidatorDoesNotExist(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrDisabledValidatorDoesNotExist,
		"Disabled validator with address=%v does not exist on the ledger",
		name,
	)
}
