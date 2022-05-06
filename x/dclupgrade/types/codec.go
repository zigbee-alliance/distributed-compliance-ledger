package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeUpgrade{}, "dclupgrade/ProposeUpgrade", nil)
	cdc.RegisterConcrete(&MsgApproveUpgrade{}, "dclupgrade/ApproveUpgrade", nil)
	cdc.RegisterConcrete(&MsgRejectUpgrade{}, "dclupgrade/RejectUpgrade", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeUpgrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveUpgrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRejectUpgrade{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
