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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateModel{}, "model/CreateModel", nil)
	cdc.RegisterConcrete(&MsgUpdateModel{}, "model/UpdateModel", nil)
	cdc.RegisterConcrete(&MsgDeleteModel{}, "model/DeleteModel", nil)
	cdc.RegisterConcrete(&MsgCreateModelVersion{}, "model/CreateModelVersion", nil)
	cdc.RegisterConcrete(&MsgUpdateModelVersion{}, "model/UpdateModelVersion", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateModel{},
		&MsgUpdateModel{},
		&MsgDeleteModel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateModelVersion{},
		&MsgUpdateModelVersion{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
