package types

const (
	// ModuleName defines the module name
	ModuleName = "compliancetest"

	// StoreKey defines the primary module store key
	StoreKey = "xxx_" + ModuleName // FIXME compliancetests conflicts with compliance as a key prefix

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "xxx_" + "mem_compliancetest" // FIXME issue 99: the same is likely here too
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
