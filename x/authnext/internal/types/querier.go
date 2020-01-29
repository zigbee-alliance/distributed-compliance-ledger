package types

import (
	"encoding/json"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountHeader struct {
	Address       sdk.AccAddress      `json:"address"`
	PubKey        string              `json:"public_key"`
	Roles         []authz.AccountRole `json:"roles"`
	Coins         sdk.Coins           `json:"coins"`
	AccountNumber uint64              `json:"account_number"`
	Sequence      uint64              `json:"sequence"`
}

// Request Payload for a account headers query
type QueryAccountHeadersParams struct {
	Skip int
	Take int
}

func NewQueryAccountHeadersParams(skip int, take int) QueryAccountHeadersParams {
	return QueryAccountHeadersParams{Skip: skip, Take: take}
}

// Result Payload for a account headers query
type QueryAccountHeadersResult struct {
	Total int             `json:"total"`
	Items []AccountHeader `json:"items"`
}

// Implement fmt.Stringer
func (n QueryAccountHeadersResult) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
