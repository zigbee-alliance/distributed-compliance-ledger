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

import sdk "github.com/cosmos/cosmos-sdk/types"

func (upgrade ProposedUpgrade) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range upgrade.Approvals {
		if approval.Address == addrStr {
			return true
		}
	}

	return false
}

func (upgrade ProposedUpgrade) HasRejectFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, reject := range upgrade.Rejects {
		if reject.Address == addrStr {
			return true
		}
	}

	return false
}
