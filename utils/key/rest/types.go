package rest

import "github.com/cosmos/cosmos-sdk/crypto/keys"

type resultKeyInfos struct {
	Total int              `json:"total"`
	Items []keys.KeyOutput `json:"items"`
}
