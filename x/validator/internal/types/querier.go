// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// Request Payload for QueryValidators (pagination and filtering) query.
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

// Response Payload for QueryValidators query.
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

// Implement fmt.Stringer.
func (n ListValidatorItems) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}
