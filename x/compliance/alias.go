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
	NewKeeper             = keeper.NewKeeper
	NewQuerier            = keeper.NewQuerier
	NewMsgAddModelInfo    = types.NewMsgAddModelInfo
	NewMsgUpdateModelInfo = types.NewMsgUpdateModelInfo
	NewMsgDeleteModelInfo = types.NewMsgDeleteModelInfo
	ModuleCdc             = types.ModuleCdc
	RegisterCodec         = types.RegisterCodec
)

type (
	Keeper             = keeper.Keeper
	MsgAddModelInfo    = types.MsgAddModelInfo
	MsgUpdateModelInfo = types.MsgUpdateModelInfo
	MsgDeleteModelInfo = types.MsgDeleteModelInfo
	QueryModelInfoIDs  = types.QueryModelInfoIDs
	ModelInfo          = types.ModelInfo
)
