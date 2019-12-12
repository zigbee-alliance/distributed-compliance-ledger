package types

import "strings"

// QueryModelInfoIDs Queries Result Payload for a names query
type QueryModelInfoIDs []string

// implement fmt.Stringer
func (n QueryModelInfoIDs) String() string {
	return strings.Join(n[:], "\n")
}
