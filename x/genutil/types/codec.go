package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// ModuleCdc defines a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	validator.RegisterCodec(ModuleCdc)
	authtypes.RegisterCodec(ModuleCdc)
	sdk.RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
