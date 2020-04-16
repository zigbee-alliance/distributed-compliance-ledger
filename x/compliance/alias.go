package compliance

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper                        = keeper.NewKeeper
	NewQuerier                       = keeper.NewQuerier
	NewMsgCertifyModel               = types.NewMsgCertifyModel
	NewDeviceCompliance              = types.NewCertifiedModel
	ModuleCdc                        = types.ModuleCdc
	RegisterCodec                    = types.RegisterCodec
	ErrDeviceComplianceAlreadyExists = types.ErrDeviceComplianceAlreadyExists
)

type (
	Keeper                  = keeper.Keeper
	MsgCertifyModel         = types.MsgCertifyModel
	CertifiedModel          = types.CertifiedModel
)
