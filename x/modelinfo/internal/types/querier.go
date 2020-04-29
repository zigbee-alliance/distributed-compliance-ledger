package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	VID   uint16         `json:"vid"`
	PID   uint16         `json:"pid"`
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
	VID uint16 `json:"vid"`
}
