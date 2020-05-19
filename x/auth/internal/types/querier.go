package types

//nolint:goimports
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
// Result Payload for accounts list query.
type ListAccountItems struct {
	Total int       `json:"total"`
	Items []Account `json:"items"`
}

// Implement fmt.Stringer.
func (n ListAccountItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for proposed accounts list query.
type ListProposedAccountItems struct {
	Total int              `json:"total"`
	Items []PendingAccount `json:"items"`
}

// Implement fmt.Stringer.
func (n ListProposedAccountItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for single account query.
// It's a hack trick for Codec so that not inserting top-level `type` filed during serialization.
type ZBAccount Account

// Implement fmt.Stringer.
func (a ZBAccount) String() string {
	res, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	return string(res)
}
