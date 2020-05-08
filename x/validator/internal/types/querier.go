package types

import "encoding/json"

// Response Payload for a list query
type ListValidatorItems struct {
	Total int         `json:"total"`
	Items []Validator `json:"items"`
}

// Implement fmt.Stringer
func (n ListValidatorItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
