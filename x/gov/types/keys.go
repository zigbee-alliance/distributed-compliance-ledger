package types

const (
	// ModuleName is the name of the module
	ModuleName = "gov"

	// StoreKey is the store key string for gov
	StoreKey = ModuleName

	// RouterKey is the message route for gov
	RouterKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName
)

// Keys for governance store
// Items are stored with the following key: values
//
// - 0x00<proposalID_Bytes>: Proposal
//
// - 0x01<proposalID_Bytes>: activeProposalID
//
// - 0x02: nextProposalID
var (
	AllProposalsKeyPrefix    = []byte{0x00}
	ActiveProposalsKeyPrefix = []byte{0x01}

	ProposalIDKey = []byte{0x02}
)
