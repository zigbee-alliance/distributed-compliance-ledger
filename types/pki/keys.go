package types

const (
	// ModuleName defines the module name.
	ModuleName = "pki"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_pki"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ApprovedRootCertificatesKeyPrefix = "ApprovedRootCertificates/value/"
	RevokedRootCertificatesKeyPrefix  = "RevokedRootCertificates/value/"
)

var (
	ApprovedRootCertificatesKey = []byte{0}
	RevokedRootCertificatesKey  = []byte{0}
)
