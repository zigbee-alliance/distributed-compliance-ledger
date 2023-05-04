package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1.
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCertifyModel{}, "compliance/CertifyModel", nil)
	cdc.RegisterConcrete(&MsgRevokeModel{}, "compliance/RevokeModel", nil)
	cdc.RegisterConcrete(&MsgProvisionModel{}, "compliance/ProvisionModel", nil)
	cdc.RegisterConcrete(&MsgUpdateComplianceInfo{}, "compliance/UpdateComplianceInfo", nil)
	cdc.RegisterConcrete(&MsgDeleteComplianceInfo{}, "compliance/DeleteComplianceInfo", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateComplianceInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteComplianceInfo{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
