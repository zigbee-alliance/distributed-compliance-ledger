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
