package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Request Payload for a list query with pagination
type PaginationParams struct {
	Skip int
	Take int
}

func NewPaginationParams(skip int, take int) PaginationParams {
	return PaginationParams{Skip: skip, Take: take}
}

// Response Payload for a list query with pagination
type LisModelInfoItems struct {
	Total int             `json:"total"`
	Items []ModelInfoItem `json:"items"`
}

// Implement fmt.Stringer
func (n LisModelInfoItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}

type ModelInfoItem struct {
	VID   int16          `json:"vid"`
	PID   int16          `json:"pid"`
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	SKU   string         `json:"sku"`
}

// Response Payload for a list query with pagination
type ListVendorItems struct {
	Total int          `json:"total"`
	Items []VendorItem `json:"items"`
}

// Implement fmt.Stringer
func (n ListVendorItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}

type VendorItem struct {
	VID int16 `json:"vid"`
}
