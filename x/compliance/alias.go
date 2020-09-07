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

package compliance

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	NewMsgCertifyModel  = types.NewMsgCertifyModel
	NewMsgRevokeModel   = types.NewMsgRevokeModel
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	CertifiedState      = types.Certified
	RevokedState        = types.Revoked
	ZbCertificationType = types.ZbCertificationType
)

type (
	Keeper                = keeper.Keeper
	MsgCertifyModel       = types.MsgCertifyModel
	MsgRevokeModel        = types.MsgRevokeModel
	ComplianceInfo        = types.ComplianceInfo
	ComplianceInfoKey     = types.ComplianceInfoKey
	ComplianceInfoInState = types.ComplianceInfoInState
	CertificationType     = types.CertificationType
)
