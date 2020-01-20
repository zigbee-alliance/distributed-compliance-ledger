package rest

import auth "github.com/cosmos/cosmos-sdk/x/auth/types"

type DecodeTxsRequest struct {
	Txs []string `json:"txs"`
}

type DecodeTxsResponse struct {
	Txs []auth.StdTx `json:"txs"`
}
