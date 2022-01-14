package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

var _ types.QueryServer = Keeper{}
