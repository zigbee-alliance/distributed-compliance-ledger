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

package auth

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	Vendor                = types.Vendor
	TestHouse             = types.TestHouse
	ZBCertificationCenter = types.ZBCertificationCenter
	Trustee               = types.Trustee
	NodeAdmin             = types.NodeAdmin
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	NewAccount    = types.NewAccount
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
	Roles         = types.Roles
)

type (
	Keeper                        = keeper.Keeper
	Account                       = types.Account
	PendingAccount                = types.PendingAccount
	PendingAccountRevocation      = types.PendingAccountRevocation
	AccountRole                   = types.AccountRole
	AccountRoles                  = types.AccountRoles
	ListAccounts                  = types.ListAccounts
	ListPendingAccounts           = types.ListPendingAccounts
	ListPendingAccountRevocations = types.ListPendingAccountRevocations
)
