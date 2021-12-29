package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCertifyModel{}, "compliance/CertifyModel", nil)
	cdc.RegisterConcrete(&MsgRevokeModel{}, "compliance/RevokeModel", nil)
	cdc.RegisterConcrete(&MsgProvisionModel{}, "compliance/ProvisionModel", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCertifyModel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeModel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvisionModel{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
