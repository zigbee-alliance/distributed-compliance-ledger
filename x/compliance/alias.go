package compliance

import (
	"github.com/askolesov/zb-ledger/x/compliance/internal/keeper"
	"github.com/askolesov/zb-ledger/x/compliance/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	NewMsgAddModelInfo = types.NewMsgAddModelInfo
	ModuleCdc          = types.ModuleCdc
	RegisterCodec      = types.RegisterCodec
)

type (
	Keeper            = keeper.Keeper
	MsgAddModelInfo   = types.MsgAddModelInfo
	QueryModelInfoIDs = types.QueryModelInfoIDs
	ModelInfo         = types.ModelInfo
)
