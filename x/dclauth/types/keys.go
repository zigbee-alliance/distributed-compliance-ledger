package types

const (
	// ModuleName defines the module name
	ModuleName = "dclauth"

	// ModuleName defines the module name to use in user interactions
	ModuleNameUser = "auth"

	// command name for the module
	CmdName = ModuleNameUser

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleNameUser

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleNameUser

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_" + ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AccountStatKey = "AccountStat-value-"
)
