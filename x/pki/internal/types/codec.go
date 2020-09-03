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
