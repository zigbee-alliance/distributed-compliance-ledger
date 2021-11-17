package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ types.QueryServer = Keeper{}
