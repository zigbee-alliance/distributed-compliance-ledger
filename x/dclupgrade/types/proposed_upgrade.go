package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (acc ProposedUpgrade) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range acc.Approvals {
		if approval == addrStr {
			return true
		}
	}

	return false
}
