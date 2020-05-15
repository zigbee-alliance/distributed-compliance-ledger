package authz

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	Administrator         = types.Administrator
	Vendor                = types.Vendor
	TestHouse             = types.TestHouse
	ZBCertificationCenter = types.ZBCertificationCenter
	Trustee               = types.Trustee
	NodeAdmin             = types.NodeAdmin
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgAssignRole = types.NewMsgAssignRole
	NewMsgRevokeRole = types.NewMsgRevokeRole
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper        = keeper.Keeper
	MsgAssignRole = types.MsgAssignRole
	MsgRevokeRole = types.MsgRevokeRole
	AccountRole   = types.AccountRole
	AccountRoles  = types.AccountRoles
)
