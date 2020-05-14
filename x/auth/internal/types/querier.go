package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	Request Payload
*/
// QueryAccountParams defines the params for querying accounts.
type QueryAccountParams struct {
	Address sdk.AccAddress
}

// NewQueryAccountParams creates a new instance of QueryAccountParams.
func NewQueryAccountParams(addr sdk.AccAddress) QueryAccountParams {
	return QueryAccountParams{Address: addr}
}

/*
	Response Payload
*/
// Result Payload for an accounts list query
type ListAccountItems struct {
	Total int       `json:"total"`
	Items []Account `json:"items"`
}

// Implement fmt.Stringer
func (n ListAccountItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
