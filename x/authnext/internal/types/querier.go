package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type AccountHeader struct {
	Address sdk.AccAddress `json:"address" yaml:"address"`
	PubKey  crypto.PubKey  `json:"public_key" yaml:"public_key"`
}

// Request Payload for a account headers query
type QueryAccountHeadersParams struct {
	Skip int
	Take int
}

func NewQueryAccountHeadersParams(skip int, take int) *QueryAccountHeadersParams {
	return &QueryAccountHeadersParams{Skip: skip, Take: take}
}

// Result Payload for a account headers query
type QueryAccountHeadersResult []AccountHeader

// Implement fmt.Stringer
func (n QueryAccountHeadersResult) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
