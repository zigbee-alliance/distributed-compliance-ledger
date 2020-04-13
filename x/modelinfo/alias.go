package modelinfo

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
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
	NewModelInfo          = types.NewModelInfo
	ModuleCdc             = types.ModuleCdc
	RegisterCodec         = types.RegisterCodec
)

type (
	Keeper             = keeper.Keeper
	MsgAddModelInfo    = types.MsgAddModelInfo
	MsgUpdateModelInfo = types.MsgUpdateModelInfo
	MsgDeleteModelInfo = types.MsgDeleteModelInfo
	ModelInfo          = types.ModelInfo
	VendorProducts     = types.VendorProducts
	ModelInfoItem      = types.ModelInfoItem
	VendorItem         = types.VendorItem
)
