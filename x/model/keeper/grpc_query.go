package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

var _ types.QueryServer = Keeper{}
