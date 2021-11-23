package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeAddAccount{}, "dclauth/ProposeAddAccount", nil)
	cdc.RegisterConcrete(&MsgApproveAddAccount{}, "dclauth/ApproveAddAccount", nil)
	cdc.RegisterConcrete(&MsgProposeRevokeAccount{}, "dclauth/ProposeRevokeAccount", nil)
	cdc.RegisterConcrete(&MsgApproveRevokeAccount{}, "dclauth/ApproveRevokeAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
