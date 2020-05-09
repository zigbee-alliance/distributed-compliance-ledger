package types

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
)

/*
	Request Payload
*/

type ValidatorState string

const (
	All    ValidatorState = ""
	Active ValidatorState = "active"
	Jailed ValidatorState = "jailed"
)

// Request Payload for QueryValidators (pagination and filtering) query
type ListValidatorsParams struct {
	Skip  int
	Take  int
	State ValidatorState
}

func NewListValidatorsParams(pagination pagination.PaginationParams, status ValidatorState) ListValidatorsParams {
	return ListValidatorsParams{
		Skip:  pagination.Skip,
		Take:  pagination.Take,
		State: status,
	}
}

/*
	Result Payload
*/

// Response Payload for QueryValidators query
type ListValidatorItems struct {
	Total int         `json:"total"`
	Items []Validator `json:"items"`
}

func NewListValidatorItems() ListValidatorItems {
	return ListValidatorItems{
		Total: 0,
		Items: []Validator{},
	}
}

// Implement fmt.Stringer
func (n ListValidatorItems) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
