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

package pki

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper                        = keeper.Keeper
	MsgProposeAddX509RootCert     = types.MsgProposeAddX509RootCert
	MsgApproveAddX509RootCert     = types.MsgApproveAddX509RootCert
	MsgAddX509Cert                = types.MsgAddX509Cert
	MsgProposeRevokeX509RootCert  = types.MsgProposeRevokeX509RootCert
	MsgApproveRevokeX509RootCert  = types.MsgApproveRevokeX509RootCert
	MsgRevokeX509Cert             = types.MsgRevokeX509Cert
	Certificate                   = types.Certificate
	Certificates                  = types.Certificates
	ProposedCertificate           = types.ProposedCertificate
	ProposedCertificateRevocation = types.ProposedCertificateRevocation
)
