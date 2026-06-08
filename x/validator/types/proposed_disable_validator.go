package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (disabledValidator ProposedDisableValidator) HasApprovalFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, approval := range disabledValidator.Approvals {
		if approval.Address == addrStr {
			return true
		}
	}

	return false
}

func (disabledValidator ProposedDisableValidator) HasRejectDisableFrom(address sdk.AccAddress) bool {
	addrStr := address.String()
	for _, rejectDisable := range disabledValidator.Rejects {
		if rejectDisable.Address == addrStr {
			return true
		}
	}

	return false
}
