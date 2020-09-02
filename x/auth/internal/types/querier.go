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
// Result Payload for accounts list query.
type ListAccounts struct {
	Total int       `json:"total"`
	Items []Account `json:"items"`
}

// Implement fmt.Stringer.
func (n ListAccounts) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for pending accounts list query.
type ListPendingAccounts struct {
	Total int              `json:"total"`
	Items []PendingAccount `json:"items"`
}

// Implement fmt.Stringer.
func (n ListPendingAccounts) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// Result Payload for pending account revocations list query.
type ListPendingAccountRevocations struct {
	Total int                        `json:"total"`
	Items []PendingAccountRevocation `json:"items"`
}

// Implement fmt.Stringer.
func (n ListPendingAccountRevocations) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// nolint:godox
// Result Payload for single account query.
// It's a hack trick for Codec so that not inserting top-level `type` filed during serialization.
// TODO: Better think regarding Pubkey representation.
type ZBAccount Account

// Implement fmt.Stringer.
func (a ZBAccount) String() string {
	res, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(res)
}
