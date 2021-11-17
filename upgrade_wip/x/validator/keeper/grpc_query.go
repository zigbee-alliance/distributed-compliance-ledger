package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ types.QueryServer = Keeper{}
