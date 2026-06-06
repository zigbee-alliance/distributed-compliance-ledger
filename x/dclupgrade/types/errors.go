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

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/dclupgrade module sentinel errors.
var (
	ErrProposedUpgradeAlreadyExists = errors.Register(ModuleName, 801, "proposed upgrade already exists")
	ErrProposedUpgradeDoesNotExist  = errors.Register(ModuleName, 802, "proposed upgrade does not exist")
	ErrApprovedUpgradeAlreadyExists = errors.Register(ModuleName, 803, "approved upgrade already exists")
)

func NewErrProposedUpgradeAlreadyExists(name interface{}) error {
	return errors.Wrapf(
		ErrProposedUpgradeAlreadyExists,
		"Proposed upgrade with name=%v already exists on the ledger",
		name,
	)
}

func NewErrProposedUpgradeDoesNotExist(name interface{}) error {
	return errors.Wrapf(
		ErrProposedUpgradeDoesNotExist,
		"Proposed upgrade with name=%v does not exist on the ledger",
		name,
	)
}

func NewErrApprovedUpgradeAlreadyExists(name interface{}) error {
	return errors.Wrapf(
		ErrApprovedUpgradeAlreadyExists,
		"Approved upgrade with name=%v already exists on the ledger",
		name,
	)
}
