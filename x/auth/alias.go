package auth

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
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
	NewAccount       = types.NewAccount
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
	Roles            = types.Roles
)

type (
	Keeper        = keeper.Keeper
	MsgAssignRole = types.MsgAssignRole
	MsgRevokeRole = types.MsgRevokeRole
	Account       = types.Account
	AccountRole   = types.AccountRole
	AccountRoles  = types.AccountRoles
)
