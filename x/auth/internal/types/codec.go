package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// ModuleCdc is the codec for the module.
var ModuleCdc = codec.New()

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterCodec registers concrete type on the Amino codec.
func RegisterCodec(cdc *codec.Codec) {
	// register types to be compatible with cosmos transactions builder and processor
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&Account{}, "cosmos-sdk/Account", nil)
	cdc.RegisterConcrete(auth.StdTx{}, "cosmos-sdk/StdTx", nil)

	// register custom types
	cdc.RegisterConcrete(MsgProposeAddAccount{}, ModuleName+"/ProposeAddAccount", nil)
	cdc.RegisterConcrete(MsgApproveAddAccount{}, ModuleName+"/ApproveAddAccount", nil)
}
