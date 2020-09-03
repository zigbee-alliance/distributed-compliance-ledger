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
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisAccounts []auth.Account

// genesis accounts contain an address.
func (gaccs GenesisAccounts) Contains(acc sdk.Address) bool {
	for _, gacc := range gaccs {
		if gacc.Address.Equals(acc) {
			return true
		}
	}

	return false
}
