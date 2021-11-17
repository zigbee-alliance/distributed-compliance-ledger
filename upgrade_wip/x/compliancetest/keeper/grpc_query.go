package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

var _ types.QueryServer = Keeper{}
