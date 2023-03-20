package types

import dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"

// TODO: 1. Move it to separate module  	2. Make it configurable		3. Save into store.
var (
	DisableValidatorPercent     = 2.0 / 3.0
	VoteForDisableValidatorRole = dclauthtypes.Trustee
	EnableDisableValidatorRole  = dclauthtypes.NodeAdmin
)
