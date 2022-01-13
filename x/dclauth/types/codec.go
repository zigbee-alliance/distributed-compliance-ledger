// Copyright 2022 DSR Corporation
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
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeAddAccount{}, "dclauth/ProposeAddAccount", nil)
	cdc.RegisterConcrete(&MsgApproveAddAccount{}, "dclauth/ApproveAddAccount", nil)
	cdc.RegisterConcrete(&MsgProposeRevokeAccount{}, "dclauth/ProposeRevokeAccount", nil)
	cdc.RegisterConcrete(&MsgApproveRevokeAccount{}, "dclauth/ApproveRevokeAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()

	cryptocodec.RegisterInterfaces(ModuleCdc.InterfaceRegistry())
}
