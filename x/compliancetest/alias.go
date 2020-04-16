package compliancetest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
	CodeTestingResultDoesNotExist = types.CodeTestingResultDoesNotExist
)

var (
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	NewMsgAddTestingResult        = types.NewMsgAddTestingResult
	ModuleCdc                     = types.ModuleCdc
	RegisterCodec                 = types.RegisterCodec
	ErrTestingResultDoesNotExist  = types.ErrTestingResultDoesNotExist
)

type (
	Keeper              = keeper.Keeper
	MsgAddTestingResult = types.MsgAddTestingResult
	TestingResults      = types.TestingResults
	TestingResult       = types.TestingResult
)
