package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVendorInfoType{}, "vendorinfo/CreateVendorInfoType", nil)
	cdc.RegisterConcrete(&MsgUpdateVendorInfoType{}, "vendorinfo/UpdateVendorInfoType", nil)
	cdc.RegisterConcrete(&MsgDeleteVendorInfoType{}, "vendorinfo/DeleteVendorInfoType", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVendorInfoType{},
		&MsgUpdateVendorInfoType{},
		&MsgDeleteVendorInfoType{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
