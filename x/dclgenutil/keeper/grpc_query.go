package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

var _ types.QueryServer = Keeper{}
