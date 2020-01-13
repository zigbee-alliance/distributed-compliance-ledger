package authnext

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	AccountKeeper             = types.AccountKeeper
	AccountHeader             = types.AccountHeader
	QueryAccountHeadersParams = types.QueryAccountHeadersParams
	QueryAccountHeadersResult = types.QueryAccountHeadersResult
)
