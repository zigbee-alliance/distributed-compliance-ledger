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
