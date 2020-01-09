package types

import "strings"

// QueryModelInfoIDsResult Queries Result Payload for a names query
type QueryModelInfoIDsResult []string

// implement fmt.Stringer
func (n QueryModelInfoIDsResult) String() string {
	return strings.Join(n[:], "\n")
}
