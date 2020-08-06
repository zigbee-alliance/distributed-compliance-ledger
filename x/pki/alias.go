package pki

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper               = keeper.NewKeeper
	NewQuerier              = keeper.NewQuerier
	ModuleCdc               = types.ModuleCdc
	RegisterCodec           = types.RegisterCodec
	RootCertificate         = types.RootCertificate
	IntermediateCertificate = types.IntermediateCertificate
)

type (
	Keeper                    = keeper.Keeper
	MsgProposeAddX509RootCert = types.MsgProposeAddX509RootCert
	MsgApproveAddX509RootCert = types.MsgApproveAddX509RootCert
	MsgAddX509Cert            = types.MsgAddX509Cert
	MsgRevokeX509Cert         = types.MsgRevokeX509Cert
	Certificate               = types.Certificate
	Certificates              = types.Certificates
	ProposedCertificate       = types.ProposedCertificate
)
