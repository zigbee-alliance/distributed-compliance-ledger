package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeAddX509RootCert{}, "pki/ProposeAddX509RootCert", nil)
	cdc.RegisterConcrete(&MsgApproveAddX509RootCert{}, "pki/ApproveAddX509RootCert", nil)
	cdc.RegisterConcrete(&MsgAddX509Cert{}, "pki/AddX509Cert", nil)
	cdc.RegisterConcrete(&MsgProposeRevokeX509RootCert{}, "pki/ProposeRevokeX509RootCert", nil)
	cdc.RegisterConcrete(&MsgApproveRevokeX509RootCert{}, "pki/ApproveRevokeX509RootCert", nil)
	cdc.RegisterConcrete(&MsgRevokeX509Cert{}, "pki/RevokeX509Cert", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddX509Cert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeX509Cert{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
