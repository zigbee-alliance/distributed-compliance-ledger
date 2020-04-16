package types

import (
	"encoding/json"
)

// Response Payload for a list query with pagination
type ListCertifiedModelItems struct {
	Total int              `json:"total"`
	Items []CertifiedModel `json:"items"`
}

// Implement fmt.Stringer
func (n ListCertifiedModelItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
