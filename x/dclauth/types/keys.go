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

var (
	PendingAccountKeyPrefix           = []byte{0x01} // prefix for each key to a pending account
	AccountKeyPrefix                  = []byte{0x02} // prefix for each key to an account
	PendingAccountRevocationKeyPrefix = []byte{0x03} // prefix for each key to a pending account revocation
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AccountStatKey = "AccountStat-value-"
)
