package settings

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// Default broadcast mode used for write transactions.
	DefaultBroadcastMode = flags.BroadcastBlock
)

// PruningStrategy of the application: Store every state. Keep last two states.
var PruningStrategy = types.NewPruningOptions(2, 1)
