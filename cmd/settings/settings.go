package settings

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// Default broadcast mode used for write transactions
	DefaultBroadcastMode = flags.BroadcastBlock
)

var (
	// Application prune strategy: Store every state. Keep last two states
	PruningStrategy = types.NewPruningOptions(2, 1)
)
