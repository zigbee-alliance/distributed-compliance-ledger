package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ types.QueryServer = Keeper{}
