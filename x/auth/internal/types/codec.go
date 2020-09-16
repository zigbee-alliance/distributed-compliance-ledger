// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	cdc.RegisterConcrete(MsgProposeRevokeAccount{}, ModuleName+"/ProposeRevokeAccount", nil)
	cdc.RegisterConcrete(MsgApproveRevokeAccount{}, ModuleName+"/ApproveRevokeAccount", nil)
}
