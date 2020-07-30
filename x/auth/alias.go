package auth

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	Vendor                = types.Vendor
	TestHouse             = types.TestHouse
	ZBCertificationCenter = types.ZBCertificationCenter
	Trustee               = types.Trustee
	NodeAdmin             = types.NodeAdmin
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	NewAccount    = types.NewAccount
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
	Roles         = types.Roles
)

type (
	Keeper         = keeper.Keeper
	Account        = types.Account
	PendingAccount = types.PendingAccount
	AccountRole    = types.AccountRole
	AccountRoles   = types.AccountRoles
)
