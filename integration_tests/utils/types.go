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

package utils

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
)

type ResponseWrapper struct {
	Height string          `json:"height"`
	Result json.RawMessage `json:"result"`
}

type KeyInfo struct {
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Address   sdk.AccAddress `json:"address"`
	PublicKey string         `json:"pubkey"`
}

type AccountInfo struct {
	Address       sdk.AccAddress    `json:"address"`
	AccountNumber string            `json:"account_number"`
	Sequence      string            `json:"sequence"`
	Roles         auth.AccountRoles `json:"roles"`
}

type PendingAccountInfo struct {
	Address   sdk.AccAddress    `json:"address"`
	Roles     auth.AccountRoles `json:"roles"`
	Approvals []sdk.AccAddress  `json:"approvals"`
}

type ModelHeadersResult struct {
	Total string            `json:"total"`
	Items []model.ModelItem `json:"items"`
}

type VendorItemHeadersResult struct {
	Total string             `json:"total"`
	Items []model.VendorItem `json:"items"`
}

type ComplianceInfosHeadersResult struct {
	Total string                      `json:"total"`
	Items []compliance.ComplianceInfo `json:"items"`
}

type ProposedCertificatesHeadersResult struct {
	Total string                    `json:"total"`
	Items []pki.ProposedCertificate `json:"items"`
}

type CertificatesHeadersResult struct {
	Total string            `json:"total"`
	Items []pki.Certificate `json:"items"`
}

type ProposedCertificateRevocationsHeadersResult struct {
	Total string                              `json:"total"`
	Items []pki.ProposedCertificateRevocation `json:"items"`
}

type AccountHeadersResult struct {
	Total string        `json:"total"`
	Items []AccountInfo `json:"items"`
}

type ProposedAccountHeadersResult struct {
	Total string               `json:"total"`
	Items []PendingAccountInfo `json:"items"`
}

type ProposedAccountToRevokeHeadersResult struct {
	Total string                          `json:"total"`
	Items []auth.PendingAccountRevocation `json:"items"`
}

type TxnResponse struct {
	Height string `json:"height"`
	TxHash string `json:"txhash"`
	Code   int    `json:"code"`
}
