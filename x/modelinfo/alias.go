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

package modelinfo

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

const (
	ModuleName                 = types.ModuleName
	RouterKey                  = types.RouterKey
	StoreKey                   = types.StoreKey
	CodeModelInfoDoesNotExist  = types.CodeModelInfoDoesNotExist
	CodeModelInfoAlreadyExists = types.CodeModelInfoAlreadyExists
)

var (
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	NewMsgAddModelInfo       = types.NewMsgAddModelInfo
	NewMsgUpdateModelInfo    = types.NewMsgUpdateModelInfo
	ModuleCdc                = types.ModuleCdc
	RegisterCodec            = types.RegisterCodec
	ErrModelInfoDoesNotExist = types.ErrModelInfoDoesNotExist
)

type (
	Keeper             = keeper.Keeper
	MsgAddModelInfo    = types.MsgAddModelInfo
	MsgUpdateModelInfo = types.MsgUpdateModelInfo
	MsgDeleteModelInfo = types.MsgDeleteModelInfo
	ModelInfo          = types.ModelInfo
	VendorProducts     = types.VendorProducts
	ModelInfoItem      = types.ModelInfoItem
	VendorItem         = types.VendorItem
	Model              = types.Model
)
