package compliancetest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper              = keeper.NewKeeper
	NewQuerier             = keeper.NewQuerier
	NewMsgAddTestingResult = types.NewMsgAddTestingResult
	NewTestingResult       = types.NewTestingResult
	ModuleCdc              = types.ModuleCdc
	RegisterCodec          = types.RegisterCodec
)

type (
	Keeper              = keeper.Keeper
	MsgAddTestingResult = types.MsgAddTestingResult
	TestingResult       = types.TestingResult
	TestingResults      = types.TestingResults
	TestingResultItem   = types.TestingResultItem
)
