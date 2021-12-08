package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateNewVendorInfo{}, "vendorinfo/CreateNewVendorInfo", nil)
	cdc.RegisterConcrete(&MsgUpdateNewVendorInfo{}, "vendorinfo/UpdateNewVendorInfo", nil)
	cdc.RegisterConcrete(&MsgDeleteNewVendorInfo{}, "vendorinfo/DeleteNewVendorInfo", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNewVendorInfo{},
		&MsgUpdateNewVendorInfo{},
		&MsgDeleteNewVendorInfo{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
