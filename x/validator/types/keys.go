package types

const (
	// ModuleName defines the module name
	ModuleName = "validator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_validator"
)

var (
	ValidatorKeyPrefix           = []byte{0x01} // prefix for each key to a validator
	ValidatorByConsAddrKeyPrefix = []byte{0x02} // prefix for each key to a validator index, by consensus address
	LastValidatorPowerKeyPrefix  = []byte{0x02} // prefix for each key to a validator index, by last power
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
