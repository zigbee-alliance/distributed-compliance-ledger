package utils

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SignMessageResponse struct {
	Height string        `json:"height"`
	Result SignedMessage `json:"result"`
}

type SignedMessage struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type GetAccountResponse struct {
	Height string      `json:"height"`
	Result AccountInfo `json:"result"`
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

type GetModelInfoResponse struct {
	Height string               `json:"height"`
	Result compliance.ModelInfo `json:"result"`
}

type GetListModelInfoResponse struct {
	Height string                 `json:"height"`
	Result ModelInfoHeadersResult `json:"result"`
}

type ModelInfoHeadersResult struct {
	Total string                       `json:"total"`
	Items []compliance.ModelInfoHeader `json:"items"`
}
