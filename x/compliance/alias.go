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
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	NewMsgCertifyModel = types.NewMsgCertifyModel
	NewMsgRevokeModel  = types.NewMsgRevokeModel
	ModuleCdc          = types.ModuleCdc
	RegisterCodec      = types.RegisterCodec
	CertifiedState     = types.Certified
	RevokedState       = types.Revoked
)

type (
	Keeper            = keeper.Keeper
	MsgCertifyModel   = types.MsgCertifyModel
	MsgRevokeModel    = types.MsgRevokeModel
	ComplianceInfo    = types.ComplianceInfo
	CertificationType = types.CertificationType
)
