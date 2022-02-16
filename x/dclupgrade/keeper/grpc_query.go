package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

var _ types.QueryServer = Keeper{}
