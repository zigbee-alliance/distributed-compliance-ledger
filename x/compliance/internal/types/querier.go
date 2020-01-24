package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModelInfoHeader struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	SKU   string         `json:"sku"`
}

// Request Payload for a ModelInfo headers query
type QueryModelInfoHeadersParams struct {
	Skip int
	Take int
}

func NewQueryModelInfoHeadersParams(skip int, take int) QueryModelInfoHeadersParams {
	return QueryModelInfoHeadersParams{Skip: skip, Take: take}
}

// Result Payload for a ModelInfo headers query
type QueryModelInfoHeadersResult []ModelInfoHeader

// Implement fmt.Stringer
func (n QueryModelInfoHeadersResult) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
