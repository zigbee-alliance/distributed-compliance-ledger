package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

var _ types.QueryServer = Keeper{}
