package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateModel{}, "model/CreateModel", nil)
	cdc.RegisterConcrete(&MsgUpdateModel{}, "model/UpdateModel", nil)
	cdc.RegisterConcrete(&MsgDeleteModel{}, "model/DeleteModel", nil)
	cdc.RegisterConcrete(&MsgCreateModelVersion{}, "model/CreateModelVersion", nil)
	cdc.RegisterConcrete(&MsgUpdateModelVersion{}, "model/UpdateModelVersion", nil)
	cdc.RegisterConcrete(&MsgDeleteModelVersion{}, "model/DeleteModelVersion", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateModel{},
		&MsgUpdateModel{},
		&MsgDeleteModel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateModelVersion{},
		&MsgUpdateModelVersion{},
		&MsgDeleteModelVersion{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
