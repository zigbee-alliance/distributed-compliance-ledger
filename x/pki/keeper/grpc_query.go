package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ types.QueryServer = Keeper{}
