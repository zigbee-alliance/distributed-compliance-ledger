package utils

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ResponseWrapper struct {
	Height string          `json:"height"`
	Result json.RawMessage `json:"result"`
}

type AccountInfo struct {
	Address       sdk.AccAddress `json:"address"`
	PublicKey     string         `json:"public_key"`
	Roles         []string       `json:"roles"`
	Coins         sdk.Coins      `json:"coins"`
	AccountNumber string         `json:"account_number"`
	Sequence      string         `json:"sequence"`
}

type KeyInfo struct {
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Address   sdk.AccAddress `json:"address"`
	PublicKey string         `json:"pubkey"`
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

type TxnResponse struct {
	Height    string `json:"height"`
	TxHash    string `json:"txhash"`
	Code      int    `json:"code"`
}
