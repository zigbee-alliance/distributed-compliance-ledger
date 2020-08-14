package utils

import (
	"encoding/json"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

type ModelInfoHeadersResult struct {
	Total string                    `json:"total"`
	Items []modelinfo.ModelInfoItem `json:"items"`
}

type VendorItemHeadersResult struct {
	Total string                 `json:"total"`
	Items []modelinfo.VendorItem `json:"items"`
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
