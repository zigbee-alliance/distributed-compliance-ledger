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

package compliancetest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
)

const (
	ModuleName                    = types.ModuleName
	RouterKey                     = types.RouterKey
	StoreKey                      = types.StoreKey
	CodeTestingResultDoesNotExist = types.CodeTestingResultsDoNotExist
)

var (
	NewKeeper                    = keeper.NewKeeper
	NewQuerier                   = keeper.NewQuerier
	NewMsgAddTestingResult       = types.NewMsgAddTestingResult
	ModuleCdc                    = types.ModuleCdc
	RegisterCodec                = types.RegisterCodec
	ErrTestingResultDoesNotExist = types.ErrTestingResultDoesNotExist
)

type (
	Keeper              = keeper.Keeper
	MsgAddTestingResult = types.MsgAddTestingResult
	TestingResults      = types.TestingResults
	TestingResult       = types.TestingResult
)
