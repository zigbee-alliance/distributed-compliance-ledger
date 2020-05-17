package validator

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier

	NewValidator  = types.NewValidator
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc
)

type (
	Keeper = keeper.Keeper

	Validator          = types.Validator
	MsgCreateValidator = types.MsgCreateValidator
)
