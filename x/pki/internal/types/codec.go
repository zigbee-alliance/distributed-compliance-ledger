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
}
