package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module.
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete type on the Amino codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgProposeAddX509RootCert{}, ModuleName+"/ProposeAddX509RootCert", nil)
	cdc.RegisterConcrete(MsgApproveAddX509RootCert{}, ModuleName+"/ApproveAddX509RootCert", nil)
	cdc.RegisterConcrete(MsgAddX509Cert{}, ModuleName+"/AddX509Cert", nil)
	cdc.RegisterConcrete(MsgProposeRevokeX509RootCert{}, ModuleName+"/ProposeRevokeX509RootCert", nil)
	cdc.RegisterConcrete(MsgApproveRevokeX509RootCert{}, ModuleName+"/ApproveRevokeX509RootCert", nil)
	cdc.RegisterConcrete(MsgRevokeX509Cert{}, ModuleName+"/RevokeX509Cert", nil)
}
