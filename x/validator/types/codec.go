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
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateValidator{}, "validator/CreateValidator", nil)
	cdc.RegisterConcrete(&MsgProposeDisableValidator{}, "validator/ProposeDisableValidator", nil)
	cdc.RegisterConcrete(&MsgApproveDisableValidator{}, "validator/ApproveDisableValidator", nil)
	cdc.RegisterConcrete(&MsgDisableValidator{}, "validator/DisableValidator", nil)
	cdc.RegisterConcrete(&MsgEnableValidator{}, "validator/EnableValidator", nil)
	cdc.RegisterConcrete(&MsgRejectDisableValidator{}, "validator/RejectDisableValidator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeDisableValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveDisableValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDisableValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnableValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRejectDisableValidator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc) //nolint:nosnakecase
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
