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
