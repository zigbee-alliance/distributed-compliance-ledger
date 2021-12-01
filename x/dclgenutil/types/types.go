package types

// DONTCOVER

const (
	// ModuleName defines the module name
	ModuleName = "dclgenutil"
)

/* FIXME issue 99
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
*/
